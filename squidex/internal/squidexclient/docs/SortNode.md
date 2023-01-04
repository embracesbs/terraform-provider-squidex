# SortNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **[]string** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] 

## Methods

### NewSortNode

`func NewSortNode() *SortNode`

NewSortNode instantiates a new SortNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortNodeWithDefaults

`func NewSortNodeWithDefaults() *SortNode`

NewSortNodeWithDefaults instantiates a new SortNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SortNode) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SortNode) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SortNode) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SortNode) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetOrder

`func (o *SortNode) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SortNode) GetOrderOk() (*SortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SortNode) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SortNode) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


