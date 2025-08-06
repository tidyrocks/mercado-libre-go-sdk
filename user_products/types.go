package user_products

import "time"

// UserProduct representa un User Product (UP) de Mercado Libre.
type UserProduct struct {
	ID                string    `json:"id"`
	SellerID          int64     `json:"seller_id"`
	SiteID            string    `json:"site_id"`
	Title             string    `json:"title"`
	CategoryID        string    `json:"category_id"`
	Condition         string    `json:"condition"`
	AvailableQuantity int       `json:"available_quantity"`
	SoldQuantity      int       `json:"sold_quantity"`
	Status            string    `json:"status"`
	DateCreated       time.Time `json:"date_created"`
	LastUpdated       time.Time `json:"last_updated"`
	FamilyID          *string   `json:"family_id"`
	FamilyName        *string   `json:"family_name"`
	DomainID          *string   `json:"domain_id"`
	CatalogProductID  *string   `json:"catalog_product_id"`
}

// ItemSearchResult representa el resultado de búsqueda de ítems.
type ItemSearchResult struct {
	Results []ItemResult `json:"results"`
	Paging  Paging       `json:"paging"`
}

// ItemResult representa un ítem en resultados de búsqueda.
type ItemResult struct {
	ID            string    `json:"id"`
	UserProductID string    `json:"user_product_id"`
	Status        string    `json:"status"`
	Price         float64   `json:"price"`
	DateCreated   time.Time `json:"date_created"`
}

// Family representa una familia de User Products.
type Family struct {
	ID       string `json:"id"`
	SiteID   string `json:"site_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// Paging representa información de paginación.
type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
