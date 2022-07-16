package main

import (
	"errors"

	"mux/internal"
)

var (
	ErrMethodMismatch = errors.New("method is not allowed")
	ErrNotFound       = errors.New("no matching route was found")
)

func main() {
	port := ":3333"

	server := internal.NewServer()
	server.Start(port)
}
