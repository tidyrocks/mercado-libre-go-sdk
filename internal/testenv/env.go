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

// loadEnv busca .env en ../env-mercado-libre-go-sdk/ y luego sube por el árbol.
func loadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}

	// Buscar subiendo por el árbol hasta encontrar el proyecto root
	currentDir := dir
	for {
		// Primero intentar ../env-mercado-libre-go-sdk/ desde este nivel
		envDir := filepath.Join(filepath.Dir(currentDir), "env-mercado-libre-go-sdk")
		envPath := filepath.Join(envDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			loadEnvFile(envPath)
			return
		}

		// Luego intentar .env en este directorio
		envPath = filepath.Join(currentDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			loadEnvFile(envPath)
			return
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			// Reached root directory
			fmt.Println("Warning: .env file not found in ../env-mercado-libre-go-sdk/ or parent directories")
			return
		}
		currentDir = parent
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
