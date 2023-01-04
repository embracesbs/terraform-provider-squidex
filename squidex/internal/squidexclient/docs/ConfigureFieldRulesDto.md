# ConfigureFieldRulesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldRules** | Pointer to [**[]FieldRuleDto**](FieldRuleDto.md) | The field rules to configure. | [optional] 

## Methods

### NewConfigureFieldRulesDto

`func NewConfigureFieldRulesDto() *ConfigureFieldRulesDto`

NewConfigureFieldRulesDto instantiates a new ConfigureFieldRulesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureFieldRulesDtoWithDefaults

`func NewConfigureFieldRulesDtoWithDefaults() *ConfigureFieldRulesDto`

NewConfigureFieldRulesDtoWithDefaults instantiates a new ConfigureFieldRulesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldRules

`func (o *ConfigureFieldRulesDto) GetFieldRules() []FieldRuleDto`

GetFieldRules returns the FieldRules field if non-nil, zero value otherwise.

### GetFieldRulesOk

`func (o *ConfigureFieldRulesDto) GetFieldRulesOk() (*[]FieldRuleDto, bool)`

GetFieldRulesOk returns a tuple with the FieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRules

`func (o *ConfigureFieldRulesDto) SetFieldRules(v []FieldRuleDto)`

SetFieldRules sets FieldRules field to given value.

### HasFieldRules

`func (o *ConfigureFieldRulesDto) HasFieldRules() bool`

HasFieldRules returns a boolean if a field has been set.

### SetFieldRulesNil

`func (o *ConfigureFieldRulesDto) SetFieldRulesNil(b bool)`

 SetFieldRulesNil sets the value for FieldRules to be an explicit nil

### UnsetFieldRules
`func (o *ConfigureFieldRulesDto) UnsetFieldRules()`

UnsetFieldRules ensures that no value is present for FieldRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


