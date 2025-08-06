package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tidyrocks/mercado-libre-go-sdk/auth"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func main() {
	fmt.Println("🔐 Mercado Libre Auth Example")
	fmt.Println("============================")

	ctx := context.Background()

	// Ejemplo 1: Validar access token actual
	fmt.Println("\n1️⃣ Validating current access token...")
	if err := auth.ValidateAccessToken(ctx, testenv.AccessToken); err != nil {
		fmt.Printf("❌ Current token is invalid: %v\n", err)
	} else {
		fmt.Printf("✅ Current token is valid: %s...\n", testenv.AccessToken[:20])
	}

	// Ejemplo 2: Refresh manual del token
	fmt.Println("\n2️⃣ Manual token refresh...")
	response, err := auth.RefreshAccessToken(ctx, testenv.ClientID, testenv.ClientSecret, testenv.RefreshToken)
	if err != nil {
		log.Printf("❌ Failed to refresh token: %v", err)
		return
	}

	fmt.Printf("✅ Token refreshed successfully!\n")
	fmt.Printf("📱 New Access Token: %s...\n", response.AccessToken[:20])
	fmt.Printf("🔄 New Refresh Token: %s...\n", response.RefreshToken[:20])
	fmt.Printf("⏰ Expires in: %d seconds (%.1f hours)\n", response.ExpiresIn, float64(response.ExpiresIn)/3600)
	fmt.Printf("👤 User ID: %d\n", response.UserID)
	fmt.Printf("🔐 Scope: %s\n", response.Scope)

	// Ejemplo 3: Refresh completo con actualización del .env
	fmt.Println("\n3️⃣ Complete refresh with .env update...")

	// Esperar un segundo para evitar rate limiting
	time.Sleep(1 * time.Second)

	updatedResponse, err := auth.RefreshAndUpdateTokens(ctx, testenv.ClientID, testenv.ClientSecret, response.RefreshToken)
	if err != nil {
		log.Printf("❌ Failed to refresh and update: %v", err)
		return
	}

	fmt.Printf("✅ Tokens refreshed and .env updated!\n")
	fmt.Printf("📱 Final Access Token: %s...\n", updatedResponse.AccessToken[:20])
	fmt.Printf("🔄 Final Refresh Token: %s...\n", updatedResponse.RefreshToken[:20])
	fmt.Printf("💾 Check your .env file - it should contain the new tokens\n")
	fmt.Printf("🔍 Previous versions backed up as .env.backup.*\n")

	// Ejemplo 4: Mostrar configuración de auth
	fmt.Println("\n4️⃣ Auth configuration summary...")
	fmt.Printf("📋 Client ID: %s\n", testenv.ClientID)
	fmt.Printf("🔑 Client Secret: %s...\n", testenv.ClientSecret[:10])
	fmt.Printf("🆔 User ID: %d\n", updatedResponse.UserID)
	fmt.Printf("📅 Token expires in: %s\n", time.Duration(updatedResponse.ExpiresIn)*time.Second)

	fmt.Println("\n🎉 Auth example completed successfully!")
	fmt.Println("💡 Your tokens are now fresh and ready to use!")
}
