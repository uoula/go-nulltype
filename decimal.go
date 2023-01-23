package nulltype

import (
	"database/sql/driver"

	"github.com/shopspring/decimal"
)

// NullDecimal is null friendly type for Decimal.
type NullDecimal struct {
	f decimal.NullDecimal
}

// NullDecimalOf return NullDecimal that he value is set.
func NullDecimalOf(value decimal.Decimal) NullDecimal {
	var s NullDecimal
	s.Set(value)
	return s
}

// Valid return the value is valid. If true, it is not null value.
func (f *NullDecimal) Valid() bool {
	return f.f.Valid
}

// Float64Value return the value.
func (f *NullDecimal) DecimalValue() decimal.Decimal {
	return f.f.Decimal
}

// Reset set nil to the value.
func (f *NullDecimal) Reset() {
	f.f.Decimal = decimal.Decimal{}
	f.f.Valid = false
}

// Set set the value.
func (f *NullDecimal) Set(value decimal.Decimal) *NullDecimal {
	f.f.Valid = true
	f.f.Decimal = value
	return f
}

// Scan is a method for database/sql.
func (f *NullDecimal) Scan(value interface{}) error {
	return f.f.Scan(value)
}

// String return string indicated the value.
func (f NullDecimal) String() string {
	if !f.f.Valid {
		return ""
	}
	return f.f.Decimal.String()
}

// MarshalJSON encode the value to JSON.
func (f NullDecimal) MarshalJSON() ([]byte, error) {
	if !f.f.Valid {
		return []byte("null"), nil
	}
	str := f.f.Decimal.String()
	return []byte(str), nil
}

// UnmarshalJSON decode data to the value.
func (f *NullDecimal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.f.Valid = false
		return nil
	}
	f.f.Valid = true
	return f.f.Decimal.UnmarshalJSON(data)
}

// Value implement driver.Valuer.
func (f NullDecimal) Value() (driver.Value, error) {
	if !f.f.Valid {
		return nil, nil
	}
	return f.f.Decimal.Value()
}
