package helper

import "strconv"

func IsValidPhpVersion(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}
