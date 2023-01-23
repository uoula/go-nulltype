package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNullDecimalStringer(t *testing.T) {
	var f NullDecimal

	want := ""
	got := fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3.14"
	f.Set(decimal.NewFromFloat(3.14))
	got = fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3.15"
	f = NullDecimalOf(decimal.NewFromFloat(3.15))
	got = fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = ""
	f.Reset()
	got = fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullDecimalMarshalJSON(t *testing.T) {
	var f NullDecimal

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(f)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	f.Set(decimal.NewFromFloat(3.14))
	err = json.NewEncoder(&buf).Encode(f)
	if err != nil {
		t.Fatal(err)
	}

	want = "3.14"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullDecimalUnmarshalJSON(t *testing.T) {
	var f NullDecimal

	err := json.NewDecoder(strings.NewReader("null")).Decode(&f)
	if err != nil {
		t.Fatal(err)
	}

	if f.Valid() {
		t.Fatalf("must be null but got %v", f)
	}

	f.Set(decimal.NewFromFloat(3.14))

	err = json.NewDecoder(strings.NewReader(`3.14`)).Decode(&f)
	if err != nil {
		t.Fatal(err)
	}

	if !f.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := decimal.NewFromFloat(3.14)
	got := f.DecimalValue()
	if !got.Equal(want) {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&f)
	if err == nil {
		t.Fatal("should be fail")
	}
}

func TestNullDecimalValueConverter(t *testing.T) {
	var f NullDecimal

	err := f.Scan("3.14")
	if err != nil {
		t.Fatal(err)
	}

	if !f.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := decimal.NewFromFloat(3.14)
	got := f.DecimalValue()
	if !got.Equal(want) {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := f.Value()
	if err != nil {
		t.Fatal(err)
	}
	wantv, err := want.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != wantv {
		t.Fatalf("want %v, but %v:", want, got)
	}

	f.Reset()

	gotv, err = f.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != nil {
		t.Fatalf("must be null but got %v", gotv)
	}
}
