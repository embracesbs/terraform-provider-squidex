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

// ConfigureUIFieldsDto struct for ConfigureUIFieldsDto
type ConfigureUIFieldsDto struct {
	FieldsInLists []string `json:"fieldsInLists,omitempty"`
	FieldsInReferences []string `json:"fieldsInReferences,omitempty"`
}

// NewConfigureUIFieldsDto instantiates a new ConfigureUIFieldsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigureUIFieldsDto() *ConfigureUIFieldsDto {
	this := ConfigureUIFieldsDto{}
	return &this
}

// NewConfigureUIFieldsDtoWithDefaults instantiates a new ConfigureUIFieldsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigureUIFieldsDtoWithDefaults() *ConfigureUIFieldsDto {
	this := ConfigureUIFieldsDto{}
	return &this
}

// GetFieldsInLists returns the FieldsInLists field value if set, zero value otherwise.
func (o *ConfigureUIFieldsDto) GetFieldsInLists() []string {
	if o == nil || isNil(o.FieldsInLists) {
		var ret []string
		return ret
	}
	return o.FieldsInLists
}

// GetFieldsInListsOk returns a tuple with the FieldsInLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureUIFieldsDto) GetFieldsInListsOk() ([]string, bool) {
	if o == nil || isNil(o.FieldsInLists) {
    return nil, false
	}
	return o.FieldsInLists, true
}

// HasFieldsInLists returns a boolean if a field has been set.
func (o *ConfigureUIFieldsDto) HasFieldsInLists() bool {
	if o != nil && !isNil(o.FieldsInLists) {
		return true
	}

	return false
}

// SetFieldsInLists gets a reference to the given []string and assigns it to the FieldsInLists field.
func (o *ConfigureUIFieldsDto) SetFieldsInLists(v []string) {
	o.FieldsInLists = v
}

// GetFieldsInReferences returns the FieldsInReferences field value if set, zero value otherwise.
func (o *ConfigureUIFieldsDto) GetFieldsInReferences() []string {
	if o == nil || isNil(o.FieldsInReferences) {
		var ret []string
		return ret
	}
	return o.FieldsInReferences
}

// GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureUIFieldsDto) GetFieldsInReferencesOk() ([]string, bool) {
	if o == nil || isNil(o.FieldsInReferences) {
    return nil, false
	}
	return o.FieldsInReferences, true
}

// HasFieldsInReferences returns a boolean if a field has been set.
func (o *ConfigureUIFieldsDto) HasFieldsInReferences() bool {
	if o != nil && !isNil(o.FieldsInReferences) {
		return true
	}

	return false
}

// SetFieldsInReferences gets a reference to the given []string and assigns it to the FieldsInReferences field.
func (o *ConfigureUIFieldsDto) SetFieldsInReferences(v []string) {
	o.FieldsInReferences = v
}

func (o ConfigureUIFieldsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FieldsInLists) {
		toSerialize["fieldsInLists"] = o.FieldsInLists
	}
	if !isNil(o.FieldsInReferences) {
		toSerialize["fieldsInReferences"] = o.FieldsInReferences
	}
	return json.Marshal(toSerialize)
}

type NullableConfigureUIFieldsDto struct {
	value *ConfigureUIFieldsDto
	isSet bool
}

func (v NullableConfigureUIFieldsDto) Get() *ConfigureUIFieldsDto {
	return v.value
}

func (v *NullableConfigureUIFieldsDto) Set(val *ConfigureUIFieldsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigureUIFieldsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigureUIFieldsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigureUIFieldsDto(val *ConfigureUIFieldsDto) *NullableConfigureUIFieldsDto {
	return &NullableConfigureUIFieldsDto{value: val, isSet: true}
}

func (v NullableConfigureUIFieldsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigureUIFieldsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


