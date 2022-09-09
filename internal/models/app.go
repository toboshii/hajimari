package models

type App struct {
	Name     string      `json:"name"`
	Icon     string      `json:"icon"`
	URL      string      `json:"url"`
	Info     string      `json:"info"`
	Replicas ReplicaInfo `json:"replicas"`
}
