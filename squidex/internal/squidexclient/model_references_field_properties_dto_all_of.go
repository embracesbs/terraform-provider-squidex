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

// ReferencesFieldPropertiesDtoAllOf struct for ReferencesFieldPropertiesDtoAllOf
type ReferencesFieldPropertiesDtoAllOf struct {
	DefaultValues *map[string][]string `json:"defaultValues,omitempty"`
	// The default value as a list of content ids.
	DefaultValue []string `json:"defaultValue,omitempty"`
	// The minimum allowed items for the field value.
	MinItems NullableInt32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems NullableInt32 `json:"maxItems,omitempty"`
	// True, if duplicate values are allowed.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// True to resolve references in the content list.
	ResolveReference *bool `json:"resolveReference,omitempty"`
	// True when all references must be published.
	MustBePublished *bool `json:"mustBePublished,omitempty"`
	Editor *ReferencesFieldEditor `json:"editor,omitempty"`
	// The id of the referenced schemas.
	SchemaIds []string `json:"schemaIds,omitempty"`
}

// NewReferencesFieldPropertiesDtoAllOf instantiates a new ReferencesFieldPropertiesDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencesFieldPropertiesDtoAllOf() *ReferencesFieldPropertiesDtoAllOf {
	this := ReferencesFieldPropertiesDtoAllOf{}
	return &this
}

