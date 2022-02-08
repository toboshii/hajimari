package tplutil

import (
	"fmt"
	"github.com/toboshii/hajimari/internal/hajimari"
)

// greet returns the greeting to be used in the h1 heading
func Greet(name string, currentHour int) (greet string) {
	switch currentHour / 6 {
	case 0:
		greet = "Good night"
	case 1:
		greet = "Good morning"
	case 2:
		greet = "Good afternoon"
	default:
		greet = "Good evening"
	}

	if name != "" {
		return fmt.Sprintf("%s, %s!", greet, name)
	}

	return fmt.Sprintf("%s!", greet)
}

// Checks if status is null indicating disabled status check
func IsStatusCheckEnabled(app hajimari.App) bool {
	return app.Status != nil
}
