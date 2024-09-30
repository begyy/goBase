package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"goBase/app/config"
	"log"
)

func main() {
	config.Load()
	direction := flag.String("direction", "up", "Specify migration direction: up or down")
	flag.Parse()

	m, err := migrate.New(
		"file://./migrations",
		config.AppConfig.DatabaseURL,
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

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Successfully migrated!")
}
