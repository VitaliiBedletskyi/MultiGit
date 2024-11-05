package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDefaultSSHKeyPath() (string, error) {
	// Common default SSH key files
	keyFiles := []string{"id_rsa", "id_dsa", "id_ecdsa", "id_ed25519"}
	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")

	// Check if SSH agent is running
	if os.Getenv("SSH_AUTH_SOCK") != "" {
		fmt.Println("SSH agent is running.")
	}

	// Search for the first existing private key file
	for _, keyFile := range keyFiles {
		keyPath := filepath.Join(sshDir, keyFile)
		if _, err := os.Stat(keyPath); err == nil {
			fmt.Printf("Using default SSH key: %s\n", keyPath)
			return keyPath, nil
		}
	}

	return "", fmt.Errorf("no default SSH key found in %s", sshDir)
}
