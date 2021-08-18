package tplutil

import "fmt"

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
