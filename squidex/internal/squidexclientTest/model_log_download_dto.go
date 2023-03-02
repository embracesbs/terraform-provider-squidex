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

// LogDownloadDto struct for LogDownloadDto
type LogDownloadDto struct {
	// The url to download the log.
	DownloadUrl NullableString `json:"downloadUrl,omitempty"`
}

// NewLogDownloadDto instantiates a new LogDownloadDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDownloadDto() *LogDownloadDto {
	this := LogDownloadDto{}
	return &this
}

// NewLogDownloadDtoWithDefaults instantiates a new LogDownloadDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDownloadDtoWithDefaults() *LogDownloadDto {
	this := LogDownloadDto{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogDownloadDto) GetDownloadUrl() string {
	if o == nil || isNil(o.DownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDownloadDto) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *LogDownloadDto) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *LogDownloadDto) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *LogDownloadDto) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *LogDownloadDto) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

func (o LogDownloadDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadUrl.IsSet() {
		toSerialize["downloadUrl"] = o.DownloadUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLogDownloadDto struct {
	value *LogDownloadDto
	isSet bool
}

func (v NullableLogDownloadDto) Get() *LogDownloadDto {
	return v.value
}

func (v *NullableLogDownloadDto) Set(val *LogDownloadDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDownloadDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDownloadDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDownloadDto(val *LogDownloadDto) *NullableLogDownloadDto {
	return &NullableLogDownloadDto{value: val, isSet: true}
}

func (v NullableLogDownloadDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDownloadDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


