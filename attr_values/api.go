package attr_values

import (
	"context"
	"fmt"
	"net/url"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetTopValues usa POST en lugar de GET según especificación de ML.
func GetTopValues(domainID, attributeID string, params url.Values, accessToken string) ([]AttributeValue, error) {
	baseURL := fmt.Sprintf("%s/catalog_domains/%s/attributes/%s/top_values", baseEndpoint, domainID, attributeID)

	request := map[string]interface{}{}

	var values []AttributeValue
	err := httpx.DoPostJSON(context.Background(), baseURL, accessToken, request, &values)
	return values, err
}

// GetTopValuesWithFilter filtra resultados usando atributos ya conocidos del producto.
func GetTopValuesWithFilter(domainID, attributeID string, knownAttributes []KnownAttribute, accessToken string) ([]AttributeValue, error) {
	url := fmt.Sprintf("%s/catalog_domains/%s/attributes/%s/top_values", baseEndpoint, domainID, attributeID)

	request := map[string]interface{}{
		"known_attributes": knownAttributes,
	}

	var values []AttributeValue
	err := httpx.DoPostJSON(context.Background(), url, accessToken, request, &values)
	return values, err
}
