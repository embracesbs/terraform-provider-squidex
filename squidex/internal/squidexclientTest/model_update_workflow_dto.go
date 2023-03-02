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

// UpdateWorkflowDto struct for UpdateWorkflowDto
type UpdateWorkflowDto struct {
	// The name of the workflow.
	Name NullableString `json:"name,omitempty"`
	// The workflow steps.
	Steps map[string]WorkflowStepDto `json:"steps"`
	// The schema ids.
	SchemaIds []string `json:"schemaIds,omitempty"`
	// The initial step.
	Initial string `json:"initial"`
}

// NewUpdateWorkflowDto instantiates a new UpdateWorkflowDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkflowDto(steps map[string]WorkflowStepDto, initial string) *UpdateWorkflowDto {
	this := UpdateWorkflowDto{}
	this.Steps = steps
	this.Initial = initial
	return &this
}

// NewUpdateWorkflowDtoWithDefaults instantiates a new UpdateWorkflowDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkflowDtoWithDefaults() *UpdateWorkflowDto {
	this := UpdateWorkflowDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateWorkflowDto) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateWorkflowDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWorkflowDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateWorkflowDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateWorkflowDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateWorkflowDto) UnsetName() {
	o.Name.Unset()
}

// GetSteps returns the Steps field value
func (o *UpdateWorkflowDto) GetSteps() map[string]WorkflowStepDto {
	if o == nil {
		var ret map[string]WorkflowStepDto
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *UpdateWorkflowDto) GetStepsOk() (*map[string]WorkflowStepDto, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value
func (o *UpdateWorkflowDto) SetSteps(v map[string]WorkflowStepDto) {
	o.Steps = v
}

// GetSchemaIds returns the SchemaIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateWorkflowDto) GetSchemaIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SchemaIds
}

// GetSchemaIdsOk returns a tuple with the SchemaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateWorkflowDto) GetSchemaIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SchemaIds) {
    return nil, false
	}
	return o.SchemaIds, true
}

// HasSchemaIds returns a boolean if a field has been set.
func (o *UpdateWorkflowDto) HasSchemaIds() bool {
	if o != nil && isNil(o.SchemaIds) {
		return true
	}

	return false
}

// SetSchemaIds gets a reference to the given []string and assigns it to the SchemaIds field.
func (o *UpdateWorkflowDto) SetSchemaIds(v []string) {
	o.SchemaIds = v
}

// GetInitial returns the Initial field value
func (o *UpdateWorkflowDto) GetInitial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Initial
}

// GetInitialOk returns a tuple with the Initial field value
// and a boolean to check if the value has been set.
func (o *UpdateWorkflowDto) GetInitialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Initial, true
}

// SetInitial sets field value
func (o *UpdateWorkflowDto) SetInitial(v string) {
	o.Initial = v
}

func (o UpdateWorkflowDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["steps"] = o.Steps
	}
	if o.SchemaIds != nil {
		toSerialize["schemaIds"] = o.SchemaIds
	}
	if true {
		toSerialize["initial"] = o.Initial
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWorkflowDto struct {
	value *UpdateWorkflowDto
	isSet bool
}

func (v NullableUpdateWorkflowDto) Get() *UpdateWorkflowDto {
	return v.value
}

func (v *NullableUpdateWorkflowDto) Set(val *UpdateWorkflowDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkflowDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkflowDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkflowDto(val *UpdateWorkflowDto) *NullableUpdateWorkflowDto {
	return &NullableUpdateWorkflowDto{value: val, isSet: true}
}

func (v NullableUpdateWorkflowDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkflowDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


