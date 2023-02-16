package config

import "os"

var (
	// DefaultSSHPublicKeyPath is the default path to look for SSH public keys.
	DefaultSSHPublicKeyPath = os.Getenv("HOME") + "/.ssh"

	// DefaultSSHPublicKeyExtension is the default file extension for SSH public keys.
	DefaultSSHPublicKeyExtension = "pub"

	// DefaultSSHEncryptedKeyPath is the default path to save encrypted SSH private keys.
	DefaultSSHEncryptedKeyPath = os.Getenv("HOME") + "/.ssh"

	// DefaultSSHEncryptedKeyExtension is the default file extension for encrypted SSH private keys.
	DefaultSSHEncryptedKeyExtension = "enc"
)
