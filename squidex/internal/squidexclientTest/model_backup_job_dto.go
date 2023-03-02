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

// BackupJobDto struct for BackupJobDto
type BackupJobDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the backup job.
	Id *string `json:"id,omitempty"`
	// The time when the job has been started.
	Started *time.Time `json:"started,omitempty"`
	// The time when the job has been stopped.
	Stopped NullableTime `json:"stopped,omitempty"`
	// The number of handled events.
	HandledEvents *int32 `json:"handledEvents,omitempty"`
	// The number of handled assets.
	HandledAssets *int32 `json:"handledAssets,omitempty"`
	Status *JobStatus `json:"status,omitempty"`
}

// NewBackupJobDto instantiates a new BackupJobDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobDto(links map[string]ResourceLink) *BackupJobDto {
	this := BackupJobDto{}
	this.Links = links
	return &this
}

// NewBackupJobDtoWithDefaults instantiates a new BackupJobDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobDtoWithDefaults() *BackupJobDto {
	this := BackupJobDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *BackupJobDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BackupJobDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupJobDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupJobDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupJobDto) SetId(v string) {
	o.Id = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *BackupJobDto) GetStarted() time.Time {
	if o == nil || isNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetStartedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Started) {
    return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *BackupJobDto) HasStarted() bool {
	if o != nil && !isNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *BackupJobDto) SetStarted(v time.Time) {
	o.Started = &v
}

// GetStopped returns the Stopped field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupJobDto) GetStopped() time.Time {
	if o == nil || isNil(o.Stopped.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Stopped.Get()
}

// GetStoppedOk returns a tuple with the Stopped field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupJobDto) GetStoppedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stopped.Get(), o.Stopped.IsSet()
}

// HasStopped returns a boolean if a field has been set.
func (o *BackupJobDto) HasStopped() bool {
	if o != nil && o.Stopped.IsSet() {
		return true
	}

	return false
}

// SetStopped gets a reference to the given NullableTime and assigns it to the Stopped field.
func (o *BackupJobDto) SetStopped(v time.Time) {
	o.Stopped.Set(&v)
}
// SetStoppedNil sets the value for Stopped to be an explicit nil
func (o *BackupJobDto) SetStoppedNil() {
	o.Stopped.Set(nil)
}

// UnsetStopped ensures that no value is present for Stopped, not even an explicit nil
func (o *BackupJobDto) UnsetStopped() {
	o.Stopped.Unset()
}

// GetHandledEvents returns the HandledEvents field value if set, zero value otherwise.
func (o *BackupJobDto) GetHandledEvents() int32 {
	if o == nil || isNil(o.HandledEvents) {
		var ret int32
		return ret
	}
	return *o.HandledEvents
}

// GetHandledEventsOk returns a tuple with the HandledEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetHandledEventsOk() (*int32, bool) {
	if o == nil || isNil(o.HandledEvents) {
    return nil, false
	}
	return o.HandledEvents, true
}

// HasHandledEvents returns a boolean if a field has been set.
func (o *BackupJobDto) HasHandledEvents() bool {
	if o != nil && !isNil(o.HandledEvents) {
		return true
	}

	return false
}

// SetHandledEvents gets a reference to the given int32 and assigns it to the HandledEvents field.
func (o *BackupJobDto) SetHandledEvents(v int32) {
	o.HandledEvents = &v
}

// GetHandledAssets returns the HandledAssets field value if set, zero value otherwise.
func (o *BackupJobDto) GetHandledAssets() int32 {
	if o == nil || isNil(o.HandledAssets) {
		var ret int32
		return ret
	}
	return *o.HandledAssets
}

// GetHandledAssetsOk returns a tuple with the HandledAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetHandledAssetsOk() (*int32, bool) {
	if o == nil || isNil(o.HandledAssets) {
    return nil, false
	}
	return o.HandledAssets, true
}

// HasHandledAssets returns a boolean if a field has been set.
func (o *BackupJobDto) HasHandledAssets() bool {
	if o != nil && !isNil(o.HandledAssets) {
		return true
	}

	return false
}

// SetHandledAssets gets a reference to the given int32 and assigns it to the HandledAssets field.
func (o *BackupJobDto) SetHandledAssets(v int32) {
	o.HandledAssets = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupJobDto) GetStatus() JobStatus {
	if o == nil || isNil(o.Status) {
		var ret JobStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobDto) GetStatusOk() (*JobStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupJobDto) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given JobStatus and assigns it to the Status field.
func (o *BackupJobDto) SetStatus(v JobStatus) {
	o.Status = &v
}

func (o BackupJobDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if o.Stopped.IsSet() {
		toSerialize["stopped"] = o.Stopped.Get()
	}
	if !isNil(o.HandledEvents) {
		toSerialize["handledEvents"] = o.HandledEvents
	}
	if !isNil(o.HandledAssets) {
		toSerialize["handledAssets"] = o.HandledAssets
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobDto struct {
	value *BackupJobDto
	isSet bool
}

func (v NullableBackupJobDto) Get() *BackupJobDto {
	return v.value
}

func (v *NullableBackupJobDto) Set(val *BackupJobDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobDto(val *BackupJobDto) *NullableBackupJobDto {
	return &NullableBackupJobDto{value: val, isSet: true}
}

func (v NullableBackupJobDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

