package articleapi

import (
	"net/http"

	"http-api-postgres/internal/pkg/storage"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db *storage.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/articles", CreateArticleHandler(db))
	r.Get("/articles", GetAllArticlesHandler(db))
	return r
}
