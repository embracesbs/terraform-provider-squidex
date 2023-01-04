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

// UsageRuleTriggerDtoAllOf struct for UsageRuleTriggerDtoAllOf
type UsageRuleTriggerDtoAllOf struct {
	// The number of monthly api calls.
	Limit *int32 `json:"limit,omitempty"`
	// The number of days to check or null for the current month.
	NumDays NullableInt32 `json:"numDays,omitempty"`
}

// NewUsageRuleTriggerDtoAllOf instantiates a new UsageRuleTriggerDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageRuleTriggerDtoAllOf() *UsageRuleTriggerDtoAllOf {
	this := UsageRuleTriggerDtoAllOf{}
	return &this
}

// NewUsageRuleTriggerDtoAllOfWithDefaults instantiates a new UsageRuleTriggerDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageRuleTriggerDtoAllOfWithDefaults() *UsageRuleTriggerDtoAllOf {
	this := UsageRuleTriggerDtoAllOf{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UsageRuleTriggerDtoAllOf) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRuleTriggerDtoAllOf) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UsageRuleTriggerDtoAllOf) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *UsageRuleTriggerDtoAllOf) SetLimit(v int32) {
	o.Limit = &v
}

// GetNumDays returns the NumDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageRuleTriggerDtoAllOf) GetNumDays() int32 {
	if o == nil || isNil(o.NumDays.Get()) {
		var ret int32
		return ret
	}
	return *o.NumDays.Get()
}

// GetNumDaysOk returns a tuple with the NumDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageRuleTriggerDtoAllOf) GetNumDaysOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumDays.Get(), o.NumDays.IsSet()
}

// HasNumDays returns a boolean if a field has been set.
func (o *UsageRuleTriggerDtoAllOf) HasNumDays() bool {
	if o != nil && o.NumDays.IsSet() {
		return true
	}

	return false
}

// SetNumDays gets a reference to the given NullableInt32 and assigns it to the NumDays field.
func (o *UsageRuleTriggerDtoAllOf) SetNumDays(v int32) {
	o.NumDays.Set(&v)
}
// SetNumDaysNil sets the value for NumDays to be an explicit nil
func (o *UsageRuleTriggerDtoAllOf) SetNumDaysNil() {
	o.NumDays.Set(nil)
}

// UnsetNumDays ensures that no value is present for NumDays, not even an explicit nil
func (o *UsageRuleTriggerDtoAllOf) UnsetNumDays() {
	o.NumDays.Unset()
}

func (o UsageRuleTriggerDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.NumDays.IsSet() {
		toSerialize["numDays"] = o.NumDays.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUsageRuleTriggerDtoAllOf struct {
	value *UsageRuleTriggerDtoAllOf
	isSet bool
}

func (v NullableUsageRuleTriggerDtoAllOf) Get() *UsageRuleTriggerDtoAllOf {
	return v.value
}

func (v *NullableUsageRuleTriggerDtoAllOf) Set(val *UsageRuleTriggerDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageRuleTriggerDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageRuleTriggerDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageRuleTriggerDtoAllOf(val *UsageRuleTriggerDtoAllOf) *NullableUsageRuleTriggerDtoAllOf {
	return &NullableUsageRuleTriggerDtoAllOf{value: val, isSet: true}
}

func (v NullableUsageRuleTriggerDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageRuleTriggerDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


