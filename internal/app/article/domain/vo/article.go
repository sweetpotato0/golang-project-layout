package vo

// Article .
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ArticleUser .
type ArticleUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
