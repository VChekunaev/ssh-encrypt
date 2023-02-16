package keys

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

// PublicKey is a struct that represents an ssh public key
type PublicKey struct {
	Name string
	Key  ssh.PublicKey
}

// LoadPublicKeys loads all public keys from a given path
func LoadPublicKeys(keysPath string) ([]PublicKey, error) {
	var publicKeys []PublicKey

	if _, err := os.Stat(keysPath); os.IsNotExist(err) {
		return publicKeys, errors.New("failed to find key directory")
	}

	files, err := ioutil.ReadDir(keysPath)
	if err != nil {
		return publicKeys, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".pub") {
			continue
		}

		filePath := filepath.Join(keysPath, file.Name())

		publicKeyBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return publicKeys, err
		}

		publicKey, err := ssh.ParsePublicKey(publicKeyBytes)
		if err != nil {
			return publicKeys, err
		}

		publicKeys = append(publicKeys, PublicKey{
			Name: strings.TrimSuffix(file.Name(), ".pub"),
			Key:  publicKey,
		})
	}

	if len(publicKeys) == 0 {
		return publicKeys, errors.New("no public keys found")
	}

	return publicKeys, nil
}

// ValidatePublicKey checks if the given public key is a valid ssh public key
func ValidatePublicKey(keyBytes []byte) error {
	_, err := ssh.ParsePublicKey(keyBytes)
	return err
}
