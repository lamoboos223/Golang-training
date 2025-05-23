package main

import (
	"log"
	"net/http"

	"structured-http-api/internal/handlers"
	"structured-http-api/internal/models"
	"structured-http-api/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize dependencies
	repo := models.NewInMemoryArticleRepository()
	articleService := service.NewArticleService(repo)
	articleHandler := handlers.NewArticleHandler(articleService)

	// Setup router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", articleHandler.HomePage).Methods("GET")
	router.HandleFunc("/articles", articleHandler.GetArticles).Methods("GET")
	router.HandleFunc("/article", articleHandler.CreateArticle).Methods("POST")
	router.HandleFunc("/article/{id}", articleHandler.GetArticle).Methods("GET")
	router.HandleFunc("/article/{id}", articleHandler.DeleteArticle).Methods("DELETE")

	// Start server
	log.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
