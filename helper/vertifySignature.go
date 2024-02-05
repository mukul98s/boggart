package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mukul98s/boggart/data"
	"io"
	"os"
)

func VerifySignature(link data.VersionLink, filepath string) error {
	fmt.Println("🦥 Verifying Signature")
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return err
	}

	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	if link.Sha256Sum != hashString {
		return fmt.Errorf("signature Verfication Failed")
	}

	return nil
}
