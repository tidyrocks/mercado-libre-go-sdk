package auth

// RefreshTokenRequest representa la petición para renovar el access token
type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`    // "refresh_token"
	ClientID     string `json:"client_id"`     // ID de la aplicación
	ClientSecret string `json:"client_secret"` // Secret de la aplicación
	RefreshToken string `json:"refresh_token"` // Token de renovación actual
}

// RefreshTokenResponse representa la respuesta del refresh token
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`  // Nuevo access token
	TokenType    string `json:"token_type"`    // "Bearer"
	ExpiresIn    int    `json:"expires_in"`    // Segundos hasta expiración (21600 = 6 horas)
	Scope        string `json:"scope"`         // "offline_access read write"
	UserID       int64  `json:"user_id"`       // ID del usuario
	RefreshToken string `json:"refresh_token"` // Nuevo refresh token
}

// AuthConfig contiene toda la configuración de autenticación
type AuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
}
