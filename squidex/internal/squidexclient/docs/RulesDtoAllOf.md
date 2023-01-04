# RulesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]RuleDto**](RuleDto.md) | The rules. | 
**RunningRuleId** | Pointer to **NullableString** | The id of the rule that is currently rerunning. | [optional] 

## Methods

### NewRulesDtoAllOf

`func NewRulesDtoAllOf(items []RuleDto, ) *RulesDtoAllOf`

NewRulesDtoAllOf instantiates a new RulesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesDtoAllOfWithDefaults

`func NewRulesDtoAllOfWithDefaults() *RulesDtoAllOf`

NewRulesDtoAllOfWithDefaults instantiates a new RulesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RulesDtoAllOf) GetItems() []RuleDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RulesDtoAllOf) GetItemsOk() (*[]RuleDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RulesDtoAllOf) SetItems(v []RuleDto)`

SetItems sets Items field to given value.


### GetRunningRuleId

`func (o *RulesDtoAllOf) GetRunningRuleId() string`

GetRunningRuleId returns the RunningRuleId field if non-nil, zero value otherwise.

### GetRunningRuleIdOk

`func (o *RulesDtoAllOf) GetRunningRuleIdOk() (*string, bool)`

GetRunningRuleIdOk returns a tuple with the RunningRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningRuleId

`func (o *RulesDtoAllOf) SetRunningRuleId(v string)`

SetRunningRuleId sets RunningRuleId field to given value.

### HasRunningRuleId

`func (o *RulesDtoAllOf) HasRunningRuleId() bool`

HasRunningRuleId returns a boolean if a field has been set.

### SetRunningRuleIdNil

`func (o *RulesDtoAllOf) SetRunningRuleIdNil(b bool)`

 SetRunningRuleIdNil sets the value for RunningRuleId to be an explicit nil

### UnsetRunningRuleId
`func (o *RulesDtoAllOf) UnsetRunningRuleId()`

UnsetRunningRuleId ensures that no value is present for RunningRuleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


