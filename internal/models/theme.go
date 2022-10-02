package models

type Theme struct {
	Name            string `json:"name"`
	BackgroundColor string `json:"backgroundColor"`
	PrimaryColor    string `json:"primaryColor"`
	AccentColor     string `json:"accentColor"`
}
