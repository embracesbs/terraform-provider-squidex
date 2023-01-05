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

// AssetFolderDto struct for AssetFolderDto
type AssetFolderDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the asset.
	Id *string `json:"id,omitempty"`
	// The id of the parent folder. Empty for files without parent.
	ParentId *string `json:"parentId,omitempty"`
	// The folder name.
	FolderName string `json:"folderName"`
	// The version of the asset folder.
	Version *int64 `json:"version,omitempty"`
}

// NewAssetFolderDto instantiates a new AssetFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetFolderDto(links map[string]ResourceLink, folderName string) *AssetFolderDto {
	this := AssetFolderDto{}
	this.Links = links
	this.FolderName = folderName
	return &this
}

// NewAssetFolderDtoWithDefaults instantiates a new AssetFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetFolderDtoWithDefaults() *AssetFolderDto {
	this := AssetFolderDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *AssetFolderDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AssetFolderDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AssetFolderDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetFolderDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFolderDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetFolderDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetFolderDto) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AssetFolderDto) GetParentId() string {
	if o == nil || isNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFolderDto) GetParentIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentId) {
    return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AssetFolderDto) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *AssetFolderDto) SetParentId(v string) {
	o.ParentId = &v
}

// GetFolderName returns the FolderName field value
func (o *AssetFolderDto) GetFolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *AssetFolderDto) GetFolderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *AssetFolderDto) SetFolderName(v string) {
	o.FolderName = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AssetFolderDto) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFolderDto) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AssetFolderDto) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *AssetFolderDto) SetVersion(v int64) {
	o.Version = &v
}

func (o AssetFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["folderName"] = o.FolderName
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableAssetFolderDto struct {
	value *AssetFolderDto
	isSet bool
}

func (v NullableAssetFolderDto) Get() *AssetFolderDto {
	return v.value
}

func (v *NullableAssetFolderDto) Set(val *AssetFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetFolderDto(val *AssetFolderDto) *NullableAssetFolderDto {
	return &NullableAssetFolderDto{value: val, isSet: true}
}

func (v NullableAssetFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

