package models

type Bookmark struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
	TargetBlank string `json:"targetBlank"`
}
