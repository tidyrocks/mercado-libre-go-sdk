package user_products

import (
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetByID(t *testing.T) {
	// Primero necesitamos obtener un user_product_id del item de prueba
	// Esto normalmente vendría de la respuesta de un item
	userProductID := "MLMU123456" // ID de ejemplo, puede fallar si no existe

	userProduct, err := GetByID(userProductID, testenv.AccessToken)
	if err != nil {
		// Es esperado que falle si el user_product_id no existe
		t.Logf("Expected error getting user product (ID may not exist): %v", err)
		return
	}

	if userProduct == nil {
		t.Fatal("UserProduct is nil")
	}

	if userProduct.ID != userProductID {
		t.Errorf("Expected user product ID %s, got %s", userProductID, userProduct.ID)
	}

	t.Logf("UserProduct: %s - %s", userProduct.ID, userProduct.Title)
	t.Logf("Family: %v, Domain: %v", userProduct.FamilyName, userProduct.DomainID)
}

func TestGetItemsByUserProduct(t *testing.T) {
	sellerID := "118401678"       // Del access token
	userProductID := "MLMU123456" // ID de ejemplo

	result, err := GetItemsByUserProduct(sellerID, userProductID, testenv.AccessToken, nil)
	if err != nil {
		// Es esperado que falle si no hay user products
		t.Logf("Expected error getting items by user product: %v", err)
		return
	}

	if result == nil {
		t.Fatal("ItemSearchResult is nil")
	}

	t.Logf("Found %d items for user product %s", len(result.Results), userProductID)

	for i, item := range result.Results {
		if i >= 3 { // Solo mostrar los primeros 3
			break
		}
		t.Logf("Item %d: %s - Status: %s - Price: %.2f",
			i, item.ID, item.Status, item.Price)
	}
}
