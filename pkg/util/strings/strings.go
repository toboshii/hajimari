package strings

import (
	"strconv"
	"strings"
)

func ParseBool(str string) bool {
	boolVal, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return boolVal
}

// ContainsBetweenDelimiter checks if the search string is present between delimiters of a full string
func ContainsBetweenDelimiter(fullString string, search string, delimiter string) bool {
	splitted := strings.Split(fullString, delimiter)
	for _, split := range splitted {
		if split == search {
			return true
		}
	}
	return false
}
