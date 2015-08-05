// Package gohook provides a simple http server receive github webhooks.
//
// Simply start a Server and receive from its EventAndTypes channel, where
// correctly unmarshalled packets will be available along with its type. The
// type is provided because of the limitations of Go's type system-
// unfortunately, it is necessary to do a type assertion based on the type
// received. You can see a commented example of this in the examples/ directory.
package gohook
