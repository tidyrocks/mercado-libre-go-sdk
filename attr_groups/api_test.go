package attr_groups

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetTechnicalSpecsInput(t *testing.T) {
	categoryID := "MLM1051" // Categoría de prueba

	response, err := GetTechnicalSpecsInput(categoryID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting technical specs input: %v", err)
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	if len(response.Groups) == 0 {
		t.Error("No technical spec groups returned")
	}

	t.Logf("Found %d technical spec groups for category %s", len(response.Groups), categoryID)

	// Verificar estructura de los grupos
	for i, group := range response.Groups {
		if i >= 3 { // Solo verificar los primeros 3
			break
		}

		if group.ID == "" || group.Label == "" {
			t.Errorf("Group %d has empty ID or Label", i)
		}

		t.Logf("Group %d: %s (%s) - Relevance: %d", i, group.Label, group.ID, group.Relevance)
	}
}

func TestGetTechnicalSpecsOutput(t *testing.T) {
	categoryID := "MLM1051" // Categoría de prueba

	response, err := GetTechnicalSpecsOutput(categoryID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting technical specs output: %v", err)
	}

	if response == nil {
		t.Fatal("Response is nil")
	}

	t.Logf("Main title: %s", response.MainTitle)
	t.Logf("Found %d output groups for category %s", len(response.Groups), categoryID)

	// Verificar estructura de los grupos de salida
	for i, group := range response.Groups {
		if i >= 3 { // Solo verificar los primeros 3
			break
		}

		if group.ID == "" || group.Label == "" {
			t.Errorf("Output group %d has empty ID or Label", i)
		}

		t.Logf("Output Group %d: %s (%s) - Relevance: %d", i, group.Label, group.ID, group.Relevance)
	}
}
