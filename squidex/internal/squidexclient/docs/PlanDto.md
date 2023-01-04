# PlanDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the plan. | 
**Name** | **string** | The name of the plan. | 
**Costs** | **string** | The monthly costs of the plan. | 
**ConfirmText** | Pointer to **NullableString** | An optional confirm text for the monthly subscription. | [optional] 
**YearlyConfirmText** | Pointer to **NullableString** | An optional confirm text for the yearly subscription. | [optional] 
**YearlyCosts** | Pointer to **NullableString** | The yearly costs of the plan. | [optional] 
**YearlyId** | Pointer to **NullableString** | The yearly id of the plan. | [optional] 
**MaxApiBytes** | Pointer to **int64** | The maximum number of API traffic. | [optional] 
**MaxApiCalls** | Pointer to **int64** | The maximum number of API calls. | [optional] 
**MaxAssetSize** | Pointer to **int64** | The maximum allowed asset size. | [optional] 
**MaxContributors** | Pointer to **int32** | The maximum number of contributors. | [optional] 

## Methods

### NewPlanDto

`func NewPlanDto(id string, name string, costs string, ) *PlanDto`

NewPlanDto instantiates a new PlanDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanDtoWithDefaults

`func NewPlanDtoWithDefaults() *PlanDto`

NewPlanDtoWithDefaults instantiates a new PlanDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PlanDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanDto) SetName(v string)`

SetName sets Name field to given value.


### GetCosts

`func (o *PlanDto) GetCosts() string`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *PlanDto) GetCostsOk() (*string, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *PlanDto) SetCosts(v string)`

SetCosts sets Costs field to given value.


### GetConfirmText

`func (o *PlanDto) GetConfirmText() string`

GetConfirmText returns the ConfirmText field if non-nil, zero value otherwise.

### GetConfirmTextOk

`func (o *PlanDto) GetConfirmTextOk() (*string, bool)`

GetConfirmTextOk returns a tuple with the ConfirmText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmText

`func (o *PlanDto) SetConfirmText(v string)`

SetConfirmText sets ConfirmText field to given value.

### HasConfirmText

`func (o *PlanDto) HasConfirmText() bool`

HasConfirmText returns a boolean if a field has been set.

### SetConfirmTextNil

`func (o *PlanDto) SetConfirmTextNil(b bool)`

 SetConfirmTextNil sets the value for ConfirmText to be an explicit nil

### UnsetConfirmText
`func (o *PlanDto) UnsetConfirmText()`

UnsetConfirmText ensures that no value is present for ConfirmText, not even an explicit nil
### GetYearlyConfirmText

`func (o *PlanDto) GetYearlyConfirmText() string`

GetYearlyConfirmText returns the YearlyConfirmText field if non-nil, zero value otherwise.

### GetYearlyConfirmTextOk

`func (o *PlanDto) GetYearlyConfirmTextOk() (*string, bool)`

GetYearlyConfirmTextOk returns a tuple with the YearlyConfirmText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyConfirmText

`func (o *PlanDto) SetYearlyConfirmText(v string)`

SetYearlyConfirmText sets YearlyConfirmText field to given value.

### HasYearlyConfirmText

`func (o *PlanDto) HasYearlyConfirmText() bool`

HasYearlyConfirmText returns a boolean if a field has been set.

### SetYearlyConfirmTextNil

`func (o *PlanDto) SetYearlyConfirmTextNil(b bool)`

 SetYearlyConfirmTextNil sets the value for YearlyConfirmText to be an explicit nil

### UnsetYearlyConfirmText
`func (o *PlanDto) UnsetYearlyConfirmText()`

UnsetYearlyConfirmText ensures that no value is present for YearlyConfirmText, not even an explicit nil
### GetYearlyCosts

`func (o *PlanDto) GetYearlyCosts() string`

GetYearlyCosts returns the YearlyCosts field if non-nil, zero value otherwise.

### GetYearlyCostsOk

`func (o *PlanDto) GetYearlyCostsOk() (*string, bool)`

GetYearlyCostsOk returns a tuple with the YearlyCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCosts

`func (o *PlanDto) SetYearlyCosts(v string)`

