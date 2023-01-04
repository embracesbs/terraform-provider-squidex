# WorkflowsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]WorkflowDto**](WorkflowDto.md) | The workflow. | 
**Errors** | **[]string** | The errros that should be fixed. | 

## Methods

### NewWorkflowsDto

`func NewWorkflowsDto(links map[string]ResourceLink, items []WorkflowDto, errors []string, ) *WorkflowsDto`

NewWorkflowsDto instantiates a new WorkflowsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDtoWithDefaults

`func NewWorkflowsDtoWithDefaults() *WorkflowsDto`

NewWorkflowsDtoWithDefaults instantiates a new WorkflowsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WorkflowsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkflowsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkflowsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *WorkflowsDto) GetItems() []WorkflowDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkflowsDto) GetItemsOk() (*[]WorkflowDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkflowsDto) SetItems(v []WorkflowDto)`

SetItems sets Items field to given value.


### GetErrors

`func (o *WorkflowsDto) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WorkflowsDto) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WorkflowsDto) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


