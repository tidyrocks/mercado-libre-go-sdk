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
// RefreshAccessToken renueva un access token usando el refresh token
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (*RefreshTokenResponse, error)

// ValidateAccessToken verifica que un access token sea válido
func ValidateAccessToken(ctx context.Context, accessToken string) error
```

**Returns:** [RefreshTokenResponse](auth/types.go#L12)

## Items

```go
// GetByID obtiene un item por su ID
func GetByID(ctx context.Context, itemID, accessToken string) (Item, error)
```

**Returns:** [Item](items/types.go#L6)

## Categories

```go
// GetByID obtiene una categoría por ID
func GetByID(id string, accessToken string) (*Category, error)

// GetBySite obtiene todas las categorías de un sitio
func GetBySite(siteID string, accessToken string) ([]Category, error)

// GetChildren obtiene las categorías hijas de una categoría
func GetChildren(categoryID string, accessToken string) ([]Category, error)

// PredictCategory predice categorías basado en el título del producto
func PredictCategory(siteID, title string, accessToken string) ([]CategoryPrediction, error)

// Search busca categorías por query
func Search(query string, accessToken string) ([]Category, error)
```

**Returns:** [Category](categories/types.go#L4), [CategoryPrediction](categories/types.go#L17)

## Variations

```go
// GetByItemID obtiene todas las variaciones de un item
func GetByItemID(itemID string, accessToken string) ([]Variation, error)

// GetByItemIDWithAttributes obtiene variaciones con atributos detallados
func GetByItemIDWithAttributes(itemID string, accessToken string) ([]Variation, error)
```

**Returns:** [Variation](variations/types.go#L6)

## Attributes

```go
// GetByCategory obtiene los atributos de una categoría
func GetByCategory(categoryID string, accessToken string) ([]Attribute, error)

// GetTechnicalSpecsInput obtiene especificaciones técnicas de entrada
func GetTechnicalSpecsInput(categoryID string, accessToken string) (*TechnicalSpecsResponse, error)

// GetTechnicalSpecsOutput obtiene especificaciones técnicas de salida
func GetTechnicalSpecsOutput(categoryID string, accessToken string) (*TechnicalSpecsOutputResponse, error)

// GetTopValues obtiene los valores más populares para un atributo
func GetTopValues(domainID, attributeID string, accessToken string) ([]AttributeValue, error)

// GetTopValuesWithFilter obtiene valores filtrados por atributos conocidos
func GetTopValuesWithFilter(domainID, attributeID string, knownAttributes []KnownAttribute, accessToken string) ([]AttributeValue, error)
```

**Returns:** [Attribute](attrs/types.go#L4), [TechnicalSpecsResponse](attr_groups/types.go#L4), [TechnicalSpecsOutputResponse](attr_groups/types.go#L23), [AttributeValue](attr_values/types.go#L4), [KnownAttribute](attr_values/types.go#L14)

## Pictures

```go
// Upload sube una imagen a los servidores de ML
func Upload(fileContent []byte, filename string, accessToken string) (*PictureUploadResponse, error)

// LinkToItem vincula una imagen a un item
func LinkToItem(itemID, pictureID string, accessToken string) error

// GetErrors obtiene errores de procesamiento de una imagen
func GetErrors(pictureID string, accessToken string) (*PictureError, error)

// UpdateItemPictures actualiza las imágenes de un item
func UpdateItemPictures(itemID string, pictures []PictureRequest, accessToken string) error
```

**Returns:** [PictureUploadResponse](pictures/types.go#L32), [PictureError](pictures/types.go#L25), [PictureRequest](pictures/types.go#L19)

## User Products

```go
// GetByID obtiene un User Product por ID
func GetByID(userProductID string, accessToken string) (*UserProduct, error)

// GetFamilyByID obtiene información de una familia de productos
func GetFamilyByID(siteID, familyID string, accessToken string) (*Family, error)

// GetUserProductsByFamily obtiene User Products de una familia
func GetUserProductsByFamily(siteID, familyID string, accessToken string) ([]UserProduct, error)

// GetItemsByUserProduct obtiene items asociados a un User Product
func GetItemsByUserProduct(sellerID, userProductID string, accessToken string) (*ItemSearchResult, error)

// GetItemsByMultipleUserProducts obtiene items de múltiples User Products
func GetItemsByMultipleUserProducts(sellerID string, userProductIDs []string, accessToken string) (*ItemSearchResult, error)

// CheckEligibility verifica elegibilidad para migrar a User Products
func CheckEligibility(itemID string, accessToken string) (bool, error)

// GetByIDWithStock obtiene User Product con información de stock
func GetByIDWithStock(userProductID string, accessToken string) (*UserProduct, error)
```

**Returns:** [UserProduct](user_products/types.go#L6), [Family](user_products/types.go#L45), [ItemSearchResult](user_products/types.go#L25)

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)