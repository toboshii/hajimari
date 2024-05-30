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

// NormalizeString trims extra spaces and changes the string to lower-case
func NormalizeString(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

// CompareNormalized compares two strings after normalizing them
func CompareNormalized(a string, b string) int {
	return strings.Compare(NormalizeString(a), NormalizeString(b))
}
