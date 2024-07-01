package response

type NewsClientResponse struct {
	Subject      string
	TotalResults int64
	Articles     []ArticleResponse
}

type ArticleResponse struct {
	Source      ArticleSourceResponse
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponse struct {
	Id   *string
	Name string
}
