package tplutil

import (
	"text/template"
)

// Map of Utility functions for use in the template
var (
	TplfuncMap = template.FuncMap {
		"deref": func(s *string) string{ return *s },
	}
)
