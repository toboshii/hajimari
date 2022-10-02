package frontend

import (
	"embed"
)

// Embed the build directory from the frontend.
//
//go:embed all:build
var BuildFs embed.FS
