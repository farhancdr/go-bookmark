package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func getOrCreateConfigDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	homeDir := currentUser.HomeDir

	configDir := filepath.Join(homeDir, ".config")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create config directory %s: %w", configDir, err)
		}
	} else if err != nil {
		return "", fmt.Errorf("failed to check config directory %s: %w", configDir, err)
	}

	return configDir, nil
}

func GetConfigDir() (string, error) {
	configDir, err := getOrCreateConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "bm"), nil
}
