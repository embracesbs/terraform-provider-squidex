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

// DateTimeCalculatedDefaultValue 
type DateTimeCalculatedDefaultValue string

// List of DateTimeCalculatedDefaultValue
const (
	NOW DateTimeCalculatedDefaultValue = "Now"
	TODAY DateTimeCalculatedDefaultValue = "Today"
)

// All allowed values of DateTimeCalculatedDefaultValue enum
var AllowedDateTimeCalculatedDefaultValueEnumValues = []DateTimeCalculatedDefaultValue{
	"Now",
	"Today",
}

func (v *DateTimeCalculatedDefaultValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DateTimeCalculatedDefaultValue(value)
	for _, existing := range AllowedDateTimeCalculatedDefaultValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DateTimeCalculatedDefaultValue", value)
}

// NewDateTimeCalculatedDefaultValueFromValue returns a pointer to a valid DateTimeCalculatedDefaultValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDateTimeCalculatedDefaultValueFromValue(v string) (*DateTimeCalculatedDefaultValue, error) {
	ev := DateTimeCalculatedDefaultValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DateTimeCalculatedDefaultValue: valid values are %v", v, AllowedDateTimeCalculatedDefaultValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DateTimeCalculatedDefaultValue) IsValid() bool {
	for _, existing := range AllowedDateTimeCalculatedDefaultValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DateTimeCalculatedDefaultValue value
func (v DateTimeCalculatedDefaultValue) Ptr() *DateTimeCalculatedDefaultValue {
	return &v
}

type NullableDateTimeCalculatedDefaultValue struct {
	value *DateTimeCalculatedDefaultValue
	isSet bool
}

func (v NullableDateTimeCalculatedDefaultValue) Get() *DateTimeCalculatedDefaultValue {
	return v.value
}

func (v *NullableDateTimeCalculatedDefaultValue) Set(val *DateTimeCalculatedDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeCalculatedDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeCalculatedDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeCalculatedDefaultValue(val *DateTimeCalculatedDefaultValue) *NullableDateTimeCalculatedDefaultValue {
	return &NullableDateTimeCalculatedDefaultValue{value: val, isSet: true}
}

func (v NullableDateTimeCalculatedDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeCalculatedDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
