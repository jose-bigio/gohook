package gohook

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestPingGithub(t *testing.T) {
	server := NewServer(8888, "secret", "/postreceive")
	server.GoListenAndServe()
	_, err := http.Post("https://api.github.com/orgs/fireside-chat/hooks/4719659/pings", "application/json", nil)
	if err != nil {
		t.Errorf("Error requesting ping: %s", err)
	}
	var packet *EventAndType
	select {
	case packet = <-server.EventAndTypes:
		if packet.Type != PingEventType {
			t.Errorf("Incorrect packet type, got %s", packet.Type)
		}
	case <-time.After(time.Duration(3) * time.Second):
		t.Error("Timeout when waiting for ping.")
	}
	payload, ok := packet.Event.(*PingEvent)
	if !ok {
		t.Error("Error asserting payload as *PingEvent.")
	}
	if !payload.Hook.Active {
		t.Error("Incorrect payload.Hook.Active value.")
	}
}

func TestPingEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_ping.json")
	if err != nil {
		t.Errorf("Error reading sample ping file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(PingEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*PingEvent)
	if !ok {
		t.Error("Error asserting payload as *PingEvent.")
	}
	if payload.Zen != "Approachable is better than simple." {
		t.Error("Incorrect PingEvent.Zen value.")
	}
	if payload.Hook.Config.ContentType != "json" {
		t.Error("Incorrect PingEvent.Hook.Config.ContentType value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling PingEvent: %s", err)
	}
}

func TestPushEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_push.json")
	if err != nil {
		t.Errorf("Error reading sample push file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(PushEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*PushEvent)
	if !ok {
		t.Error("Error asserting payload as *PushEvent.")
	}
	if payload.Ref != "refs/heads/master" {
		t.Error("Incorrect PushEvent.Ref value.")
	}
	if payload.Commits[0].Author.Name != "M. Cameron Palone" {
		t.Error("Incorrect payload.Commits[0].Author.Name value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling PushEvent: %s", err)
	}
}

func TestWatchEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_watch.json")
	if err != nil {
		t.Errorf("Error reading sample watch file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(WatchEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*WatchEvent)
	if !ok {
		t.Error("Error asserting payload as *WatchEvent.")
	}
	if payload.Action != "started" {
		t.Error("Incorrect payload.Action value.")
	}
	if payload.Sender.ID != 16893 {
		t.Error("Incorrect payload.Sender.ID value.")
	}
	if payload.Organization.ID != 12191882 {
		t.Error("Incorrect payload.Organization.ID value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling WatchEvent: %s", err)
	}
}

func TestCommitCommentEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_commit_comment.json")
	if err != nil {
		t.Errorf("Error reading sample commit_comment file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(CommitCommentEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*CommitCommentEvent)
	if !ok {
		t.Error("Error asserting payload as *CommitCommentEvent.")
	}
	if payload.Comment.Body != "This is a really good change! :+1:" {
		t.Error("Incorrect payload.Comment.Body value.")
	}
	if payload.Comment.User.Login != "baxterthehacker" {
		t.Error("Incorrect payload.Comment.User.Login value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling CommitCommentEvent: %s", err)
	}
}

func TestTeamAddEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_team_add.json")
	if err != nil {
		t.Errorf("Error reading sample team_add file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(TeamAddEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*TeamAddEvent)
	if !ok {
		t.Error("Error asserting payload as *TeamAddEvent.")
	}
	if payload.Team.ID != 1474248 {
		t.Error("Incorrect payload.Team.ID value.")
	}
	if payload.Team.Slug != "owners" {
		t.Error("Incorrect payload.Team.Slug value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling TeamAddEvent: %s", err)
	}
}

func TestCreateEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_create.json")
	if err != nil {
		t.Errorf("Error reading sample create file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(CreateEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*CreateEvent)
	if !ok {
		t.Error("Error asserting payload as *CreateEvent.")
	}
	if payload.Description !=
		"Stripped down version of euphoria.io's heim written in Go." {
		t.Error("Incorrect payload.Description value.")
	}
	if !bytes.Equal(payload.Repository.CreatedAt, []byte("\"2015-05-01T18:43:26Z\"")) {
		t.Errorf("Incorrect payload.Repository.CreatedAt value. Got %s, expected %s", payload.Repository.CreatedAt, []byte("2015-05-01T18:43:26Z"))
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling CreatedEvent: %s", err)
	}
}

func TestRepositoryEvent(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/sample_repository.json")
	if err != nil {
		t.Errorf("Error reading sample create file: %s", err)
	}
	server := NewServer(8888, "secret", "path")
	server.processPacket(RepositoryEventType, raw)
	packet := <-server.EventAndTypes
	payload, ok := packet.Event.(*RepositoryEvent)
	if !ok {
		t.Error("Error asserting payload as *CreateEvent.")
	}
	if payload.Action != "created" {
		t.Error("Incorrect payload.Action value.")
	}
	if payload.Organization.Login != "fireside-chat" {
		t.Error("Incorrect payload.Organization.Login value.")
	}
	_, err = json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshalling TeamAddEvent: %s", err)
	}
}
