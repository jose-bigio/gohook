# gohook
Introduction
---
gohook is a small library for working with Github webhooks in Go. Currently, ping events and push events are processed, but all others result in a 500 status code and an explanation that the type is not yet implemented.

gohook is loosely based on phayes' [hookserve.](https://github.com/phayes/hookserve)

