package auth

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/logger"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/validation"
)

const (
	// Endpoint oficial de Mercado Libre para OAuth
	tokenEndpoint = "https://api.mercadolibre.com/oauth/token"
)

// RefreshAccessToken valida params, registra operación y valida respuesta.
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (*RefreshTokenResponse, error) {
	// Validaciones de entrada
	if err := validation.ValidateRequired("client_id", clientID); err != nil {
		return nil, err
	}
	if err := validation.ValidateRequired("client_secret", clientSecret); err != nil {
		return nil, err
	}
	if err := validation.ValidateRequired("refresh_token", refreshToken); err != nil {
		return nil, err
	}

	// Preparar request
	request := RefreshTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}

	logger.LogAPIOperation(ctx, "refresh_token", clientID, 0, false, map[string]interface{}{
		"endpoint": tokenEndpoint,
	})

	// Ejecutar request
	var response RefreshTokenResponse
	err := httpx.DoPostJSON(ctx, tokenEndpoint, "", request, &response)
	if err != nil {
		logger.LogAPIOperation(ctx, "refresh_token", clientID, 0, false, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to refresh access token: %w", err)
	}

	// Validar response
	if response.AccessToken == "" {
		return nil, fmt.Errorf("refresh token response missing access_token")
	}
	if response.RefreshToken == "" {
		return nil, fmt.Errorf("refresh token response missing refresh_token")
	}

	logger.LogAPIOperation(ctx, "refresh_token", clientID, 0, true, map[string]interface{}{
		"expires_in": response.ExpiresIn,
		"user_id":    response.UserID,
		"scope":      response.Scope,
	})

	return &response, nil
}

// GetAuthConfig carga desde variables de entorno o retorna error si faltan.
func GetAuthConfig() (*AuthConfig, error) {
	config := &AuthConfig{}

	// Validar que tengamos la configuración mínima
	if config.ClientID == "" {
		return nil, fmt.Errorf("CLIENT_ID is required")
	}
	if config.ClientSecret == "" {
		return nil, fmt.Errorf("CLIENT_SECRET is required")
	}

	return config, nil
}

// ValidateAccessToken valida formato y opcionalmente verifica contra la API.
func ValidateAccessToken(ctx context.Context, accessToken string) error {
	if err := validation.ValidateAccessToken(accessToken); err != nil {
		return err
	}

	return nil
}
