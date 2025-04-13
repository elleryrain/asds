package dto

type BasicResponse[T any] struct {
	Embedded T `json:"_embedded"`
}

type CustomFieldsValue struct {
	FieldID int     `json:"field_id"`
	Values  []Value `json:"values"`
}

type Value struct {
	Value    interface{} `json:"value,omitempty"`
	EnumID   int         `json:"enum_id,omitempty"`
	EnumCode string      `json:"enum_code,omitempty"`
}
