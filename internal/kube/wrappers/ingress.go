package wrappers

import (
	"net/url"
	"strings"

	"github.com/ullbergm/hajimari/internal/annotations"
	utilStrings "github.com/ullbergm/hajimari/internal/util/strings"
	v1 "k8s.io/api/networking/v1"
)

// IngressWrapper struct wraps a kubernetes ingress object
type IngressWrapper struct {
	ingress *v1.Ingress
}

// NewIngressWrapper func creates an instance of IngressWrapper
func NewIngressWrapper(ingress *v1.Ingress) *IngressWrapper {
	return &IngressWrapper{
		ingress: ingress,
	}
}

// GetAnnotationValue extracts an annotation's value present on the ingress wrapped by the object
func (iw *IngressWrapper) GetAnnotationValue(annotationKey string) string {
	if value, ok := iw.ingress.Annotations[annotationKey]; ok {
		return value
	}
	return ""
}

// GetName func extracts name of the ingress wrapped by the object
func (iw *IngressWrapper) GetName() string {
	if nameFromAnnotation := iw.GetAnnotationValue(annotations.HajimariAppNameAnnotation); nameFromAnnotation != "" {
		return nameFromAnnotation
	}
	return iw.ingress.ObjectMeta.Name
}

// GetNamespace func extracts namespace of the ingress wrapped by the object
func (iw *IngressWrapper) GetNamespace() string {
	return iw.ingress.ObjectMeta.Namespace
}

// GetGroup func extracts group name from the ingress
func (iw *IngressWrapper) GetGroup() string {
	if groupFromAnnotation := iw.GetAnnotationValue(annotations.HajimariGroupAnnotation); groupFromAnnotation != "" {
		return utilStrings.NormalizeString(groupFromAnnotation)
	}
	return utilStrings.NormalizeString(iw.GetNamespace())
}

// GetGroup func extracts group name from the ingress
func (iw *IngressWrapper) GetInfo() string {
	if infoFromAnnotation := iw.GetAnnotationValue(annotations.HajimariInfoAnnotation); infoFromAnnotation != "" {
		return infoFromAnnotation
	}
	return ""
}

// GetStatusCheckEnabled func extracts statusCheck feature gate from the ingress
// @default true
func (iw *IngressWrapper) GetStatusCheckEnabled() bool {
	if statusCheckEnabledFromAnnotation := iw.GetAnnotationValue(annotations.HajimariStatusCheckEnabledAnnotation); statusCheckEnabledFromAnnotation != "" {
		return utilStrings.ParseBool(statusCheckEnabledFromAnnotation)
	}
	return true
}

// GetTargetBlank func extracts open in new window feature gate from the ingress
// @default false
func (iw *IngressWrapper) GetTargetBlank() bool {
	if targetBlankFromAnnotation := iw.GetAnnotationValue(annotations.HajimariTargetBlankAnnotation); targetBlankFromAnnotation != "" {
		return utilStrings.ParseBool(targetBlankFromAnnotation)
	}
	return false
}

// GetURL func extracts url of the ingress wrapped by the object
func (iw *IngressWrapper) GetURL() string {

	if urlFromAnnotation := iw.GetAnnotationValue(annotations.HajimariURLAnnotation); urlFromAnnotation != "" {
		parsedURL, err := url.ParseRequestURI(urlFromAnnotation)
		if err != nil {
			logger.Warn(err)
			return ""
		}
		return parsedURL.String()
	}

	if !iw.rulesExist() {
		logger.Warn("No rules exist in ingress: ", iw.ingress.GetName())
		return ""
	}

	var url string

	if host, exists := iw.tryGetTLSHost(); exists { // Get TLS Host if it exists
		url = host
	} else {
		url = iw.getHost() // Fallback for normal Host
	}

	// Append port + ingressSubPath
	url += iw.getIngressSubPath()

	return url
}

func (iw *IngressWrapper) rulesExist() bool {
	if iw.ingress.Spec.Rules != nil && len(iw.ingress.Spec.Rules) > 0 {
		return true
	}
	return false
}

func (iw *IngressWrapper) tryGetTLSHost() (string, bool) {
	if iw.supportsTLS() {
		return "https://" + iw.ingress.Spec.TLS[0].Hosts[0], true
	}

	return "", false
}

func (iw *IngressWrapper) supportsTLS() bool {
	if iw.ingress.Spec.TLS != nil && len(iw.ingress.Spec.TLS) > 0 {
		return true
	}
	return false
}

func (iw *IngressWrapper) getHost() string {
	return "http://" + iw.ingress.Spec.Rules[0].Host
}

func (iw *IngressWrapper) getIngressSubPath() string {
	rule := iw.ingress.Spec.Rules[0]
	if rule.HTTP != nil {
		if rule.HTTP.Paths != nil && len(rule.HTTP.Paths) > 0 {
			return strings.TrimRight(rule.HTTP.Paths[0].Path, "/")
		}
	}
	return ""
}
