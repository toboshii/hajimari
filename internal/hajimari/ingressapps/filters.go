package ingressapps

import (
	"github.com/toboshii/hajimari/internal/annotations"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/util/strings"
	v1 "k8s.io/api/networking/v1"
)

// For filtering ingresses having hajimari enable annotation set to true
func byHajimariEnableAnnotation(ingress v1.Ingress, appConfig config.Config) bool {
	if appConfig.DefaultEnable {
		if val, ok := ingress.Annotations[annotations.HajimariEnableAnnotation]; ok {
			// Has Hajimari annotation and is enabled
			if val == "false" {
				return false
			}
		}
		return true
	} else {
		if val, ok := ingress.Annotations[annotations.HajimariEnableAnnotation]; ok {
			// Has Hajimari annotation and is enabled
			if val == "true" {
				return true
			}
		}
		return false
	}
}

// For filtering ingresses by hajimari instance
func byHajimariInstanceAnnotation(ingress v1.Ingress, appConfig config.Config) bool {
	if val, ok := ingress.Annotations[annotations.HajimariInstanceAnnotation]; ok {
		return strings.ContainsBetweenDelimiter(val, appConfig.InstanceName, ",")
	}
	return false
}
