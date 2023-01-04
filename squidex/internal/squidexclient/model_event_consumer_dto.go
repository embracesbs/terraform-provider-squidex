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

// EventConsumerDto struct for EventConsumerDto
type EventConsumerDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	IsStopped *bool `json:"isStopped,omitempty"`
	IsResetting *bool `json:"isResetting,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Name *string `json:"name,omitempty"`
	Error NullableString `json:"error,omitempty"`
	Position NullableString `json:"position,omitempty"`
}

// NewEventConsumerDto instantiates a new EventConsumerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventConsumerDto(links map[string]ResourceLink) *EventConsumerDto {
	this := EventConsumerDto{}
	this.Links = links
	return &this
}

// NewEventConsumerDtoWithDefaults instantiates a new EventConsumerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventConsumerDtoWithDefaults() *EventConsumerDto {
	this := EventConsumerDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *EventConsumerDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EventConsumerDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EventConsumerDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetIsStopped returns the IsStopped field value if set, zero value otherwise.
func (o *EventConsumerDto) GetIsStopped() bool {
	if o == nil || isNil(o.IsStopped) {
		var ret bool
		return ret
	}
	return *o.IsStopped
}

// GetIsStoppedOk returns a tuple with the IsStopped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventConsumerDto) GetIsStoppedOk() (*bool, bool) {
	if o == nil || isNil(o.IsStopped) {
    return nil, false
	}
	return o.IsStopped, true
}

// HasIsStopped returns a boolean if a field has been set.
func (o *EventConsumerDto) HasIsStopped() bool {
	if o != nil && !isNil(o.IsStopped) {
		return true
	}

	return false
}

// SetIsStopped gets a reference to the given bool and assigns it to the IsStopped field.
func (o *EventConsumerDto) SetIsStopped(v bool) {
	o.IsStopped = &v
}

// GetIsResetting returns the IsResetting field value if set, zero value otherwise.
func (o *EventConsumerDto) GetIsResetting() bool {
	if o == nil || isNil(o.IsResetting) {
		var ret bool
		return ret
	}
	return *o.IsResetting
}

// GetIsResettingOk returns a tuple with the IsResetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventConsumerDto) GetIsResettingOk() (*bool, bool) {
	if o == nil || isNil(o.IsResetting) {
    return nil, false
	}
	return o.IsResetting, true
}

// HasIsResetting returns a boolean if a field has been set.
func (o *EventConsumerDto) HasIsResetting() bool {
	if o != nil && !isNil(o.IsResetting) {
		return true
	}

	return false
}

// SetIsResetting gets a reference to the given bool and assigns it to the IsResetting field.
func (o *EventConsumerDto) SetIsResetting(v bool) {
	o.IsResetting = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EventConsumerDto) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventConsumerDto) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EventConsumerDto) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *EventConsumerDto) SetCount(v int32) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventConsumerDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventConsumerDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventConsumerDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventConsumerDto) SetName(v string) {
	o.Name = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventConsumerDto) GetError() string {
	if o == nil || isNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventConsumerDto) GetErrorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *EventConsumerDto) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *EventConsumerDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *EventConsumerDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *EventConsumerDto) UnsetError() {
	o.Error.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventConsumerDto) GetPosition() string {
	if o == nil || isNil(o.Position.Get()) {
		var ret string
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventConsumerDto) GetPositionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *EventConsumerDto) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableString and assigns it to the Position field.
func (o *EventConsumerDto) SetPosition(v string) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *EventConsumerDto) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *EventConsumerDto) UnsetPosition() {
	o.Position.Unset()
}

func (o EventConsumerDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.IsStopped) {
		toSerialize["isStopped"] = o.IsStopped
	}
	if !isNil(o.IsResetting) {
		toSerialize["isResetting"] = o.IsResetting
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEventConsumerDto struct {
	value *EventConsumerDto
	isSet bool
}

func (v NullableEventConsumerDto) Get() *EventConsumerDto {
	return v.value
}

func (v *NullableEventConsumerDto) Set(val *EventConsumerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEventConsumerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEventConsumerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventConsumerDto(val *EventConsumerDto) *NullableEventConsumerDto {
	return &NullableEventConsumerDto{value: val, isSet: true}
}

func (v NullableEventConsumerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventConsumerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


