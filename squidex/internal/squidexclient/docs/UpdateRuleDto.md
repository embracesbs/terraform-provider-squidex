# UpdateRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Optional rule name. | [optional] 
**Trigger** | Pointer to [**RuleTriggerDto**](RuleTriggerDto.md) |  | [optional] 
**Action** | Pointer to [**RuleAction**](RuleAction.md) |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Enable or disable the rule. | [optional] 

## Methods

### NewUpdateRuleDto

`func NewUpdateRuleDto() *UpdateRuleDto`

NewUpdateRuleDto instantiates a new UpdateRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRuleDtoWithDefaults

`func NewUpdateRuleDtoWithDefaults() *UpdateRuleDto`

NewUpdateRuleDtoWithDefaults instantiates a new UpdateRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRuleDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRuleDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRuleDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRuleDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateRuleDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateRuleDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTrigger

`func (o *UpdateRuleDto) GetTrigger() RuleTriggerDto`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *UpdateRuleDto) GetTriggerOk() (*RuleTriggerDto, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *UpdateRuleDto) SetTrigger(v RuleTriggerDto)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *UpdateRuleDto) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetAction

`func (o *UpdateRuleDto) GetAction() RuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateRuleDto) GetActionOk() (*RuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateRuleDto) SetAction(v RuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateRuleDto) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsEnabled

`func (o *UpdateRuleDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateRuleDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateRuleDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateRuleDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UpdateRuleDto) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UpdateRuleDto) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


