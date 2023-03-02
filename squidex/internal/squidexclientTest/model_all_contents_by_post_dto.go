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

// AllContentsByPostDto struct for AllContentsByPostDto
type AllContentsByPostDto struct {
	// The list of ids to query.
	Ids []string `json:"ids,omitempty"`
	// The start of the schedule.
	ScheduledFrom NullableTime `json:"scheduledFrom,omitempty"`
	// The end of the schedule.
	ScheduledTo NullableTime `json:"scheduledTo,omitempty"`
}

// NewAllContentsByPostDto instantiates a new AllContentsByPostDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllContentsByPostDto() *AllContentsByPostDto {
	this := AllContentsByPostDto{}
	return &this
}

// NewAllContentsByPostDtoWithDefaults instantiates a new AllContentsByPostDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllContentsByPostDtoWithDefaults() *AllContentsByPostDto {
	this := AllContentsByPostDto{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllContentsByPostDto) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllContentsByPostDto) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AllContentsByPostDto) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *AllContentsByPostDto) SetIds(v []string) {
	o.Ids = v
}

// GetScheduledFrom returns the ScheduledFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllContentsByPostDto) GetScheduledFrom() time.Time {
	if o == nil || isNil(o.ScheduledFrom.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFrom.Get()
}

// GetScheduledFromOk returns a tuple with the ScheduledFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllContentsByPostDto) GetScheduledFromOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ScheduledFrom.Get(), o.ScheduledFrom.IsSet()
}

// HasScheduledFrom returns a boolean if a field has been set.
func (o *AllContentsByPostDto) HasScheduledFrom() bool {
	if o != nil && o.ScheduledFrom.IsSet() {
		return true
	}

	return false
}

// SetScheduledFrom gets a reference to the given NullableTime and assigns it to the ScheduledFrom field.
func (o *AllContentsByPostDto) SetScheduledFrom(v time.Time) {
	o.ScheduledFrom.Set(&v)
}
// SetScheduledFromNil sets the value for ScheduledFrom to be an explicit nil
func (o *AllContentsByPostDto) SetScheduledFromNil() {
	o.ScheduledFrom.Set(nil)
}

// UnsetScheduledFrom ensures that no value is present for ScheduledFrom, not even an explicit nil
func (o *AllContentsByPostDto) UnsetScheduledFrom() {
	o.ScheduledFrom.Unset()
}

// GetScheduledTo returns the ScheduledTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllContentsByPostDto) GetScheduledTo() time.Time {
	if o == nil || isNil(o.ScheduledTo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledTo.Get()
}

// GetScheduledToOk returns a tuple with the ScheduledTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllContentsByPostDto) GetScheduledToOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ScheduledTo.Get(), o.ScheduledTo.IsSet()
}

// HasScheduledTo returns a boolean if a field has been set.
func (o *AllContentsByPostDto) HasScheduledTo() bool {
	if o != nil && o.ScheduledTo.IsSet() {
		return true
	}

	return false
}

// SetScheduledTo gets a reference to the given NullableTime and assigns it to the ScheduledTo field.
func (o *AllContentsByPostDto) SetScheduledTo(v time.Time) {
	o.ScheduledTo.Set(&v)
}
// SetScheduledToNil sets the value for ScheduledTo to be an explicit nil
func (o *AllContentsByPostDto) SetScheduledToNil() {
	o.ScheduledTo.Set(nil)
}

// UnsetScheduledTo ensures that no value is present for ScheduledTo, not even an explicit nil
func (o *AllContentsByPostDto) UnsetScheduledTo() {
	o.ScheduledTo.Unset()
}

func (o AllContentsByPostDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.ScheduledFrom.IsSet() {
		toSerialize["scheduledFrom"] = o.ScheduledFrom.Get()
	}
	if o.ScheduledTo.IsSet() {
		toSerialize["scheduledTo"] = o.ScheduledTo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAllContentsByPostDto struct {
	value *AllContentsByPostDto
	isSet bool
}

func (v NullableAllContentsByPostDto) Get() *AllContentsByPostDto {
	return v.value
}

func (v *NullableAllContentsByPostDto) Set(val *AllContentsByPostDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAllContentsByPostDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAllContentsByPostDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllContentsByPostDto(val *AllContentsByPostDto) *NullableAllContentsByPostDto {
	return &NullableAllContentsByPostDto{value: val, isSet: true}
}

func (v NullableAllContentsByPostDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllContentsByPostDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

