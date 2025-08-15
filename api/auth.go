package api

import (
	"context"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const tokenEndpoint = "https://api.mercadolibre.com/oauth/token"

// Token representa un token de acceso de la API de Mercado Libre
type Token struct {
	AccessToken  string `json:"access_token"`  // Token de acceso
	TokenType    string `json:"token_type"`    // Tipo de token ("Bearer")
	ExpiresIn    int    `json:"expires_in"`    // Segundos hasta expiración (21600 = 6 horas)
	Scope        string `json:"scope"`         // Permisos ("offline_access read write")
	UserID       int64  `json:"user_id"`       // ID del usuario
	RefreshToken string `json:"refresh_token"` // Token para renovar
}

// refreshTokenRequest representa la solicitud de renovación de token (uso interno)
type refreshTokenRequest struct {
	GrantType    string `json:"grant_type"`    // "refresh_token"
	ClientID     string `json:"client_id"`     // ID de la aplicación
	ClientSecret string `json:"client_secret"` // Secret de la aplicación
	RefreshToken string `json:"refresh_token"` // Token de renovación actual
}

// RefreshAccessToken renueva un access token usando el refresh token
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (Token, error) {
	request := refreshTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}

	var token Token
	err := http.DoPostJSON(ctx, tokenEndpoint, "", request, &token)
	return token, err
}