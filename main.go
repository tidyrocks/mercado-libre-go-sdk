package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tidyrocks/mercado-libre-go-sdk/api"
)

func main() {
	// Cargar variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: No se pudo cargar el archivo .env: %v", err)
	}

	// Obtener configuración de variables de entorno
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("ACCESS_TOKEN no está definido en .env")
	}
	
	itemID := os.Getenv("TEST_ITEM_ID")
	if itemID == "" {
		itemID = "MLM710148475" // fallback
	}

	ctx := context.Background()
	r, err := api.GetItem(ctx, itemID, accessToken)
	if err != nil {
		log.Fatalf("Error obteniendo item %s: %v", itemID, err)
	}

	// Generate JSON file with pretty format
	jsonData, _ := json.MarshalIndent(r, "", "  ")
	err = os.WriteFile("r.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Archivo r.json generado exitosamente")
}
