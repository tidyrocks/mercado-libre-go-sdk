package attrs

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetByCategory(t *testing.T) {
	categoryID := "MLM1051" // Categoría de prueba de México

	attributes, err := GetByCategory(categoryID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting attributes: %v", err)
	}

	if len(attributes) == 0 {
		t.Error("No attributes returned")
	}

	t.Logf("Found %d attributes for category %s", len(attributes), categoryID)

	// Verificar estructura de los primeros atributos
	for i, attr := range attributes {
		if i >= 5 { // Solo verificar los primeros 5
			break
		}

		if attr.ID == "" || attr.Name == "" {
			t.Errorf("Attribute %d has empty ID or Name", i)
		}

		if attr.ValueType == "" {
			t.Errorf("Attribute %d has empty ValueType", i)
		}

		t.Logf("Attribute %d: %s (%s) - Type: %s", i, attr.Name, attr.ID, attr.ValueType)

		// Mostrar algunos tags si existen
		if len(attr.Tags) > 0 {
			tagCount := 0
			tagInfo := "  Tags:"
			for tag, value := range attr.Tags {
				if tagCount >= 3 { // Solo mostrar primeros 3 tags
					break
				}
				if value {
					tagInfo += " " + tag
					tagCount++
				}
			}
			if tagCount > 0 {
				t.Log(tagInfo)
			}
		}
	}
}
