# WorkflowTransitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **NullableString** | The optional expression. | [optional] 
**Roles** | Pointer to **[]string** | The optional restricted role. | [optional] 

## Methods

### NewWorkflowTransitionDto

`func NewWorkflowTransitionDto() *WorkflowTransitionDto`

NewWorkflowTransitionDto instantiates a new WorkflowTransitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionDtoWithDefaults

`func NewWorkflowTransitionDtoWithDefaults() *WorkflowTransitionDto`

NewWorkflowTransitionDtoWithDefaults instantiates a new WorkflowTransitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *WorkflowTransitionDto) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *WorkflowTransitionDto) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *WorkflowTransitionDto) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *WorkflowTransitionDto) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *WorkflowTransitionDto) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *WorkflowTransitionDto) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil
### GetRoles

`func (o *WorkflowTransitionDto) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *WorkflowTransitionDto) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *WorkflowTransitionDto) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *WorkflowTransitionDto) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *WorkflowTransitionDto) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *WorkflowTransitionDto) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


