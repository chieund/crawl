package crawl

type DataTag struct {
	Title string
	Slug  string
}

type DataArticle struct {
	Title string
	Slug  string
	Link  string
	Image string
	Tags  []DataTag
}
