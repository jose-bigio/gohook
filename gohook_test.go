package gohook

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

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
		t.Error("Error asserting payload as *PingEvent")
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
