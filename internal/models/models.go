package models

type Startpage struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
