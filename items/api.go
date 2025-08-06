package items

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const endpoint = "https://api.mercadolibre.com/items"

// GetByID obtiene un ítem por ID usando la API de Mercado Libre.
func GetByID(ctx context.Context, itemID, accessToken string) (Item, error) {
	url := fmt.Sprintf("%s/%s", endpoint, itemID)
	var item Item
	err := httpx.DoGetJSON(ctx, url, accessToken, &item)
	return item, err
}
