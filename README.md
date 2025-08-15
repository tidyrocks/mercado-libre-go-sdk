# Mercado Libre Go SDK

SDK Go para la API REST de MercadoLibre con OAuth 2.0, gestión de errores y alto rendimiento.

> **⚠️ Estado de Desarrollo**  
> Esta librería está en desarrollo activo. La API puede cambiar entre versiones.  
> Contribuciones y feedback son bienvenidos en la rama `develop`.

## Instalación

```bash
go get github.com/tidyrocks/mercado-libre-go-sdk
```

## Ejemplo de uso rápido

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/tidyrocks/mercado-libre-go-sdk/api"
)

func main() {
    ctx := context.Background()
    accessToken := "APP_USR-your-access-token"
    
    // Obtener un ítem
    item, err := api.GetItem(ctx, "MLM123456789", accessToken)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Item: %s - $%.2f\n", item.Title, item.Price)
}
```

## Recursos de la API

### Autenticación

```go
// RefreshAccessToken renueva un access token usando el refresh token
func RefreshAccessToken(ctx context.Context, clientID, clientSecret, refreshToken string) (Token, error)
```

**Retorna:** [Token](api/auth.go#L12)

### Items

```go
// GetItem obtiene un ítem por su ID
func GetItem(ctx context.Context, itemID, accessToken string) (Item, error)
```

**Retorna:** [Item](api/items.go#L14)

### Categorías

```go
// GetCategoryByID obtiene información detallada de una categoría por su ID
func GetCategoryByID(ctx context.Context, categoryID, accessToken string) (Category, error)

// GetCategoriesBySite obtiene el árbol de categorías de un sitio específico
func GetCategoriesBySite(ctx context.Context, siteID, accessToken string) ([]CategorySummary, error)

// GetCategoryAttributes obtiene los atributos disponibles para una categoría
func GetCategoryAttributes(ctx context.Context, categoryID, accessToken string) ([]Attr, error)
```

**Retorna:** [Category](api/categories.go#L14), [CategorySummary](api/categories.go#L85), [Attr](api/attrs.go#L4)

### Dominios

```go
// GetDomainByID obtiene información detallada de un dominio por su ID
func GetDomainByID(ctx context.Context, domainID, accessToken string) (Domain, error)

// GetDomainShippingAttributes obtiene los atributos de envío requeridos para un dominio
func GetDomainShippingAttributes(ctx context.Context, domainID, accessToken string) (DomainShippingAttributes, error)
```

**Retorna:** [Domain](api/domains.go#L12), [DomainShippingAttributes](api/domains.go#L21)

### Sitios

```go
// GetSites obtiene la lista de todos los sitios disponibles en Mercado Libre
func GetSites(ctx context.Context, accessToken string) ([]Site, error)
```

**Retorna:** [Site](api/sites.go#L10)

### Variaciones

```go
// Definidos en api/variations.go (parte del modelo de Item)
```

**Retorna:** [Variation](api/variations.go#L6)

### Atributos

```go
// Definidos en api/attrs.go (utilizados por múltiples recursos)
```

**Retorna:** [Attr](api/attrs.go#L4), [AttrVal](api/attrs.go#L26), [MeasuredValue](api/attrs.go#L43)

### Imágenes

```go
// Definidos en api/pictures.go (parte del modelo de Item y UserProduct)
```

**Retorna:** [Picture](api/pictures.go#L4)

### User Products

```go
// GetUserProductByID obtiene información detallada de un User Product por su ID
func GetUserProductByID(ctx context.Context, userProductID, accessToken string) (UserProduct, error)

// GetUserProductFamilyByID obtiene todos los User Products de una familia específica
func GetUserProductFamilyByID(ctx context.Context, siteID, familyID, accessToken string) (UserProductFamily, error)

// GetUserProductStock obtiene información del stock de un User Product
func GetUserProductStock(ctx context.Context, userProductID, accessToken string) (UserProductStock, error)

// ValidateItemEligibility valida si un ítem es elegible para migración al modelo User Products
func ValidateItemEligibility(ctx context.Context, itemID, accessToken string) (UserProductEligibility, error)
```

**Retorna:** [UserProduct](api/user_products.go#L16), [UserProductFamily](api/user_products.go#L32), [UserProductStock](api/user_products.go#L42), [UserProductEligibility](api/user_products.go#L75)

## Arquitectura

- **`api/`**: Funciones públicas y tipos de datos
- **`internal/http/`**: Cliente HTTP con autenticación Bearer
- **`docs/`**: Documentación de la API de MELI
- **`CLAUDE.md`**: Instrucciones para el desarrollo

## Principios de diseño

- **Funciones simples**: `func(ctx, params..., accessToken) (Type, error)`
- **Context-first**: Soporte completo para cancelación y timeouts
- **Tipos idiomáticos**: Structs Go con tags JSON descriptivos
- **Sin sorpresas**: Campos consistentes, arrays siempre presentes
- **Nomenclatura clara**: Nombres descriptivos en español en comentarios

Licencia MIT - Creado por [Gus Salazar](https://www.linkedin.com/in/gussalazar/)