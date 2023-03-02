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

// MoveAssetFolderDto struct for MoveAssetFolderDto
type MoveAssetFolderDto struct {
	// The parent folder id.
	ParentId *string `json:"parentId,omitempty"`
}

// NewMoveAssetFolderDto instantiates a new MoveAssetFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveAssetFolderDto() *MoveAssetFolderDto {
	this := MoveAssetFolderDto{}
	return &this
}

// NewMoveAssetFolderDtoWithDefaults instantiates a new MoveAssetFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveAssetFolderDtoWithDefaults() *MoveAssetFolderDto {
	this := MoveAssetFolderDto{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *MoveAssetFolderDto) GetParentId() string {
	if o == nil || isNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveAssetFolderDto) GetParentIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentId) {
    return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *MoveAssetFolderDto) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *MoveAssetFolderDto) SetParentId(v string) {
	o.ParentId = &v
}

func (o MoveAssetFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	return json.Marshal(toSerialize)
}

type NullableMoveAssetFolderDto struct {
	value *MoveAssetFolderDto
	isSet bool
}

func (v NullableMoveAssetFolderDto) Get() *MoveAssetFolderDto {
	return v.value
}

func (v *NullableMoveAssetFolderDto) Set(val *MoveAssetFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveAssetFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveAssetFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveAssetFolderDto(val *MoveAssetFolderDto) *NullableMoveAssetFolderDto {
	return &NullableMoveAssetFolderDto{value: val, isSet: true}
}

func (v NullableMoveAssetFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveAssetFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

