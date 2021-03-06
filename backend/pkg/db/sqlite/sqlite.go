package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(migration bool, dbName string) *sql.DB {
	// create db file if not exists
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		_, err = os.Create(dbName)
		if err != nil {
			log.Fatalf("Cannot create db file, err: %s", err)
		}
		fmt.Println("Created db file")
	}

	// open database with foreign keys on
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbName))
	if err != nil {
		log.Fatalf("Cannot open db, err: %s", err)
	}

	if !migration {
		return db
	}

	// create sqilte driver
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Cannot create driver, err: %s", err)
	}

	// migration
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations/",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Cannot create instance, err: %s", err)
	}

	// migration down
	err = m.Down()
	if err != nil {
		fmt.Println("Cannot migrate down")
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Cannot migrate down, err: %s", err)
		}
	}

	// migration up
	err = m.Up()
	if err != nil {
		fmt.Println("Cannot migrate up")
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Cannot migrate up, err: %s", err)
		}
	}

	return db
}
