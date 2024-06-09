package models

type Bookmark struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
	TargetBlank bool   `json:"targetBlank"`
	Location    int    `json:"location"`
}
