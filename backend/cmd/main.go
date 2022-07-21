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
	runEnv := os.Getenv("RUN_ENV")
	withMigration := false

	if migration == "true" {
		withMigration = true
	}

	if runEnv == "test" {
		defer func() {
			os.Remove(dbName)
		}()
	}

	server := internal.NewServer(withMigration, dbName)
	err := server.Start(port)
	if err != nil {
		log.Println(err)
		return
	}
}
