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

// TagsFieldPropertiesDto struct for TagsFieldPropertiesDto
type TagsFieldPropertiesDto struct {
	FieldPropertiesDto
	DefaultValues *map[string][]string `json:"defaultValues,omitempty"`
	// The default value.
	DefaultValue []string `json:"defaultValue,omitempty"`
	// The minimum allowed items for the field value.
	MinItems NullableInt32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems NullableInt32 `json:"maxItems,omitempty"`
	// The allowed values for the field value.
	AllowedValues []string `json:"allowedValues,omitempty"`
	// Indicates whether GraphQL Enum should be created.
	CreateEnum *bool `json:"createEnum,omitempty"`
	Editor *TagsFieldEditor `json:"editor,omitempty"`
}

// NewTagsFieldPropertiesDto instantiates a new TagsFieldPropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsFieldPropertiesDto(fieldType string) *TagsFieldPropertiesDto {
	this := TagsFieldPropertiesDto{}
	this.FieldType = fieldType
	return &this
}

// NewTagsFieldPropertiesDtoWithDefaults instantiates a new TagsFieldPropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsFieldPropertiesDtoWithDefaults() *TagsFieldPropertiesDto {
	this := TagsFieldPropertiesDto{}
	return &this
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *TagsFieldPropertiesDto) GetDefaultValues() map[string][]string {
	if o == nil || isNil(o.DefaultValues) {
		var ret map[string][]string
		return ret
	}
	return *o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsFieldPropertiesDto) GetDefaultValuesOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.DefaultValues) {
    return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasDefaultValues() bool {
	if o != nil && !isNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given map[string][]string and assigns it to the DefaultValues field.
func (o *TagsFieldPropertiesDto) SetDefaultValues(v map[string][]string) {
	o.DefaultValues = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagsFieldPropertiesDto) GetDefaultValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagsFieldPropertiesDto) GetDefaultValueOk() ([]string, bool) {
	if o == nil || isNil(o.DefaultValue) {
    return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasDefaultValue() bool {
	if o != nil && isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given []string and assigns it to the DefaultValue field.
func (o *TagsFieldPropertiesDto) SetDefaultValue(v []string) {
	o.DefaultValue = v
}

// GetMinItems returns the MinItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagsFieldPropertiesDto) GetMinItems() int32 {
	if o == nil || isNil(o.MinItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MinItems.Get()
}

// GetMinItemsOk returns a tuple with the MinItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagsFieldPropertiesDto) GetMinItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinItems.Get(), o.MinItems.IsSet()
}

// HasMinItems returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasMinItems() bool {
	if o != nil && o.MinItems.IsSet() {
		return true
	}

	return false
}

// SetMinItems gets a reference to the given NullableInt32 and assigns it to the MinItems field.
func (o *TagsFieldPropertiesDto) SetMinItems(v int32) {
	o.MinItems.Set(&v)
}
// SetMinItemsNil sets the value for MinItems to be an explicit nil
func (o *TagsFieldPropertiesDto) SetMinItemsNil() {
	o.MinItems.Set(nil)
}

// UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
func (o *TagsFieldPropertiesDto) UnsetMinItems() {
	o.MinItems.Unset()
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagsFieldPropertiesDto) GetMaxItems() int32 {
	if o == nil || isNil(o.MaxItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItems.Get()
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagsFieldPropertiesDto) GetMaxItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItems.Get(), o.MaxItems.IsSet()
}

// HasMaxItems returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasMaxItems() bool {
	if o != nil && o.MaxItems.IsSet() {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given NullableInt32 and assigns it to the MaxItems field.
func (o *TagsFieldPropertiesDto) SetMaxItems(v int32) {
	o.MaxItems.Set(&v)
}
// SetMaxItemsNil sets the value for MaxItems to be an explicit nil
func (o *TagsFieldPropertiesDto) SetMaxItemsNil() {
	o.MaxItems.Set(nil)
}

// UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
func (o *TagsFieldPropertiesDto) UnsetMaxItems() {
	o.MaxItems.Unset()
}

// GetAllowedValues returns the AllowedValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagsFieldPropertiesDto) GetAllowedValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagsFieldPropertiesDto) GetAllowedValuesOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedValues) {
    return nil, false
	}
	return o.AllowedValues, true
}

// HasAllowedValues returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasAllowedValues() bool {
	if o != nil && isNil(o.AllowedValues) {
		return true
	}

	return false
}

