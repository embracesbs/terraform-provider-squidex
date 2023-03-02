# BulkUpdateContentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]BulkUpdateContentsJobDto**](BulkUpdateContentsJobDto.md) | The contents to update or insert. | 
**Publish** | Pointer to **bool** | True to automatically publish the content. | [optional] 
**DoNotScript** | Pointer to **bool** | True to turn off scripting for faster inserts. Default: true. | [optional] 
**DoNotValidate** | Pointer to **bool** | True to turn off validation for faster inserts. Default: false. | [optional] 
**DoNotValidateWorkflow** | Pointer to **bool** | True to turn off validation of workflow rules. Default: false. | [optional] 
**CheckReferrers** | Pointer to **bool** | True to check referrers of deleted contents. | [optional] 
**OptimizeValidation** | Pointer to **bool** | True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true. | [optional] 

## Methods

### NewBulkUpdateContentsDto

`func NewBulkUpdateContentsDto(jobs []BulkUpdateContentsJobDto, ) *BulkUpdateContentsDto`

NewBulkUpdateContentsDto instantiates a new BulkUpdateContentsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateContentsDtoWithDefaults

`func NewBulkUpdateContentsDtoWithDefaults() *BulkUpdateContentsDto`

NewBulkUpdateContentsDtoWithDefaults instantiates a new BulkUpdateContentsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *BulkUpdateContentsDto) GetJobs() []BulkUpdateContentsJobDto`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BulkUpdateContentsDto) GetJobsOk() (*[]BulkUpdateContentsJobDto, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BulkUpdateContentsDto) SetJobs(v []BulkUpdateContentsJobDto)`

SetJobs sets Jobs field to given value.


### GetPublish

`func (o *BulkUpdateContentsDto) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *BulkUpdateContentsDto) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *BulkUpdateContentsDto) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *BulkUpdateContentsDto) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetDoNotScript

`func (o *BulkUpdateContentsDto) GetDoNotScript() bool`

GetDoNotScript returns the DoNotScript field if non-nil, zero value otherwise.

### GetDoNotScriptOk

`func (o *BulkUpdateContentsDto) GetDoNotScriptOk() (*bool, bool)`

GetDoNotScriptOk returns a tuple with the DoNotScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotScript

`func (o *BulkUpdateContentsDto) SetDoNotScript(v bool)`

SetDoNotScript sets DoNotScript field to given value.

### HasDoNotScript

`func (o *BulkUpdateContentsDto) HasDoNotScript() bool`

HasDoNotScript returns a boolean if a field has been set.

### GetDoNotValidate

`func (o *BulkUpdateContentsDto) GetDoNotValidate() bool`

GetDoNotValidate returns the DoNotValidate field if non-nil, zero value otherwise.

### GetDoNotValidateOk

`func (o *BulkUpdateContentsDto) GetDoNotValidateOk() (*bool, bool)`

GetDoNotValidateOk returns a tuple with the DoNotValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidate

`func (o *BulkUpdateContentsDto) SetDoNotValidate(v bool)`

SetDoNotValidate sets DoNotValidate field to given value.

### HasDoNotValidate

`func (o *BulkUpdateContentsDto) HasDoNotValidate() bool`

HasDoNotValidate returns a boolean if a field has been set.

### GetDoNotValidateWorkflow

`func (o *BulkUpdateContentsDto) GetDoNotValidateWorkflow() bool`

GetDoNotValidateWorkflow returns the DoNotValidateWorkflow field if non-nil, zero value otherwise.

### GetDoNotValidateWorkflowOk

`func (o *BulkUpdateContentsDto) GetDoNotValidateWorkflowOk() (*bool, bool)`

GetDoNotValidateWorkflowOk returns a tuple with the DoNotValidateWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateWorkflow

`func (o *BulkUpdateContentsDto) SetDoNotValidateWorkflow(v bool)`

SetDoNotValidateWorkflow sets DoNotValidateWorkflow field to given value.

### HasDoNotValidateWorkflow

`func (o *BulkUpdateContentsDto) HasDoNotValidateWorkflow() bool`

HasDoNotValidateWorkflow returns a boolean if a field has been set.

### GetCheckReferrers

`func (o *BulkUpdateContentsDto) GetCheckReferrers() bool`

GetCheckReferrers returns the CheckReferrers field if non-nil, zero value otherwise.

### GetCheckReferrersOk

`func (o *BulkUpdateContentsDto) GetCheckReferrersOk() (*bool, bool)`

GetCheckReferrersOk returns a tuple with the CheckReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckReferrers

`func (o *BulkUpdateContentsDto) SetCheckReferrers(v bool)`

SetCheckReferrers sets CheckReferrers field to given value.

### HasCheckReferrers

`func (o *BulkUpdateContentsDto) HasCheckReferrers() bool`

HasCheckReferrers returns a boolean if a field has been set.

### GetOptimizeValidation

`func (o *BulkUpdateContentsDto) GetOptimizeValidation() bool`

GetOptimizeValidation returns the OptimizeValidation field if non-nil, zero value otherwise.

### GetOptimizeValidationOk

`func (o *BulkUpdateContentsDto) GetOptimizeValidationOk() (*bool, bool)`

GetOptimizeValidationOk returns a tuple with the OptimizeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeValidation

`func (o *BulkUpdateContentsDto) SetOptimizeValidation(v bool)`

SetOptimizeValidation sets OptimizeValidation field to given value.

### HasOptimizeValidation

`func (o *BulkUpdateContentsDto) HasOptimizeValidation() bool`

HasOptimizeValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


