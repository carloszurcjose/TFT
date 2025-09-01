package main

import (
	"log"

	"github.com/carloszurcjose/TFT/internal/httpx"
	"github.com/carloszurcjose/TFT/internal/storage"
)

func main() {
	db := storage.OpenPostgres() //connects using env vars
	router := httpx.NewRouter(db)

	log.Println("Listening on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
