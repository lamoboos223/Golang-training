package storage

func (db *DB) SaveArticle(title, content string) error {
	_, err := db.conn.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", title, content)
	return err
}

func (db *DB) GetAllArticles() ([]Article, error) {
	rows, err := db.conn.Query("SELECT id, title, content FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
