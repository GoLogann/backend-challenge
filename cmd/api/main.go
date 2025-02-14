package main

import (
	"backend-challenge/internal/config"
	"backend-challenge/internal/repository"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	mongoDB, err := repository.NewMongoDB(cfg.MongoURI, cfg.Database)
	if err != nil {
		log.Fatalf("❌ Error connecting to MongoDB: %v", err)
	}
	defer mongoDB.Close()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
