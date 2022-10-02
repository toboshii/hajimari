package models

type ReplicaInfo struct {
	Total     int     `json:"total"`
	Available int     `json:"available"`
	PctReady  float64 `json:"pctReady"`
}
