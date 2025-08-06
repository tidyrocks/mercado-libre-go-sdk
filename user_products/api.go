package user_products

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
	"gitlab.com/tidyrocks/tidy-go-common/shared"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetByID obtiene un User Product por su ID.
func GetByID(userProductID, accessToken string) (*UserProduct, error) {
	url := fmt.Sprintf("%s/user-products/%s", baseEndpoint, userProductID)
	var userProduct UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProduct)
	if err != nil {
		return nil, err
	}
	return &userProduct, nil
}

// GetFamilyByID obtiene información de una familia por su ID.
func GetFamilyByID(siteID, familyID, accessToken string) (*Family, error) {
	url := fmt.Sprintf("%s/sites/%s/user-products-families/%s", baseEndpoint, siteID, familyID)
	var family Family
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &family)
	if err != nil {
		return nil, err
	}
	return &family, nil
}

// GetUserProductsByFamily obtiene todos los User Products de una familia.
func GetUserProductsByFamily(siteID, familyID, accessToken string) ([]UserProduct, error) {
	url := fmt.Sprintf("%s/sites/%s/user-products-families/%s/user-products", baseEndpoint, siteID, familyID)
	var userProducts []UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProducts)
	return userProducts, err
}

// GetItemsByUserProduct obtiene todos los ítems asociados a un User Product.
func GetItemsByUserProduct(sellerID, userProductID, accessToken string, params []shared.KeyValue) (*ItemSearchResult, error) {
	url := fmt.Sprintf("%s/users/%s/items/search", baseEndpoint, sellerID)

	// Agregar user_product_id como parámetro
	userProductParam := shared.KeyValue{Key: "user_product_id", Value: userProductID}
	allParams := append([]shared.KeyValue{userProductParam}, params...)

	var result ItemSearchResult
	err := httpx.DoGetJSONWithParams(context.Background(), url, accessToken, allParams, &result)
	return &result, err
}

// GetItemsByMultipleUserProducts acepta lista vacía y retorna resultado vacío sin error.
func GetItemsByMultipleUserProducts(sellerID string, userProductIDs []string, accessToken string, params []shared.KeyValue) (*ItemSearchResult, error) {
	if len(userProductIDs) == 0 {
		return &ItemSearchResult{}, nil
	}

	url := fmt.Sprintf("%s/users/%s/items/search", baseEndpoint, sellerID)

	// Crear lista de user_product_ids separados por coma
	var userProductsList string
	for i, id := range userProductIDs {
		if i > 0 {
			userProductsList += ","
		}
		userProductsList += id
	}

	// Agregar user_product_ids como parámetro
	userProductParam := shared.KeyValue{Key: "user_product_id", Value: userProductsList}
	allParams := append([]shared.KeyValue{userProductParam}, params...)

	var result ItemSearchResult
	err := httpx.DoGetJSONWithParams(context.Background(), url, accessToken, allParams, &result)
	return &result, err
}

// CheckEligibility verifica si un ítem es elegible para migrar al modelo User Products.
func CheckEligibility(itemID, accessToken string) (bool, error) {
	url := fmt.Sprintf("%s/items/%s/uptin-eligibility", baseEndpoint, itemID)
	var response struct {
		Eligible bool `json:"eligible"`
	}
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	return response.Eligible, err
}

// GetByIDWithStock incluye información de stock por ubicación geográfica.
func GetByIDWithStock(userProductID, accessToken string) (*UserProduct, error) {
	url := fmt.Sprintf("%s/user-products/%s?include_stock_locations=true", baseEndpoint, userProductID)
	var userProduct UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProduct)
	if err != nil {
		return nil, err
	}
	return &userProduct, nil
}
