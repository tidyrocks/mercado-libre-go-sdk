package attr_groups

// TechnicalSpec representa especificaciones técnicas de categoría.
type TechnicalSpec struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Relevance int    `json:"relevance"`
}
