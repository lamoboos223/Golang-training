package models

// Article represents a blog article
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ArticleRepository defines the interface for article data operations
type ArticleRepository interface {
	GetAll() []Article
	GetByID(id string) (Article, bool)
	Create(article Article) Article
	Delete(id string) bool
}

// InMemoryArticleRepository implements ArticleRepository with in-memory storage
type InMemoryArticleRepository struct {
	articles []Article
}

// NewInMemoryArticleRepository creates a new in-memory article repository
func NewInMemoryArticleRepository() *InMemoryArticleRepository {
	return &InMemoryArticleRepository{
		articles: []Article{
			{ID: "1", Title: "Hello", Content: "Article Content"},
			{ID: "2", Title: "Hello 2", Content: "Article Content 2"},
		},
	}
}

// GetAll returns all articles
func (r *InMemoryArticleRepository) GetAll() []Article {
	return r.articles
}

// GetByID returns an article by its ID
func (r *InMemoryArticleRepository) GetByID(id string) (Article, bool) {
	for _, article := range r.articles {
		if article.ID == id {
			return article, true
		}
	}
	return Article{}, false
}

// Create adds a new article
func (r *InMemoryArticleRepository) Create(article Article) Article {
	r.articles = append(r.articles, article)
	return article
}

// Delete removes an article by its ID
func (r *InMemoryArticleRepository) Delete(id string) bool {
	for i, article := range r.articles {
		if article.ID == id {
			r.articles = append(r.articles[:i], r.articles[i+1:]...)
			return true
		}
	}
	return false
}
