package categories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint       = "https://api.mercadolibre.com"
	categoriesEndpoint = baseEndpoint + "/categories"
	sitesEndpoint      = baseEndpoint + "/sites"
)

// GetByID obtiene una categoría por ID usando la API de Mercado Libre.
func GetByID(id string, accessToken string) (*Category, error) {
	url := fmt.Sprintf("%s/%s", categoriesEndpoint, id)
	var cat Category
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &cat)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

// GetBySite obtiene las categorías de un sitio específico.
func GetBySite(siteID string, params url.Values, accessToken string) ([]Category, error) {
	baseURL := fmt.Sprintf("%s/%s/categories", sitesEndpoint, siteID)
	var categories []Category
	err := httpx.DoGetJSONWithParams(context.Background(), baseURL, accessToken, params, &categories)
	return categories, err
}

// GetChildren obtiene las categorías hijas de una categoría.
func GetChildren(categoryID string, accessToken string) ([]Category, error) {
	url := fmt.Sprintf("%s/%s", categoriesEndpoint, categoryID)
	var response struct {
		ChildrenCategories []Category `json:"children_categories"`
	}
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	return response.ChildrenCategories, err
}

// PredictCategory utiliza algoritmos ML para sugerir categorías desde el título.
func PredictCategory(siteID, title string, params url.Values, accessToken string) ([]CategoryPrediction, error) {
	baseURL := fmt.Sprintf("%s/%s/category_predictor/predict", sitesEndpoint, siteID)

	// Agregar el título como parámetro
	if params == nil {
		params = url.Values{}
	}
	params.Add("title", title)

	var predictions []CategoryPrediction
	err := httpx.DoGetJSONWithParams(context.Background(), baseURL, accessToken, params, &predictions)
	return predictions, err
}

// Search busca categorías por query.
func Search(query string, params url.Values, accessToken string) ([]Category, error) {
	baseURL := fmt.Sprintf("%s/search", categoriesEndpoint)

	// Agregar el query como parámetro
	if params == nil {
		params = url.Values{}
	}
	params.Add("q", query)

	var response struct {
		Results []Category `json:"results"`
	}
	err := httpx.DoGetJSONWithParams(context.Background(), baseURL, accessToken, params, &response)
	return response.Results, err
}

// GetCategoriesByDomain obtiene categorías por dominio.
func GetCategoriesByDomain(domainID string, accessToken string) ([]Category, error) {
	url := fmt.Sprintf("%s/domains/%s/categories", baseEndpoint, domainID)
	var categories []Category
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &categories)
	return categories, err
}
