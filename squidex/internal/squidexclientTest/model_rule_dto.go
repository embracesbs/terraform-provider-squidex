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

// RuleDto struct for RuleDto
type RuleDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the rule.
	Id *string `json:"id,omitempty"`
	// The user that has created the rule.
	CreatedBy string `json:"createdBy"`
	// The user that has updated the rule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The date and time when the rule has been created.
	Created *time.Time `json:"created,omitempty"`
	// The date and time when the rule has been modified last.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// The version of the rule.
	Version *int64 `json:"version,omitempty"`
	// Determines if the rule is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Optional rule name.
	Name NullableString `json:"name,omitempty"`
	Trigger RuleTriggerDto `json:"trigger"`
	Action RuleAction `json:"action"`
	// The number of completed executions.
	NumSucceeded *int32 `json:"numSucceeded,omitempty"`
	// The number of failed executions.
	NumFailed *int32 `json:"numFailed,omitempty"`
	// The date and time when the rule was executed the last time.
	LastExecuted NullableTime `json:"lastExecuted,omitempty"`
}

// NewRuleDto instantiates a new RuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDto(links map[string]ResourceLink, createdBy string, lastModifiedBy string, trigger RuleTriggerDto, action RuleAction) *RuleDto {
	this := RuleDto{}
	this.Links = links
	this.CreatedBy = createdBy
	this.LastModifiedBy = lastModifiedBy
	this.Trigger = trigger
	this.Action = action
	return &this
}

// NewRuleDtoWithDefaults instantiates a new RuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDtoWithDefaults() *RuleDto {
	this := RuleDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *RuleDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RuleDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleDto) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RuleDto) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RuleDto) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLastModifiedBy returns the LastModifiedBy field value
func (o *RuleDto) GetLastModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value
func (o *RuleDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RuleDto) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RuleDto) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RuleDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *RuleDto) GetLastModified() time.Time {
	if o == nil || isNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModified) {
    return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *RuleDto) HasLastModified() bool {
	if o != nil && !isNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *RuleDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RuleDto) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RuleDto) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *RuleDto) SetVersion(v int64) {
	o.Version = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *RuleDto) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *RuleDto) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *RuleDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleDto) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RuleDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RuleDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RuleDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RuleDto) UnsetName() {
	o.Name.Unset()
}

// GetTrigger returns the Trigger field value
func (o *RuleDto) GetTrigger() RuleTriggerDto {
	if o == nil {
		var ret RuleTriggerDto
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetTriggerOk() (*RuleTriggerDto, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *RuleDto) SetTrigger(v RuleTriggerDto) {
	o.Trigger = v
}

// GetAction returns the Action field value
func (o *RuleDto) GetAction() RuleAction {
	if o == nil {
		var ret RuleAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetActionOk() (*RuleAction, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RuleDto) SetAction(v RuleAction) {
	o.Action = v
}

// GetNumSucceeded returns the NumSucceeded field value if set, zero value otherwise.
func (o *RuleDto) GetNumSucceeded() int32 {
	if o == nil || isNil(o.NumSucceeded) {
		var ret int32
		return ret
	}
	return *o.NumSucceeded
}

// GetNumSucceededOk returns a tuple with the NumSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetNumSucceededOk() (*int32, bool) {
	if o == nil || isNil(o.NumSucceeded) {
    return nil, false
	}
	return o.NumSucceeded, true
}

// HasNumSucceeded returns a boolean if a field has been set.
func (o *RuleDto) HasNumSucceeded() bool {
	if o != nil && !isNil(o.NumSucceeded) {
		return true
	}

	return false
}

// SetNumSucceeded gets a reference to the given int32 and assigns it to the NumSucceeded field.
func (o *RuleDto) SetNumSucceeded(v int32) {
	o.NumSucceeded = &v
}

// GetNumFailed returns the NumFailed field value if set, zero value otherwise.
func (o *RuleDto) GetNumFailed() int32 {
	if o == nil || isNil(o.NumFailed) {
		var ret int32
		return ret
	}
	return *o.NumFailed
}

// GetNumFailedOk returns a tuple with the NumFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetNumFailedOk() (*int32, bool) {
	if o == nil || isNil(o.NumFailed) {
    return nil, false
	}
	return o.NumFailed, true
}

// HasNumFailed returns a boolean if a field has been set.
func (o *RuleDto) HasNumFailed() bool {
	if o != nil && !isNil(o.NumFailed) {
		return true
	}

	return false
}

// SetNumFailed gets a reference to the given int32 and assigns it to the NumFailed field.
func (o *RuleDto) SetNumFailed(v int32) {
	o.NumFailed = &v
}

// GetLastExecuted returns the LastExecuted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleDto) GetLastExecuted() time.Time {
	if o == nil || isNil(o.LastExecuted.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastExecuted.Get()
}

// GetLastExecutedOk returns a tuple with the LastExecuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleDto) GetLastExecutedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastExecuted.Get(), o.LastExecuted.IsSet()
}

// HasLastExecuted returns a boolean if a field has been set.
func (o *RuleDto) HasLastExecuted() bool {
	if o != nil && o.LastExecuted.IsSet() {
		return true
	}

	return false
}

// SetLastExecuted gets a reference to the given NullableTime and assigns it to the LastExecuted field.
func (o *RuleDto) SetLastExecuted(v time.Time) {
	o.LastExecuted.Set(&v)
}
// SetLastExecutedNil sets the value for LastExecuted to be an explicit nil
func (o *RuleDto) SetLastExecutedNil() {
	o.LastExecuted.Set(nil)
}

// UnsetLastExecuted ensures that no value is present for LastExecuted, not even an explicit nil
func (o *RuleDto) UnsetLastExecuted() {
	o.LastExecuted.Unset()
}

func (o RuleDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.NumSucceeded) {
		toSerialize["numSucceeded"] = o.NumSucceeded
	}
	if !isNil(o.NumFailed) {
		toSerialize["numFailed"] = o.NumFailed
	}
	if o.LastExecuted.IsSet() {
		toSerialize["lastExecuted"] = o.LastExecuted.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRuleDto struct {
	value *RuleDto
	isSet bool
}

func (v NullableRuleDto) Get() *RuleDto {
	return v.value
}

func (v *NullableRuleDto) Set(val *RuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDto(val *RuleDto) *NullableRuleDto {
	return &NullableRuleDto{value: val, isSet: true}
}

func (v NullableRuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

