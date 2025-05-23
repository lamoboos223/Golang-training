package service

import (
	"structured-http-api/internal/models"
)

// ArticleService handles business logic for articles
type ArticleService struct {
	repo models.ArticleRepository
}

// NewArticleService creates a new article service
func NewArticleService(repo models.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

// GetAllArticles returns all articles
func (s *ArticleService) GetAllArticles() []models.Article {
	return s.repo.GetAll()
}

// GetArticleByID returns an article by its ID
func (s *ArticleService) GetArticleByID(id string) (models.Article, bool) {
	return s.repo.GetByID(id)
}

// CreateArticle creates a new article
func (s *ArticleService) CreateArticle(article models.Article) models.Article {
	return s.repo.Create(article)
}

// DeleteArticle deletes an article by its ID
func (s *ArticleService) DeleteArticle(id string) bool {
	return s.repo.Delete(id)
}
