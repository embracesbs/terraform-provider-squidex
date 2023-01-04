# CreateRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**RuleTriggerDto**](RuleTriggerDto.md) |  | 
**Action** | [**RuleAction**](RuleAction.md) |  | 

## Methods

### NewCreateRuleDto

`func NewCreateRuleDto(trigger RuleTriggerDto, action RuleAction, ) *CreateRuleDto`

NewCreateRuleDto instantiates a new CreateRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleDtoWithDefaults

`func NewCreateRuleDtoWithDefaults() *CreateRuleDto`

NewCreateRuleDtoWithDefaults instantiates a new CreateRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *CreateRuleDto) GetTrigger() RuleTriggerDto`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateRuleDto) GetTriggerOk() (*RuleTriggerDto, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateRuleDto) SetTrigger(v RuleTriggerDto)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *CreateRuleDto) GetAction() RuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateRuleDto) GetActionOk() (*RuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateRuleDto) SetAction(v RuleAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


