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

// RulesDtoAllOf struct for RulesDtoAllOf
type RulesDtoAllOf struct {
	// The rules.
	Items []RuleDto `json:"items"`
	// The id of the rule that is currently rerunning.
	RunningRuleId NullableString `json:"runningRuleId,omitempty"`
}

// NewRulesDtoAllOf instantiates a new RulesDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesDtoAllOf(items []RuleDto) *RulesDtoAllOf {
	this := RulesDtoAllOf{}
	this.Items = items
	return &this
}

// NewRulesDtoAllOfWithDefaults instantiates a new RulesDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesDtoAllOfWithDefaults() *RulesDtoAllOf {
	this := RulesDtoAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *RulesDtoAllOf) GetItems() []RuleDto {
	if o == nil {
		var ret []RuleDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RulesDtoAllOf) GetItemsOk() ([]RuleDto, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RulesDtoAllOf) SetItems(v []RuleDto) {
	o.Items = v
}

// GetRunningRuleId returns the RunningRuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RulesDtoAllOf) GetRunningRuleId() string {
	if o == nil || isNil(o.RunningRuleId.Get()) {
		var ret string
		return ret
	}
	return *o.RunningRuleId.Get()
}

// GetRunningRuleIdOk returns a tuple with the RunningRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RulesDtoAllOf) GetRunningRuleIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RunningRuleId.Get(), o.RunningRuleId.IsSet()
}

// HasRunningRuleId returns a boolean if a field has been set.
func (o *RulesDtoAllOf) HasRunningRuleId() bool {
	if o != nil && o.RunningRuleId.IsSet() {
		return true
	}

	return false
}

// SetRunningRuleId gets a reference to the given NullableString and assigns it to the RunningRuleId field.
func (o *RulesDtoAllOf) SetRunningRuleId(v string) {
	o.RunningRuleId.Set(&v)
}
// SetRunningRuleIdNil sets the value for RunningRuleId to be an explicit nil
func (o *RulesDtoAllOf) SetRunningRuleIdNil() {
	o.RunningRuleId.Set(nil)
}

// UnsetRunningRuleId ensures that no value is present for RunningRuleId, not even an explicit nil
func (o *RulesDtoAllOf) UnsetRunningRuleId() {
	o.RunningRuleId.Unset()
}

func (o RulesDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.RunningRuleId.IsSet() {
		toSerialize["runningRuleId"] = o.RunningRuleId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRulesDtoAllOf struct {
	value *RulesDtoAllOf
	isSet bool
}

func (v NullableRulesDtoAllOf) Get() *RulesDtoAllOf {
	return v.value
}

func (v *NullableRulesDtoAllOf) Set(val *RulesDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesDtoAllOf(val *RulesDtoAllOf) *NullableRulesDtoAllOf {
	return &NullableRulesDtoAllOf{value: val, isSet: true}
}

func (v NullableRulesDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

