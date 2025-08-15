package api

// Attr representa un atributo (usado en items, variations, categories, etc.)
type Attr struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	ValueType       string              `json:"value_type"`
	ValueName       string              `json:"value_name"`
	Values          []AttrVal           `json:"values"`
	AllowedUnits    []UnitOfMeasurement `json:"allowed_units"`
	SuggestedValues []AttrVal           `json:"suggested_values"`
	AttrGroupID     string              `json:"attribute_group_id"`
	AttrGroupName   string              `json:"attribute_group_name"`
	Tags            map[string]bool     `json:"tags"`
	Type            *string             `json:"type,omitempty"`
	DefaultUnit     *string             `json:"default_unit,omitempty"`
	Hierarchy       *string             `json:"hierarchy,omitempty"`
	Relevance       *int                `json:"relevance,omitempty"`
	Tooltip         *string             `json:"tooltip,omitempty"`
	ValueMaxLength  *int                `json:"value_max_length,omitempty"`
	ValueID         *string             `json:"value_id,omitempty"`
	MeasuredValue   *MeasuredValue      `json:"value_struct,omitempty"`
}

// AttrVal representa un valor dentro de un atributo
type AttrVal struct {
	ID     *string        `json:"id"`
	Name   *string        `json:"name"`
	Struct *MeasuredValue `json:"struct"`
	Metric *int           `json:"metric"`
	//Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// UnitOfMeasurement representa una unidad de medida permitida
// Ejemplo: {ID: "kg", Name: "kg"} o {ID: "cm", Name: "cm"}
type UnitOfMeasurement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MeasuredValue representa estructura de valores con unidades/números
// Ejemplo: {Number: 2.5, Unit: "kg"} o {Number: 150, Unit: "cm"}
type MeasuredValue struct {
	Number *float64 `json:"number"`
	Unit   *string  `json:"unit"`
}
