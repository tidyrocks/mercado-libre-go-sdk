# Mercado Libre Go SDK

SDK Go para la API REST de MercadoLibre con OAuth 2.0, gestión de errores y alto rendimiento.

## Instalación

```bash
go get github.com/tidyrocks/mercado-libre-go-sdk
```

## Autenticación

**Refresh access token**
```go
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (*RefreshTokenResponse, error)
```
→ Returns: [RefreshTokenResponse](auth/types.go#L12)

**Validate access token**
```go
func ValidateAccessToken(ctx context.Context, accessToken string) error
```

## Items

**Get item by ID**
```go
func GetByID(ctx context.Context, itemID, accessToken string) (Item, error)
```
→ Returns: [Item](items/types.go#L6)

## Categories

```go
// Get category by ID
func GetByID(id, accessToken string) (*[Category](categories/types.go#L4), error)

// Get categories by site
func GetBySite(siteID, accessToken string, params []shared.KeyValue) ([][Category](categories/types.go#L4), error)

// Get child categories
func GetChildren(categoryID, accessToken string) ([][Category](categories/types.go#L4), error)

// Predict category from title using ML
func PredictCategory(siteID, title, accessToken string, params []shared.KeyValue) ([][CategoryPrediction](categories/types.go#L17), error)

// Search categories
func Search(query, accessToken string, params []shared.KeyValue) ([][Category](categories/types.go#L4), error)
```

## Variations

```go
// Get all variations for an item
func GetByItemID(itemID, accessToken string) ([][Variation](variations/types.go#L6), error)

// Get specific variation
func GetByID(itemID, variationID, accessToken string) (*[Variation](variations/types.go#L6), error)
```

## Attributes

```go
// Get attributes by category
func GetByCategoryID(categoryID, accessToken string) ([][Attribute](attrs/types.go#L4), error)

// Get attribute groups
func GetByCategoryID(categoryID, accessToken string) ([][TechnicalSpec](attr_groups/types.go#L4), error)

// Get attribute values
func GetByAttrID(attrID, accessToken string) ([][AttributeValue](attr_values/types.go#L4), error)
```

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)
