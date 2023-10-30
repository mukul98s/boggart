package helper

import "strings"

func RemoveWhitespace(inputString string) string {
	return strings.Trim(inputString, "\t\n\r")
}
