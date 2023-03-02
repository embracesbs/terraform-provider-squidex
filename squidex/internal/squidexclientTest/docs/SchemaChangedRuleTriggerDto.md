# SchemaChangedRuleTriggerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **NullableString** | Javascript condition when to trigger. | [optional] 

## Methods

### NewSchemaChangedRuleTriggerDto

`func NewSchemaChangedRuleTriggerDto() *SchemaChangedRuleTriggerDto`

NewSchemaChangedRuleTriggerDto instantiates a new SchemaChangedRuleTriggerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaChangedRuleTriggerDtoWithDefaults

`func NewSchemaChangedRuleTriggerDtoWithDefaults() *SchemaChangedRuleTriggerDto`

NewSchemaChangedRuleTriggerDtoWithDefaults instantiates a new SchemaChangedRuleTriggerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *SchemaChangedRuleTriggerDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SchemaChangedRuleTriggerDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SchemaChangedRuleTriggerDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SchemaChangedRuleTriggerDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *SchemaChangedRuleTriggerDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *SchemaChangedRuleTriggerDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

