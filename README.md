# Mercado Libre Go SDK

SDK Go para la API REST de MercadoLibre con OAuth 2.0, gestión de errores y alto rendimiento.

> **⚠️ Estado de Desarrollo**  
> Esta librería está en desarrollo activo. La API puede cambiar entre versiones.  
> Contribuciones y feedback son bienvenidos en la rama `develop`.

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
func GetByID(id string, accessToken string) (*Category, error)
func GetBySite(siteID string, params url.Values, accessToken string) ([]Category, error)
func GetChildren(categoryID string, accessToken string) ([]Category, error)
func PredictCategory(siteID, title string, params url.Values, accessToken string) ([]CategoryPrediction, error)
func Search(query string, params url.Values, accessToken string) ([]Category, error)
```

**Returns:** [Category](categories/types.go#L4), [CategoryPrediction](categories/types.go#L17)

## Variations

```go
func GetByItemID(itemID string, accessToken string) ([]Variation, error)
func GetByItemIDWithAttributes(itemID string, accessToken string) ([]Variation, error)
```

**Returns:** [Variation](variations/types.go#L6)

## Attributes

```go
func GetByCategoryID(categoryID string, accessToken string) ([]Attribute, error)                      // attrs
func GetTechnicalSpecsInput(categoryID string, accessToken string) (*TechnicalSpecsResponse, error)    // attr_groups
func GetTechnicalSpecsOutput(categoryID string, accessToken string) (*TechnicalSpecsOutputResponse, error) // attr_groups
func GetTopValues(domainID, attributeID string, params url.Values, accessToken string) ([]AttributeValue, error) // attr_values
func GetTopValuesWithFilter(domainID, attributeID string, knownAttributes []KnownAttribute, accessToken string) ([]AttributeValue, error) // attr_values
```

**Returns:** [Attribute](attrs/types.go#L4), [TechnicalSpecsResponse](attr_groups/types.go#L4), [TechnicalSpecsOutputResponse](attr_groups/types.go#L23), [AttributeValue](attr_values/types.go#L4), [KnownAttribute](attr_values/types.go#L14)

## Pictures

```go
func Upload(fileContent []byte, filename string, accessToken string) (*PictureUploadResponse, error)
func LinkToItem(itemID, pictureID string, accessToken string) error
func GetErrors(pictureID string, accessToken string) (*PictureError, error)
func UpdateItemPictures(itemID string, pictures []PictureRequest, accessToken string) error
```

**Returns:** [PictureUploadResponse](pictures/types.go#L32), [PictureError](pictures/types.go#L25), [PictureRequest](pictures/types.go#L19)

## User Products

```go
func GetByID(userProductID string, accessToken string) (*UserProduct, error)
func GetFamilyByID(siteID, familyID string, accessToken string) (*Family, error)
func GetUserProductsByFamily(siteID, familyID string, accessToken string) ([]UserProduct, error)
func GetItemsByUserProduct(sellerID, userProductID string, params url.Values, accessToken string) (*ItemSearchResult, error)
func GetItemsByMultipleUserProducts(sellerID string, userProductIDs []string, params url.Values, accessToken string) (*ItemSearchResult, error)
func CheckEligibility(itemID string, accessToken string) (bool, error)
func GetByIDWithStock(userProductID string, accessToken string) (*UserProduct, error)
```

**Params:** [url.Values](https://pkg.go.dev/net/url#Values)  
**Returns:** [UserProduct](user_products/types.go#L6), [Family](user_products/types.go#L45), [ItemSearchResult](user_products/types.go#L25)

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)
