package api

import (
	"context"
	"fmt"
	"time"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const itemsEndpoint = "https://api.mercadolibre.com/items"

// Item representa una publicación individual visible en el sitio de Mercado Libre.
type Item struct {
	ID                 string       `json:"id"`                  // ID del ítem (ej. MLA123456)
	SiteID             string       `json:"site_id"`             // ID del sitio (MLM, MLA, etc.)
	Title              string       `json:"title"`               // Título visible del ítem
	Price              float64      `json:"price"`               // Precio publicado
	BasePrice          float64      `json:"base_price"`          // Precio base
	CurrencyID         string       `json:"currency_id"`         // Moneda del precio (ej. MXN)
	InitialQuantity    int          `json:"initial_quantity"`    // Cantidad inicial
	AvailableQuantity  int          `json:"available_quantity"`  // Stock disponible
	SoldQuantity       int          `json:"sold_quantity"`       // Cantidad vendida
	BuyingMode         string       `json:"buying_mode"`         // Modo de compra: "buy_it_now", etc.
	ListingTypeID      string       `json:"listing_type_id"`     // Tipo de publicación: "gold_special", etc.
	Condition          string       `json:"condition"`           // Condición del ítem: "new", "used", etc.
	Status             string       `json:"status"`              // Estado del ítem: "active", "paused", "under_review", etc.
	DateCreated        time.Time    `json:"date_created"`        // Fecha de creación del ítem
	LastUpdated        time.Time    `json:"last_updated"`        // Fecha de última actualización
	CategoryID         string       `json:"category_id"`         // ID de la categoría del ítem
	SellerID           int64        `json:"seller_id"`           // ID del vendedor
	DomainID           string       `json:"domain_id"`           // ID del dominio
	Permalink          string       `json:"permalink"`           // URL permanente
	ThumbnailID        string       `json:"thumbnail_id"`        // ID de thumbnail
	Thumbnail          string       `json:"thumbnail"`           // URL del thumbnail
	Health             float64      `json:"health"`              // Salud del ítem (0-1)
	CatalogListing     bool         `json:"catalog_listing"`     // Si es listing de catálogo
	AcceptsMercadoPago bool         `json:"accepts_mercadopago"` // Si acepta MercadoPago
	ShippingConfig     ItemShipping `json:"shipping"`            // Configuración de envío
	// Arrays y slices
	Tags       []string    `json:"tags"`       // Tags del ítem
	Channels   []string    `json:"channels"`   // Canales de venta
	Variations []Variation `json:"variations"` // Variaciones del ítem
	Pictures   []Picture   `json:"pictures"`   // Imágenes del ítem
	Attrs      []Attr      `json:"attributes"` // Atributos del ítem
	// Campos opcionales (punteros)
	OriginalPrice     *float64   `json:"original_price,omitempty"`      // Precio original (si hay descuento)
	StartTime         *time.Time `json:"start_time,omitempty"`          // Fecha de inicio
	StopTime          *time.Time `json:"stop_time,omitempty"`           // Fecha de fin
	EndTime           *time.Time `json:"end_time,omitempty"`            // Fecha de finalización
	ExpirationTime    *time.Time `json:"expiration_time,omitempty"`     // Fecha de expiración
	OfficialStoreID   *int64     `json:"official_store_id,omitempty"`   // ID de tienda oficial
	UserProductID     *string    `json:"user_product_id,omitempty"`     // ID del UP asociado (si aplica)
	SellerCustomField *string    `json:"seller_custom_field,omitempty"` // SKU privado del vendedor
	InventoryID       *string    `json:"inventory_id,omitempty"`        // ID del inventario asociado (si aplica)
	CatalogProductID  *string    `json:"catalog_product_id,omitempty"`  // ID de producto de catálogo
	ParentItemID      *string    `json:"parent_item_id,omitempty"`      // ID del ítem padre
}

type ItemShipping struct {
	Mode         string   `json:"mode"` // "me2", etc.
	Methods      []string `json:"methods"`
	Tags         []string `json:"tags"`
	Dimensions   *string  `json:"dimensions"`
	LocalPickup  bool     `json:"local_pick_up"`
	FreeShipping bool     `json:"free_shipping"`
	LogisticType string   `json:"logistic_type"`
	StorePickup  bool     `json:"store_pick_up"`
}

// GetItem obtiene un ítem por ID usando la API de Mercado Libre.
func GetItem(ctx context.Context, itemID, accessToken string) (Item, error) {
	url := fmt.Sprintf("%s/%s", itemsEndpoint, itemID)
	var item Item
	err := http.DoGetJSON(ctx, url, accessToken, &item)
	return item, err
}
