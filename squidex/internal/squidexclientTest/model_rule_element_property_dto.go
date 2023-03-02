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

// RuleElementPropertyDto struct for RuleElementPropertyDto
type RuleElementPropertyDto struct {
	Editor RuleFieldEditor `json:"editor"`
	// The name of the editor.
	Name string `json:"name"`
	// The label to use.
	Display string `json:"display"`
	// The options, if the editor is a dropdown.
	Options []string `json:"options,omitempty"`
	// The optional description.
	Description NullableString `json:"description,omitempty"`
	// Indicates if the property is formattable.
	IsFormattable *bool `json:"isFormattable,omitempty"`
	// Indicates if the property is required.
	IsRequired *bool `json:"isRequired,omitempty"`
}

// NewRuleElementPropertyDto instantiates a new RuleElementPropertyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleElementPropertyDto(editor RuleFieldEditor, name string, display string) *RuleElementPropertyDto {
	this := RuleElementPropertyDto{}
	this.Editor = editor
	this.Name = name
	this.Display = display
	return &this
}

// NewRuleElementPropertyDtoWithDefaults instantiates a new RuleElementPropertyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleElementPropertyDtoWithDefaults() *RuleElementPropertyDto {
	this := RuleElementPropertyDto{}
	return &this
}

// GetEditor returns the Editor field value
func (o *RuleElementPropertyDto) GetEditor() RuleFieldEditor {
	if o == nil {
		var ret RuleFieldEditor
		return ret
	}

	return o.Editor
}

// GetEditorOk returns a tuple with the Editor field value
// and a boolean to check if the value has been set.
func (o *RuleElementPropertyDto) GetEditorOk() (*RuleFieldEditor, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Editor, true
}

// SetEditor sets field value
func (o *RuleElementPropertyDto) SetEditor(v RuleFieldEditor) {
	o.Editor = v
}

// GetName returns the Name field value
func (o *RuleElementPropertyDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleElementPropertyDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleElementPropertyDto) SetName(v string) {
	o.Name = v
}

// GetDisplay returns the Display field value
func (o *RuleElementPropertyDto) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *RuleElementPropertyDto) GetDisplayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *RuleElementPropertyDto) SetDisplay(v string) {
	o.Display = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleElementPropertyDto) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleElementPropertyDto) GetOptionsOk() ([]string, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *RuleElementPropertyDto) HasOptions() bool {
	if o != nil && isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *RuleElementPropertyDto) SetOptions(v []string) {
	o.Options = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleElementPropertyDto) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleElementPropertyDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleElementPropertyDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RuleElementPropertyDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RuleElementPropertyDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RuleElementPropertyDto) UnsetDescription() {
	o.Description.Unset()
}

// GetIsFormattable returns the IsFormattable field value if set, zero value otherwise.
func (o *RuleElementPropertyDto) GetIsFormattable() bool {
	if o == nil || isNil(o.IsFormattable) {
		var ret bool
		return ret
	}
	return *o.IsFormattable
}

// GetIsFormattableOk returns a tuple with the IsFormattable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleElementPropertyDto) GetIsFormattableOk() (*bool, bool) {
	if o == nil || isNil(o.IsFormattable) {
    return nil, false
	}
	return o.IsFormattable, true
}

// HasIsFormattable returns a boolean if a field has been set.
func (o *RuleElementPropertyDto) HasIsFormattable() bool {
	if o != nil && !isNil(o.IsFormattable) {
		return true
	}

	return false
}

// SetIsFormattable gets a reference to the given bool and assigns it to the IsFormattable field.
func (o *RuleElementPropertyDto) SetIsFormattable(v bool) {
	o.IsFormattable = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *RuleElementPropertyDto) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleElementPropertyDto) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
    return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *RuleElementPropertyDto) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *RuleElementPropertyDto) SetIsRequired(v bool) {
	o.IsRequired = &v
}

func (o RuleElementPropertyDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["editor"] = o.Editor
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["display"] = o.Display
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.IsFormattable) {
		toSerialize["isFormattable"] = o.IsFormattable
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	return json.Marshal(toSerialize)
}

type NullableRuleElementPropertyDto struct {
	value *RuleElementPropertyDto
	isSet bool
}

func (v NullableRuleElementPropertyDto) Get() *RuleElementPropertyDto {
	return v.value
}

func (v *NullableRuleElementPropertyDto) Set(val *RuleElementPropertyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleElementPropertyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleElementPropertyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleElementPropertyDto(val *RuleElementPropertyDto) *NullableRuleElementPropertyDto {
	return &NullableRuleElementPropertyDto{value: val, isSet: true}
}

func (v NullableRuleElementPropertyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleElementPropertyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

