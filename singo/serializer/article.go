package serializer

import "singo/model"


// Article 文章序列化器
type Article struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	UserName  string `json:"username"`
	Markdown string `json:"markdown"`
	CreatedAt int64  `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(item model.Article) Article {
	return Article{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		UserName:  item.UserName,
		Markdown:	item.Markdown,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildArticles 序列化文章列表
func BuildArticles(items []model.Article) (Articles []Article) {
	for _, item := range items {
		Article := BuildArticle(item)
		Articles = append(Articles, Article)
	}
	return Articles
}
