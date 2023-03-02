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

// BulkUpdateContentsDto struct for BulkUpdateContentsDto
type BulkUpdateContentsDto struct {
	// The contents to update or insert.
	Jobs []BulkUpdateContentsJobDto `json:"jobs"`
	// True to automatically publish the content.
	// Deprecated
	Publish *bool `json:"publish,omitempty"`
	// True to turn off scripting for faster inserts. Default: true.
	DoNotScript *bool `json:"doNotScript,omitempty"`
	// True to turn off validation for faster inserts. Default: false.
	DoNotValidate *bool `json:"doNotValidate,omitempty"`
	// True to turn off validation of workflow rules. Default: false.
	DoNotValidateWorkflow *bool `json:"doNotValidateWorkflow,omitempty"`
	// True to check referrers of deleted contents.
	CheckReferrers *bool `json:"checkReferrers,omitempty"`
	// True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true.
	OptimizeValidation *bool `json:"optimizeValidation,omitempty"`
}

// NewBulkUpdateContentsDto instantiates a new BulkUpdateContentsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateContentsDto(jobs []BulkUpdateContentsJobDto) *BulkUpdateContentsDto {
	this := BulkUpdateContentsDto{}
	this.Jobs = jobs
	return &this
}

// NewBulkUpdateContentsDtoWithDefaults instantiates a new BulkUpdateContentsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateContentsDtoWithDefaults() *BulkUpdateContentsDto {
	this := BulkUpdateContentsDto{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *BulkUpdateContentsDto) GetJobs() []BulkUpdateContentsJobDto {
	if o == nil {
		var ret []BulkUpdateContentsJobDto
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetJobsOk() ([]BulkUpdateContentsJobDto, bool) {
	if o == nil {
    return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *BulkUpdateContentsDto) SetJobs(v []BulkUpdateContentsJobDto) {
	o.Jobs = v
}

// GetPublish returns the Publish field value if set, zero value otherwise.
// Deprecated
func (o *BulkUpdateContentsDto) GetPublish() bool {
	if o == nil || isNil(o.Publish) {
		var ret bool
		return ret
	}
	return *o.Publish
}

// GetPublishOk returns a tuple with the Publish field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BulkUpdateContentsDto) GetPublishOk() (*bool, bool) {
	if o == nil || isNil(o.Publish) {
    return nil, false
	}
	return o.Publish, true
}

// HasPublish returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasPublish() bool {
	if o != nil && !isNil(o.Publish) {
		return true
	}

	return false
}

// SetPublish gets a reference to the given bool and assigns it to the Publish field.
// Deprecated
func (o *BulkUpdateContentsDto) SetPublish(v bool) {
	o.Publish = &v
}

// GetDoNotScript returns the DoNotScript field value if set, zero value otherwise.
func (o *BulkUpdateContentsDto) GetDoNotScript() bool {
	if o == nil || isNil(o.DoNotScript) {
		var ret bool
		return ret
	}
	return *o.DoNotScript
}

// GetDoNotScriptOk returns a tuple with the DoNotScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetDoNotScriptOk() (*bool, bool) {
	if o == nil || isNil(o.DoNotScript) {
    return nil, false
	}
	return o.DoNotScript, true
}

// HasDoNotScript returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasDoNotScript() bool {
	if o != nil && !isNil(o.DoNotScript) {
		return true
	}

	return false
}

// SetDoNotScript gets a reference to the given bool and assigns it to the DoNotScript field.
func (o *BulkUpdateContentsDto) SetDoNotScript(v bool) {
	o.DoNotScript = &v
}

// GetDoNotValidate returns the DoNotValidate field value if set, zero value otherwise.
func (o *BulkUpdateContentsDto) GetDoNotValidate() bool {
	if o == nil || isNil(o.DoNotValidate) {
		var ret bool
		return ret
	}
	return *o.DoNotValidate
}

// GetDoNotValidateOk returns a tuple with the DoNotValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetDoNotValidateOk() (*bool, bool) {
	if o == nil || isNil(o.DoNotValidate) {
    return nil, false
	}
	return o.DoNotValidate, true
}

// HasDoNotValidate returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasDoNotValidate() bool {
	if o != nil && !isNil(o.DoNotValidate) {
		return true
	}

	return false
}

// SetDoNotValidate gets a reference to the given bool and assigns it to the DoNotValidate field.
func (o *BulkUpdateContentsDto) SetDoNotValidate(v bool) {
	o.DoNotValidate = &v
}

