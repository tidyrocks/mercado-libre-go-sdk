package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrEmptyAccessToken = errors.New("access token cannot be empty")
	ErrInvalidItemID    = errors.New("invalid item ID format")
	ErrInvalidSiteID    = errors.New("invalid site ID format")
	ErrInvalidEmail     = errors.New("invalid email format")
)

// Patrones regex para validación
var (
	itemIDPattern     = regexp.MustCompile(`^ML[A-Z]\d+$`) // MLA123456, MLM789012
	siteIDPattern     = regexp.MustCompile(`^ML[A-Z]$`)    // MLA, MLM, MLB
	categoryIDPattern = regexp.MustCompile(`^ML[A-Z]\d+$`) // MLM1051
	emailPattern      = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
)

// ValidateAccessToken valida el formato del access token
func ValidateAccessToken(token string) error {
	if strings.TrimSpace(token) == "" {
		return ErrEmptyAccessToken
	}

	// Validar formato APP_USR-*
	if !strings.HasPrefix(token, "APP_USR-") {
		return fmt.Errorf("invalid access token format: must start with APP_USR-")
	}

	if len(token) < 50 {
		return fmt.Errorf("access token too short: expected at least 50 characters")
	}

	return nil
}

// ValidateItemID valida el formato de ID de ítem
func ValidateItemID(itemID string) error {
	if strings.TrimSpace(itemID) == "" {
		return fmt.Errorf("item ID cannot be empty")
	}

	if !itemIDPattern.MatchString(itemID) {
		return fmt.Errorf("%w: must match pattern ML[A-Z]\\d+ (e.g., MLA123456)", ErrInvalidItemID)
	}

	return nil
}

// ValidateSiteID valida el formato de ID de sitio
func ValidateSiteID(siteID string) error {
	if strings.TrimSpace(siteID) == "" {
		return fmt.Errorf("site ID cannot be empty")
	}

	if !siteIDPattern.MatchString(siteID) {
		return fmt.Errorf("%w: must match pattern ML[A-Z] (e.g., MLA, MLM)", ErrInvalidSiteID)
	}

	return nil
}

// ValidateCategoryID valida el formato de ID de categoría
func ValidateCategoryID(categoryID string) error {
	if strings.TrimSpace(categoryID) == "" {
		return fmt.Errorf("category ID cannot be empty")
	}

	if !categoryIDPattern.MatchString(categoryID) {
		return fmt.Errorf("invalid category ID format: must match pattern ML[A-Z]\\d+ (e.g., MLM1051)")
	}

	return nil
}

// ValidateEmail valida formato de email básico
func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if !emailPattern.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

// ValidateRequired valida que campos requeridos no estén vacíos
func ValidateRequired(fieldName, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required and cannot be empty", fieldName)
	}
	return nil
}

// ValidateStringLength valida longitud de string
func ValidateStringLength(fieldName, value string, minLen, maxLen int) error {
	length := len(strings.TrimSpace(value))

	if length < minLen {
		return fmt.Errorf("%s must be at least %d characters long", fieldName, minLen)
	}

	if length > maxLen {
		return fmt.Errorf("%s must be at most %d characters long", fieldName, maxLen)
	}

	return nil
}

// ValidatePositiveNumber valida que un número sea positivo
func ValidatePositiveNumber(fieldName string, value float64) error {
	if value <= 0 {
		return fmt.Errorf("%s must be a positive number", fieldName)
	}
	return nil
}

// ValidateQuantity valida cantidad (entero no negativo)
func ValidateQuantity(fieldName string, value int) error {
	if value < 0 {
		return fmt.Errorf("%s cannot be negative", fieldName)
	}
	return nil
}
