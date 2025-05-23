package articleapi

import (
	"encoding/json"
	"net/http"

	"http-api-postgres/internal/pkg/article"
	"http-api-postgres/internal/pkg/storage"
)

func CreateArticleHandler(db *storage.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		svc := article.NewService(db)
		err := svc.Create(req.Title, req.Content)
		if err != nil {
			http.Error(w, "Could not save article", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"ok"}`))
	}
}

func GetAllArticlesHandler(db *storage.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		svc := article.NewService(db)
		articles, err := svc.GetAll()
		if err != nil {
			http.Error(w, "Could not get articles", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(articles)
	}
}
