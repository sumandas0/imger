package effect

import (
	"fmt"
	"image/color"

	"github.com/RexterR/imger/errors"
	"github.com/disintegration/imaging"
)

var colorsList = []string{"black", "opaque", "transparent", "white"}
var colorMapping = map[string]color.Color{
	"black":       color.Black,
	"opaque":      color.Opaque,
	"transparent": color.Transparent,
	"white":       color.White,
}

var filtersList = []string{"lanczos", "catmull-rom", "mitchell-netravali", "bs-pline", "linear", "box", "nearest-neighbor"}
var filterMapping = map[string]imaging.ResampleFilter{
	"lanczos":            imaging.Lanczos,
	"catmull-rom":        imaging.CatmullRom,
	"mitchell-netravali": imaging.MitchellNetravali,
	"bs-pline":           imaging.BSpline,
	"linear":             imaging.Linear,
	"box":                imaging.Box,
	"nearest-neighbor":   imaging.NearestNeighbor,
}

func extractParameter(key string, params map[string]interface{}) (interface{}, error) {
	if value, ok := params[key]; ok {
		return value, nil
	}

	return nil, errors.EValidation(fmt.Sprintf("Parameter %s required", key), nil)
}

//check if integer array is of corect size helper

func integerArrayBinder(key string, array interface{}, expectedLen int) ([]int, error) {

	genericArray, ok := array.([]interface{})
	intArray := make([]int, 0, expectedLen)

	if !ok || len(genericArray) != expectedLen {
		return nil, errors.EValidation(fmt.Sprintf("Parameter %s needs to be an array of integers with length of 4", key), nil)
	}

	for _, number := range genericArray {
		n, ok := number.(float64)

		if !ok {
			return nil, errors.EValidation(fmt.Sprintf("Parameter %s needs to be array of integers", key), nil)
		}

		intArray = append(intArray, int(n))
	}

	return intArray, nil
}

func integerBinder(key string, params map[string]interface{}) (int, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return 0, err
	}

	valueInt, ok := value.(float64)

	if !ok {
		return 0, errors.EValidation(fmt.Sprintf("Parameter %s needs to be an integer", key), nil)
	}

	return int(valueInt), nil
}

// If in any field float isa nedeed then use this binder

func floatBinder(key string, params map[string]interface{}) (float64, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return 0, err
	}

	valueFloat, ok := value.(float64)

	if !ok {
		return 0, errors.EValidation(fmt.Sprintf("Parameter %s needs to be a float", key), nil)
	}

	return valueFloat, nil
}

// Checking if specified color binding is good

func colorBinder(key string, params map[string]interface{}) (color.Color, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return nil, err
	}

	colorKey, ok := value.(string)

	if !ok {
		return nil, errors.EValidation(fmt.Sprintf("Parameter %s needs to be a string", key), nil)
	}

	color, ok := colorMapping[colorKey]

	if !ok {
		return nil, errors.EValidation(fmt.Sprintf("Value %s not supported", colorKey), nil)
	}

	return color, nil
}

// Check if filter could be binded

func filterBinder(key string, params map[string]interface{}) (imaging.ResampleFilter, error) {
	value, err := extractParameter(key, params)

	if err != nil {
		return imaging.ResampleFilter{}, err
	}

	filterKey, ok := value.(string)

	if !ok {
		return imaging.ResampleFilter{}, errors.EValidation(fmt.Sprintf("Parameter %s needs to be a string", key), nil)
	}

	filter, ok := filterMapping[filterKey]

	if !ok {
		return imaging.ResampleFilter{}, errors.EValidation(fmt.Sprintf("Value %s not supported", filterKey), nil)
	}

	return filter, nil
}
