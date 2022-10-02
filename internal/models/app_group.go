package models

type AppGroup struct {
	Group string `json:"group"`
	Apps  []App  `json:"apps"`
}
