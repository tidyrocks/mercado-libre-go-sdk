package items

import "time"

// Item representa una publicación individual visible en el sitio de Mercado Libre.
type Item struct {
	ID                string       `json:"id"`                            // ID del ítem (ej. MLA123456)
	DateCreated       time.Time    `json:"date_created"`                  // Fecha de creación del ítem
	Title             string       `json:"title"`                         // Título visible del ítem
	Price             float64      `json:"price"`                         // Precio publicado
	CurrencyID        string       `json:"currency_id"`                   // Moneda del precio (ej. MXN)
	AvailableQuantity int          `json:"available_quantity"`            // Stock disponible
	BuyingMode        string       `json:"buying_mode"`                   // Modo de compra: "buy_it_now", etc.
	ListingTypeID     string       `json:"listing_type_id"`               // Tipo de publicación: "gold_special", etc.
	ItemCondition     string       `json:"condition"`                     // Condición del ítem: "new", "used", etc.
	Status            string       `json:"status"`                        // Estado del ítem: "active", "paused", "under_review", etc.
	Shipping          ItemShipping `json:"shipping"`                      // Información de envío
	UserProductID     *string      `json:"user_product_id,omitempty"`     // ID del UP asociado (si aplica)
	SellerCustomField *string      `json:"seller_custom_field,omitempty"` // SKU privado del vendedor
	InventoryID       *string      `json:"inventory_id,omitempty"`        // ID del inventario asociado (si aplica)
	// TODO: ver si category_id es soportada en ScanOptions y PaginateOptions
	CategoryID string `json:"category_id"` // ID de la categoría del ítem
}

type ItemShipping struct {
	Mode         string `json:"mode"`          // Modo de envío: "me2", "not_specified", etc.
	LogisticType string `json:"logistic_type"` // Tipo de logística: "drop_off", "fulfillment", etc.
}
