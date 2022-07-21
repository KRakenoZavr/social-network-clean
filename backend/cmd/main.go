package main

import (
	"os"

	"mux/internal"
)

func main() {
	port := os.Getenv("port")

	server := internal.NewServer()
	server.Start(port)
}
