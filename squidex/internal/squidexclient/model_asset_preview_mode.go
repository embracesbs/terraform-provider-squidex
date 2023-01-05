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

// AssetPreviewMode 
type AssetPreviewMode string

// List of AssetPreviewMode
const (
	AssetPreviewModeIMAGE_AND_FILE_NAME AssetPreviewMode = "ImageAndFileName"
	AssetPreviewModeIMAGE AssetPreviewMode = "Image"
	AssetPreviewModeFILE_NAME AssetPreviewMode = "FileName"
)

// All allowed values of AssetPreviewMode enum
var AllowedAssetPreviewModeEnumValues = []AssetPreviewMode{
	"ImageAndFileName",
	"Image",
	"FileName",
}

func (v *AssetPreviewMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetPreviewMode(value)
	for _, existing := range AllowedAssetPreviewModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetPreviewMode", value)
}

// NewAssetPreviewModeFromValue returns a pointer to a valid AssetPreviewMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetPreviewModeFromValue(v string) (*AssetPreviewMode, error) {
	ev := AssetPreviewMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetPreviewMode: valid values are %v", v, AllowedAssetPreviewModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetPreviewMode) IsValid() bool {
	for _, existing := range AllowedAssetPreviewModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetPreviewMode value
func (v AssetPreviewMode) Ptr() *AssetPreviewMode {
	return &v
}

type NullableAssetPreviewMode struct {
	value *AssetPreviewMode
	isSet bool
}

func (v NullableAssetPreviewMode) Get() *AssetPreviewMode {
	return v.value
}

func (v *NullableAssetPreviewMode) Set(val *AssetPreviewMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetPreviewMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetPreviewMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetPreviewMode(val *AssetPreviewMode) *NullableAssetPreviewMode {
	return &NullableAssetPreviewMode{value: val, isSet: true}
}

func (v NullableAssetPreviewMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetPreviewMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

