package response

type NewsReponse struct {
	Subject      string            `json:"subject"`
	TotalResults int64             `json:"totalResults"`
	Articles     []ArticleResponse `json:"articles"`
}

type ArticleResponse struct {
	Source      ArticleSourceReponse `json:"source"`
	Author      string               `json:"author"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	URL         string               `json:"url"`
	URLToImage  string               `json:"urlToImage"`
	PublishedAt string               `json:"publishedAt"`
	Content     string               `json:"content"`
}

type ArticleSourceReponse struct {
	Id   *string `json:"id"`
	Name string  `json:"name"`
}
