package helper

import (
	"fmt"
	"github.com/mukul98s/boggart/data"
	"io"
	"net/http"
	"os"
)

func DownloadPhp(link data.VersionLink, filename string) (string, error) {
	fmt.Println("🦥 Downloading selected version...")
	downloadPath := fmt.Sprintf("./%s", filename)
	if _, err := os.Stat(downloadPath); err == nil {
		return downloadPath, nil
	}

	response, err := http.Get(link.DownloadLink)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad Response: %v", response.StatusCode)
	}

	file, err := os.Create(downloadPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		deleteDownloadedBinary(downloadPath)
		return "", err
	}

	return downloadPath, nil
}

func deleteDownloadedBinary(path string) {
	err := os.Remove(path)
	if err != nil {
		return
	}
	fmt.Println("DELETED DOWNLOADED BINARY")
}
