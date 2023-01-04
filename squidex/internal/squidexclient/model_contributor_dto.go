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

// ContributorDto struct for ContributorDto
type ContributorDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the user that contributes to the app.
	ContributorId string `json:"contributorId"`
	// The display name.
	ContributorName string `json:"contributorName"`
	// The email address.
	ContributorEmail string `json:"contributorEmail"`
	// The role of the contributor.
	Role NullableString `json:"role,omitempty"`
}

// NewContributorDto instantiates a new ContributorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContributorDto(links map[string]ResourceLink, contributorId string, contributorName string, contributorEmail string) *ContributorDto {
	this := ContributorDto{}
	this.Links = links
	this.ContributorId = contributorId
	this.ContributorName = contributorName
	this.ContributorEmail = contributorEmail
	return &this
}

// NewContributorDtoWithDefaults instantiates a new ContributorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContributorDtoWithDefaults() *ContributorDto {
	this := ContributorDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *ContributorDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ContributorDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ContributorDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetContributorId returns the ContributorId field value
func (o *ContributorDto) GetContributorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorId
}

// GetContributorIdOk returns a tuple with the ContributorId field value
// and a boolean to check if the value has been set.
func (o *ContributorDto) GetContributorIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorId, true
}

// SetContributorId sets field value
func (o *ContributorDto) SetContributorId(v string) {
	o.ContributorId = v
}

// GetContributorName returns the ContributorName field value
func (o *ContributorDto) GetContributorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorName
}

// GetContributorNameOk returns a tuple with the ContributorName field value
// and a boolean to check if the value has been set.
func (o *ContributorDto) GetContributorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorName, true
}

// SetContributorName sets field value
func (o *ContributorDto) SetContributorName(v string) {
	o.ContributorName = v
}

// GetContributorEmail returns the ContributorEmail field value
func (o *ContributorDto) GetContributorEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContributorEmail
}

// GetContributorEmailOk returns a tuple with the ContributorEmail field value
// and a boolean to check if the value has been set.
func (o *ContributorDto) GetContributorEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContributorEmail, true
}

// SetContributorEmail sets field value
func (o *ContributorDto) SetContributorEmail(v string) {
	o.ContributorEmail = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContributorDto) GetRole() string {
	if o == nil || isNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContributorDto) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *ContributorDto) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *ContributorDto) SetRole(v string) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *ContributorDto) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *ContributorDto) UnsetRole() {
	o.Role.Unset()
}

func (o ContributorDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
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

type NullableContributorDto struct {
	value *ContributorDto
	isSet bool
}

func (v NullableContributorDto) Get() *ContributorDto {
	return v.value
}

func (v *NullableContributorDto) Set(val *ContributorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContributorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContributorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContributorDto(val *ContributorDto) *NullableContributorDto {
	return &NullableContributorDto{value: val, isSet: true}
}

func (v NullableContributorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContributorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


