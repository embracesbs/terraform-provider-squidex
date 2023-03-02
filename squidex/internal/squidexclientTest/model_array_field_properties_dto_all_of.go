/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
)

// ArrayFieldPropertiesDtoAllOf struct for ArrayFieldPropertiesDtoAllOf
type ArrayFieldPropertiesDtoAllOf struct {
	// The minimum allowed items for the field value.
	MinItems NullableInt32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems NullableInt32 `json:"maxItems,omitempty"`
	// The fields that must be unique.
	UniqueFields []string `json:"uniqueFields,omitempty"`
}

// NewArrayFieldPropertiesDtoAllOf instantiates a new ArrayFieldPropertiesDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayFieldPropertiesDtoAllOf() *ArrayFieldPropertiesDtoAllOf {
	this := ArrayFieldPropertiesDtoAllOf{}
	return &this
}

// NewArrayFieldPropertiesDtoAllOfWithDefaults instantiates a new ArrayFieldPropertiesDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayFieldPropertiesDtoAllOfWithDefaults() *ArrayFieldPropertiesDtoAllOf {
	this := ArrayFieldPropertiesDtoAllOf{}
	return &this
}

// GetMinItems returns the MinItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArrayFieldPropertiesDtoAllOf) GetMinItems() int32 {
	if o == nil || isNil(o.MinItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MinItems.Get()
}

// GetMinItemsOk returns a tuple with the MinItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArrayFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinItems.Get(), o.MinItems.IsSet()
}

// HasMinItems returns a boolean if a field has been set.
func (o *ArrayFieldPropertiesDtoAllOf) HasMinItems() bool {
	if o != nil && o.MinItems.IsSet() {
		return true
	}

	return false
}

// SetMinItems gets a reference to the given NullableInt32 and assigns it to the MinItems field.
func (o *ArrayFieldPropertiesDtoAllOf) SetMinItems(v int32) {
	o.MinItems.Set(&v)
}
// SetMinItemsNil sets the value for MinItems to be an explicit nil
func (o *ArrayFieldPropertiesDtoAllOf) SetMinItemsNil() {
	o.MinItems.Set(nil)
}

// UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
func (o *ArrayFieldPropertiesDtoAllOf) UnsetMinItems() {
	o.MinItems.Unset()
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArrayFieldPropertiesDtoAllOf) GetMaxItems() int32 {
	if o == nil || isNil(o.MaxItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItems.Get()
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArrayFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItems.Get(), o.MaxItems.IsSet()
}

// HasMaxItems returns a boolean if a field has been set.
func (o *ArrayFieldPropertiesDtoAllOf) HasMaxItems() bool {
	if o != nil && o.MaxItems.IsSet() {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given NullableInt32 and assigns it to the MaxItems field.
func (o *ArrayFieldPropertiesDtoAllOf) SetMaxItems(v int32) {
	o.MaxItems.Set(&v)
}
// SetMaxItemsNil sets the value for MaxItems to be an explicit nil
func (o *ArrayFieldPropertiesDtoAllOf) SetMaxItemsNil() {
	o.MaxItems.Set(nil)
}

// UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
func (o *ArrayFieldPropertiesDtoAllOf) UnsetMaxItems() {
	o.MaxItems.Unset()
}

// GetUniqueFields returns the UniqueFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArrayFieldPropertiesDtoAllOf) GetUniqueFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UniqueFields
}

// GetUniqueFieldsOk returns a tuple with the UniqueFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArrayFieldPropertiesDtoAllOf) GetUniqueFieldsOk() ([]string, bool) {
	if o == nil || isNil(o.UniqueFields) {
    return nil, false
	}
	return o.UniqueFields, true
}

// HasUniqueFields returns a boolean if a field has been set.
func (o *ArrayFieldPropertiesDtoAllOf) HasUniqueFields() bool {
	if o != nil && isNil(o.UniqueFields) {
		return true
	}

	return false
}

// SetUniqueFields gets a reference to the given []string and assigns it to the UniqueFields field.
func (o *ArrayFieldPropertiesDtoAllOf) SetUniqueFields(v []string) {
	o.UniqueFields = v
}

func (o ArrayFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinItems.IsSet() {
		toSerialize["minItems"] = o.MinItems.Get()
	}
	if o.MaxItems.IsSet() {
		toSerialize["maxItems"] = o.MaxItems.Get()
	}
	if o.UniqueFields != nil {
		toSerialize["uniqueFields"] = o.UniqueFields
	}
	return json.Marshal(toSerialize)
}

type NullableArrayFieldPropertiesDtoAllOf struct {
	value *ArrayFieldPropertiesDtoAllOf
	isSet bool
}

func (v NullableArrayFieldPropertiesDtoAllOf) Get() *ArrayFieldPropertiesDtoAllOf {
	return v.value
}

func (v *NullableArrayFieldPropertiesDtoAllOf) Set(val *ArrayFieldPropertiesDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayFieldPropertiesDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayFieldPropertiesDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayFieldPropertiesDtoAllOf(val *ArrayFieldPropertiesDtoAllOf) *NullableArrayFieldPropertiesDtoAllOf {
	return &NullableArrayFieldPropertiesDtoAllOf{value: val, isSet: true}
}

func (v NullableArrayFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayFieldPropertiesDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

