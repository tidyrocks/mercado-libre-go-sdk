package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func main() {
	fmt.Println("🚀 Mercado Libre Go SDK - Test Runner")
	fmt.Println("====================================")
	fmt.Println("⚠️  SOLO EJECUTANDO TESTS GET (sin modificaciones)")
	fmt.Println()

	// Verificar que las variables de entorno estén cargadas
	if testenv.AccessToken == "" {
		fmt.Println("❌ Error: ACCESS_TOKEN no encontrado en .env")
		os.Exit(1)
	}
	if testenv.TestItemID == "" {
		fmt.Println("❌ Error: TEST_ITEM_ID no encontrado en .env")
		os.Exit(1)
	}
	fmt.Printf("✅ Variables de entorno cargadas (Token: %s..., Item: %s)\n\n", testenv.AccessToken[:10], testenv.TestItemID)

	// Modules to test
	modules := []string{
		"categories",
		"items",
		"variations",
		"user_products",
		"attrs",
		"attr_values",
		"attr_groups",
		"auth",
	}

	totalPassed := 0
	totalFailed := 0

	for _, module := range modules {
		fmt.Printf("📦 Testing module: %s\n", module)
		fmt.Println(strings.Repeat("-", 40))

		cmd := exec.Command("go", "test", "-v", fmt.Sprintf("./%s", module))
		cmd.Dir = "."

		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("❌ Module %s failed: %v\n", module, err)
			totalFailed++
		} else {
			fmt.Printf("✅ Module %s passed\n", module)
			totalPassed++
		}

		// Show output
		fmt.Println(string(output))
		fmt.Println()
	}

	fmt.Println("📊 Test Summary")
	fmt.Println("===============")
	fmt.Printf("✅ Passed: %d\n", totalPassed)
	fmt.Printf("❌ Failed: %d\n", totalFailed)
	fmt.Printf("📈 Total: %d\n", totalPassed+totalFailed)

	if totalFailed > 0 {
		os.Exit(1)
	}
}
