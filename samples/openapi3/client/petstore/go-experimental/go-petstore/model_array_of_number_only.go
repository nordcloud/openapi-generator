/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// ArrayOfNumberOnly struct for ArrayOfNumberOnly
type ArrayOfNumberOnly struct {
	ArrayNumber *[]float32 `json:"ArrayNumber,omitempty"`
}

// GetArrayNumber returns the ArrayNumber field value if set, zero value otherwise.
func (o *ArrayOfNumberOnly) GetArrayNumber() []float32 {
	if o == nil || o.ArrayNumber == nil {
		var ret []float32
		return ret
	}
	return *o.ArrayNumber
}

// GetArrayNumberOk returns a tuple with the ArrayNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ArrayOfNumberOnly) GetArrayNumberOk() ([]float32, bool) {
	if o == nil || o.ArrayNumber == nil {
		var ret []float32
		return ret, false
	}
	return *o.ArrayNumber, true
}

// HasArrayNumber returns a boolean if a field has been set.
func (o *ArrayOfNumberOnly) HasArrayNumber() bool {
	if o != nil && o.ArrayNumber != nil {
		return true
	}

	return false
}

// SetArrayNumber gets a reference to the given []float32 and assigns it to the ArrayNumber field.
func (o *ArrayOfNumberOnly) SetArrayNumber(v []float32) {
	o.ArrayNumber = &v
}

type NullableArrayOfNumberOnly struct {
	Value ArrayOfNumberOnly
	ExplicitNull bool
}

func (v NullableArrayOfNumberOnly) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableArrayOfNumberOnly) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
