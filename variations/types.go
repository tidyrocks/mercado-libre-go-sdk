package variations

import "time"

// Variation representa una variación de un ítem de Mercado Libre.
type Variation struct {
	ID                    int64                  `json:"id"`
	Price                 float64                `json:"price"`
	AvailableQuantity     int                    `json:"available_quantity"`
	SoldQuantity          int                    `json:"sold_quantity"`
	SellerCustomField     *string                `json:"seller_custom_field"`
	CatalogProductID      *string                `json:"catalog_product_id"`
	InventoryID           *string                `json:"inventory_id"`
	UserProductID         *string                `json:"user_product_id"`
	PictureIDs            []string               `json:"picture_ids"`
	AttributeCombinations []AttributeCombination `json:"attribute_combinations"`
	Attributes            []VariationAttribute   `json:"attributes"`
	DateCreated           *time.Time             `json:"date_created,omitempty"`
}

// AttributeCombination representa una combinación de atributos que define la variación.
type AttributeCombination struct {
	ID          *string               `json:"id"`
	Name        string                `json:"name"`
	ValueID     *string               `json:"value_id"`
	ValueName   string                `json:"value_name"`
	ValueStruct *AttributeValueStruct `json:"value_struct"`
	Values      []AttributeValue      `json:"values"`
}

// VariationAttribute representa un atributo específico de una variación.
type VariationAttribute struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	ValueID     *string               `json:"value_id"`
	ValueName   string                `json:"value_name"`
	ValueStruct *AttributeValueStruct `json:"value_struct"`
	Values      []AttributeValue      `json:"values"`
}

// AttributeValue representa un valor de atributo.
type AttributeValue struct {
	ID     *string               `json:"id"`
	Name   *string               `json:"name"`
	Struct *AttributeValueStruct `json:"struct"`
}

// AttributeValueStruct representa la estructura de un valor de atributo con unidades.
type AttributeValueStruct struct {
	Unit   *string  `json:"unit"`
	Number *float64 `json:"number"`
}

// VariationRequest representa el payload para crear o actualizar variaciones.
type VariationRequest struct {
	ID                    *int64                 `json:"id,omitempty"`
	Price                 *float64               `json:"price,omitempty"`
	AvailableQuantity     *int                   `json:"available_quantity,omitempty"`
	SellerCustomField     *string                `json:"seller_custom_field,omitempty"`
	CatalogProductID      *string                `json:"catalog_product_id,omitempty"`
	InventoryID           *string                `json:"inventory_id,omitempty"`
	UserProductID         *string                `json:"user_product_id,omitempty"`
	PictureIDs            []string               `json:"picture_ids,omitempty"`
	AttributeCombinations []AttributeCombination `json:"attribute_combinations,omitempty"`
	Attributes            []VariationAttribute   `json:"attributes,omitempty"`
}
