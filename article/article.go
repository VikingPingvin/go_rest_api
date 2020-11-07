package article

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func InitializeArticles() *[]Article {
	Articles = []Article{
		Article{Title: "Title 1", Desc: "Description 1", Content: "Content 1"},
		Article{Title: "Title 2", Desc: "Description 2", Content: "Content 2"},
	}

	return &Articles
}

func ReturnAllArticles() *[]Article {
	return &Articles
}

func ReturnArticle(id int) *Article {
	return &Articles[id]
}

func AddArticle(art Article) bool {
	Articles = append(Articles, art)
	return true
}
