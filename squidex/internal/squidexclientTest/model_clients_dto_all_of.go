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

// ClientsDtoAllOf struct for ClientsDtoAllOf
type ClientsDtoAllOf struct {
	// The clients.
	Items []ClientDto `json:"items"`
}

// NewClientsDtoAllOf instantiates a new ClientsDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientsDtoAllOf(items []ClientDto) *ClientsDtoAllOf {
	this := ClientsDtoAllOf{}
	this.Items = items
	return &this
}

// NewClientsDtoAllOfWithDefaults instantiates a new ClientsDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientsDtoAllOfWithDefaults() *ClientsDtoAllOf {
	this := ClientsDtoAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ClientsDtoAllOf) GetItems() []ClientDto {
	if o == nil {
		var ret []ClientDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ClientsDtoAllOf) GetItemsOk() ([]ClientDto, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ClientsDtoAllOf) SetItems(v []ClientDto) {
	o.Items = v
}

func (o ClientsDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableClientsDtoAllOf struct {
	value *ClientsDtoAllOf
	isSet bool
}

func (v NullableClientsDtoAllOf) Get() *ClientsDtoAllOf {
	return v.value
}

func (v *NullableClientsDtoAllOf) Set(val *ClientsDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClientsDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClientsDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientsDtoAllOf(val *ClientsDtoAllOf) *NullableClientsDtoAllOf {
	return &NullableClientsDtoAllOf{value: val, isSet: true}
}

func (v NullableClientsDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientsDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

