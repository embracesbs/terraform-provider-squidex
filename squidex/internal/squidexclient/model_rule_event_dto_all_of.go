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

// RuleEventDtoAllOf struct for RuleEventDtoAllOf
type RuleEventDtoAllOf struct {
	// The id of the event.
	Id *string `json:"id,omitempty"`
	// The time when the event has been created.
	Created *time.Time `json:"created,omitempty"`
	// The description.
	Description string `json:"description"`
	// The name of the event.
	EventName string `json:"eventName"`
	// The last dump.
	LastDump NullableString `json:"lastDump,omitempty"`
	// The number of calls.
	NumCalls *int32 `json:"numCalls,omitempty"`
	// The next attempt.
	NextAttempt NullableTime `json:"nextAttempt,omitempty"`
	Result *RuleResult `json:"result,omitempty"`
	JobResult *RuleJobResult `json:"jobResult,omitempty"`
}

// NewRuleEventDtoAllOf instantiates a new RuleEventDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleEventDtoAllOf(description string, eventName string) *RuleEventDtoAllOf {
	this := RuleEventDtoAllOf{}
	this.Description = description
	this.EventName = eventName
	return &this
}

// NewRuleEventDtoAllOfWithDefaults instantiates a new RuleEventDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleEventDtoAllOfWithDefaults() *RuleEventDtoAllOf {
	this := RuleEventDtoAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleEventDtoAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleEventDtoAllOf) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RuleEventDtoAllOf) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RuleEventDtoAllOf) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value
func (o *RuleEventDtoAllOf) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RuleEventDtoAllOf) SetDescription(v string) {
	o.Description = v
}

// GetEventName returns the EventName field value
func (o *RuleEventDtoAllOf) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetEventNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *RuleEventDtoAllOf) SetEventName(v string) {
	o.EventName = v
}

// GetLastDump returns the LastDump field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleEventDtoAllOf) GetLastDump() string {
	if o == nil || isNil(o.LastDump.Get()) {
		var ret string
		return ret
	}
	return *o.LastDump.Get()
}

// GetLastDumpOk returns a tuple with the LastDump field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleEventDtoAllOf) GetLastDumpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastDump.Get(), o.LastDump.IsSet()
}

// HasLastDump returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasLastDump() bool {
	if o != nil && o.LastDump.IsSet() {
		return true
	}

	return false
}

// SetLastDump gets a reference to the given NullableString and assigns it to the LastDump field.
func (o *RuleEventDtoAllOf) SetLastDump(v string) {
	o.LastDump.Set(&v)
}
// SetLastDumpNil sets the value for LastDump to be an explicit nil
func (o *RuleEventDtoAllOf) SetLastDumpNil() {
	o.LastDump.Set(nil)
}

// UnsetLastDump ensures that no value is present for LastDump, not even an explicit nil
func (o *RuleEventDtoAllOf) UnsetLastDump() {
	o.LastDump.Unset()
}

// GetNumCalls returns the NumCalls field value if set, zero value otherwise.
func (o *RuleEventDtoAllOf) GetNumCalls() int32 {
	if o == nil || isNil(o.NumCalls) {
		var ret int32
		return ret
	}
	return *o.NumCalls
}

// GetNumCallsOk returns a tuple with the NumCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetNumCallsOk() (*int32, bool) {
	if o == nil || isNil(o.NumCalls) {
    return nil, false
	}
	return o.NumCalls, true
}

// HasNumCalls returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasNumCalls() bool {
	if o != nil && !isNil(o.NumCalls) {
		return true
	}

	return false
}

// SetNumCalls gets a reference to the given int32 and assigns it to the NumCalls field.
func (o *RuleEventDtoAllOf) SetNumCalls(v int32) {
	o.NumCalls = &v
}

// GetNextAttempt returns the NextAttempt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleEventDtoAllOf) GetNextAttempt() time.Time {
	if o == nil || isNil(o.NextAttempt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextAttempt.Get()
}

// GetNextAttemptOk returns a tuple with the NextAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleEventDtoAllOf) GetNextAttemptOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextAttempt.Get(), o.NextAttempt.IsSet()
}

// HasNextAttempt returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasNextAttempt() bool {
	if o != nil && o.NextAttempt.IsSet() {
		return true
	}

	return false
}

// SetNextAttempt gets a reference to the given NullableTime and assigns it to the NextAttempt field.
func (o *RuleEventDtoAllOf) SetNextAttempt(v time.Time) {
	o.NextAttempt.Set(&v)
}
// SetNextAttemptNil sets the value for NextAttempt to be an explicit nil
func (o *RuleEventDtoAllOf) SetNextAttemptNil() {
	o.NextAttempt.Set(nil)
}

// UnsetNextAttempt ensures that no value is present for NextAttempt, not even an explicit nil
func (o *RuleEventDtoAllOf) UnsetNextAttempt() {
	o.NextAttempt.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RuleEventDtoAllOf) GetResult() RuleResult {
	if o == nil || isNil(o.Result) {
		var ret RuleResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetResultOk() (*RuleResult, bool) {
	if o == nil || isNil(o.Result) {
    return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RuleResult and assigns it to the Result field.
func (o *RuleEventDtoAllOf) SetResult(v RuleResult) {
	o.Result = &v
}

// GetJobResult returns the JobResult field value if set, zero value otherwise.
func (o *RuleEventDtoAllOf) GetJobResult() RuleJobResult {
	if o == nil || isNil(o.JobResult) {
		var ret RuleJobResult
		return ret
	}
	return *o.JobResult
}

// GetJobResultOk returns a tuple with the JobResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleEventDtoAllOf) GetJobResultOk() (*RuleJobResult, bool) {
	if o == nil || isNil(o.JobResult) {
    return nil, false
	}
	return o.JobResult, true
}

// HasJobResult returns a boolean if a field has been set.
func (o *RuleEventDtoAllOf) HasJobResult() bool {
	if o != nil && !isNil(o.JobResult) {
		return true
	}

	return false
}

// SetJobResult gets a reference to the given RuleJobResult and assigns it to the JobResult field.
func (o *RuleEventDtoAllOf) SetJobResult(v RuleJobResult) {
	o.JobResult = &v
}

func (o RuleEventDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if o.LastDump.IsSet() {
		toSerialize["lastDump"] = o.LastDump.Get()
	}
	if !isNil(o.NumCalls) {
		toSerialize["numCalls"] = o.NumCalls
	}
	if o.NextAttempt.IsSet() {
		toSerialize["nextAttempt"] = o.NextAttempt.Get()
	}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !isNil(o.JobResult) {
		toSerialize["jobResult"] = o.JobResult
	}
	return json.Marshal(toSerialize)
}

type NullableRuleEventDtoAllOf struct {
	value *RuleEventDtoAllOf
	isSet bool
}

func (v NullableRuleEventDtoAllOf) Get() *RuleEventDtoAllOf {
	return v.value
}

func (v *NullableRuleEventDtoAllOf) Set(val *RuleEventDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleEventDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleEventDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleEventDtoAllOf(val *RuleEventDtoAllOf) *NullableRuleEventDtoAllOf {
	return &NullableRuleEventDtoAllOf{value: val, isSet: true}
}

func (v NullableRuleEventDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleEventDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


