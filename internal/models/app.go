package models

type App struct {
	Name        string      `json:"name"`
	Icon        string      `json:"icon"`
	IconColor   string      `json:"iconColor"`
	URL         string      `json:"url"`
	Info        string      `json:"info"`
	TargetBlank bool        `json:"targetBlank"`
	Replicas    ReplicaInfo `json:"replicas"`
	Location    int         `json:"location"`
}
