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

// UserProperty struct for UserProperty
type UserProperty struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewUserProperty instantiates a new UserProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProperty(name string, value string) *UserProperty {
	this := UserProperty{}
	this.Name = name
	this.Value = value
	return &this
}

// NewUserPropertyWithDefaults instantiates a new UserProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPropertyWithDefaults() *UserProperty {
	this := UserProperty{}
	return &this
}

// GetName returns the Name field value
func (o *UserProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserProperty) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserProperty) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *UserProperty) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserProperty) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserProperty) SetValue(v string) {
	o.Value = v
}

func (o UserProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUserProperty struct {
	value *UserProperty
	isSet bool
}

func (v NullableUserProperty) Get() *UserProperty {
	return v.value
}

func (v *NullableUserProperty) Set(val *UserProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProperty(val *UserProperty) *NullableUserProperty {
	return &NullableUserProperty{value: val, isSet: true}
}

func (v NullableUserProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


