package main

import (
	"mux/internal"
)

func main() {
	port := ":3333"

	server := internal.NewServer()
	server.Start(port)
}
