# WorkflowsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WorkflowDto**](WorkflowDto.md) | The workflow. | 
**Errors** | **[]string** | The errros that should be fixed. | 

## Methods

### NewWorkflowsDtoAllOf

`func NewWorkflowsDtoAllOf(items []WorkflowDto, errors []string, ) *WorkflowsDtoAllOf`

NewWorkflowsDtoAllOf instantiates a new WorkflowsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDtoAllOfWithDefaults

`func NewWorkflowsDtoAllOfWithDefaults() *WorkflowsDtoAllOf`

NewWorkflowsDtoAllOfWithDefaults instantiates a new WorkflowsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WorkflowsDtoAllOf) GetItems() []WorkflowDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkflowsDtoAllOf) GetItemsOk() (*[]WorkflowDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkflowsDtoAllOf) SetItems(v []WorkflowDto)`

SetItems sets Items field to given value.


### GetErrors

`func (o *WorkflowsDtoAllOf) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WorkflowsDtoAllOf) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WorkflowsDtoAllOf) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


