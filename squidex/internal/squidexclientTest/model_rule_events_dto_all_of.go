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

// RuleEventsDtoAllOf struct for RuleEventsDtoAllOf
type RuleEventsDtoAllOf struct {
	// The rule events.
	Items []RuleEventDto `json:"items"`
	// The total number of rule events.
	Total *int64 `json:"total,omitempty"`
}

// NewRuleEventsDtoAllOf instantiates a new RuleEventsDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleEventsDtoAllOf(items []RuleEventDto) *RuleEventsDtoAllOf {
	this := RuleEventsDtoAllOf{}
	this.Items = items
	return &this
}

// NewRuleEventsDtoAllOfWithDefaults instantiates a new RuleEventsDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleEventsDtoAllOfWithDefaults() *RuleEventsDtoAllOf {
	this := RuleEventsDtoAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *RuleEventsDtoAllOf) GetItems() []RuleEventDto {
	if o == nil {
		var ret []RuleEventDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RuleEventsDtoAllOf) GetItemsOk() ([]RuleEventDto, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RuleEventsDtoAllOf) SetItems(v []RuleEventDto) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RuleEventsDtoAllOf) GetTotal() int64 {
	if o == nil || isNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventsDtoAllOf) GetTotalOk() (*int64, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RuleEventsDtoAllOf) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RuleEventsDtoAllOf) SetTotal(v int64) {
	o.Total = &v
}

func (o RuleEventsDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableRuleEventsDtoAllOf struct {
	value *RuleEventsDtoAllOf
	isSet bool
}

func (v NullableRuleEventsDtoAllOf) Get() *RuleEventsDtoAllOf {
	return v.value
}

func (v *NullableRuleEventsDtoAllOf) Set(val *RuleEventsDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleEventsDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleEventsDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleEventsDtoAllOf(val *RuleEventsDtoAllOf) *NullableRuleEventsDtoAllOf {
	return &NullableRuleEventsDtoAllOf{value: val, isSet: true}
}

func (v NullableRuleEventsDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleEventsDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

