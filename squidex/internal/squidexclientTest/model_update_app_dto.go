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

// UpdateAppDto struct for UpdateAppDto
type UpdateAppDto struct {
	// The optional label of your app.
	Label NullableString `json:"label,omitempty"`
	// The optional description of your app.
	Description NullableString `json:"description,omitempty"`
}

// NewUpdateAppDto instantiates a new UpdateAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAppDto() *UpdateAppDto {
	this := UpdateAppDto{}
	return &this
}

// NewUpdateAppDtoWithDefaults instantiates a new UpdateAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppDtoWithDefaults() *UpdateAppDto {
	this := UpdateAppDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppDto) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppDto) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateAppDto) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *UpdateAppDto) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *UpdateAppDto) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *UpdateAppDto) UnsetLabel() {
	o.Label.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppDto) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAppDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateAppDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateAppDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateAppDto) UnsetDescription() {
	o.Description.Unset()
}

func (o UpdateAppDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAppDto struct {
	value *UpdateAppDto
	isSet bool
}

func (v NullableUpdateAppDto) Get() *UpdateAppDto {
	return v.value
}

func (v *NullableUpdateAppDto) Set(val *UpdateAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAppDto(val *UpdateAppDto) *NullableUpdateAppDto {
	return &NullableUpdateAppDto{value: val, isSet: true}
}

func (v NullableUpdateAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

