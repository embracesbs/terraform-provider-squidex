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

// AssetChangedRuleTriggerDto struct for AssetChangedRuleTriggerDto
type AssetChangedRuleTriggerDto struct {
	RuleTriggerDto
	// Javascript condition when to trigger.
	Condition NullableString `json:"condition,omitempty"`
}

// NewAssetChangedRuleTriggerDto instantiates a new AssetChangedRuleTriggerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetChangedRuleTriggerDto(triggerType string) *AssetChangedRuleTriggerDto {
	this := AssetChangedRuleTriggerDto{}
	this.TriggerType = triggerType
	return &this
}

// NewAssetChangedRuleTriggerDtoWithDefaults instantiates a new AssetChangedRuleTriggerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetChangedRuleTriggerDtoWithDefaults() *AssetChangedRuleTriggerDto {
	this := AssetChangedRuleTriggerDto{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetChangedRuleTriggerDto) GetCondition() string {
	if o == nil || isNil(o.Condition.Get()) {
		var ret string
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetChangedRuleTriggerDto) GetConditionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *AssetChangedRuleTriggerDto) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableString and assigns it to the Condition field.
func (o *AssetChangedRuleTriggerDto) SetCondition(v string) {
	o.Condition.Set(&v)
}
// SetConditionNil sets the value for Condition to be an explicit nil
func (o *AssetChangedRuleTriggerDto) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *AssetChangedRuleTriggerDto) UnsetCondition() {
	o.Condition.Unset()
}

func (o AssetChangedRuleTriggerDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRuleTriggerDto, errRuleTriggerDto := json.Marshal(o.RuleTriggerDto)
	if errRuleTriggerDto != nil {
		return []byte{}, errRuleTriggerDto
	}
	errRuleTriggerDto = json.Unmarshal([]byte(serializedRuleTriggerDto), &toSerialize)
	if errRuleTriggerDto != nil {
		return []byte{}, errRuleTriggerDto
	}
	if o.Condition.IsSet() {
		toSerialize["condition"] = o.Condition.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetChangedRuleTriggerDto struct {
	value *AssetChangedRuleTriggerDto
	isSet bool
}

func (v NullableAssetChangedRuleTriggerDto) Get() *AssetChangedRuleTriggerDto {
	return v.value
}

func (v *NullableAssetChangedRuleTriggerDto) Set(val *AssetChangedRuleTriggerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetChangedRuleTriggerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetChangedRuleTriggerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetChangedRuleTriggerDto(val *AssetChangedRuleTriggerDto) *NullableAssetChangedRuleTriggerDto {
	return &NullableAssetChangedRuleTriggerDto{value: val, isSet: true}
}

func (v NullableAssetChangedRuleTriggerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetChangedRuleTriggerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


