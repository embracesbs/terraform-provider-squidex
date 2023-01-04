# SearchResultDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the search result. | 
**Type** | [**SearchResultType**](SearchResultType.md) |  | 
**Label** | Pointer to **NullableString** | An optional label. | [optional] 

## Methods

### NewSearchResultDtoAllOf

`func NewSearchResultDtoAllOf(name string, type_ SearchResultType, ) *SearchResultDtoAllOf`

NewSearchResultDtoAllOf instantiates a new SearchResultDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultDtoAllOfWithDefaults

`func NewSearchResultDtoAllOfWithDefaults() *SearchResultDtoAllOf`

NewSearchResultDtoAllOfWithDefaults instantiates a new SearchResultDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchResultDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchResultDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchResultDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SearchResultDtoAllOf) GetType() SearchResultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResultDtoAllOf) GetTypeOk() (*SearchResultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResultDtoAllOf) SetType(v SearchResultType)`

SetType sets Type field to given value.


### GetLabel

`func (o *SearchResultDtoAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SearchResultDtoAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SearchResultDtoAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SearchResultDtoAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *SearchResultDtoAllOf) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *SearchResultDtoAllOf) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


