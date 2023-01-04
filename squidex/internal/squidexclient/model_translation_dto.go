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

// TranslationDto struct for TranslationDto
type TranslationDto struct {
	Result *TranslationResultCode `json:"result,omitempty"`
	// The translated text.
	Text NullableString `json:"text,omitempty"`
}

// NewTranslationDto instantiates a new TranslationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationDto() *TranslationDto {
	this := TranslationDto{}
	return &this
}

// NewTranslationDtoWithDefaults instantiates a new TranslationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationDtoWithDefaults() *TranslationDto {
	this := TranslationDto{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TranslationDto) GetResult() TranslationResultCode {
	if o == nil || isNil(o.Result) {
		var ret TranslationResultCode
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationDto) GetResultOk() (*TranslationResultCode, bool) {
	if o == nil || isNil(o.Result) {
    return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TranslationDto) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TranslationResultCode and assigns it to the Result field.
func (o *TranslationDto) SetResult(v TranslationResultCode) {
	o.Result = &v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranslationDto) GetText() string {
	if o == nil || isNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranslationDto) GetTextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *TranslationDto) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *TranslationDto) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *TranslationDto) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *TranslationDto) UnsetText() {
	o.Text.Unset()
}

func (o TranslationDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTranslationDto struct {
	value *TranslationDto
	isSet bool
}

func (v NullableTranslationDto) Get() *TranslationDto {
	return v.value
}

func (v *NullableTranslationDto) Set(val *TranslationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationDto(val *TranslationDto) *NullableTranslationDto {
	return &NullableTranslationDto{value: val, isSet: true}
}

func (v NullableTranslationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


