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

// UISettingsDto struct for UISettingsDto
type UISettingsDto struct {
	// True when the user can create apps.
	CanCreateApps *bool `json:"canCreateApps,omitempty"`
}

// NewUISettingsDto instantiates a new UISettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUISettingsDto() *UISettingsDto {
	this := UISettingsDto{}
	return &this
}

// NewUISettingsDtoWithDefaults instantiates a new UISettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUISettingsDtoWithDefaults() *UISettingsDto {
	this := UISettingsDto{}
	return &this
}

// GetCanCreateApps returns the CanCreateApps field value if set, zero value otherwise.
func (o *UISettingsDto) GetCanCreateApps() bool {
	if o == nil || isNil(o.CanCreateApps) {
		var ret bool
		return ret
	}
	return *o.CanCreateApps
}

// GetCanCreateAppsOk returns a tuple with the CanCreateApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UISettingsDto) GetCanCreateAppsOk() (*bool, bool) {
	if o == nil || isNil(o.CanCreateApps) {
    return nil, false
	}
	return o.CanCreateApps, true
}

// HasCanCreateApps returns a boolean if a field has been set.
func (o *UISettingsDto) HasCanCreateApps() bool {
	if o != nil && !isNil(o.CanCreateApps) {
		return true
	}

	return false
}

// SetCanCreateApps gets a reference to the given bool and assigns it to the CanCreateApps field.
func (o *UISettingsDto) SetCanCreateApps(v bool) {
	o.CanCreateApps = &v
}

func (o UISettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CanCreateApps) {
		toSerialize["canCreateApps"] = o.CanCreateApps
	}
	return json.Marshal(toSerialize)
}

type NullableUISettingsDto struct {
	value *UISettingsDto
	isSet bool
}

func (v NullableUISettingsDto) Get() *UISettingsDto {
	return v.value
}

func (v *NullableUISettingsDto) Set(val *UISettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUISettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUISettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUISettingsDto(val *UISettingsDto) *NullableUISettingsDto {
	return &NullableUISettingsDto{value: val, isSet: true}
}

func (v NullableUISettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUISettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

