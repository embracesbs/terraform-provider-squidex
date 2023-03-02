# QueryJsonDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **interface{}** |  | [optional] 
**FullText** | Pointer to **NullableString** |  | [optional] 
**Skip** | Pointer to **int64** |  | [optional] 
**Take** | Pointer to **int64** |  | [optional] 
**Top** | Pointer to **int64** |  | [optional] 
**Sort** | Pointer to [**[]SortNode**](SortNode.md) |  | [optional] 

## Methods

### NewQueryJsonDto

`func NewQueryJsonDto() *QueryJsonDto`

NewQueryJsonDto instantiates a new QueryJsonDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryJsonDtoWithDefaults

`func NewQueryJsonDtoWithDefaults() *QueryJsonDto`

NewQueryJsonDtoWithDefaults instantiates a new QueryJsonDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *QueryJsonDto) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryJsonDto) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryJsonDto) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueryJsonDto) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *QueryJsonDto) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *QueryJsonDto) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFullText

`func (o *QueryJsonDto) GetFullText() string`

GetFullText returns the FullText field if non-nil, zero value otherwise.

### GetFullTextOk

`func (o *QueryJsonDto) GetFullTextOk() (*string, bool)`

GetFullTextOk returns a tuple with the FullText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullText

`func (o *QueryJsonDto) SetFullText(v string)`

SetFullText sets FullText field to given value.

### HasFullText

`func (o *QueryJsonDto) HasFullText() bool`

HasFullText returns a boolean if a field has been set.

### SetFullTextNil

`func (o *QueryJsonDto) SetFullTextNil(b bool)`

 SetFullTextNil sets the value for FullText to be an explicit nil

### UnsetFullText
`func (o *QueryJsonDto) UnsetFullText()`

UnsetFullText ensures that no value is present for FullText, not even an explicit nil
### GetSkip

`func (o *QueryJsonDto) GetSkip() int64`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *QueryJsonDto) GetSkipOk() (*int64, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *QueryJsonDto) SetSkip(v int64)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *QueryJsonDto) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *QueryJsonDto) GetTake() int64`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *QueryJsonDto) GetTakeOk() (*int64, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *QueryJsonDto) SetTake(v int64)`

SetTake sets Take field to given value.

### HasTake

`func (o *QueryJsonDto) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetTop

`func (o *QueryJsonDto) GetTop() int64`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *QueryJsonDto) GetTopOk() (*int64, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *QueryJsonDto) SetTop(v int64)`

SetTop sets Top field to given value.

### HasTop

`func (o *QueryJsonDto) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetSort

`func (o *QueryJsonDto) GetSort() []SortNode`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryJsonDto) GetSortOk() (*[]SortNode, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryJsonDto) SetSort(v []SortNode)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryJsonDto) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


