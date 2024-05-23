package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joshuahenriques/go-ecom/config"
	"github.com/joshuahenriques/go-ecom/db"
)

func main() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBAddress, config.Envs.DBName)
	db, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"posgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
