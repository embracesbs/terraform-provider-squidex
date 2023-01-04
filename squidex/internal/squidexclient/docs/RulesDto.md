# RulesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]RuleDto**](RuleDto.md) | The rules. | 
**RunningRuleId** | Pointer to **NullableString** | The id of the rule that is currently rerunning. | [optional] 

## Methods

### NewRulesDto

`func NewRulesDto(links map[string]ResourceLink, items []RuleDto, ) *RulesDto`

NewRulesDto instantiates a new RulesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesDtoWithDefaults

`func NewRulesDtoWithDefaults() *RulesDto`

NewRulesDtoWithDefaults instantiates a new RulesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RulesDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RulesDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RulesDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *RulesDto) GetItems() []RuleDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RulesDto) GetItemsOk() (*[]RuleDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RulesDto) SetItems(v []RuleDto)`

SetItems sets Items field to given value.


### GetRunningRuleId

`func (o *RulesDto) GetRunningRuleId() string`

GetRunningRuleId returns the RunningRuleId field if non-nil, zero value otherwise.

### GetRunningRuleIdOk

`func (o *RulesDto) GetRunningRuleIdOk() (*string, bool)`

GetRunningRuleIdOk returns a tuple with the RunningRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningRuleId

`func (o *RulesDto) SetRunningRuleId(v string)`

SetRunningRuleId sets RunningRuleId field to given value.

### HasRunningRuleId

`func (o *RulesDto) HasRunningRuleId() bool`

HasRunningRuleId returns a boolean if a field has been set.

### SetRunningRuleIdNil

`func (o *RulesDto) SetRunningRuleIdNil(b bool)`

 SetRunningRuleIdNil sets the value for RunningRuleId to be an explicit nil

### UnsetRunningRuleId
`func (o *RulesDto) UnsetRunningRuleId()`

UnsetRunningRuleId ensures that no value is present for RunningRuleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


