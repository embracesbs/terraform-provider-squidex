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

// UpdateSettingDto struct for UpdateSettingDto
type UpdateSettingDto struct {
	// The value for the setting.
	Value interface{} `json:"value,omitempty"`
}

// NewUpdateSettingDto instantiates a new UpdateSettingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingDto() *UpdateSettingDto {
	this := UpdateSettingDto{}
	return &this
}

// NewUpdateSettingDtoWithDefaults instantiates a new UpdateSettingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingDtoWithDefaults() *UpdateSettingDto {
	this := UpdateSettingDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSettingDto) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSettingDto) GetValueOk() (*interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateSettingDto) HasValue() bool {
	if o != nil && isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *UpdateSettingDto) SetValue(v interface{}) {
	o.Value = v
}

func (o UpdateSettingDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSettingDto struct {
	value *UpdateSettingDto
	isSet bool
}

func (v NullableUpdateSettingDto) Get() *UpdateSettingDto {
	return v.value
}

func (v *NullableUpdateSettingDto) Set(val *UpdateSettingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingDto(val *UpdateSettingDto) *NullableUpdateSettingDto {
	return &NullableUpdateSettingDto{value: val, isSet: true}
}

func (v NullableUpdateSettingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