// NewReferencesFieldPropertiesDtoAllOfWithDefaults instantiates a new ReferencesFieldPropertiesDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencesFieldPropertiesDtoAllOfWithDefaults() *ReferencesFieldPropertiesDtoAllOf {
	this := ReferencesFieldPropertiesDtoAllOf{}
	return &this
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValues() map[string][]string {
	if o == nil || isNil(o.DefaultValues) {
		var ret map[string][]string
		return ret
	}
	return *o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.DefaultValues) {
    return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasDefaultValues() bool {
	if o != nil && !isNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given map[string][]string and assigns it to the DefaultValues field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetDefaultValues(v map[string][]string) {
	o.DefaultValues = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValueOk() ([]string, bool) {
	if o == nil || isNil(o.DefaultValue) {
    return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasDefaultValue() bool {
	if o != nil && isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given []string and assigns it to the DefaultValue field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetDefaultValue(v []string) {
	o.DefaultValue = v
}

// GetMinItems returns the MinItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferencesFieldPropertiesDtoAllOf) GetMinItems() int32 {
	if o == nil || isNil(o.MinItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MinItems.Get()
}

// GetMinItemsOk returns a tuple with the MinItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferencesFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinItems.Get(), o.MinItems.IsSet()
}

// HasMinItems returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasMinItems() bool {
	if o != nil && o.MinItems.IsSet() {
		return true
	}

	return false
}

// SetMinItems gets a reference to the given NullableInt32 and assigns it to the MinItems field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetMinItems(v int32) {
	o.MinItems.Set(&v)
}
// SetMinItemsNil sets the value for MinItems to be an explicit nil
func (o *ReferencesFieldPropertiesDtoAllOf) SetMinItemsNil() {
	o.MinItems.Set(nil)
}

// UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
func (o *ReferencesFieldPropertiesDtoAllOf) UnsetMinItems() {
	o.MinItems.Unset()
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferencesFieldPropertiesDtoAllOf) GetMaxItems() int32 {
	if o == nil || isNil(o.MaxItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItems.Get()
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferencesFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItems.Get(), o.MaxItems.IsSet()
}

// HasMaxItems returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasMaxItems() bool {
	if o != nil && o.MaxItems.IsSet() {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given NullableInt32 and assigns it to the MaxItems field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetMaxItems(v int32) {
	o.MaxItems.Set(&v)
}
// SetMaxItemsNil sets the value for MaxItems to be an explicit nil
func (o *ReferencesFieldPropertiesDtoAllOf) SetMaxItemsNil() {
	o.MaxItems.Set(nil)
}

// UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
func (o *ReferencesFieldPropertiesDtoAllOf) UnsetMaxItems() {
	o.MaxItems.Unset()
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *ReferencesFieldPropertiesDtoAllOf) GetAllowDuplicates() bool {
	if o == nil || isNil(o.AllowDuplicates) {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || isNil(o.AllowDuplicates) {
    return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasAllowDuplicates() bool {
	if o != nil && !isNil(o.AllowDuplicates) {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetResolveReference returns the ResolveReference field value if set, zero value otherwise.
func (o *ReferencesFieldPropertiesDtoAllOf) GetResolveReference() bool {
	if o == nil || isNil(o.ResolveReference) {
		var ret bool
		return ret
	}
	return *o.ResolveReference
}

// GetResolveReferenceOk returns a tuple with the ResolveReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) GetResolveReferenceOk() (*bool, bool) {
	if o == nil || isNil(o.ResolveReference) {
    return nil, false
	}
	return o.ResolveReference, true
}

// HasResolveReference returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasResolveReference() bool {
	if o != nil && !isNil(o.ResolveReference) {
		return true
	}

	return false
}

// SetResolveReference gets a reference to the given bool and assigns it to the ResolveReference field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetResolveReference(v bool) {
	o.ResolveReference = &v
}

// GetMustBePublished returns the MustBePublished field value if set, zero value otherwise.
func (o *ReferencesFieldPropertiesDtoAllOf) GetMustBePublished() bool {
	if o == nil || isNil(o.MustBePublished) {
		var ret bool
		return ret
	}
	return *o.MustBePublished
}

// GetMustBePublishedOk returns a tuple with the MustBePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) GetMustBePublishedOk() (*bool, bool) {
	if o == nil || isNil(o.MustBePublished) {
    return nil, false
	}
	return o.MustBePublished, true
}

// HasMustBePublished returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasMustBePublished() bool {
	if o != nil && !isNil(o.MustBePublished) {
		return true
	}

	return false
}

// SetMustBePublished gets a reference to the given bool and assigns it to the MustBePublished field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetMustBePublished(v bool) {
	o.MustBePublished = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *ReferencesFieldPropertiesDtoAllOf) GetEditor() ReferencesFieldEditor {
	if o == nil || isNil(o.Editor) {
		var ret ReferencesFieldEditor
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) GetEditorOk() (*ReferencesFieldEditor, bool) {
	if o == nil || isNil(o.Editor) {
    return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasEditor() bool {
	if o != nil && !isNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given ReferencesFieldEditor and assigns it to the Editor field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetEditor(v ReferencesFieldEditor) {
	o.Editor = &v
}

// GetSchemaIds returns the SchemaIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferencesFieldPropertiesDtoAllOf) GetSchemaIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SchemaIds
}

// GetSchemaIdsOk returns a tuple with the SchemaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferencesFieldPropertiesDtoAllOf) GetSchemaIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SchemaIds) {
    return nil, false
	}
	return o.SchemaIds, true
}

// HasSchemaIds returns a boolean if a field has been set.
func (o *ReferencesFieldPropertiesDtoAllOf) HasSchemaIds() bool {
	if o != nil && isNil(o.SchemaIds) {
		return true
	}

	return false
}

// SetSchemaIds gets a reference to the given []string and assigns it to the SchemaIds field.
func (o *ReferencesFieldPropertiesDtoAllOf) SetSchemaIds(v []string) {
	o.SchemaIds = v
}

func (o ReferencesFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.AllowDuplicates) {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
	}
	if !isNil(o.ResolveReference) {
		toSerialize["resolveReference"] = o.ResolveReference
	}
	if !isNil(o.MustBePublished) {
		toSerialize["mustBePublished"] = o.MustBePublished
	}
	if !isNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	if o.SchemaIds != nil {
		toSerialize["schemaIds"] = o.SchemaIds
	}
	return json.Marshal(toSerialize)
}

type NullableReferencesFieldPropertiesDtoAllOf struct {
	value *ReferencesFieldPropertiesDtoAllOf
	isSet bool
}

func (v NullableReferencesFieldPropertiesDtoAllOf) Get() *ReferencesFieldPropertiesDtoAllOf {
	return v.value
}

func (v *NullableReferencesFieldPropertiesDtoAllOf) Set(val *ReferencesFieldPropertiesDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencesFieldPropertiesDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencesFieldPropertiesDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencesFieldPropertiesDtoAllOf(val *ReferencesFieldPropertiesDtoAllOf) *NullableReferencesFieldPropertiesDtoAllOf {
	return &NullableReferencesFieldPropertiesDtoAllOf{value: val, isSet: true}
}

func (v NullableReferencesFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencesFieldPropertiesDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


