package api

import (
	"context"
	"fmt"
	"time"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const userProductsEndpoint = "https://api.mercadolibre.com/user-products"
const siteUserProductFamiliesEndpoint = "https://api.mercadolibre.com/sites"
const itemsEligibilityEndpoint = "https://api.mercadolibre.com/items"

// UserProduct representa un producto físico que un vendedor posee en el nuevo modelo de User Products
type UserProduct struct {
	ID                string       `json:"id"`                  // ID del User Product (ej. MLBU22012)
	Name              string       `json:"name"`                // Nombre del producto
	UserID            int64        `json:"user_id"`             // ID del vendedor
	DomainID          string       `json:"domain_id"`           // ID del dominio (ej. MLB-CELLPHONES)
	FamilyID          int64        `json:"family_id"`           // ID de la familia
	DateCreated       time.Time    `json:"date_created"`        // Fecha de creación
	LastUpdated       time.Time    `json:"last_updated"`        // Fecha de última actualización
	// Arrays y slices
	Attrs             []Attr       `json:"attributes"`          // Atributos del producto
	Pictures          []Picture    `json:"pictures"`            // Imágenes del producto
	Tags              []string     `json:"tags"`                // Tags del producto
	// Campos opcionales
	CatalogProductID  *string      `json:"catalog_product_id,omitempty"`  // ID del producto de catálogo
	Thumbnail         *Picture     `json:"thumbnail,omitempty"`           // Imagen thumbnail
}

// UserProductFamily representa una familia de User Products
type UserProductFamily struct {
	ID              int64          `json:"id"`                // ID de la familia
	Name            string         `json:"name"`              // Nombre de la familia
	SiteID          string         `json:"site_id"`           // ID del sitio
	UserProducts    []UserProduct  `json:"user_products"`     // User Products de la familia
	DateCreated     time.Time      `json:"date_created"`      // Fecha de creación
	LastUpdated     time.Time      `json:"last_updated"`      // Fecha de última actualización
}

// UserProductStock representa información de stock de un User Product
type UserProductStock struct {
	UserProductID     string                   `json:"user_product_id"`      // ID del User Product
	TotalStock        int                      `json:"total_stock"`          // Stock total
	ReservedStock     int                      `json:"reserved_stock"`       // Stock reservado
	AvailableStock    int                      `json:"available_stock"`      // Stock disponible
	// Arrays y slices
	Locations         []UserProductLocation    `json:"locations"`            // Ubicaciones del stock
}

// UserProductLocation representa una ubicación de stock específica
type UserProductLocation struct {
	LocationID        string    `json:"location_id"`          // ID de la ubicación
	LocationType      string    `json:"location_type"`        // Tipo de ubicación (selling_address, meli_facility, etc.)
	Quantity          int       `json:"quantity"`             // Cantidad en esta ubicación
	ReservedQuantity  int       `json:"reserved_quantity"`    // Cantidad reservada
	AvailableQuantity int       `json:"available_quantity"`   // Cantidad disponible
	// Campos opcionales
	Address           *string   `json:"address,omitempty"`    // Dirección de la ubicación
	StoreID           *string   `json:"store_id,omitempty"`   // ID de la tienda
}

// UserProductStockUpdate representa una actualización de stock
type UserProductStockUpdate struct {
	Locations []UserProductLocationUpdate `json:"locations"` // Ubicaciones a actualizar
}

// UserProductLocationUpdate representa la actualización de stock en una ubicación específica
type UserProductLocationUpdate struct {
	LocationID string `json:"location_id"` // ID de la ubicación
	Quantity   int    `json:"quantity"`    // Nueva cantidad
}

// UserProductEligibility representa la elegibilidad de migración de un ítem
type UserProductEligibility struct {
	ItemID      string `json:"item_id"`       // ID del ítem a validar
	IsEligible  bool   `json:"is_eligible"`   // Si es elegible para migración
	// Campos opcionales
	Reason      *string `json:"reason,omitempty"`      // Razón de no elegibilidad
	Message     *string `json:"message,omitempty"`     // Mensaje descriptivo
}

// UserProductMigrationRequest representa una solicitud de migración
type UserProductMigrationRequest struct {
	ItemID string `json:"item_id"` // ID del ítem a migrar
}

// GetUserProductByID obtiene información detallada de un User Product por su ID
func GetUserProductByID(ctx context.Context, userProductID, accessToken string) (UserProduct, error) {
	url := fmt.Sprintf("%s/%s", userProductsEndpoint, userProductID)
	var userProduct UserProduct
	err := http.DoGetJSON(ctx, url, accessToken, &userProduct)
	return userProduct, err
}

// GetUserProductFamilyByID obtiene todos los User Products de una familia específica
func GetUserProductFamilyByID(ctx context.Context, siteID, familyID, accessToken string) (UserProductFamily, error) {
	url := fmt.Sprintf("%s/%s/user-products-families/%s", siteUserProductFamiliesEndpoint, siteID, familyID)
	var family UserProductFamily
	err := http.DoGetJSON(ctx, url, accessToken, &family)
	return family, err
}

// GetUserProductStock obtiene información del stock de un User Product
func GetUserProductStock(ctx context.Context, userProductID, accessToken string) (UserProductStock, error) {
	url := fmt.Sprintf("%s/%s/stock", userProductsEndpoint, userProductID)
	var stock UserProductStock
	err := http.DoGetJSON(ctx, url, accessToken, &stock)
	return stock, err
}

// ValidateItemEligibility valida si un ítem es elegible para migración al modelo User Products
func ValidateItemEligibility(ctx context.Context, itemID, accessToken string) (UserProductEligibility, error) {
	url := fmt.Sprintf("%s/%s/user_product_listings/validate", itemsEligibilityEndpoint, itemID)
	var eligibility UserProductEligibility
	err := http.DoGetJSON(ctx, url, accessToken, &eligibility)
	return eligibility, err
}