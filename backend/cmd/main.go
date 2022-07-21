package main

import (
	"mux/internal"
	"os"
)

func main() {
	port := os.Getenv("port")

	server := internal.NewServer()
	server.Start(port)
}
