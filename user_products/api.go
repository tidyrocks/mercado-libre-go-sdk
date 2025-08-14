package user_products

import (
	"context"
	"fmt"
	"net/url"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetByID obtiene un User Product por su ID.
func GetByID(userProductID string, accessToken string) (*UserProduct, error) {
	url := fmt.Sprintf("%s/user-products/%s", baseEndpoint, userProductID)
	var userProduct UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProduct)
	if err != nil {
		return nil, err
	}
	return &userProduct, nil
}

// GetFamilyByID obtiene información de una familia por su ID.
func GetFamilyByID(siteID, familyID string, accessToken string) (*Family, error) {
	url := fmt.Sprintf("%s/sites/%s/user-products-families/%s", baseEndpoint, siteID, familyID)
	var family Family
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &family)
	if err != nil {
		return nil, err
	}
	return &family, nil
}

// GetUserProductsByFamily obtiene todos los User Products de una familia.
func GetUserProductsByFamily(siteID, familyID string, accessToken string) ([]UserProduct, error) {
	url := fmt.Sprintf("%s/sites/%s/user-products-families/%s/user-products", baseEndpoint, siteID, familyID)
	var userProducts []UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProducts)
	return userProducts, err
}

// GetItemsByUserProduct obtiene todos los ítems asociados a un User Product.
func GetItemsByUserProduct(sellerID, userProductID string, params url.Values, accessToken string) (*ItemSearchResult, error) {
	baseURL := fmt.Sprintf("%s/users/%s/items/search", baseEndpoint, sellerID)

	// Agregar user_product_id como parámetro
	if params == nil {
		params = url.Values{}
	}
	params.Add("user_product_id", userProductID)

	var result ItemSearchResult
	err := httpx.DoGetJSONWithParams(context.Background(), baseURL, accessToken, params, &result)
	return &result, err
}

// GetItemsByMultipleUserProducts acepta lista vacía y retorna resultado vacío sin error.
func GetItemsByMultipleUserProducts(sellerID string, userProductIDs []string, params url.Values, accessToken string) (*ItemSearchResult, error) {
	if len(userProductIDs) == 0 {
		return &ItemSearchResult{}, nil
	}

	baseURL := fmt.Sprintf("%s/users/%s/items/search", baseEndpoint, sellerID)

	// Crear lista de user_product_ids separados por coma
	var userProductsList string
	for i, id := range userProductIDs {
		if i > 0 {
			userProductsList += ","
		}
		userProductsList += id
	}

	// Agregar user_product_ids como parámetro
	if params == nil {
		params = url.Values{}
	}
	params.Add("user_product_id", userProductsList)

	var result ItemSearchResult
	err := httpx.DoGetJSONWithParams(context.Background(), baseURL, accessToken, params, &result)
	return &result, err
}

// CheckEligibility verifica si un ítem es elegible para migrar al modelo User Products.
func CheckEligibility(itemID string, accessToken string) (bool, error) {
	url := fmt.Sprintf("%s/items/%s/uptin-eligibility", baseEndpoint, itemID)
	var response struct {
		Eligible bool `json:"eligible"`
	}
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	return response.Eligible, err
}

// GetByIDWithStock incluye información de stock por ubicación geográfica.
func GetByIDWithStock(userProductID string, accessToken string) (*UserProduct, error) {
	url := fmt.Sprintf("%s/user-products/%s?include_stock_locations=true", baseEndpoint, userProductID)
	var userProduct UserProduct
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &userProduct)
	if err != nil {
		return nil, err
	}
	return &userProduct, nil
}
