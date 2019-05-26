package utils

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Parse map to struct for graphql
func Parse(data map[string]interface{}, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:   nil,
		Result:     result,
		TagName:    "graphql",
		DecodeHook: stringToTimeHookFunc(time.RFC3339),
	}

	dec, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return dec.Decode(data)
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func stringToTimeHookFunc(layout string) mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// Check source is string or not
		if f.Kind() != reflect.String {
			return data, nil
		}
		// Check destination is Time or not
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		// Convert it by parsing
		return time.Parse(layout, data.(string))
	}
}

// CalculateOffsetForPage map to struct for graphql
func CalculateOffsetForPage(page, pageSize int64) int64 {
	offset := (page - 1) * pageSize
	return offset
}
