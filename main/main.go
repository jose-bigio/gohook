package main

import (
	"fmt"

	"github.com/cpalone/gohook"
)

func main() {
	// rawJSON, err := ioutil.ReadFile("../sample_push.json")
	// if err != nil {
	// 	panic(err)
	// }
	// var packet gohook.PushEvent
	// if err := json.Unmarshal(rawJSON, &packet); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(packet.Repository.Organization)
	server := gohook.NewServer(8000, "secret", "/postreceive")
	server.GoListenAndServe()
	for {
		eAndT := <-server.EventAndTypes
		switch eAndT.Type {
		case gohook.PingEventType:
			packet, ok := eAndT.Event.(*gohook.PingEvent)
			if !ok {
				panic("Could not assert PingEvent as such.")
			}
			fmt.Println(packet.Organization.Login)
		case gohook.PushEventType:
			packet, ok := eAndT.Event.(*gohook.PushEvent)
			if !ok {
				panic("Could not assert PushEvent as such.")
			}
			fmt.Println(packet.Organization.Login)
		}
	}
}
