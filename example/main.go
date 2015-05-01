package main

import (
	"fmt"

	"github.com/cpalone/gohook"
)

func main() {
	// Create and start the server.
	server := gohook.NewServer(8000, "secret", "/postreceive")
	server.GoListenAndServe()

	// Loop continuously, blocking on a receive from server.EventAndTypes
	for {
		eAndT := <-server.EventAndTypes

		// switch on the packet type and use the appropriate type assertion
		// to access the data in the event.
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
