# WorkflowDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The workflow id. | [optional] 
**Name** | Pointer to **NullableString** | The name of the workflow. | [optional] 
**Steps** | [**map[string]WorkflowStepDto**](WorkflowStepDto.md) | The workflow steps. | 
**SchemaIds** | Pointer to **[]string** | The schema ids. | [optional] 
**Initial** | Pointer to **string** | The initial step. | [optional] 

## Methods

### NewWorkflowDtoAllOf

`func NewWorkflowDtoAllOf(steps map[string]WorkflowStepDto, ) *WorkflowDtoAllOf`

NewWorkflowDtoAllOf instantiates a new WorkflowDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDtoAllOfWithDefaults

`func NewWorkflowDtoAllOfWithDefaults() *WorkflowDtoAllOf`

NewWorkflowDtoAllOfWithDefaults instantiates a new WorkflowDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowDtoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowDtoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkflowDtoAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkflowDtoAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSteps

`func (o *WorkflowDtoAllOf) GetSteps() map[string]WorkflowStepDto`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowDtoAllOf) GetStepsOk() (*map[string]WorkflowStepDto, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowDtoAllOf) SetSteps(v map[string]WorkflowStepDto)`

SetSteps sets Steps field to given value.


### GetSchemaIds

`func (o *WorkflowDtoAllOf) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *WorkflowDtoAllOf) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *WorkflowDtoAllOf) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *WorkflowDtoAllOf) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *WorkflowDtoAllOf) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *WorkflowDtoAllOf) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil
### GetInitial

`func (o *WorkflowDtoAllOf) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *WorkflowDtoAllOf) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *WorkflowDtoAllOf) SetInitial(v string)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *WorkflowDtoAllOf) HasInitial() bool`

HasInitial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


