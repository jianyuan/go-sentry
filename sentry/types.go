package sentry

import (
	"encoding/json"
	"fmt"
)

// BoolOrStringSlice is a type that can be unmarshaled from either a bool or a
// string slice.
type BoolOrStringSlice struct {
	IsBool   bool
	BoolVal  bool
	SliceVal []string
}

var _ json.Unmarshaler = (*BoolOrStringSlice)(nil)

// UnmarshalJSON implements json.Unmarshaler.
func (bos *BoolOrStringSlice) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a bool
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		bos.IsBool = true
		bos.BoolVal = boolVal
		return nil
	}

	// Try to unmarshal as a string slice
	var sliceVal []string
	if err := json.Unmarshal(data, &sliceVal); err == nil {
		bos.IsBool = false
		bos.SliceVal = sliceVal
		return nil
	}

	// If neither worked, return an error
	return fmt.Errorf("unable to unmarshal as bool or string slice: %s", string(data))
}
