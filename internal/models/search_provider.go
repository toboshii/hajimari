package models

type SearchProvider struct {
	Name      string `json:"name"`
	Token     string `json:"token"`
	Icon      string `json:"icon"`
	SearchUrl string `json:"searchUrl"`
	URL       string `json:"url"`
}
