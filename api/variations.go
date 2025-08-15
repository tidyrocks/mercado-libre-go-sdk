package api

// Variation representa una variación de un ítem
type Variation struct {
	ID                int64   `json:"id"`
	Price             float64 `json:"price"`
	AttrCombinations  []Attr  `json:"attribute_combinations"`
	AvailableQuantity int     `json:"available_quantity"`
	SoldQuantity      int     `json:"sold_quantity"`
	// field > SaleTerms
	PictureIDs []string `json:"picture_ids"`

	//Attrs             []Attr     `json:"attributes"`
	SellerCustomField *string `json:"seller_custom_field,omitempty"`
	CatalogProductID  *string `json:"catalog_product_id,omitempty"`
	InventoryID       *string `json:"inventory_id,omitempty"`
	// field > ItemRelations
	UserProductID *string `json:"user_product_id,omitempty"`
}
