package attrs

// Attribute representa un atributo de categoría.
type Attribute struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	ValueType string          `json:"value_type"`
	Tags      map[string]bool `json:"tags"`
}

// AttributeRequest representa un atributo en requests.
type AttributeRequest struct {
	ID        string  `json:"id"`
	ValueName string  `json:"value_name"`
	ValueID   *string `json:"value_id,omitempty"`
}
