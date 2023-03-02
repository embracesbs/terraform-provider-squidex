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

// AssetsDtoAllOf struct for AssetsDtoAllOf
type AssetsDtoAllOf struct {
	// The total number of assets.
	Total *int64 `json:"total,omitempty"`
	// The assets.
	Items []AssetDto `json:"items"`
}

// NewAssetsDtoAllOf instantiates a new AssetsDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsDtoAllOf(items []AssetDto) *AssetsDtoAllOf {
	this := AssetsDtoAllOf{}
	this.Items = items
	return &this
}

// NewAssetsDtoAllOfWithDefaults instantiates a new AssetsDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsDtoAllOfWithDefaults() *AssetsDtoAllOf {
	this := AssetsDtoAllOf{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AssetsDtoAllOf) GetTotal() int64 {
	if o == nil || isNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsDtoAllOf) GetTotalOk() (*int64, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AssetsDtoAllOf) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *AssetsDtoAllOf) SetTotal(v int64) {
	o.Total = &v
}

// GetItems returns the Items field value
func (o *AssetsDtoAllOf) GetItems() []AssetDto {
	if o == nil {
		var ret []AssetDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AssetsDtoAllOf) GetItemsOk() ([]AssetDto, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AssetsDtoAllOf) SetItems(v []AssetDto) {
	o.Items = v
}

func (o AssetsDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAssetsDtoAllOf struct {
	value *AssetsDtoAllOf
	isSet bool
}

func (v NullableAssetsDtoAllOf) Get() *AssetsDtoAllOf {
	return v.value
}

func (v *NullableAssetsDtoAllOf) Set(val *AssetsDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsDtoAllOf(val *AssetsDtoAllOf) *NullableAssetsDtoAllOf {
	return &NullableAssetsDtoAllOf{value: val, isSet: true}
}

func (v NullableAssetsDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