SetYearlyCosts sets YearlyCosts field to given value.

### HasYearlyCosts

`func (o *PlanDto) HasYearlyCosts() bool`

HasYearlyCosts returns a boolean if a field has been set.

### SetYearlyCostsNil

`func (o *PlanDto) SetYearlyCostsNil(b bool)`

 SetYearlyCostsNil sets the value for YearlyCosts to be an explicit nil

### UnsetYearlyCosts
`func (o *PlanDto) UnsetYearlyCosts()`

UnsetYearlyCosts ensures that no value is present for YearlyCosts, not even an explicit nil
### GetYearlyId

`func (o *PlanDto) GetYearlyId() string`

GetYearlyId returns the YearlyId field if non-nil, zero value otherwise.

### GetYearlyIdOk

`func (o *PlanDto) GetYearlyIdOk() (*string, bool)`

GetYearlyIdOk returns a tuple with the YearlyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyId

`func (o *PlanDto) SetYearlyId(v string)`

SetYearlyId sets YearlyId field to given value.

### HasYearlyId

`func (o *PlanDto) HasYearlyId() bool`

HasYearlyId returns a boolean if a field has been set.

### SetYearlyIdNil

`func (o *PlanDto) SetYearlyIdNil(b bool)`

 SetYearlyIdNil sets the value for YearlyId to be an explicit nil

### UnsetYearlyId
`func (o *PlanDto) UnsetYearlyId()`

UnsetYearlyId ensures that no value is present for YearlyId, not even an explicit nil
### GetMaxApiBytes

`func (o *PlanDto) GetMaxApiBytes() int64`

GetMaxApiBytes returns the MaxApiBytes field if non-nil, zero value otherwise.

### GetMaxApiBytesOk

`func (o *PlanDto) GetMaxApiBytesOk() (*int64, bool)`

GetMaxApiBytesOk returns a tuple with the MaxApiBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxApiBytes

`func (o *PlanDto) SetMaxApiBytes(v int64)`

SetMaxApiBytes sets MaxApiBytes field to given value.

### HasMaxApiBytes

`func (o *PlanDto) HasMaxApiBytes() bool`

HasMaxApiBytes returns a boolean if a field has been set.

### GetMaxApiCalls

`func (o *PlanDto) GetMaxApiCalls() int64`

GetMaxApiCalls returns the MaxApiCalls field if non-nil, zero value otherwise.

### GetMaxApiCallsOk

`func (o *PlanDto) GetMaxApiCallsOk() (*int64, bool)`

GetMaxApiCallsOk returns a tuple with the MaxApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxApiCalls

`func (o *PlanDto) SetMaxApiCalls(v int64)`

SetMaxApiCalls sets MaxApiCalls field to given value.

### HasMaxApiCalls

`func (o *PlanDto) HasMaxApiCalls() bool`

HasMaxApiCalls returns a boolean if a field has been set.

### GetMaxAssetSize

`func (o *PlanDto) GetMaxAssetSize() int64`

GetMaxAssetSize returns the MaxAssetSize field if non-nil, zero value otherwise.

### GetMaxAssetSizeOk

`func (o *PlanDto) GetMaxAssetSizeOk() (*int64, bool)`

GetMaxAssetSizeOk returns a tuple with the MaxAssetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAssetSize

`func (o *PlanDto) SetMaxAssetSize(v int64)`

SetMaxAssetSize sets MaxAssetSize field to given value.

### HasMaxAssetSize

`func (o *PlanDto) HasMaxAssetSize() bool`

HasMaxAssetSize returns a boolean if a field has been set.

### GetMaxContributors

`func (o *PlanDto) GetMaxContributors() int32`

GetMaxContributors returns the MaxContributors field if non-nil, zero value otherwise.

### GetMaxContributorsOk

`func (o *PlanDto) GetMaxContributorsOk() (*int32, bool)`

GetMaxContributorsOk returns a tuple with the MaxContributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContributors

`func (o *PlanDto) SetMaxContributors(v int32)`

SetMaxContributors sets MaxContributors field to given value.

### HasMaxContributors

`func (o *PlanDto) HasMaxContributors() bool`

HasMaxContributors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


