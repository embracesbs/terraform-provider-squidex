# UpdateWorkflowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the workflow. | [optional] 
**Steps** | [**map[string]WorkflowStepDto**](WorkflowStepDto.md) | The workflow steps. | 
**SchemaIds** | Pointer to **[]string** | The schema ids. | [optional] 
**Initial** | **string** | The initial step. | 

## Methods

### NewUpdateWorkflowDto

`func NewUpdateWorkflowDto(steps map[string]WorkflowStepDto, initial string, ) *UpdateWorkflowDto`

NewUpdateWorkflowDto instantiates a new UpdateWorkflowDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflowDtoWithDefaults

`func NewUpdateWorkflowDtoWithDefaults() *UpdateWorkflowDto`

NewUpdateWorkflowDtoWithDefaults instantiates a new UpdateWorkflowDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWorkflowDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkflowDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkflowDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWorkflowDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateWorkflowDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateWorkflowDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSteps

`func (o *UpdateWorkflowDto) GetSteps() map[string]WorkflowStepDto`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *UpdateWorkflowDto) GetStepsOk() (*map[string]WorkflowStepDto, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *UpdateWorkflowDto) SetSteps(v map[string]WorkflowStepDto)`

SetSteps sets Steps field to given value.


### GetSchemaIds

`func (o *UpdateWorkflowDto) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *UpdateWorkflowDto) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *UpdateWorkflowDto) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *UpdateWorkflowDto) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *UpdateWorkflowDto) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *UpdateWorkflowDto) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil
### GetInitial

`func (o *UpdateWorkflowDto) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *UpdateWorkflowDto) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *UpdateWorkflowDto) SetInitial(v string)`

SetInitial sets Initial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