// GetDoNotValidateWorkflow returns the DoNotValidateWorkflow field value if set, zero value otherwise.
func (o *BulkUpdateContentsDto) GetDoNotValidateWorkflow() bool {
	if o == nil || isNil(o.DoNotValidateWorkflow) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateWorkflow
}

// GetDoNotValidateWorkflowOk returns a tuple with the DoNotValidateWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetDoNotValidateWorkflowOk() (*bool, bool) {
	if o == nil || isNil(o.DoNotValidateWorkflow) {
    return nil, false
	}
	return o.DoNotValidateWorkflow, true
}

// HasDoNotValidateWorkflow returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasDoNotValidateWorkflow() bool {
	if o != nil && !isNil(o.DoNotValidateWorkflow) {
		return true
	}

	return false
}

// SetDoNotValidateWorkflow gets a reference to the given bool and assigns it to the DoNotValidateWorkflow field.
func (o *BulkUpdateContentsDto) SetDoNotValidateWorkflow(v bool) {
	o.DoNotValidateWorkflow = &v
}

// GetCheckReferrers returns the CheckReferrers field value if set, zero value otherwise.
func (o *BulkUpdateContentsDto) GetCheckReferrers() bool {
	if o == nil || isNil(o.CheckReferrers) {
		var ret bool
		return ret
	}
	return *o.CheckReferrers
}

// GetCheckReferrersOk returns a tuple with the CheckReferrers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetCheckReferrersOk() (*bool, bool) {
	if o == nil || isNil(o.CheckReferrers) {
    return nil, false
	}
	return o.CheckReferrers, true
}

// HasCheckReferrers returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasCheckReferrers() bool {
	if o != nil && !isNil(o.CheckReferrers) {
		return true
	}

	return false
}

// SetCheckReferrers gets a reference to the given bool and assigns it to the CheckReferrers field.
func (o *BulkUpdateContentsDto) SetCheckReferrers(v bool) {
	o.CheckReferrers = &v
}

// GetOptimizeValidation returns the OptimizeValidation field value if set, zero value otherwise.
func (o *BulkUpdateContentsDto) GetOptimizeValidation() bool {
	if o == nil || isNil(o.OptimizeValidation) {
		var ret bool
		return ret
	}
	return *o.OptimizeValidation
}

// GetOptimizeValidationOk returns a tuple with the OptimizeValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateContentsDto) GetOptimizeValidationOk() (*bool, bool) {
	if o == nil || isNil(o.OptimizeValidation) {
    return nil, false
	}
	return o.OptimizeValidation, true
}

// HasOptimizeValidation returns a boolean if a field has been set.
func (o *BulkUpdateContentsDto) HasOptimizeValidation() bool {
	if o != nil && !isNil(o.OptimizeValidation) {
		return true
	}

	return false
}

// SetOptimizeValidation gets a reference to the given bool and assigns it to the OptimizeValidation field.
func (o *BulkUpdateContentsDto) SetOptimizeValidation(v bool) {
	o.OptimizeValidation = &v
}

func (o BulkUpdateContentsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	if !isNil(o.Publish) {
		toSerialize["publish"] = o.Publish
	}
	if !isNil(o.DoNotScript) {
		toSerialize["doNotScript"] = o.DoNotScript
	}
	if !isNil(o.DoNotValidate) {
		toSerialize["doNotValidate"] = o.DoNotValidate
	}
	if !isNil(o.DoNotValidateWorkflow) {
		toSerialize["doNotValidateWorkflow"] = o.DoNotValidateWorkflow
	}
	if !isNil(o.CheckReferrers) {
		toSerialize["checkReferrers"] = o.CheckReferrers
	}
	if !isNil(o.OptimizeValidation) {
		toSerialize["optimizeValidation"] = o.OptimizeValidation
	}
	return json.Marshal(toSerialize)
}

type NullableBulkUpdateContentsDto struct {
	value *BulkUpdateContentsDto
	isSet bool
}

func (v NullableBulkUpdateContentsDto) Get() *BulkUpdateContentsDto {
	return v.value
}

func (v *NullableBulkUpdateContentsDto) Set(val *BulkUpdateContentsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateContentsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateContentsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateContentsDto(val *BulkUpdateContentsDto) *NullableBulkUpdateContentsDto {
	return &NullableBulkUpdateContentsDto{value: val, isSet: true}
}

func (v NullableBulkUpdateContentsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateContentsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

