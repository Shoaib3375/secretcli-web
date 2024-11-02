package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// validatePayload checks if JSON payload fields match the struct fields
func ValidatePayload(data []byte, target interface{}) error {
	// Decode JSON into a map to get the payload's keys
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return errors.New("invalid JSON format")
	}

	// Use reflection to get the struct's field names
	val := reflect.TypeOf(target).Elem()
	expectedFields := make(map[string]bool)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			expectedFields[jsonTag] = true
		}
	}

	// Check for any unexpected fields
	for key := range input {
		if _, found := expectedFields[key]; !found {
			return fmt.Errorf("unexpected field: %s", key)
		}
	}
	return nil
}
