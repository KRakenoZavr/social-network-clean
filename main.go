package main

import (
	"mux/internal"
)

func main() {
	port := ":3000"

	server := internal.NewServer()
	server.Start(port)
}