// SetAllowedValues gets a reference to the given []string and assigns it to the AllowedValues field.
func (o *TagsFieldPropertiesDto) SetAllowedValues(v []string) {
	o.AllowedValues = v
}

// GetCreateEnum returns the CreateEnum field value if set, zero value otherwise.
func (o *TagsFieldPropertiesDto) GetCreateEnum() bool {
	if o == nil || isNil(o.CreateEnum) {
		var ret bool
		return ret
	}
	return *o.CreateEnum
}

// GetCreateEnumOk returns a tuple with the CreateEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsFieldPropertiesDto) GetCreateEnumOk() (*bool, bool) {
	if o == nil || isNil(o.CreateEnum) {
    return nil, false
	}
	return o.CreateEnum, true
}

// HasCreateEnum returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasCreateEnum() bool {
	if o != nil && !isNil(o.CreateEnum) {
		return true
	}

	return false
}

// SetCreateEnum gets a reference to the given bool and assigns it to the CreateEnum field.
func (o *TagsFieldPropertiesDto) SetCreateEnum(v bool) {
	o.CreateEnum = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *TagsFieldPropertiesDto) GetEditor() TagsFieldEditor {
	if o == nil || isNil(o.Editor) {
		var ret TagsFieldEditor
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsFieldPropertiesDto) GetEditorOk() (*TagsFieldEditor, bool) {
	if o == nil || isNil(o.Editor) {
    return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *TagsFieldPropertiesDto) HasEditor() bool {
	if o != nil && !isNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given TagsFieldEditor and assigns it to the Editor field.
func (o *TagsFieldPropertiesDto) SetEditor(v TagsFieldEditor) {
	o.Editor = &v
}

func (o TagsFieldPropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFieldPropertiesDto, errFieldPropertiesDto := json.Marshal(o.FieldPropertiesDto)
	if errFieldPropertiesDto != nil {
		return []byte{}, errFieldPropertiesDto
	}
	errFieldPropertiesDto = json.Unmarshal([]byte(serializedFieldPropertiesDto), &toSerialize)
	if errFieldPropertiesDto != nil {
		return []byte{}, errFieldPropertiesDto
	}
	if !isNil(o.DefaultValues) {
		toSerialize["defaultValues"] = o.DefaultValues
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.MinItems.IsSet() {
		toSerialize["minItems"] = o.MinItems.Get()
	}
	if o.MaxItems.IsSet() {
		toSerialize["maxItems"] = o.MaxItems.Get()
	}
	if o.AllowedValues != nil {
		toSerialize["allowedValues"] = o.AllowedValues
	}
	if !isNil(o.CreateEnum) {
		toSerialize["createEnum"] = o.CreateEnum
	}
	if !isNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	return json.Marshal(toSerialize)
}

type NullableTagsFieldPropertiesDto struct {
	value *TagsFieldPropertiesDto
	isSet bool
}

func (v NullableTagsFieldPropertiesDto) Get() *TagsFieldPropertiesDto {
	return v.value
}

func (v *NullableTagsFieldPropertiesDto) Set(val *TagsFieldPropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsFieldPropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsFieldPropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsFieldPropertiesDto(val *TagsFieldPropertiesDto) *NullableTagsFieldPropertiesDto {
	return &NullableTagsFieldPropertiesDto{value: val, isSet: true}
}

func (v NullableTagsFieldPropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsFieldPropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


