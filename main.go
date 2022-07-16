package main

import (
	"mux/internal"
)

func main() {
	port := ":3001"

	server := internal.NewServer()
	server.Start(port)
}
