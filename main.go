package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidyrocks/mercado-libre-go-sdk/api"
)

func main() {
	ctx := context.Background()
	r, _ := api.GetItem(ctx, "MLM710148475", "APP_USR-")

	// Generate JSON file with pretty format
	jsonData, _ := json.MarshalIndent(r, "", "  ")
	err := os.WriteFile("r.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Archivo r.json generado exitosamente")
}
