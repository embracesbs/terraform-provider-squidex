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

// BulkUpdateContentType 
type BulkUpdateContentType string

// List of BulkUpdateContentType
const (
	BulkUpdateContentTypeUPSERT BulkUpdateContentType = "Upsert"
	BulkUpdateContentTypeCHANGE_STATUS BulkUpdateContentType = "ChangeStatus"
	BulkUpdateContentTypeCREATE BulkUpdateContentType = "Create"
	BulkUpdateContentTypeDELETE BulkUpdateContentType = "Delete"
	BulkUpdateContentTypePATCH BulkUpdateContentType = "Patch"
	BulkUpdateContentTypeUPDATE BulkUpdateContentType = "Update"
	BulkUpdateContentTypeVALIDATE BulkUpdateContentType = "Validate"
)

// All allowed values of BulkUpdateContentType enum
var AllowedBulkUpdateContentTypeEnumValues = []BulkUpdateContentType{
	"Upsert",
	"ChangeStatus",
	"Create",
	"Delete",
	"Patch",
	"Update",
	"Validate",
}

func (v *BulkUpdateContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkUpdateContentType(value)
	for _, existing := range AllowedBulkUpdateContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkUpdateContentType", value)
}

// NewBulkUpdateContentTypeFromValue returns a pointer to a valid BulkUpdateContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkUpdateContentTypeFromValue(v string) (*BulkUpdateContentType, error) {
	ev := BulkUpdateContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkUpdateContentType: valid values are %v", v, AllowedBulkUpdateContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkUpdateContentType) IsValid() bool {
	for _, existing := range AllowedBulkUpdateContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BulkUpdateContentType value
func (v BulkUpdateContentType) Ptr() *BulkUpdateContentType {
	return &v
}

type NullableBulkUpdateContentType struct {
	value *BulkUpdateContentType
	isSet bool
}

func (v NullableBulkUpdateContentType) Get() *BulkUpdateContentType {
	return v.value
}

func (v *NullableBulkUpdateContentType) Set(val *BulkUpdateContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateContentType(val *BulkUpdateContentType) *NullableBulkUpdateContentType {
	return &NullableBulkUpdateContentType{value: val, isSet: true}
}

func (v NullableBulkUpdateContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

