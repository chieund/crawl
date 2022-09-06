package typesense

type ArticleJson struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
	Link  string `json:"link"`
	//Tags []TagJson `json:"tags"`
	//Website	WebsiteJson `json:"website"`
}

type TagJson struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type WebsiteJson struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
