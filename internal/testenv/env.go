package testenv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// OAuth Configuration
	ClientID     string
	ClientSecret string
	RefreshToken string

	// Current tokens
	AccessToken string

	// Test data
	TestItemID string
)

func init() {
	loadEnv()
}

// loadEnv busca .env subiendo por el árbol de directorios en init().
func loadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			loadEnvFile(envPath)
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			fmt.Println("Warning: .env file not found")
			return
		}
		dir = parent
	}
}

// loadEnvFile parsea cada línea KEY=VALUE e ignora comentarios.
func loadEnvFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening .env file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "CLIENT_ID":
			ClientID = value
		case "CLIENT_SECRET":
			ClientSecret = value
		case "REFRESH_TOKEN":
			RefreshToken = value
		case "ACCESS_TOKEN":
			AccessToken = value
		case "TEST_ITEM_ID":
			TestItemID = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading .env file: %v\n", err)
	}
}
