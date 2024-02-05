package helper

import "strings"

func ExtractMajorPhpVersion(version string) string {
	numbers := strings.Split(version, ".")
	return strings.Join(numbers[:2], ".")
}
