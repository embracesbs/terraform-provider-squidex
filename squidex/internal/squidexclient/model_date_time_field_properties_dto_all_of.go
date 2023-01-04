/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
	"time"
)

// DateTimeFieldPropertiesDtoAllOf struct for DateTimeFieldPropertiesDtoAllOf
type DateTimeFieldPropertiesDtoAllOf struct {
	DefaultValues *map[string]time.Time `json:"defaultValues,omitempty"`
	// The default value for the field value.
	DefaultValue NullableTime `json:"defaultValue,omitempty"`
	// The maximum allowed value for the field value.
	MaxValue NullableTime `json:"maxValue,omitempty"`
	// The minimum allowed value for the field value.
	MinValue NullableTime `json:"minValue,omitempty"`
	// The format pattern when displayed in the UI.
	Format NullableString `json:"format,omitempty"`
	Editor *DateTimeFieldEditor `json:"editor,omitempty"`
	CalculatedDefaultValue *DateTimeCalculatedDefaultValue `json:"calculatedDefaultValue,omitempty"`
}

// NewDateTimeFieldPropertiesDtoAllOf instantiates a new DateTimeFieldPropertiesDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeFieldPropertiesDtoAllOf() *DateTimeFieldPropertiesDtoAllOf {
	this := DateTimeFieldPropertiesDtoAllOf{}
	return &this
}

// NewDateTimeFieldPropertiesDtoAllOfWithDefaults instantiates a new DateTimeFieldPropertiesDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeFieldPropertiesDtoAllOfWithDefaults() *DateTimeFieldPropertiesDtoAllOf {
	this := DateTimeFieldPropertiesDtoAllOf{}
	return &this
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValues() map[string]time.Time {
	if o == nil || isNil(o.DefaultValues) {
		var ret map[string]time.Time
		return ret
	}
	return *o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string]time.Time, bool) {
	if o == nil || isNil(o.DefaultValues) {
    return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasDefaultValues() bool {
	if o != nil && !isNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given map[string]time.Time and assigns it to the DefaultValues field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValues(v map[string]time.Time) {
	o.DefaultValues = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValue() time.Time {
	if o == nil || isNil(o.DefaultValue.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValueOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableTime and assigns it to the DefaultValue field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValue(v time.Time) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFieldPropertiesDtoAllOf) GetMaxValue() time.Time {
	if o == nil || isNil(o.MaxValue.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MaxValue.Get()
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFieldPropertiesDtoAllOf) GetMaxValueOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxValue.Get(), o.MaxValue.IsSet()
}

// HasMaxValue returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasMaxValue() bool {
	if o != nil && o.MaxValue.IsSet() {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given NullableTime and assigns it to the MaxValue field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetMaxValue(v time.Time) {
	o.MaxValue.Set(&v)
}
// SetMaxValueNil sets the value for MaxValue to be an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) SetMaxValueNil() {
	o.MaxValue.Set(nil)
}

// UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) UnsetMaxValue() {
	o.MaxValue.Unset()
}

// GetMinValue returns the MinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFieldPropertiesDtoAllOf) GetMinValue() time.Time {
	if o == nil || isNil(o.MinValue.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MinValue.Get()
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFieldPropertiesDtoAllOf) GetMinValueOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinValue.Get(), o.MinValue.IsSet()
}

// HasMinValue returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasMinValue() bool {
	if o != nil && o.MinValue.IsSet() {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given NullableTime and assigns it to the MinValue field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetMinValue(v time.Time) {
	o.MinValue.Set(&v)
}
// SetMinValueNil sets the value for MinValue to be an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) SetMinValueNil() {
	o.MinValue.Set(nil)
}

// UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) UnsetMinValue() {
	o.MinValue.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFieldPropertiesDtoAllOf) GetFormat() string {
	if o == nil || isNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFieldPropertiesDtoAllOf) GetFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *DateTimeFieldPropertiesDtoAllOf) UnsetFormat() {
	o.Format.Unset()
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *DateTimeFieldPropertiesDtoAllOf) GetEditor() DateTimeFieldEditor {
	if o == nil || isNil(o.Editor) {
		var ret DateTimeFieldEditor
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) GetEditorOk() (*DateTimeFieldEditor, bool) {
	if o == nil || isNil(o.Editor) {
    return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasEditor() bool {
	if o != nil && !isNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given DateTimeFieldEditor and assigns it to the Editor field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetEditor(v DateTimeFieldEditor) {
	o.Editor = &v
}

// GetCalculatedDefaultValue returns the CalculatedDefaultValue field value if set, zero value otherwise.
func (o *DateTimeFieldPropertiesDtoAllOf) GetCalculatedDefaultValue() DateTimeCalculatedDefaultValue {
	if o == nil || isNil(o.CalculatedDefaultValue) {
		var ret DateTimeCalculatedDefaultValue
		return ret
	}
	return *o.CalculatedDefaultValue
}

// GetCalculatedDefaultValueOk returns a tuple with the CalculatedDefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) GetCalculatedDefaultValueOk() (*DateTimeCalculatedDefaultValue, bool) {
	if o == nil || isNil(o.CalculatedDefaultValue) {
    return nil, false
	}
	return o.CalculatedDefaultValue, true
}

// HasCalculatedDefaultValue returns a boolean if a field has been set.
func (o *DateTimeFieldPropertiesDtoAllOf) HasCalculatedDefaultValue() bool {
	if o != nil && !isNil(o.CalculatedDefaultValue) {
		return true
	}

	return false
}

// SetCalculatedDefaultValue gets a reference to the given DateTimeCalculatedDefaultValue and assigns it to the CalculatedDefaultValue field.
func (o *DateTimeFieldPropertiesDtoAllOf) SetCalculatedDefaultValue(v DateTimeCalculatedDefaultValue) {
	o.CalculatedDefaultValue = &v
}

func (o DateTimeFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultValues) {
		toSerialize["defaultValues"] = o.DefaultValues
	}
	if o.DefaultValue.IsSet() {
		toSerialize["defaultValue"] = o.DefaultValue.Get()
	}
	if o.MaxValue.IsSet() {
		toSerialize["maxValue"] = o.MaxValue.Get()
	}
	if o.MinValue.IsSet() {
		toSerialize["minValue"] = o.MinValue.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	if !isNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	if !isNil(o.CalculatedDefaultValue) {
		toSerialize["calculatedDefaultValue"] = o.CalculatedDefaultValue
	}
	return json.Marshal(toSerialize)
}

type NullableDateTimeFieldPropertiesDtoAllOf struct {
	value *DateTimeFieldPropertiesDtoAllOf
	isSet bool
}

func (v NullableDateTimeFieldPropertiesDtoAllOf) Get() *DateTimeFieldPropertiesDtoAllOf {
	return v.value
}

func (v *NullableDateTimeFieldPropertiesDtoAllOf) Set(val *DateTimeFieldPropertiesDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeFieldPropertiesDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeFieldPropertiesDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeFieldPropertiesDtoAllOf(val *DateTimeFieldPropertiesDtoAllOf) *NullableDateTimeFieldPropertiesDtoAllOf {
	return &NullableDateTimeFieldPropertiesDtoAllOf{value: val, isSet: true}
}

func (v NullableDateTimeFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeFieldPropertiesDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


