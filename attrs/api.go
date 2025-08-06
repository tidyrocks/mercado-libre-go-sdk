package attrs

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// GetByCategory obtiene todos los atributos de una categoría.
func GetByCategory(categoryID, accessToken string) ([]Attribute, error) {
	url := fmt.Sprintf("%s/categories/%s/attributes", baseEndpoint, categoryID)
	var attributes []Attribute
	err := httpx.DoGetJSON(context.Background(), url, accessToken, &attributes)
	return attributes, err
}
