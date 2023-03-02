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

// UpdateRoleDto struct for UpdateRoleDto
type UpdateRoleDto struct {
	// Associated list of permissions.
	Permissions []string `json:"permissions"`
	// Associated list of UI properties.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewUpdateRoleDto instantiates a new UpdateRoleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleDto(permissions []string) *UpdateRoleDto {
	this := UpdateRoleDto{}
	this.Permissions = permissions
	return &this
}

// NewUpdateRoleDtoWithDefaults instantiates a new UpdateRoleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleDtoWithDefaults() *UpdateRoleDto {
	this := UpdateRoleDto{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *UpdateRoleDto) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDto) GetPermissionsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UpdateRoleDto) SetPermissions(v []string) {
	o.Permissions = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UpdateRoleDto) GetProperties() map[string]interface{} {
	if o == nil || isNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleDto) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Properties) {
    return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UpdateRoleDto) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *UpdateRoleDto) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o UpdateRoleDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRoleDto struct {
	value *UpdateRoleDto
	isSet bool
}

func (v NullableUpdateRoleDto) Get() *UpdateRoleDto {
	return v.value
}

func (v *NullableUpdateRoleDto) Set(val *UpdateRoleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleDto(val *UpdateRoleDto) *NullableUpdateRoleDto {
	return &NullableUpdateRoleDto{value: val, isSet: true}
}

func (v NullableUpdateRoleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

