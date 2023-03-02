# FieldRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**FieldRuleAction**](FieldRuleAction.md) |  | 
**Field** | **string** | The field to update. | 
**Condition** | Pointer to **NullableString** | The condition. | [optional] 

## Methods

### NewFieldRuleDto

`func NewFieldRuleDto(action FieldRuleAction, field string, ) *FieldRuleDto`

NewFieldRuleDto instantiates a new FieldRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldRuleDtoWithDefaults

`func NewFieldRuleDtoWithDefaults() *FieldRuleDto`

NewFieldRuleDtoWithDefaults instantiates a new FieldRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FieldRuleDto) GetAction() FieldRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FieldRuleDto) GetActionOk() (*FieldRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FieldRuleDto) SetAction(v FieldRuleAction)`

SetAction sets Action field to given value.


### GetField

`func (o *FieldRuleDto) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldRuleDto) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldRuleDto) SetField(v string)`

SetField sets Field field to given value.


### GetCondition

`func (o *FieldRuleDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *FieldRuleDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *FieldRuleDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *FieldRuleDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *FieldRuleDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *FieldRuleDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


