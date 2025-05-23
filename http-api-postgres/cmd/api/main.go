package main

import (
	"log"
	"net/http"

	"http-api-postgres/internal/app/articleapi"
	"http-api-postgres/internal/pkg/storage"
)

func main() {
	db, err := storage.NewPostgresDB("postgres://postgres:postgres@localhost:5432/articles?sslmode=disable")
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	router := articleapi.NewRouter(db)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", router)
}
