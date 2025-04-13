package models

type Optional[T any] struct {
	Value T    `json:"value"`
	Set   bool `json:"set"`
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		Value: value,
		Set:   true,
	}
}
