package variations

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetByItemID obtiene todas las variaciones de un ítem.
func GetByItemID(itemID, accessToken string) ([]Variation, error) {
	url := fmt.Sprintf("%s/items/%s/variations", baseEndpoint, itemID)
	var variations []Variation
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &variations)
	return variations, err
}

// GetByID obtiene una variación específica de un ítem.
func GetByID(itemID, variationID, accessToken string) (*Variation, error) {
	url := fmt.Sprintf("%s/items/%s/variations/%s", baseEndpoint, itemID, variationID)
	var variation Variation
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &variation)
	if err != nil {
		return nil, err
	}
	return &variation, nil
}

// GetByItemIDWithAttributes incluye atributos completos en la respuesta.
func GetByItemIDWithAttributes(itemID, accessToken string) ([]Variation, error) {
	url := fmt.Sprintf("%s/items/%s/variations?include_attributes=all", baseEndpoint, itemID)
	var variations []Variation
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &variations)
	return variations, err
}

// GetByIDWithAttributes incluye atributos completos en la respuesta.
func GetByIDWithAttributes(itemID, variationID, accessToken string) (*Variation, error) {
	url := fmt.Sprintf("%s/items/%s/variations/%s?include_attributes=all", baseEndpoint, itemID, variationID)
	var variation Variation
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &variation)
	if err != nil {
		return nil, err
	}
	return &variation, nil
}
