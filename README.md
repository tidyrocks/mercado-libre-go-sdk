# Mercado Libre Go SDK

SDK Go para la API REST de MercadoLibre con OAuth 2.0, gestión de errores y alto rendimiento.

> **⚠️ Estado de Desarrollo**  
> Esta librería está en desarrollo activo. La API puede cambiar entre versiones.  
> Contribuciones y feedback son bienvenidos en la rama `develop`.

## Instalación

```bash
go get github.com/tidyrocks/mercado-libre-go-sdk
```

## 🔐 Autenticación

**RefreshAccessToken** - Renueva tu access token cuando expira
```go
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (*RefreshTokenResponse, error)
```

**ValidateAccessToken** - Verifica que tu token sea válido
```go
func ValidateAccessToken(ctx context.Context, accessToken string) error
```

## 📦 Productos (Items)

**GetByID** - Busca un producto por su ID (ej: MLM123456789)
```go
func GetByID(ctx context.Context, itemID, accessToken string) (Item, error)
```

## 🏷️ Categorías

**GetByID** - Obtiene info de una categoría (ej: MLM1051 = Celulares)  
**GetBySite** - Lista todas las categorías de un país (MLM, MLA, etc)  
**GetChildren** - Busca subcategorías dentro de una categoría  
**PredictCategory** - Adivina la categoría por el título del producto  
**Search** - Busca categorías por palabra clave

```go
func GetByID(id string, accessToken string) (*Category, error)
func GetBySite(siteID string, accessToken string) ([]Category, error)  
func GetChildren(categoryID string, accessToken string) ([]Category, error)
func PredictCategory(siteID, title string, accessToken string) ([]CategoryPrediction, error)
func Search(query string, accessToken string) ([]Category, error)
```

## 🎨 Variaciones (Colores, Tallas, etc)

**GetByItemID** - Obtiene todas las variaciones de un producto (tallas, colores)  
**GetByItemIDWithAttributes** - Igual pero con más detalles de cada variación

```go
func GetByItemID(itemID string, accessToken string) ([]Variation, error)
func GetByItemIDWithAttributes(itemID string, accessToken string) ([]Variation, error)
```

## 📋 Atributos y Especificaciones

**GetByCategory** - Lista qué datos necesitas para publicar en una categoría (marca, modelo, etc)  
**GetTechnicalSpecsInput** - Especificaciones técnicas que puedes agregar al producto  
**GetTechnicalSpecsOutput** - Cómo se muestran las especificaciones al comprador  
**GetTopValues** - Valores más populares para un atributo (ej: marcas más vendidas)  
**GetTopValuesWithFilter** - Igual pero filtrado por otros atributos ya seleccionados

```go
func GetByCategory(categoryID string, accessToken string) ([]Attribute, error)                      
func GetTechnicalSpecsInput(categoryID string, accessToken string) (*TechnicalSpecsResponse, error)    
func GetTechnicalSpecsOutput(categoryID string, accessToken string) (*TechnicalSpecsOutputResponse, error) 
func GetTopValues(domainID, attributeID string, accessToken string) ([]AttributeValue, error) 
func GetTopValuesWithFilter(domainID, attributeID string, knownAttributes []KnownAttribute, accessToken string) ([]AttributeValue, error)
```

## 📸 Imágenes

**Upload** - Sube una imagen a los servidores de ML  
**LinkToItem** - Conecta una imagen ya subida con tu producto  
**GetErrors** - Revisa si hubo errores al procesar la imagen  
**UpdateItemPictures** - Cambia todas las fotos de un producto

```go
func Upload(fileContent []byte, filename string, accessToken string) (*PictureUploadResponse, error)
func LinkToItem(itemID, pictureID string, accessToken string) error
func GetErrors(pictureID string, accessToken string) (*PictureError, error)
func UpdateItemPictures(itemID string, pictures []PictureRequest, accessToken string) error
```

## 👤 User Products (Productos del Usuario)

**GetByID** - Busca un User Product específico  
**GetFamilyByID** - Info de una familia de productos  
**GetUserProductsByFamily** - Todos los User Products de una familia  
**GetItemsByUserProduct** - Todos los ítems publicados usando un User Product  
**GetItemsByMultipleUserProducts** - Ítems de varios User Products a la vez  
**CheckEligibility** - Verifica si un ítem puede migrar al modelo User Products  
**GetByIDWithStock** - User Product con info de inventario por ubicación

```go
func GetByID(userProductID string, accessToken string) (*UserProduct, error)
func GetFamilyByID(siteID, familyID string, accessToken string) (*Family, error)
func GetUserProductsByFamily(siteID, familyID string, accessToken string) ([]UserProduct, error)
func GetItemsByUserProduct(sellerID, userProductID string, accessToken string) (*ItemSearchResult, error)
func GetItemsByMultipleUserProducts(sellerID string, userProductIDs []string, accessToken string) (*ItemSearchResult, error)
func CheckEligibility(itemID string, accessToken string) (bool, error)
func GetByIDWithStock(userProductID string, accessToken string) (*UserProduct, error)
```

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)