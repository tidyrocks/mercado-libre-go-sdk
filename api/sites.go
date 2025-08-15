package api

import (
	"context"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const sitesEndpoint = "https://api.mercadolibre.com/sites"

// Site representa un sitio de Mercado Libre (país)
type Site struct {
	ID                 string `json:"id"`                   // ID del sitio (ej. MLA, MLM, MLB)
	Name               string `json:"name"`                 // Nombre del país (ej. Argentina, Mexico, Brasil)
	CountryID          string `json:"country_id"`           // ID del país
	DefaultCurrencyID  string `json:"default_currency_id"`  // Moneda por defecto (ej. ARS, MXN, BRL)
}

// GetSites obtiene la lista de todos los sitios disponibles en Mercado Libre
func GetSites(ctx context.Context, accessToken string) ([]Site, error) {
	var sites []Site
	err := http.DoGetJSON(ctx, sitesEndpoint, accessToken, &sites)
	return sites, err
}