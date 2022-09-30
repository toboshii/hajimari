package models

type App struct {
	Name        string      `json:"name"`
	Icon        string      `json:"icon"`
	URL         string      `json:"url"`
	Info        string      `json:"info"`
	TargetBlank string      `json:"targetBlank"`
	Replicas    ReplicaInfo `json:"replicas"`
}
