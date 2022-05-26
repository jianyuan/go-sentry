package sentry

import (
	"encoding/json"
	"fmt"
)

// APIError represents a Sentry API Error response.
// Should look like:
//
// 	type apiError struct {
// 		Detail string `json:"detail"`
// 	}
//
type APIError struct {
	f interface{} // unknown
}

func (e *APIError) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &e.f); err != nil {
		e.f = string(b)
	}
	return nil
}

func (e *APIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.f)
}

func (e APIError) Error() string {
	switch v := e.f.(type) {
	case map[string]interface{}:
		if len(v) == 1 {
			if detail, ok := v["detail"].(string); ok {
				return fmt.Sprintf("sentry: %s", detail)
			}
		}
		return fmt.Sprintf("sentry: %v", v)
	default:
		return fmt.Sprintf("sentry: %v", v)
	}
}

// Empty returns true if empty.
func (e APIError) Empty() bool {
	return e.f == nil
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if !apiError.Empty() {
		return apiError
	}
	return nil
}
