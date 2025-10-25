package model

type News struct {
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	PublishedOn string  `json:"publishedOn"`
	ImageUrl    *string `json:"imageUrl"`
	NewsUrl     string  `json:"newsUrl"`
	Content     string  `json:"content"`
	Source      string  `json:"source"`
}

type Headline struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}
