package helper

import (
	"fmt"
)

func InstallPhp(version string) (string, error) {
	// get the download link
	link, err := GetDownloadLink(version)
	if err != nil {
		return "", err
	}

	// download the file
	filename := fmt.Sprintf("php-%s.tar.gz", version)
	downloadedPath, err := DownloadPhp(link, filename)
	if err != nil {
		return "", err
	}

	err = VerifySignature(link, downloadedPath)
	if err != nil {
		return "", err
	}

	err = ExtractDownloadedBinary(downloadedPath)
	if err != nil {
		return "", err
	}

	// TODO: install the binary
	InstallPhpBinary()
	// TODO: remove the downloaded files

	return "", nil
}
