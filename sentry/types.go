package sentry

import (
	"encoding/json"
	"fmt"
)

// BoolOrStringSlice is a type that can be unmarshaled from either a bool or a
// string slice.
type BoolOrStringSlice struct {
	IsBool         bool
	IsStringSlice  bool
	BoolVal        bool
	StringSliceVal []string
}

var _ json.Unmarshaler = (*BoolOrStringSlice)(nil)
var _ json.Marshaler = (*BoolOrStringSlice)(nil)

// UnmarshalJSON implements json.Unmarshaler.
func (bos *BoolOrStringSlice) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a bool
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		bos.IsBool = true
		bos.IsStringSlice = false
		bos.BoolVal = boolVal
		return nil
	}

	// Try to unmarshal as a string slice
	var sliceVal []string
	if err := json.Unmarshal(data, &sliceVal); err == nil {
		bos.IsBool = false
		bos.IsStringSlice = true
		bos.StringSliceVal = sliceVal
		return nil
	}

	// If neither worked, return an error
	return fmt.Errorf("unable to unmarshal as bool or string slice: %s", string(data))
}

func (bos BoolOrStringSlice) MarshalJSON() ([]byte, error) {
	if bos.IsBool {
		return json.Marshal(bos.BoolVal)
	}
	return json.Marshal(bos.StringSliceVal)
}

// Int64OrString is a type that can be unmarshaled from either an int64 or a
// string.
type Int64OrString struct {
	IsInt64   bool
	IsString  bool
	Int64Val  int64
	StringVal string
}

var _ json.Unmarshaler = (*Int64OrString)(nil)
var _ json.Marshaler = (*Int64OrString)(nil)

func (ios *Int64OrString) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an int64
	var int64Val int64
	if err := json.Unmarshal(data, &int64Val); err == nil {
		ios.IsInt64 = true
		ios.IsString = false
		ios.Int64Val = int64Val
		return nil
	}

	// Try to unmarshal as a string
	var stringVal string
	if err := json.Unmarshal(data, &stringVal); err == nil {
		ios.IsInt64 = false
		ios.IsString = true
		ios.StringVal = stringVal
		return nil
	}

	// If neither worked, return an error
	return fmt.Errorf("unable to unmarshal as int64 or string: %s", string(data))
}

func (ios Int64OrString) MarshalJSON() ([]byte, error) {
	if ios.IsInt64 {
		return json.Marshal(ios.Int64Val)
	}
	return json.Marshal(ios.StringVal)
}
