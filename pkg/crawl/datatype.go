package crawl

type DataTag struct {
	Name string
}

type DataArticle struct {
	Title string
	Slug  string
	Link  string
	Image string
	Tags  []DataTag
}
