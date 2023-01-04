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

// ContributorDtoAllOf struct for ContributorDtoAllOf
type ContributorDtoAllOf struct {
	// The id of the user that contributes to the app.
	ContributorId string `json:"contributorId"`
	// The display name.
	ContributorName string `json:"contributorName"`
	// The email address.
	ContributorEmail string `json:"contributorEmail"`
	// The role of the contributor.
	Role NullableString `json:"role,omitempty"`
}

// NewContributorDtoAllOf instantiates a new ContributorDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContributorDtoAllOf(contributorId string, contributorName string, contributorEmail string) *ContributorDtoAllOf {
	this := ContributorDtoAllOf{}
	this.ContributorId = contributorId
	this.ContributorName = contributorName
	this.ContributorEmail = contributorEmail
	return &this
}

// NewContributorDtoAllOfWithDefaults instantiates a new ContributorDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContributorDtoAllOfWithDefaults() *ContributorDtoAllOf {
	this := ContributorDtoAllOf{}
	return &this
}

// GetContributorId returns the ContributorId field value
func (o *ContributorDtoAllOf) GetContributorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorId
}

// GetContributorIdOk returns a tuple with the ContributorId field value
// and a boolean to check if the value has been set.
func (o *ContributorDtoAllOf) GetContributorIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorId, true
}

// SetContributorId sets field value
func (o *ContributorDtoAllOf) SetContributorId(v string) {
	o.ContributorId = v
}

// GetContributorName returns the ContributorName field value
func (o *ContributorDtoAllOf) GetContributorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorName
}

// GetContributorNameOk returns a tuple with the ContributorName field value
// and a boolean to check if the value has been set.
func (o *ContributorDtoAllOf) GetContributorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorName, true
}

// SetContributorName sets field value
func (o *ContributorDtoAllOf) SetContributorName(v string) {
	o.ContributorName = v
}

// GetContributorEmail returns the ContributorEmail field value
func (o *ContributorDtoAllOf) GetContributorEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorEmail
}

// GetContributorEmailOk returns a tuple with the ContributorEmail field value
// and a boolean to check if the value has been set.
func (o *ContributorDtoAllOf) GetContributorEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorEmail, true
}

// SetContributorEmail sets field value
func (o *ContributorDtoAllOf) SetContributorEmail(v string) {
	o.ContributorEmail = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContributorDtoAllOf) GetRole() string {
	if o == nil || isNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContributorDtoAllOf) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *ContributorDtoAllOf) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *ContributorDtoAllOf) SetRole(v string) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *ContributorDtoAllOf) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *ContributorDtoAllOf) UnsetRole() {
	o.Role.Unset()
}

func (o ContributorDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contributorId"] = o.ContributorId
	}
	if true {
		toSerialize["contributorName"] = o.ContributorName
	}
	if true {
		toSerialize["contributorEmail"] = o.ContributorEmail
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContributorDtoAllOf struct {
	value *ContributorDtoAllOf
	isSet bool
}

func (v NullableContributorDtoAllOf) Get() *ContributorDtoAllOf {
	return v.value
}

func (v *NullableContributorDtoAllOf) Set(val *ContributorDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContributorDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContributorDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContributorDtoAllOf(val *ContributorDtoAllOf) *NullableContributorDtoAllOf {
	return &NullableContributorDtoAllOf{value: val, isSet: true}
}

func (v NullableContributorDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContributorDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


