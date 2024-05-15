package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joshuahenriques/go-ecom/cmd/api"
	"github.com/joshuahenriques/go-ecom/config"
	"github.com/joshuahenriques/go-ecom/db"
)

func main() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBAddress, config.Envs.DBName)
	db, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
