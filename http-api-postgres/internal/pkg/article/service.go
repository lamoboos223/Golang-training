package article

import (
	"http-api-postgres/internal/pkg/storage"
)

type Service struct {
	db *storage.DB
}

func NewService(db *storage.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(title, content string) error {
	return s.db.SaveArticle(title, content)
}

func (s *Service) GetAll() ([]storage.Article, error) {
	return s.db.GetAllArticles()
}
