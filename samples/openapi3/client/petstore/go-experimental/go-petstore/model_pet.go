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

// Pet struct for Pet
type Pet struct {
	Id *int64 `json:"id,omitempty"`
	Category *Category `json:"category,omitempty"`
	Name string `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags *[]Tag `json:"tags,omitempty"`
	// pet status in the store
	Status *string `json:"status,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Pet) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Pet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Pet) SetId(v int64) {
	o.Id = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Pet) GetCategory() Category {
	if o == nil || o.Category == nil {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetCategoryOk() (Category, bool) {
	if o == nil || o.Category == nil {
		var ret Category
		return ret, false
	}
	return *o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Pet) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *Pet) SetCategory(v Category) {
	o.Category = &v
}

// GetName returns the Name field value
func (o *Pet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Pet) SetName(v string) {
	o.Name = v
}

// GetPhotoUrls returns the PhotoUrls field value
func (o *Pet) GetPhotoUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PhotoUrls
}

// SetPhotoUrls sets field value
func (o *Pet) SetPhotoUrls(v []string) {
	o.PhotoUrls = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Pet) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetTagsOk() ([]Tag, bool) {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Pet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Pet) SetTags(v []Tag) {
	o.Tags = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Pet) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Pet) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Pet) SetStatus(v string) {
	o.Status = &v
}

type NullablePet struct {
	Value Pet
	ExplicitNull bool
}

func (v NullablePet) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
