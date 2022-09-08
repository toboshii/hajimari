package models

type ReplicaInfo struct {
	Total     int     `json:"total,omitempty"`
	Available int     `json:"available,omitempty"`
	PctReady  float64 `json:"pctReady,omitempty"`
}
