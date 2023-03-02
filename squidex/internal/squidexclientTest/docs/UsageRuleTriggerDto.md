# UsageRuleTriggerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The number of monthly api calls. | [optional] 
**NumDays** | Pointer to **NullableInt32** | The number of days to check or null for the current month. | [optional] 

## Methods

### NewUsageRuleTriggerDto

`func NewUsageRuleTriggerDto() *UsageRuleTriggerDto`

NewUsageRuleTriggerDto instantiates a new UsageRuleTriggerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageRuleTriggerDtoWithDefaults

`func NewUsageRuleTriggerDtoWithDefaults() *UsageRuleTriggerDto`

NewUsageRuleTriggerDtoWithDefaults instantiates a new UsageRuleTriggerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *UsageRuleTriggerDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UsageRuleTriggerDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UsageRuleTriggerDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UsageRuleTriggerDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNumDays

`func (o *UsageRuleTriggerDto) GetNumDays() int32`

GetNumDays returns the NumDays field if non-nil, zero value otherwise.

### GetNumDaysOk

`func (o *UsageRuleTriggerDto) GetNumDaysOk() (*int32, bool)`

GetNumDaysOk returns a tuple with the NumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDays

`func (o *UsageRuleTriggerDto) SetNumDays(v int32)`

SetNumDays sets NumDays field to given value.

### HasNumDays

`func (o *UsageRuleTriggerDto) HasNumDays() bool`

HasNumDays returns a boolean if a field has been set.

### SetNumDaysNil

`func (o *UsageRuleTriggerDto) SetNumDaysNil(b bool)`

 SetNumDaysNil sets the value for NumDays to be an explicit nil

### UnsetNumDays
`func (o *UsageRuleTriggerDto) UnsetNumDays()`

UnsetNumDays ensures that no value is present for NumDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


