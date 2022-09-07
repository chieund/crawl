package typesense

type ArticleJson struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	Image           string   `json:"image"`
	Link            string   `json:"link"`
	IsUpdateContent int32    `json:"is_update_content"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	Tags            []string `json:"tags"`
	Website         string   `json:"website"`
}

type TagJson struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type WebsiteJson struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type ArticleUpdateJson struct {
	IsUpdateContent int32  `json:"is_update_content"`
	UpdatedAt       string `json:"updated_at"`
}
