package models

type Template struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	HtmlBody string `json:"htmlBody"`
}
