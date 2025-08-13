package categories

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
	"gitlab.com/tidyrocks/tidy-go-common/shared"
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
func GetBySite(siteID string, params []shared.KeyValue, accessToken string) ([]Category, error) {
	url := fmt.Sprintf("%s/%s/categories", sitesEndpoint, siteID)
	var categories []Category
	err := httpx.DoGetJSONWithParams(context.Background(), url, accessToken, params, &categories)
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
func PredictCategory(siteID, title string, params []shared.KeyValue, accessToken string) ([]CategoryPrediction, error) {
	url := fmt.Sprintf("%s/%s/category_predictor/predict", sitesEndpoint, siteID)

	// Agregar el título como parámetro
	titleParam := shared.KeyValue{Key: "title", Value: title}
	allParams := append([]shared.KeyValue{titleParam}, params...)

	var predictions []CategoryPrediction
	err := httpx.DoGetJSONWithParams(context.Background(), url, accessToken, allParams, &predictions)
	return predictions, err
}

// Search busca categorías por query.
func Search(query string, params []shared.KeyValue, accessToken string) ([]Category, error) {
	url := fmt.Sprintf("%s/search", categoriesEndpoint)

	// Agregar el query como parámetro
	queryParam := shared.KeyValue{Key: "q", Value: query}
	allParams := append([]shared.KeyValue{queryParam}, params...)

	var response struct {
		Results []Category `json:"results"`
	}
	err := httpx.DoGetJSONWithParams(context.Background(), url, accessToken, allParams, &response)
	return response.Results, err
}

// GetCategoriesByDomain obtiene categorías por dominio.
func GetCategoriesByDomain(domainID string, accessToken string) ([]Category, error) {
	url := fmt.Sprintf("%s/domains/%s/categories", baseEndpoint, domainID)
	var categories []Category
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &categories)
	return categories, err
}
