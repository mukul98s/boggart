package helper

import (
	"errors"
	"fmt"
	"github.com/mukul98s/boggart/data"
)

func GetDownloadLink(version string) (data.VersionLink, error) {
	majorVersion := ExtractMajorPhpVersion(version)
	link, exists := data.Versions[majorVersion]
	if !exists {
		return data.VersionLink{}, errors.New(fmt.Sprintf("Download link not found for PHP version %s", version))
	}
	return link, nil
}
