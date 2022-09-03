package models

type App struct {
	Name				string `json:"name"`
	Icon				string `json:"icon"`
	URL					string `json:"url"`
	Replicas			int `json:"replicas"`
	AvailableReplicas	int `json:"availableReplicas"`
}

type ReplicaInfo struct {
	Total		int
	Available	int
}
