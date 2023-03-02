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

// PlanChangedDto struct for PlanChangedDto
type PlanChangedDto struct {
	// Optional redirect uri.
	RedirectUri NullableString `json:"redirectUri,omitempty"`
}

// NewPlanChangedDto instantiates a new PlanChangedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanChangedDto() *PlanChangedDto {
	this := PlanChangedDto{}
	return &this
}

// NewPlanChangedDtoWithDefaults instantiates a new PlanChangedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanChangedDtoWithDefaults() *PlanChangedDto {
	this := PlanChangedDto{}
	return &this
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanChangedDto) GetRedirectUri() string {
	if o == nil || isNil(o.RedirectUri.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUri.Get()
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanChangedDto) GetRedirectUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RedirectUri.Get(), o.RedirectUri.IsSet()
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *PlanChangedDto) HasRedirectUri() bool {
	if o != nil && o.RedirectUri.IsSet() {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given NullableString and assigns it to the RedirectUri field.
func (o *PlanChangedDto) SetRedirectUri(v string) {
	o.RedirectUri.Set(&v)
}
// SetRedirectUriNil sets the value for RedirectUri to be an explicit nil
func (o *PlanChangedDto) SetRedirectUriNil() {
	o.RedirectUri.Set(nil)
}

// UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
func (o *PlanChangedDto) UnsetRedirectUri() {
	o.RedirectUri.Unset()
}

func (o PlanChangedDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUri.IsSet() {
		toSerialize["redirectUri"] = o.RedirectUri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePlanChangedDto struct {
	value *PlanChangedDto
	isSet bool
}

func (v NullablePlanChangedDto) Get() *PlanChangedDto {
	return v.value
}

func (v *NullablePlanChangedDto) Set(val *PlanChangedDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanChangedDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanChangedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanChangedDto(val *PlanChangedDto) *NullablePlanChangedDto {
	return &NullablePlanChangedDto{value: val, isSet: true}
}

func (v NullablePlanChangedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanChangedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

