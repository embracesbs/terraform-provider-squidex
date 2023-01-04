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

// UpdateUserDto struct for UpdateUserDto
type UpdateUserDto struct {
	// The email of the user. Unique value.
	Email string `json:"email"`
	// The display name (usually first name and last name) of the user.
	DisplayName string `json:"displayName"`
	// The password of the user.
	Password NullableString `json:"password,omitempty"`
	// Additional permissions for the user.
	Permissions []string `json:"permissions"`
}

// NewUpdateUserDto instantiates a new UpdateUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserDto(email string, displayName string, permissions []string) *UpdateUserDto {
	this := UpdateUserDto{}
	this.Email = email
	this.DisplayName = displayName
	this.Permissions = permissions
	return &this
}

// NewUpdateUserDtoWithDefaults instantiates a new UpdateUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserDtoWithDefaults() *UpdateUserDto {
	this := UpdateUserDto{}
	return &this
}

// GetEmail returns the Email field value
func (o *UpdateUserDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UpdateUserDto) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UpdateUserDto) SetEmail(v string) {
	o.Email = v
}

// GetDisplayName returns the DisplayName field value
func (o *UpdateUserDto) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *UpdateUserDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *UpdateUserDto) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserDto) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserDto) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateUserDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UpdateUserDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UpdateUserDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UpdateUserDto) UnsetPassword() {
	o.Password.Unset()
}

// GetPermissions returns the Permissions field value
func (o *UpdateUserDto) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UpdateUserDto) GetPermissionsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UpdateUserDto) SetPermissions(v []string) {
	o.Permissions = v
}

func (o UpdateUserDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUserDto struct {
	value *UpdateUserDto
	isSet bool
}

func (v NullableUpdateUserDto) Get() *UpdateUserDto {
	return v.value
}

func (v *NullableUpdateUserDto) Set(val *UpdateUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserDto(val *UpdateUserDto) *NullableUpdateUserDto {
	return &NullableUpdateUserDto{value: val, isSet: true}
}

func (v NullableUpdateUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


