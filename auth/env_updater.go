package auth

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/logger"
)

// UpdateEnvFile busca .env desde directorio actual, crea backup y actualiza tokens.
func UpdateEnvFile(ctx context.Context, newAccessToken, newRefreshToken string) error {
	logger.LogAPIOperation(ctx, "update_env_file", "local", 0, false, map[string]interface{}{
		"updating_tokens": true,
	})

	// Encontrar el archivo .env
	envPath, err := findEnvFile()
	if err != nil {
		return fmt.Errorf("failed to find .env file: %w", err)
	}

	// Leer el archivo actual
	lines, err := readEnvFile(envPath)
	if err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	// Crear backup del archivo original
	backupPath := envPath + ".backup." + time.Now().Format("20060102-150405")
	if err := createBackup(envPath, backupPath); err != nil {
		logger.LogAPIOperation(ctx, "update_env_file", "local", 0, false, map[string]interface{}{
			"backup_error": err.Error(),
		})
		// Continuar sin backup, pero log el error
	}

	// Actualizar las líneas
	updatedLines := updateTokenLines(lines, newAccessToken, newRefreshToken)

	// Escribir el archivo actualizado
	if err := writeEnvFile(envPath, updatedLines); err != nil {
		return fmt.Errorf("failed to write updated .env file: %w", err)
	}

	logger.LogAPIOperation(ctx, "update_env_file", "local", 0, true, map[string]interface{}{
		"env_file_path": envPath,
		"backup_path":   backupPath,
	})

	return nil
}

// findEnvFile busca .env en ../env-mercado-libre-go-sdk/ y luego sube por el árbol.
func findEnvFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Buscar subiendo por el árbol hasta encontrar el archivo
	currentDir := dir
	for {
		// Primero intentar ../env-mercado-libre-go-sdk/ desde este nivel
		envDir := filepath.Join(filepath.Dir(currentDir), "env-mercado-libre-go-sdk")
		envPath := filepath.Join(envDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		// Luego intentar .env en este directorio
		envPath = filepath.Join(currentDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			// Reached root directory
			return "", fmt.Errorf(".env file not found in ../env-mercado-libre-go-sdk/ or parent directories")
		}
		currentDir = parent
	}
}

// readEnvFile lee todas las líneas del archivo .env.
func readEnvFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// updateTokenLines reemplaza líneas existentes o agrega al final si no existen.
func updateTokenLines(lines []string, newAccessToken, newRefreshToken string) []string {
	var updatedLines []string
	accessTokenUpdated := false
	refreshTokenUpdated := false

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "ACCESS_TOKEN="):
			updatedLines = append(updatedLines, "ACCESS_TOKEN="+newAccessToken)
			accessTokenUpdated = true
		case strings.HasPrefix(line, "REFRESH_TOKEN="):
			updatedLines = append(updatedLines, "REFRESH_TOKEN="+newRefreshToken)
			refreshTokenUpdated = true
		default:
			updatedLines = append(updatedLines, line)
		}
	}

	// Si no existían las variables, agregarlas al final
	if !accessTokenUpdated {
		updatedLines = append(updatedLines, "ACCESS_TOKEN="+newAccessToken)
	}
	if !refreshTokenUpdated {
		updatedLines = append(updatedLines, "REFRESH_TOKEN="+newRefreshToken)
	}

	return updatedLines
}

// writeEnvFile escribe las líneas actualizadas al archivo .env.
func writeEnvFile(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	return nil
}

// createBackup crea una copia de seguridad del archivo .env.
func createBackup(originalPath, backupPath string) error {
	originalFile, err := os.Open(originalPath)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	backupFile, err := os.Create(backupPath)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	scanner := bufio.NewScanner(originalFile)
	writer := bufio.NewWriter(backupFile)
	defer writer.Flush()

	for scanner.Scan() {
		if _, err := writer.WriteString(scanner.Text() + "\n"); err != nil {
			return err
		}
	}

	return scanner.Err()
}

// RefreshAndUpdateTokens retorna tokens renovados aunque falle la actualización del .env.
func RefreshAndUpdateTokens(ctx context.Context, clientID, clientSecret, currentRefreshToken string) (*RefreshTokenResponse, error) {
	// Hacer refresh del token
	response, err := RefreshAccessToken(ctx, clientID, clientSecret, currentRefreshToken)
	if err != nil {
		return nil, err
	}

	// Actualizar el archivo .env
	if err := UpdateEnvFile(ctx, response.AccessToken, response.RefreshToken); err != nil {
		logger.LogAPIOperation(ctx, "refresh_and_update", clientID, 0, false, map[string]interface{}{
			"env_update_error": err.Error(),
		})
		// Retornar el response exitoso pero log el error de actualización del .env
		return response, fmt.Errorf("token refreshed successfully but failed to update .env: %w", err)
	}

	return response, nil
}
