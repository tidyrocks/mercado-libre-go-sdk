package auth

import (
	"context"
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestRefreshAccessToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping auth test in short mode")
	}

	ctx := context.Background()

	t.Logf("🔑 Testing OAuth refresh token flow")
	t.Logf("📋 Using Client ID: %s", testenv.ClientID)
	t.Logf("🔄 Current refresh token: %s...", testenv.RefreshToken[:20])

	// Validar que tenemos las credenciales necesarias
	if testenv.ClientID == "" {
		t.Fatal("❌ CLIENT_ID not found in .env")
	}
	if testenv.ClientSecret == "" {
		t.Fatal("❌ CLIENT_SECRET not found in .env")
	}
	if testenv.RefreshToken == "" {
		t.Fatal("❌ REFRESH_TOKEN not found in .env")
	}

	// Ejecutar refresh del token
	response, err := RefreshAccessToken(ctx, testenv.ClientID, testenv.ClientSecret, testenv.RefreshToken)
	if err != nil {
		t.Fatalf("❌ Failed to refresh access token: %v", err)
	}

	// Validar response
	if response == nil {
		t.Fatal("❌ Response is nil")
	}

	if response.AccessToken == "" {
		t.Error("❌ New access_token is empty")
	}

	if response.RefreshToken == "" {
		t.Error("❌ New refresh_token is empty")
	}

	if response.TokenType != "Bearer" {
		t.Errorf("❌ Expected token_type 'Bearer', got: %s", response.TokenType)
	}

	if response.ExpiresIn <= 0 {
		t.Errorf("❌ Invalid expires_in: %d", response.ExpiresIn)
	}

	if response.UserID <= 0 {
		t.Errorf("❌ Invalid user_id: %d", response.UserID)
	}

	// Log successful response details
	t.Logf("✅ Token refresh successful!")
	t.Logf("📱 New access token: %s...", response.AccessToken[:20])
	t.Logf("🔄 New refresh token: %s...", response.RefreshToken[:20])
	t.Logf("⏰ Expires in: %d seconds (%.1f hours)", response.ExpiresIn, float64(response.ExpiresIn)/3600)
	t.Logf("👤 User ID: %d", response.UserID)
	t.Logf("🔐 Scope: %s", response.Scope)

	// Verificar que los tokens son diferentes (renovados)
	if response.AccessToken == testenv.AccessToken {
		t.Logf("⚠️  Warning: New access token is same as old one")
	} else {
		t.Logf("✅ Access token was renewed successfully")
	}

	if response.RefreshToken == testenv.RefreshToken {
		t.Logf("⚠️  Warning: New refresh token is same as old one")
	} else {
		t.Logf("✅ Refresh token was renewed successfully")
	}
}

func TestRefreshAndUpdateTokens(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping .env update test in short mode")
	}

	ctx := context.Background()

	t.Logf("🔄 Testing complete refresh and .env update flow")

	// Validar que tenemos las credenciales necesarias
	if testenv.ClientID == "" || testenv.ClientSecret == "" || testenv.RefreshToken == "" {
		t.Skip("❌ Missing auth credentials in .env, skipping .env update test")
	}

	// Capturar tokens originales
	originalAccessToken := testenv.AccessToken
	originalRefreshToken := testenv.RefreshToken

	t.Logf("📝 Original access token: %s...", originalAccessToken[:20])
	t.Logf("📝 Original refresh token: %s...", originalRefreshToken[:20])

	// Ejecutar refresh completo con actualización del .env
	response, err := RefreshAndUpdateTokens(ctx, testenv.ClientID, testenv.ClientSecret, testenv.RefreshToken)
	if err != nil {
		t.Fatalf("❌ Failed to refresh and update tokens: %v", err)
	}

	// Validar response
	if response.AccessToken == "" || response.RefreshToken == "" {
		t.Fatal("❌ Response missing required tokens")
	}

	t.Logf("✅ Tokens refreshed and .env updated successfully!")
	t.Logf("📱 Updated access token: %s...", response.AccessToken[:20])
	t.Logf("🔄 Updated refresh token: %s...", response.RefreshToken[:20])

	// Nota: En un test real podrías verificar que el archivo .env se actualizó
	// leyendo el archivo y comparando los valores

	t.Logf("💾 .env file should now contain the new tokens")
	t.Logf("🔍 Check .env.backup.* files for previous versions")
}

func TestValidateAccessToken(t *testing.T) {
	ctx := context.Background()

	t.Run("valid_token_format", func(t *testing.T) {
		err := ValidateAccessToken(ctx, testenv.AccessToken)
		if err != nil {
			t.Errorf("❌ Valid token failed validation: %v", err)
		} else {
			t.Logf("✅ Access token format is valid")
		}
	})

	t.Run("invalid_token_format", func(t *testing.T) {
		invalidToken := "invalid-token-format"
		err := ValidateAccessToken(ctx, invalidToken)
		if err == nil {
			t.Error("❌ Invalid token passed validation")
		} else {
			t.Logf("✅ Invalid token correctly rejected: %v", err)
		}
	})

	t.Run("empty_token", func(t *testing.T) {
		err := ValidateAccessToken(ctx, "")
		if err == nil {
			t.Error("❌ Empty token passed validation")
		} else {
			t.Logf("✅ Empty token correctly rejected: %v", err)
		}
	})
}
