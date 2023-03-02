/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
	"fmt"
)

// ResizeMode 
type ResizeMode string

// List of ResizeMode
const (
	CROP ResizeMode = "Crop"
	CROP_UPSIZE ResizeMode = "CropUpsize"
	PAD ResizeMode = "Pad"
	BOX_PAD ResizeMode = "BoxPad"
	MAX ResizeMode = "Max"
	MIN ResizeMode = "Min"
	STRETCH ResizeMode = "Stretch"
)

// All allowed values of ResizeMode enum
var AllowedResizeModeEnumValues = []ResizeMode{
	"Crop",
	"CropUpsize",
	"Pad",
	"BoxPad",
	"Max",
	"Min",
	"Stretch",
}

func (v *ResizeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResizeMode(value)
	for _, existing := range AllowedResizeModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResizeMode", value)
}

// NewResizeModeFromValue returns a pointer to a valid ResizeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResizeModeFromValue(v string) (*ResizeMode, error) {
	ev := ResizeMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResizeMode: valid values are %v", v, AllowedResizeModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResizeMode) IsValid() bool {
	for _, existing := range AllowedResizeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResizeMode value
func (v ResizeMode) Ptr() *ResizeMode {
	return &v
}

type NullableResizeMode struct {
	value *ResizeMode
	isSet bool
}

func (v NullableResizeMode) Get() *ResizeMode {
	return v.value
}

func (v *NullableResizeMode) Set(val *ResizeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableResizeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableResizeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResizeMode(val *ResizeMode) *NullableResizeMode {
	return &NullableResizeMode{value: val, isSet: true}
}

func (v NullableResizeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResizeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
