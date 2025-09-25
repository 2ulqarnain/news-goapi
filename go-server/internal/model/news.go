package model

type News struct {
	Slug        string  `json:"slug"`
	Headline    string  `json:"headline"`
	PublishedOn string  `json:"publishedOn"`
	ImageUrl    *string `json:"imageUrl"`
	NewsUrl     string  `json:"newsUrl"`
	Body        string  `json:"body"`
}
