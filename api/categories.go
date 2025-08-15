package api

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/http"
)

const categoriesEndpoint = "https://api.mercadolibre.com/categories"
const siteCategoriesEndpoint = "https://api.mercadolibre.com/sites"

// Category representa una categoría de productos en Mercado Libre
type Category struct {
	ID                   string           `json:"id"`                     // ID de la categoría (ej. MLA5725)
	Name                 string           `json:"name"`                   // Nombre de la categoría
	Picture              string           `json:"picture"`                // URL de imagen
	Permalink            string           `json:"permalink"`              // URL permanente
	TotalItemsInCategory int64            `json:"total_items_in_this_category"`
	PathFromRoot         []CategoryPath   `json:"path_from_root"`         // Ruta desde la raíz
	ChildrenCategories   []CategorySummary `json:"children_categories"`   // Subcategorías
	AttributeTypes       string           `json:"attribute_types"`        // Tipos de atributos
	Settings             CategorySettings `json:"settings"`               // Configuración
	ChannelsSettings     []ChannelSetting `json:"channels_settings"`      // Configuración por canal
	MetaCategoryID       string           `json:"meta_categ_id"`
}

// CategoryPath representa un elemento en la ruta de categoría
type CategoryPath struct {
	ID   string `json:"id"`   // ID de la categoría padre
	Name string `json:"name"` // Nombre de la categoría padre
}

// CategorySettings representa la configuración de una categoría
type CategorySettings struct {
	AdultContent                    bool     `json:"adult_content"`
	BuyingAllowed                   bool     `json:"buying_allowed"`
	BuyingModes                     []string `json:"buying_modes"`
	CatalogDomain                   string   `json:"catalog_domain"`
	CoverageAreas                   string   `json:"coverage_areas"`
	CurrenciesAllowed               []string `json:"currencies_allowed"`
	FragileCategoryFlag             bool     `json:"fragile_category_flag"`
	ImmediatePayment                string   `json:"immediate_payment"`
	ItemConditions                  []string `json:"item_conditions"`
	ItemsReviewsAllowed             bool     `json:"items_reviews_allowed"`
	ListingAllowed                  bool     `json:"listing_allowed"`
	MaxDescriptionLength            int      `json:"max_description_length"`
	MaxPicturesPerItem              int      `json:"max_pictures_per_item"`
	MaxPicturesPerItemVar           int      `json:"max_pictures_per_item_var"`
	MaxSubTitle                     int      `json:"max_sub_title"`
	MaxTitle                        int      `json:"max_title"`
	MaximumPrice                    string   `json:"maximum_price"`
	MaximumPriceCurrency            string   `json:"maximum_price_currency"`
	MinimumPrice                    float64  `json:"minimum_price"`
	MinimumPriceCurrency            string   `json:"minimum_price_currency"`
	Price                           string   `json:"price"`
	RestrictionType                 string   `json:"restriction_type"`
	RoundedAddress                  bool     `json:"rounded_address"`
	SellerContact                   string   `json:"seller_contact"`
	ShippingModes                   []string `json:"shipping_modes"`
	ShippingOptions                 []string `json:"shipping_options"`
	ShippingProfile                 string   `json:"shipping_profile"`
	ShowContactInformation          bool     `json:"show_contact_information"`
	SimpleShipping                  string   `json:"simple_shipping"`
	Stock                           string   `json:"stock"`
	SubVertical                     string   `json:"sub_vertical"`
	TagsRequired                    bool     `json:"tags_required"`
	Vertical                        string   `json:"vertical"`
	VariationsAllowed               bool     `json:"variations_allowed"`
	VariationsAttributesAllowed     []string `json:"variations_attributes_allowed"`
	VipSubdomain                    string   `json:"vip_subdomain"`
	// Campos opcionales
	MirrorCategory       *string `json:"mirror_category,omitempty"`
	MirrorMasterCategory *string `json:"mirror_master_category,omitempty"`
	MirrorSlaveCategory  *string `json:"mirror_slave_category,omitempty"`
}

// ChannelSetting representa configuración por canal
type ChannelSetting struct {
	Channel  string                `json:"channel"`
	Settings CategoryChannelConfig `json:"settings"`
}

// CategoryChannelConfig representa configuración específica de canal
type CategoryChannelConfig struct {
	BuyingAllowed    bool     `json:"buying_allowed"`
	BuyingModes      []string `json:"buying_modes"`
	Catalogs         []string `json:"catalogs"`
	ExposureByRules  []string `json:"exposure_by_rules"`
	ListingAllowed   bool     `json:"listing_allowed"`
	ListingTypes     []string `json:"listing_types"`
	Price            string   `json:"price"`
	Stock            string   `json:"stock"`
}

// CategorySummary representa una categoría resumida (para listas)
type CategorySummary struct {
	ID   string `json:"id"`   // ID de la categoría
	Name string `json:"name"` // Nombre de la categoría
}

// GetCategoryByID obtiene información detallada de una categoría por su ID
func GetCategoryByID(ctx context.Context, categoryID, accessToken string) (Category, error) {
	url := fmt.Sprintf("%s/%s", categoriesEndpoint, categoryID)
	var category Category
	err := http.DoGetJSON(ctx, url, accessToken, &category)
	return category, err
}

// GetCategoriesBySite obtiene el árbol de categorías de un sitio específico
func GetCategoriesBySite(ctx context.Context, siteID, accessToken string) ([]CategorySummary, error) {
	url := fmt.Sprintf("%s/%s/categories", siteCategoriesEndpoint, siteID)
	var categories []CategorySummary
	err := http.DoGetJSON(ctx, url, accessToken, &categories)
	return categories, err
}

// GetCategoryAttributes obtiene los atributos disponibles para una categoría
func GetCategoryAttributes(ctx context.Context, categoryID, accessToken string) ([]Attr, error) {
	url := fmt.Sprintf("%s/%s/attributes", categoriesEndpoint, categoryID)
	var attrs []Attr
	err := http.DoGetJSON(ctx, url, accessToken, &attrs)
	return attrs, err
}