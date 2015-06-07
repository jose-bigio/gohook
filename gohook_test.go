package gohook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SendRequest(method, url, secret, eventHeader string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Github-Event", eventHeader)
	if secret != "" {
		mac := hmac.New(sha1.New, []byte(secret))
		_, err := mac.Write(body)
		if err != nil {
			return nil, err
		}
		macVal := mac.Sum(nil)
		sig := "sha1=" + hex.EncodeToString(macVal)
		req.Header.Add("X-Hub-Signature", sig)
	}
	client := http.Client{}
	return client.Do(req)
}

func TestCreateEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_create.json")
	if err != nil {
		t.Errorf("Error reading sample create file: %s", err)
	}
	hookServer := NewServer(8888, "secret", "/path")
	server := httptest.NewServer(hookServer)
	defer server.Close()
	resp, err := SendRequest("POST", server.URL+"/path", "secret", "create", raw)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Bad status code: %s", resp.StatusCode)
	}
	select {
	case et := <-hookServer.EventAndTypes:
		if et.Type != CreateEventType {
			t.Errorf("Incorrect event type from server: %s", et.Type)
		}
		_, ok := et.Event.(*CreateEvent)
		if !ok {
			t.Error("Could not assert event as *CreateEvent.")
		}
		return
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event from server.")
	}
}

func TestInvalidRequests(t *testing.T) {
	correctBody, err := ioutil.ReadFile("testdata/sample_create.json")
	if err != nil {
		t.Errorf("Error reading sample create file: %s", err)
	}
	hookServer := NewServer(8888, "secret", "/path")
	server := httptest.NewServer(hookServer)
	defer server.Close()

	// check wrong method
	resp, err := SendRequest("GET", server.URL+"/path", "secret", "create", correctBody)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 405 {
		t.Errorf("Expected 405 Method Not Allowed, got: %s", resp.StatusCode)
	}

	// check wrong path
	resp, err = SendRequest("POST", server.URL+"/wrongpath", "secret", "create", correctBody)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 404 {
		t.Errorf("Expected 404 Not Found, got: %s", resp.StatusCode)
	}

	// check wrong secret
	resp, err = SendRequest("POST", server.URL+"/path", "wrongsecret", "create", correctBody)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 403 {
		t.Errorf("Expected 403 Unauthorized, got: %s", resp.StatusCode)
	}

	// check missing event type
	resp, err = SendRequest("POST", server.URL+"/path", "secret", "", correctBody)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Errorf("Expected 400 Bad Request, got: %s", resp.StatusCode)
	}

	// check invalid body
	invalidBody := []byte("{{}")

	resp, err = SendRequest("POST", server.URL+"/path", "secret", "create", correctBody)
	if err != nil {
		t.Errorf("SendRequest: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Errorf("Expected 400 Bad Request, got: %s", resp.StatusCode)
	}
}
