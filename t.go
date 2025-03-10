package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type T[Type any] struct {
	t Type
	v bool
}

func TOf[Type any](value Type) T[Type] {
	var t T[Type]
	t.Set(value)
	return t
}

// Set set the value.
func (t *T[Type]) Set(value Type) {
	t.v = true
	t.t = value
}

// Reset set nil to the value.
func (t *T[Type]) Reset() {
	t.t = *new(Type)
	t.v = false
}

// Valid return the value is valid. If true, it is not null value.
func (t *T[T]) Valid() bool {
	return t.v
}

// TimeValue return the value.
func (t *T[Type]) TValue() Type {
	if t.v {
		return t.t
	}

	return *new(Type)
}

// implement json.Marshaler.
func (t T[Type]) MarshalJSON() ([]byte, error) {
	if !t.v {
		return []byte{'n', 'u', 'l', 'l'}, nil
	}
	return json.Marshal(t.t)
}

// implement json.Unmarshaler.
func (t *T[Type]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.v = false
		return nil
	}
	t.v = true
	return json.Unmarshal(data, t.t)
}

// Value implement driver.Valuer.
func (t T[Type]) Value() (driver.Value, error) {

	return nil, fmt.Errorf("valuer doesn't support for %T type", t)
}
