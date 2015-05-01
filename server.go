package gohook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Server is the basic type that listens for webhook events and pushes them onto
// the EventAndTypes channel. You should create a Server through the NewServer
// function.
type Server struct {
	Port          int
	Secret        string
	Path          string
	EventAndTypes chan *EventAndType
}

// NewServer creates a new server with the given settings. If secret is the
// zero-length string, no security verification will be done. Path should be
// specified with a leading slash.
func NewServer(port int, secret string, path string) *Server {
	return &Server{
		Port:          port,
		Secret:        secret,
		Path:          path,
		EventAndTypes: make(chan *EventAndType, 5),
	}
}

// ServeHTTP implements http.Handler interface on Server. You should never need
// to call this yourself.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != "POST" {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
	}

	if req.URL.Path != s.Path {
		http.Error(w, "404 Not found- expected "+s.Path, http.StatusNotFound)
	}

	eventType := req.Header.Get("X-Github-Event")
	if eventType == "" {
		http.Error(w, "400 Bad Request - Missing X-Github-Event Header", http.StatusBadRequest)
		return
	}
	if eventType != "ping" && eventType != "push" {
		http.Error(w, "500 Internal Server Error - Event type not yet implemented.", http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "500 Internal Error - Could not read from request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if s.Secret != "" {
		signature := req.Header.Get("X-Hub-Signature")
		if signature == "" {
			http.Error(w, "403 Forbidden- missing secret.", http.StatusForbidden)
			return
		}
		mac := hmac.New(sha1.New, []byte(s.Secret))
		mac.Write(body)
		expectedMAC := mac.Sum(nil)
		expectedSig := "sha1=" + hex.EncodeToString(expectedMAC)
		if !hmac.Equal([]byte(expectedSig), []byte(signature)) {
			http.Error(w, "403 Forbidden - Verification of secret failed.", http.StatusForbidden)
			return
		}
	}
	var payload interface{}
	switch EventType(eventType) {
	case PushEventType:
		payload = &PushEvent{}
	case PingEventType:
		payload = &PingEvent{}
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "500 Internal Error - Could not unmarshal request: "+err.Error(), http.StatusInternalServerError)
	}
	go func() {
		et := &EventAndType{
			Event: payload,
			Type:  EventType(eventType),
		}
		s.EventAndTypes <- et
	}()
}

// ListenAndServe simply returns the error received from a call of
// http.ListenAndServe() with the correct parameters. Use this function if you
// want the call to block.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", s.Port), s)
}

// GoListenAndServe is a convenience function that runs Server.ListenAndServe()
// in a goroutine and panics upon receipt of an error. Use this function if
// you do not want the call to block.
func (s *Server) GoListenAndServe() {
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}
