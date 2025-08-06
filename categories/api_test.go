package categories

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetByID(t *testing.T) {
	categoryID := "MLM1051" // Categoría de prueba

	category, err := GetByID(categoryID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting category: %v", err)
	}

	if category == nil {
		t.Fatal("Category is nil")
	}

	if category.ID != categoryID {
		t.Errorf("Expected category ID %s, got %s", categoryID, category.ID)
	}

	if category.Name == "" {
		t.Error("Category name is empty")
	}

	t.Logf("Category: %s - %s", category.ID, category.Name)
}

func TestGetBySite(t *testing.T) {
	siteID := "MLM"

	categories, err := GetBySite(siteID, testenv.AccessToken, nil)
	if err != nil {
		t.Fatalf("Error getting categories by site: %v", err)
	}

	if len(categories) == 0 {
		t.Error("No categories returned")
	}

	t.Logf("Found %d categories for site %s", len(categories), siteID)

	// Verificar primera categoría
	if len(categories) > 0 {
		cat := categories[0]
		if cat.ID == "" || cat.Name == "" {
			t.Error("First category has empty ID or Name")
		}
		t.Logf("First category: %s - %s", cat.ID, cat.Name)
	}
}

func TestGetChildren(t *testing.T) {
	categoryID := "MLM1051" // Categoría padre

	children, err := GetChildren(categoryID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting category children: %v", err)
	}

	t.Logf("Found %d child categories for %s", len(children), categoryID)

	// Verificar estructura de las categorías hijas si existen
	for i, child := range children {
		if i >= 3 { // Solo verificar las primeras 3
			break
		}
		if child.ID == "" || child.Name == "" {
			t.Errorf("Child category %d has empty ID or Name", i)
		}
		t.Logf("Child %d: %s - %s", i, child.ID, child.Name)
	}
}
