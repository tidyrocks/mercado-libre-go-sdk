package attr_values

// AttributeValue representa un valor de atributo.
type AttributeValue struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Metric *int   `json:"metric"`
}

// KnownAttribute representa un atributo conocido.
type KnownAttribute struct {
	ID      string `json:"id"`
	ValueID string `json:"value_id"`
}
