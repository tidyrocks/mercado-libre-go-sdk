package categories

// Category representa una categoría de Mercado Libre.
type Category struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	PathFromRoot []PathNode `json:"path_from_root"`
}

// PathNode representa un nodo en el path de categorías.
type PathNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CategoryPrediction representa una predicción de categoría.
type CategoryPrediction struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	DomainID     string `json:"domain_id"`
	DomainName   string `json:"domain_name"`
}
