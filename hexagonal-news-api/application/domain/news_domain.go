package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Subject      string
	TotalResults int64
	Articles     []Article
}

type Article struct {
	Source      ArticleSource
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type ArticleSource struct {
	Id   *string
	Name string
}
