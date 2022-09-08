package models

type AppGroup struct {
	Name string `json:"name"`
	Apps  []App  `json:"apps"`
}
