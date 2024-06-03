package annotations

const (
	// HajimariIconAnnotation const used for hajimari icon
	HajimariIconAnnotation = "hajimari.io/icon"
	// HajimariEnableAnnotation const used for checking whether an ingress is exposed to hajimari
	HajimariEnableAnnotation = "hajimari.io/enable"
	// HajimariAppNameAnnotation const used for overriding the name of the app
	HajimariAppNameAnnotation = "hajimari.io/appName"
	// HajimariGroupAnnotation const used for overriding group
	HajimariGroupAnnotation = "hajimari.io/group"
	// HajimariInstanceAnnotation const used for defining which instance of hajimari to use
	HajimariInstanceAnnotation = "hajimari.io/instance"
	// HajimariURLAnnotation const used for specifying the URL for the hajimari app
	HajimariURLAnnotation = "hajimari.io/url"
	// HajimariInfoAnnotation const used for specifying the info line for the hajimari app
	HajimariInfoAnnotation = "hajimari.io/info"
	// HajimariStatusCheckAnnotation boolean used for enabling status indicators.
	HajimariStatusCheckEnabledAnnotation = "hajimari.io/statusCheckEnabled"
	// HajimariTargetBlankAnnotation boolean used for making links open in a new window.
	HajimariTargetBlankAnnotation = "hajimari.io/targetBlank"
	// HajimariLocationAnnotation int used for specifying the location of the app in the list
	HajimariLocationAnnotation = "hajimari.io/location"
)
