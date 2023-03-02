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

// ContentChangedRuleTriggerSchemaDto struct for ContentChangedRuleTriggerSchemaDto
type ContentChangedRuleTriggerSchemaDto struct {
	// The id of the schema.
	SchemaId *string `json:"schemaId,omitempty"`
	// Javascript condition when to trigger.
	Condition NullableString `json:"condition,omitempty"`
}

// NewContentChangedRuleTriggerSchemaDto instantiates a new ContentChangedRuleTriggerSchemaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentChangedRuleTriggerSchemaDto() *ContentChangedRuleTriggerSchemaDto {
	this := ContentChangedRuleTriggerSchemaDto{}
	return &this
}

// NewContentChangedRuleTriggerSchemaDtoWithDefaults instantiates a new ContentChangedRuleTriggerSchemaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentChangedRuleTriggerSchemaDtoWithDefaults() *ContentChangedRuleTriggerSchemaDto {
	this := ContentChangedRuleTriggerSchemaDto{}
	return &this
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *ContentChangedRuleTriggerSchemaDto) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentChangedRuleTriggerSchemaDto) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *ContentChangedRuleTriggerSchemaDto) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *ContentChangedRuleTriggerSchemaDto) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentChangedRuleTriggerSchemaDto) GetCondition() string {
	if o == nil || isNil(o.Condition.Get()) {
		var ret string
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentChangedRuleTriggerSchemaDto) GetConditionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *ContentChangedRuleTriggerSchemaDto) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableString and assigns it to the Condition field.
func (o *ContentChangedRuleTriggerSchemaDto) SetCondition(v string) {
	o.Condition.Set(&v)
}
// SetConditionNil sets the value for Condition to be an explicit nil
func (o *ContentChangedRuleTriggerSchemaDto) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *ContentChangedRuleTriggerSchemaDto) UnsetCondition() {
	o.Condition.Unset()
}

func (o ContentChangedRuleTriggerSchemaDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.Condition.IsSet() {
		toSerialize["condition"] = o.Condition.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContentChangedRuleTriggerSchemaDto struct {
	value *ContentChangedRuleTriggerSchemaDto
	isSet bool
}

func (v NullableContentChangedRuleTriggerSchemaDto) Get() *ContentChangedRuleTriggerSchemaDto {
	return v.value
}

func (v *NullableContentChangedRuleTriggerSchemaDto) Set(val *ContentChangedRuleTriggerSchemaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContentChangedRuleTriggerSchemaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContentChangedRuleTriggerSchemaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentChangedRuleTriggerSchemaDto(val *ContentChangedRuleTriggerSchemaDto) *NullableContentChangedRuleTriggerSchemaDto {
	return &NullableContentChangedRuleTriggerSchemaDto{value: val, isSet: true}
}

func (v NullableContentChangedRuleTriggerSchemaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentChangedRuleTriggerSchemaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

