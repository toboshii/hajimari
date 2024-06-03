package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ApplicationSpec `json:"spec"`
}

type ApplicationSpec struct {
	Name        string `json:"name"`
	Group       string `json:"group,omitempty"`
	Icon        string `json:"icon,omitempty"`
	URL         string `json:"url"`
	Info        string `json:"info,omitempty"`
	TargetBlank bool   `json:"targetBlank,omitempty"`
	Location    int    `json:"location"`
}

type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Application `json:"items"`
}
