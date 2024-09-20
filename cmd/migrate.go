package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	direction := flag.String("direction", "up", "Specify migration direction: up or down")
	flag.Parse()

	m, err := migrate.New(
		"file://./migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if *direction == "up" {
		fmt.Println("Upgrade")
		err = m.Up()
	} else {
		err = m.Down()
		fmt.Println("Down")
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Successfully migrated!")
}
