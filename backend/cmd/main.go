package main

import (
	"log"
	"os"

	"mux/internal"
)

func main() {
	port := os.Getenv("BACKEND_PORT")
	migration := os.Getenv("MIGRATION")
	dbName := os.Getenv("DB_NAME")
	withMigration := false

	if migration == "true" {
		withMigration = true
	}

	server := internal.NewServer(withMigration, dbName)
	err := server.Start(port)
	if err != nil {
		log.Println(err)
		return
	}
}
