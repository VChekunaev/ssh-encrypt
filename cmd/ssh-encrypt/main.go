package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/username/ssh-encrypt/encryption"
	"github.com/username/ssh-encrypt/keys"
)

func main() {
	var keysFile string
	flag.StringVar(&keysFile, "keys", "", "path to file with ssh keys")
	flag.Parse()

	if keysFile == "" {
		keysFile = filepath.Join(os.Getenv("HOME"), ".ssh", "authorized_keys")
		if _, err := os.Stat(keysFile); err != nil {
			fmt.Println("error: cannot find ssh keys file")
			os.Exit(1)
		}
	}

	keyData, err := keys.Parse(keysFile)
	if err != nil {
		fmt.Printf("error: could not parse ssh keys: %v\n", err)
		os.Exit(1)
	}

	encryptedKeys, err := encryption.EncryptKeys(keyData)
	if err != nil {
		fmt.Printf("error: could not encrypt ssh keys: %v\n", err)
		os.Exit(1)
	}

	encryptedFile := fmt.Sprintf("%s.enc", keysFile)

	err = encryption.WriteToFile(encryptedFile, encryptedKeys)
	if err != nil {
		fmt.Printf("error: could not write encrypted keys to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("SSH keys successfully encrypted and saved to %s\n", encryptedFile)
}
