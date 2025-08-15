package api

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const domainsEndpoint = "https://api.mercadolibre.com/catalog_domains"

// Domain representa un dominio de productos en Mercado Libre
type Domain struct {
	ID          string   `json:"id"`          // ID del dominio (ej. MLA-CELLPHONES)
	Name        string   `json:"name"`        // Nombre del dominio
	CategoryIDs []string `json:"category_ids"` // IDs de categorías asociadas
	Picture     string   `json:"picture"`     // URL de imagen del dominio
	Permalink   string   `json:"permalink"`   // URL permanente del dominio
	// Campos opcionales
	Description *string `json:"description,omitempty"` // Descripción del dominio
}

// DomainShippingAttributes representa atributos de envío de un dominio
type DomainShippingAttributes struct {
	DomainID             string                 `json:"domain_id"`
	ShippingAttributes   []ShippingAttribute    `json:"shipping_attributes"`
	RequiredAttributes   []string              `json:"required_attributes"`
}

// ShippingAttribute representa un atributo de envío específico
type ShippingAttribute struct {
	ID           string   `json:"id"`           // ID del atributo
	Name         string   `json:"name"`         // Nombre del atributo
	Type         string   `json:"type"`         // Tipo del atributo
	Required     bool     `json:"required"`     // Si es obligatorio
	ValueType    string   `json:"value_type"`   // Tipo de valor
	AllowedUnits []string `json:"allowed_units"` // Unidades permitidas
	// Campos opcionales
	DefaultValue *string `json:"default_value,omitempty"` // Valor por defecto
	Description  *string `json:"description,omitempty"`   // Descripción del atributo
}

// GetDomainByID obtiene información detallada de un dominio por su ID
func GetDomainByID(ctx context.Context, domainID, accessToken string) (Domain, error) {
	url := fmt.Sprintf("%s/%s", domainsEndpoint, domainID)
	var domain Domain
	err := http.DoGetJSON(ctx, url, accessToken, &domain)
	return domain, err
}

// GetDomainShippingAttributes obtiene los atributos de envío requeridos para un dominio
func GetDomainShippingAttributes(ctx context.Context, domainID, accessToken string) (DomainShippingAttributes, error) {
	url := fmt.Sprintf("%s/%s/shipping_attributes", domainsEndpoint, domainID)
	var shippingAttrs DomainShippingAttributes
	err := http.DoGetJSON(ctx, url, accessToken, &shippingAttrs)
	return shippingAttrs, err
}