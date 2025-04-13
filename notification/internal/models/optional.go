package models

type Optional[T any] struct {
	Value T    `json:"value,omitempty"`
	Set   bool `json:"set,omitempty"`
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		Value: value,
		Set:   true,
	}
}