# BulkUpdateAssetsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]BulkUpdateAssetsJobDto**](BulkUpdateAssetsJobDto.md) | The contents to update or insert. | 
**CheckReferrers** | Pointer to **bool** | True to check referrers of deleted assets. | [optional] 
**OptimizeValidation** | Pointer to **bool** | True to turn off costly validation: Folder checks. Default: true. | [optional] 
**DoNotScript** | Pointer to **bool** | True to turn off scripting for faster inserts. Default: true. | [optional] 

## Methods

### NewBulkUpdateAssetsDto

`func NewBulkUpdateAssetsDto(jobs []BulkUpdateAssetsJobDto, ) *BulkUpdateAssetsDto`

NewBulkUpdateAssetsDto instantiates a new BulkUpdateAssetsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAssetsDtoWithDefaults

`func NewBulkUpdateAssetsDtoWithDefaults() *BulkUpdateAssetsDto`

NewBulkUpdateAssetsDtoWithDefaults instantiates a new BulkUpdateAssetsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *BulkUpdateAssetsDto) GetJobs() []BulkUpdateAssetsJobDto`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BulkUpdateAssetsDto) GetJobsOk() (*[]BulkUpdateAssetsJobDto, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BulkUpdateAssetsDto) SetJobs(v []BulkUpdateAssetsJobDto)`

SetJobs sets Jobs field to given value.


### GetCheckReferrers

`func (o *BulkUpdateAssetsDto) GetCheckReferrers() bool`

GetCheckReferrers returns the CheckReferrers field if non-nil, zero value otherwise.

### GetCheckReferrersOk

`func (o *BulkUpdateAssetsDto) GetCheckReferrersOk() (*bool, bool)`

GetCheckReferrersOk returns a tuple with the CheckReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckReferrers

`func (o *BulkUpdateAssetsDto) SetCheckReferrers(v bool)`

SetCheckReferrers sets CheckReferrers field to given value.

### HasCheckReferrers

`func (o *BulkUpdateAssetsDto) HasCheckReferrers() bool`

HasCheckReferrers returns a boolean if a field has been set.

### GetOptimizeValidation

`func (o *BulkUpdateAssetsDto) GetOptimizeValidation() bool`

GetOptimizeValidation returns the OptimizeValidation field if non-nil, zero value otherwise.

### GetOptimizeValidationOk

`func (o *BulkUpdateAssetsDto) GetOptimizeValidationOk() (*bool, bool)`

GetOptimizeValidationOk returns a tuple with the OptimizeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeValidation

`func (o *BulkUpdateAssetsDto) SetOptimizeValidation(v bool)`

SetOptimizeValidation sets OptimizeValidation field to given value.

### HasOptimizeValidation

`func (o *BulkUpdateAssetsDto) HasOptimizeValidation() bool`

HasOptimizeValidation returns a boolean if a field has been set.

### GetDoNotScript

`func (o *BulkUpdateAssetsDto) GetDoNotScript() bool`

GetDoNotScript returns the DoNotScript field if non-nil, zero value otherwise.

### GetDoNotScriptOk

`func (o *BulkUpdateAssetsDto) GetDoNotScriptOk() (*bool, bool)`

GetDoNotScriptOk returns a tuple with the DoNotScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotScript

`func (o *BulkUpdateAssetsDto) SetDoNotScript(v bool)`

SetDoNotScript sets DoNotScript field to given value.

### HasDoNotScript

`func (o *BulkUpdateAssetsDto) HasDoNotScript() bool`

HasDoNotScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


