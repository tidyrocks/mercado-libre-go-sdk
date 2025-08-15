package api

// Picture representa una imagen de un producto en Mercado Libre.
type Picture struct {
	ID        string `json:"id"`         // ID único de la imagen
	URL       string `json:"url"`        // URL de la imagen
	SecureURL string `json:"secure_url"` // URL segura (HTTPS) de la imagen
	Size      string `json:"size"`       // Dimensiones de la imagen (ej: "461x500")
	MaxSize   string `json:"max_size"`   // Dimensiones máximas disponibles
	Quality   string `json:"quality"`    // Calidad de la imagen
}
