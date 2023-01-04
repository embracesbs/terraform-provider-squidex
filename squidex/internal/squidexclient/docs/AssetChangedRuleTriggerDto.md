# AssetChangedRuleTriggerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **NullableString** | Javascript condition when to trigger. | [optional] 

## Methods

### NewAssetChangedRuleTriggerDto

`func NewAssetChangedRuleTriggerDto() *AssetChangedRuleTriggerDto`

NewAssetChangedRuleTriggerDto instantiates a new AssetChangedRuleTriggerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetChangedRuleTriggerDtoWithDefaults

`func NewAssetChangedRuleTriggerDtoWithDefaults() *AssetChangedRuleTriggerDto`

NewAssetChangedRuleTriggerDtoWithDefaults instantiates a new AssetChangedRuleTriggerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *AssetChangedRuleTriggerDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AssetChangedRuleTriggerDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AssetChangedRuleTriggerDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AssetChangedRuleTriggerDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *AssetChangedRuleTriggerDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *AssetChangedRuleTriggerDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


