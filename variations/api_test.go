package variations

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetByItemID(t *testing.T) {
	variations, err := GetByItemID(testenv.TestItemID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting variations: %v", err)
	}

	t.Logf("Found %d variations for item %s", len(variations), testenv.TestItemID)

	// Si hay variaciones, verificar la estructura
	for i, variation := range variations {
		if i >= 3 { // Solo verificar las primeras 3
			break
		}

		if variation.ID == 0 {
			t.Errorf("Variation %d has invalid ID", i)
		}

		if variation.Price <= 0 {
			t.Errorf("Variation %d has invalid price: %.2f", i, variation.Price)
		}

		t.Logf("Variation %d: ID=%d, Price=%.2f, Available=%d",
			i, variation.ID, variation.Price, variation.AvailableQuantity)

		// Verificar combinaciones de atributos
		if len(variation.AttributeCombinations) > 0 {
			t.Logf("  Attributes: %d combinations", len(variation.AttributeCombinations))
			for j, attr := range variation.AttributeCombinations {
				if j >= 2 { // Solo mostrar las primeras 2
					break
				}
				t.Logf("    %s: %s", attr.Name, attr.ValueName)
			}
		}
	}
}

func TestGetByItemIDWithAttributes(t *testing.T) {
	variations, err := GetByItemIDWithAttributes(testenv.TestItemID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting variations with attributes: %v", err)
	}

	t.Logf("Found %d variations with attributes for item %s", len(variations), testenv.TestItemID)

	// Verificar que incluye atributos detallados
	for i, variation := range variations {
		if i >= 1 { // Solo verificar la primera
			break
		}

		t.Logf("Variation %d attributes: %d", i, len(variation.Attributes))
		for j, attr := range variation.Attributes {
			if j >= 3 { // Solo mostrar los primeros 3
				break
			}
			t.Logf("  %s (%s): %s", attr.Name, attr.ID, attr.ValueName)
		}
	}
}
