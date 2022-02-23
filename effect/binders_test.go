package effect

import (
	"testing"

	"github.com/RexterR/imger/errors"
)

func TestIntegerBinder(t *testing.T) {
	params := map[string]interface{}{"key": 1.0}
	_, err := integerBinder("key", params)

	if err != nil {
		t.Error("Should return an int", err)
	}
}
func TestIntegerBinderWrongInteger(t *testing.T) {
	params := map[string]interface{}{"key": "sumandas"}
	_, err := integerBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}

func TestFloatBinder(t *testing.T) {
	params := map[string]interface{}{"key": 10.5}
	_, err := floatBinder("key", params)

	if err != nil {
		t.Error("Should return a float", err)
	}
}

func TestFloatBinderWrongFloat(t *testing.T) {
	params := map[string]interface{}{"key": "Imger"}
	_, err := floatBinder("key", params)

	if !errors.Is(errors.Validation, err) {
		t.Error("Should be a validation error", err)
	}
}
