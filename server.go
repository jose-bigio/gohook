package gohook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (s *Server) verifyAuth(body []byte, req *http.Request) bool {
	signature := req.Header.Get("X-Hub-Signature")
	if signature == "" {
		return false
	}
	mac := hmac.New(sha1.New, []byte(s.Secret))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	expectedSig := "sha1=" + hex.EncodeToString(expectedMAC)
	if !hmac.Equal([]byte(expectedSig), []byte(signature)) {
		return false
	}
	return true
}

func (s *Server) processPacket(eventType EventType, respBody []byte) {
	var payload interface{}
	switch eventType {
	case PushEventType:
		payload = &PushEvent{}
	case PingEventType:
		payload = &PingEvent{}
	case CommitCommentEventType:
		payload = &CommitCommentEvent{}
	case IssueCommentEventType:
		payload = &IssueCommentEvent{}
	case IssuesEventType:
		payload = &IssuesEvent{}
	case CreateEventType:
		payload = &CreateEvent{}
	case DeleteEventType:
		payload = &DeleteEvent{}
	case RepositoryEventType:
		payload = &RepositoryEvent{}
	default:
		log.Printf("Attempt to process unknown packet type: %s", eventType)
		return
	}
	if err := json.Unmarshal(respBody, &payload); err != nil {
		panic(err)
	}
	et := &EventAndType{
		Event: payload,
		Type:  eventType,
	}
	s.EventAndTypes <- et
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

	respBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "500 Internal Error - Could not read from request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if s.Secret != "" && !s.verifyAuth(respBody, req) {
		http.Error(w, "403 Forbidden - Verification of secret failed.", http.StatusForbidden)
		return
	}
	go s.processPacket(EventType(eventType), respBody)
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
