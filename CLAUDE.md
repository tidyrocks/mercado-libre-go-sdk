# CLAUDE.md

Este archivo proporciona orientación a Claude Code (claude.ai/code) cuando trabaja con código en este repositorio.

## Comandos Comunes

### Desarrollo

```bash
go run .           # Ejecutar el proyecto (si hay main)
go build .         # Compilar el proyecto
go mod tidy        # Limpiar dependencias
go fmt ./...       # Formatear código
```

### Testing

```bash
go test ./...      # Ejecutar todos los tests
go test -v ./...   # Tests con output verboso
go test ./categories -v  # Test de un paquete específico
```

## Arquitectura del Proyecto

Este es un SDK de Go para interactuar con la API de Mercado Libre.

### Estructura Principal

- **`internal/httpx/`**: Contiene utilidades HTTP compartidas. El paquete `httpx` proporciona la función genérica `DoGetJSON` que maneja las llamadas GET con autenticación Bearer token y deserialización JSON.

- **`categories/` y `items/`**: Módulos principales del SDK que siguen el mismo patrón:
  - `api.go`: Funciones públicas que llaman a la API de Mercado Libre
  - `types.go`: Structs que representan los modelos de datos de la API

### Patrones de Diseño

- Todas las funciones de API requieren `context.Context` como primer parámetro
- Las funciones de API retornan `(tipo, error)` siguiendo convenciones de Go
- Los endpoints de la API están definidos como constantes en cada módulo
- El token de acceso es opcional y se pasa como string vacío cuando no se necesita
- Los modelos JSON usan tags descriptivos y comentarios en español
- Se usa generics en `httpx.DoGetJSON` para type safety

### Endpoints Disponibles

- `categories.GetByID()`: Obtiene información de una categoría por su ID
- `items.GetByID()`: Obtiene información de un ítem por su ID

### Convenciones del Código

- Comentarios y documentación en español
- Nombres de funciones en inglés siguiendo convenciones de Go
- Structs con tags JSON descriptivos
- Manejo de
