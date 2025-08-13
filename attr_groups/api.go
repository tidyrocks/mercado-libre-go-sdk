package attr_groups

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetTechnicalSpecsInput obtiene especificaciones técnicas de entrada para una categoría.
func GetTechnicalSpecsInput(categoryID string, accessToken string) (*struct {
	Groups []TechnicalSpec `json:"groups"`
}, error) {
	url := fmt.Sprintf("%s/categories/%s/technical_specs/input", baseEndpoint, categoryID)

	var response struct {
		Groups []TechnicalSpec `json:"groups"`
	}
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTechnicalSpecsOutput obtiene especificaciones técnicas de salida para una categoría.
func GetTechnicalSpecsOutput(categoryID string, accessToken string) (*struct {
	MainTitle string          `json:"main_title"`
	Groups    []TechnicalSpec `json:"groups"`
}, error) {
	url := fmt.Sprintf("%s/categories/%s/technical_specs/output", baseEndpoint, categoryID)

	var response struct {
		MainTitle string          `json:"main_title"`
		Groups    []TechnicalSpec `json:"groups"`
	}
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
