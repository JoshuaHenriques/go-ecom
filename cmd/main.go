package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/joshuahenriques/go-ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":3000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
