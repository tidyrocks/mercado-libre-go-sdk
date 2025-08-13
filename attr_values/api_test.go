package attr_values

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetTopValues(t *testing.T) {
	domainID := "MLM-CELLPHONES" // Dominio de celulares en México
	attributeID := "BRAND"       // Atributo marca

	values, err := GetTopValues(domainID, attributeID, nil, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting top values: %v", err)
	}

	if len(values) == 0 {
		t.Error("No values returned")
	}

	t.Logf("Found %d top values for %s in domain %s", len(values), attributeID, domainID)

	// Verificar estructura de los primeros valores
	for i, value := range values {
		if i >= 10 { // Solo verificar los primeros 10
			break
		}

		if value.ID == "" || value.Name == "" {
			t.Errorf("Value %d has empty ID or Name", i)
		}

		metricInfo := "N/A"
		if value.Metric != nil {
			metricInfo = string(rune(*value.Metric))
		}

		t.Logf("Value %d: %s (%s) - Metric: %s", i, value.Name, value.ID, metricInfo)
	}
}

func TestGetTopValuesWithFilter(t *testing.T) {
	domainID := "MLM-CELLPHONES"
	attributeID := "MODEL"

	// Filtrar por marca Samsung
	knownAttributes := []KnownAttribute{
		{
			ID:      "BRAND",
			ValueID: "206", // Samsung
		},
	}

	values, err := GetTopValuesWithFilter(domainID, attributeID, knownAttributes, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting filtered top values: %v", err)
	}

	t.Logf("Found %d top models for Samsung in domain %s", len(values), domainID)

	// Verificar algunos modelos de Samsung
	for i, value := range values {
		if i >= 5 { // Solo mostrar los primeros 5
			break
		}

		metricInfo := "N/A"
		if value.Metric != nil {
			metricInfo = string(rune(*value.Metric))
		}

		t.Logf("Samsung Model %d: %s (%s) - Metric: %s", i, value.Name, value.ID, metricInfo)
	}
}
