# WorkflowStepDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transitions** | [**map[string]WorkflowTransitionDto**](WorkflowTransitionDto.md) | The transitions. | 
**Color** | Pointer to **NullableString** | The optional color. | [optional] 
**Validate** | Pointer to **bool** | True if the content should be validated when moving to this step. | [optional] 
**NoUpdate** | Pointer to **bool** | Indicates if updates should not be allowed. | [optional] 
**NoUpdateExpression** | Pointer to **NullableString** | Optional expression that must evaluate to true when you want to prevent updates. | [optional] 
**NoUpdateRoles** | Pointer to **[]string** | Optional list of roles to restrict the updates for users with these roles. | [optional] 

## Methods

### NewWorkflowStepDto

`func NewWorkflowStepDto(transitions map[string]WorkflowTransitionDto, ) *WorkflowStepDto`

NewWorkflowStepDto instantiates a new WorkflowStepDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepDtoWithDefaults

`func NewWorkflowStepDtoWithDefaults() *WorkflowStepDto`

NewWorkflowStepDtoWithDefaults instantiates a new WorkflowStepDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransitions

`func (o *WorkflowStepDto) GetTransitions() map[string]WorkflowTransitionDto`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowStepDto) GetTransitionsOk() (*map[string]WorkflowTransitionDto, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowStepDto) SetTransitions(v map[string]WorkflowTransitionDto)`

SetTransitions sets Transitions field to given value.


### GetColor

`func (o *WorkflowStepDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkflowStepDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkflowStepDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WorkflowStepDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *WorkflowStepDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *WorkflowStepDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetValidate

`func (o *WorkflowStepDto) GetValidate() bool`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *WorkflowStepDto) GetValidateOk() (*bool, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *WorkflowStepDto) SetValidate(v bool)`

SetValidate sets Validate field to given value.

### HasValidate

`func (o *WorkflowStepDto) HasValidate() bool`

HasValidate returns a boolean if a field has been set.

### GetNoUpdate

`func (o *WorkflowStepDto) GetNoUpdate() bool`

GetNoUpdate returns the NoUpdate field if non-nil, zero value otherwise.

### GetNoUpdateOk

`func (o *WorkflowStepDto) GetNoUpdateOk() (*bool, bool)`

GetNoUpdateOk returns a tuple with the NoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUpdate

`func (o *WorkflowStepDto) SetNoUpdate(v bool)`

SetNoUpdate sets NoUpdate field to given value.

### HasNoUpdate

`func (o *WorkflowStepDto) HasNoUpdate() bool`

HasNoUpdate returns a boolean if a field has been set.

### GetNoUpdateExpression

`func (o *WorkflowStepDto) GetNoUpdateExpression() string`

GetNoUpdateExpression returns the NoUpdateExpression field if non-nil, zero value otherwise.

### GetNoUpdateExpressionOk

`func (o *WorkflowStepDto) GetNoUpdateExpressionOk() (*string, bool)`

GetNoUpdateExpressionOk returns a tuple with the NoUpdateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUpdateExpression

`func (o *WorkflowStepDto) SetNoUpdateExpression(v string)`

SetNoUpdateExpression sets NoUpdateExpression field to given value.

### HasNoUpdateExpression

`func (o *WorkflowStepDto) HasNoUpdateExpression() bool`

HasNoUpdateExpression returns a boolean if a field has been set.

### SetNoUpdateExpressionNil

`func (o *WorkflowStepDto) SetNoUpdateExpressionNil(b bool)`

 SetNoUpdateExpressionNil sets the value for NoUpdateExpression to be an explicit nil

### UnsetNoUpdateExpression
`func (o *WorkflowStepDto) UnsetNoUpdateExpression()`

UnsetNoUpdateExpression ensures that no value is present for NoUpdateExpression, not even an explicit nil
### GetNoUpdateRoles

`func (o *WorkflowStepDto) GetNoUpdateRoles() []string`

GetNoUpdateRoles returns the NoUpdateRoles field if non-nil, zero value otherwise.

### GetNoUpdateRolesOk

`func (o *WorkflowStepDto) GetNoUpdateRolesOk() (*[]string, bool)`

GetNoUpdateRolesOk returns a tuple with the NoUpdateRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUpdateRoles

`func (o *WorkflowStepDto) SetNoUpdateRoles(v []string)`

SetNoUpdateRoles sets NoUpdateRoles field to given value.

### HasNoUpdateRoles

`func (o *WorkflowStepDto) HasNoUpdateRoles() bool`

HasNoUpdateRoles returns a boolean if a field has been set.

### SetNoUpdateRolesNil

`func (o *WorkflowStepDto) SetNoUpdateRolesNil(b bool)`

 SetNoUpdateRolesNil sets the value for NoUpdateRoles to be an explicit nil

### UnsetNoUpdateRoles
`func (o *WorkflowStepDto) UnsetNoUpdateRoles()`

UnsetNoUpdateRoles ensures that no value is present for NoUpdateRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


