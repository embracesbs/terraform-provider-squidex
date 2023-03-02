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

// CreateSchemaDtoAllOf struct for CreateSchemaDtoAllOf
type CreateSchemaDtoAllOf struct {
	// The name of the schema.
	Name string `json:"name"`
	Type *SchemaType `json:"type,omitempty"`
	// Set to true to allow a single content item only.
	// Deprecated
	IsSingleton *bool `json:"isSingleton,omitempty"`
}

// NewCreateSchemaDtoAllOf instantiates a new CreateSchemaDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSchemaDtoAllOf(name string) *CreateSchemaDtoAllOf {
	this := CreateSchemaDtoAllOf{}
	this.Name = name
	return &this
}

// NewCreateSchemaDtoAllOfWithDefaults instantiates a new CreateSchemaDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSchemaDtoAllOfWithDefaults() *CreateSchemaDtoAllOf {
	this := CreateSchemaDtoAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSchemaDtoAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSchemaDtoAllOf) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSchemaDtoAllOf) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateSchemaDtoAllOf) GetType() SchemaType {
	if o == nil || isNil(o.Type) {
		var ret SchemaType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDtoAllOf) GetTypeOk() (*SchemaType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateSchemaDtoAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SchemaType and assigns it to the Type field.
func (o *CreateSchemaDtoAllOf) SetType(v SchemaType) {
	o.Type = &v
}

// GetIsSingleton returns the IsSingleton field value if set, zero value otherwise.
// Deprecated
func (o *CreateSchemaDtoAllOf) GetIsSingleton() bool {
	if o == nil || isNil(o.IsSingleton) {
		var ret bool
		return ret
	}
	return *o.IsSingleton
}

// GetIsSingletonOk returns a tuple with the IsSingleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateSchemaDtoAllOf) GetIsSingletonOk() (*bool, bool) {
	if o == nil || isNil(o.IsSingleton) {
    return nil, false
	}
	return o.IsSingleton, true
}

// HasIsSingleton returns a boolean if a field has been set.
func (o *CreateSchemaDtoAllOf) HasIsSingleton() bool {
	if o != nil && !isNil(o.IsSingleton) {
		return true
	}

	return false
}

// SetIsSingleton gets a reference to the given bool and assigns it to the IsSingleton field.
// Deprecated
func (o *CreateSchemaDtoAllOf) SetIsSingleton(v bool) {
	o.IsSingleton = &v
}

func (o CreateSchemaDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.IsSingleton) {
		toSerialize["isSingleton"] = o.IsSingleton
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSchemaDtoAllOf struct {
	value *CreateSchemaDtoAllOf
	isSet bool
}

func (v NullableCreateSchemaDtoAllOf) Get() *CreateSchemaDtoAllOf {
	return v.value
}

func (v *NullableCreateSchemaDtoAllOf) Set(val *CreateSchemaDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSchemaDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSchemaDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSchemaDtoAllOf(val *CreateSchemaDtoAllOf) *NullableCreateSchemaDtoAllOf {
	return &NullableCreateSchemaDtoAllOf{value: val, isSet: true}
}

func (v NullableCreateSchemaDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSchemaDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

