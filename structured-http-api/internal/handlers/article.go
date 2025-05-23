package handlers

import (
	"encoding/json"
	"net/http"

	"structured-http-api/internal/models"
	"structured-http-api/internal/service"

	"github.com/gorilla/mux"
)

// ArticleHandler handles HTTP requests for articles
type ArticleHandler struct {
	service *service.ArticleService
}

// NewArticleHandler creates a new article handler
func NewArticleHandler(service *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

// HomePage handles the root endpoint
func (h *ArticleHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the HomePage!")
}

// GetArticles handles GET /articles
func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := h.service.GetAllArticles()
	json.NewEncoder(w).Encode(articles)
}

// GetArticle handles GET /article/{id}
func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	article, found := h.service.GetArticleByID(params["id"])
	if !found {
		json.NewEncoder(w).Encode(models.Article{})
		return
	}
	json.NewEncoder(w).Encode(article)
}

// CreateArticle handles POST /article
func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	createdArticle := h.service.CreateArticle(article)
	json.NewEncoder(w).Encode(createdArticle)
}

// DeleteArticle handles DELETE /article/{id}
func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	h.service.DeleteArticle(params["id"])
	articles := h.service.GetAllArticles()
	json.NewEncoder(w).Encode(articles)
}
