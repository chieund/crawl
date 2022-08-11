package crawl

type DataTag struct {
	Title string
	Slug  string
}

type DataArticle struct {
	Title       string
	Slug        string
	Link        string
	Image       string
	Description string
	Content     string
	WebsiteId   int
	WebsiteSlug string
	Tags        []DataTag
}
