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

// CommentsDto struct for CommentsDto
type CommentsDto struct {
	// The created comments including the updates.
	CreatedComments []CommentDto `json:"createdComments,omitempty"`
	// The updates comments since the last version.
	UpdatedComments []CommentDto `json:"updatedComments,omitempty"`
	// The deleted comments since the last version.
	DeletedComments []string `json:"deletedComments,omitempty"`
	// The current version.
	Version *int64 `json:"version,omitempty"`
}

// NewCommentsDto instantiates a new CommentsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentsDto() *CommentsDto {
	this := CommentsDto{}
	return &this
}

// NewCommentsDtoWithDefaults instantiates a new CommentsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentsDtoWithDefaults() *CommentsDto {
	this := CommentsDto{}
	return &this
}

// GetCreatedComments returns the CreatedComments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentsDto) GetCreatedComments() []CommentDto {
	if o == nil {
		var ret []CommentDto
		return ret
	}
	return o.CreatedComments
}

// GetCreatedCommentsOk returns a tuple with the CreatedComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsDto) GetCreatedCommentsOk() ([]CommentDto, bool) {
	if o == nil || isNil(o.CreatedComments) {
    return nil, false
	}
	return o.CreatedComments, true
}

// HasCreatedComments returns a boolean if a field has been set.
func (o *CommentsDto) HasCreatedComments() bool {
	if o != nil && isNil(o.CreatedComments) {
		return true
	}

	return false
}

// SetCreatedComments gets a reference to the given []CommentDto and assigns it to the CreatedComments field.
func (o *CommentsDto) SetCreatedComments(v []CommentDto) {
	o.CreatedComments = v
}

// GetUpdatedComments returns the UpdatedComments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentsDto) GetUpdatedComments() []CommentDto {
	if o == nil {
		var ret []CommentDto
		return ret
	}
	return o.UpdatedComments
}

// GetUpdatedCommentsOk returns a tuple with the UpdatedComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsDto) GetUpdatedCommentsOk() ([]CommentDto, bool) {
	if o == nil || isNil(o.UpdatedComments) {
    return nil, false
	}
	return o.UpdatedComments, true
}

// HasUpdatedComments returns a boolean if a field has been set.
func (o *CommentsDto) HasUpdatedComments() bool {
	if o != nil && isNil(o.UpdatedComments) {
		return true
	}

	return false
}

// SetUpdatedComments gets a reference to the given []CommentDto and assigns it to the UpdatedComments field.
func (o *CommentsDto) SetUpdatedComments(v []CommentDto) {
	o.UpdatedComments = v
}

// GetDeletedComments returns the DeletedComments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentsDto) GetDeletedComments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeletedComments
}

// GetDeletedCommentsOk returns a tuple with the DeletedComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsDto) GetDeletedCommentsOk() ([]string, bool) {
	if o == nil || isNil(o.DeletedComments) {
    return nil, false
	}
	return o.DeletedComments, true
}

// HasDeletedComments returns a boolean if a field has been set.
func (o *CommentsDto) HasDeletedComments() bool {
	if o != nil && isNil(o.DeletedComments) {
		return true
	}

	return false
}

// SetDeletedComments gets a reference to the given []string and assigns it to the DeletedComments field.
func (o *CommentsDto) SetDeletedComments(v []string) {
	o.DeletedComments = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CommentsDto) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsDto) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CommentsDto) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CommentsDto) SetVersion(v int64) {
	o.Version = &v
}

func (o CommentsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedComments != nil {
		toSerialize["createdComments"] = o.CreatedComments
	}
	if o.UpdatedComments != nil {
		toSerialize["updatedComments"] = o.UpdatedComments
	}
	if o.DeletedComments != nil {
		toSerialize["deletedComments"] = o.DeletedComments
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCommentsDto struct {
	value *CommentsDto
	isSet bool
}

func (v NullableCommentsDto) Get() *CommentsDto {
	return v.value
}

func (v *NullableCommentsDto) Set(val *CommentsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentsDto(val *CommentsDto) *NullableCommentsDto {
	return &NullableCommentsDto{value: val, isSet: true}
}

func (v NullableCommentsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

