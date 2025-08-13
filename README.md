# Mercado Libre Go SDK

SDK Go para la API REST de MercadoLibre con OAuth 2.0, gestión de errores y alto rendimiento.

## Instalación

```bash
go get github.com/tidyrocks/mercado-libre-go-sdk
```

## Autenticación

```go
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (*RefreshTokenResponse, error)
func ValidateAccessToken(ctx context.Context, accessToken string) error
```

**Returns:** [RefreshTokenResponse](auth/types.go#L12)

## Items

```go
func GetByID(ctx context.Context, itemID, accessToken string) (Item, error)
```

**Returns:** [Item](items/types.go#L6)

## Categories

```go
func GetByID(id, accessToken string) (*Category, error)
func GetBySite(siteID, accessToken string, params []shared.KeyValue) ([]Category, error)
func GetChildren(categoryID, accessToken string) ([]Category, error)
func PredictCategory(siteID, title, accessToken string, params []shared.KeyValue) ([]CategoryPrediction, error)
func Search(query, accessToken string, params []shared.KeyValue) ([]Category, error)
```

**Params:** [shared.KeyValue](https://pkg.go.dev/gitlab.com/tidyrocks/tidy-go-common/shared#KeyValue)  
**Returns:** [Category](categories/types.go#L4), [CategoryPrediction](categories/types.go#L17)

## Variations

```go
func GetByItemID(itemID, accessToken string) ([]Variation, error)
func GetByID(itemID, variationID, accessToken string) (*Variation, error)
```

**Returns:** [Variation](variations/types.go#L6)

## Attributes

```go
func GetByCategoryID(categoryID, accessToken string) ([]Attribute, error)     // attrs
func GetByCategoryID(categoryID, accessToken string) ([]TechnicalSpec, error) // attr_groups  
func GetByAttrID(attrID, accessToken string) ([]AttributeValue, error)       // attr_values
```

**Returns:** [Attribute](attrs/types.go#L4), [TechnicalSpec](attr_groups/types.go#L4), [AttributeValue](attr_values/types.go#L4)

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)
