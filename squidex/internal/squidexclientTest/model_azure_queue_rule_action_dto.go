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

// AzureQueueRuleActionDto struct for AzureQueueRuleActionDto
type AzureQueueRuleActionDto struct {
	RuleAction
	// The connection string to the storage account.
	ConnectionString string `json:"connectionString"`
	// The name of the queue.
	Queue string `json:"queue"`
	// Leave it empty to use the full event as body.
	Payload NullableString `json:"payload,omitempty"`
}

// NewAzureQueueRuleActionDto instantiates a new AzureQueueRuleActionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureQueueRuleActionDto(connectionString string, queue string, actionType NullableString) *AzureQueueRuleActionDto {
	this := AzureQueueRuleActionDto{}
	this.ActionType = actionType
	this.ConnectionString = connectionString
	this.Queue = queue
	return &this
}

// NewAzureQueueRuleActionDtoWithDefaults instantiates a new AzureQueueRuleActionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureQueueRuleActionDtoWithDefaults() *AzureQueueRuleActionDto {
	this := AzureQueueRuleActionDto{}
	return &this
}

// GetConnectionString returns the ConnectionString field value
func (o *AzureQueueRuleActionDto) GetConnectionString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value
// and a boolean to check if the value has been set.
func (o *AzureQueueRuleActionDto) GetConnectionStringOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionString, true
}

// SetConnectionString sets field value
func (o *AzureQueueRuleActionDto) SetConnectionString(v string) {
	o.ConnectionString = v
}

// GetQueue returns the Queue field value
func (o *AzureQueueRuleActionDto) GetQueue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Queue
}

// GetQueueOk returns a tuple with the Queue field value
// and a boolean to check if the value has been set.
func (o *AzureQueueRuleActionDto) GetQueueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Queue, true
}

// SetQueue sets field value
func (o *AzureQueueRuleActionDto) SetQueue(v string) {
	o.Queue = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureQueueRuleActionDto) GetPayload() string {
	if o == nil || isNil(o.Payload.Get()) {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureQueueRuleActionDto) GetPayloadOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *AzureQueueRuleActionDto) HasPayload() bool {
	if o != nil && o.Payload.IsSet() {
		return true
	}

	return false
}

// SetPayload gets a reference to the given NullableString and assigns it to the Payload field.
func (o *AzureQueueRuleActionDto) SetPayload(v string) {
	o.Payload.Set(&v)
}
// SetPayloadNil sets the value for Payload to be an explicit nil
func (o *AzureQueueRuleActionDto) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil
func (o *AzureQueueRuleActionDto) UnsetPayload() {
	o.Payload.Unset()
}

func (o AzureQueueRuleActionDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRuleAction, errRuleAction := json.Marshal(o.RuleAction)
	if errRuleAction != nil {
		return []byte{}, errRuleAction
	}
	errRuleAction = json.Unmarshal([]byte(serializedRuleAction), &toSerialize)
	if errRuleAction != nil {
		return []byte{}, errRuleAction
	}
	if true {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if true {
		toSerialize["queue"] = o.Queue
	}
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureQueueRuleActionDto struct {
	value *AzureQueueRuleActionDto
	isSet bool
}

func (v NullableAzureQueueRuleActionDto) Get() *AzureQueueRuleActionDto {
	return v.value
}

func (v *NullableAzureQueueRuleActionDto) Set(val *AzureQueueRuleActionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureQueueRuleActionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureQueueRuleActionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureQueueRuleActionDto(val *AzureQueueRuleActionDto) *NullableAzureQueueRuleActionDto {
	return &NullableAzureQueueRuleActionDto{value: val, isSet: true}
}

func (v NullableAzureQueueRuleActionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureQueueRuleActionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

