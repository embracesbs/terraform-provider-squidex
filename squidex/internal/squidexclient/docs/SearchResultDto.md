# SearchResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Name** | **string** | The name of the search result. | 
**Type** | [**SearchResultType**](SearchResultType.md) |  | 
**Label** | Pointer to **NullableString** | An optional label. | [optional] 

## Methods

### NewSearchResultDto

`func NewSearchResultDto(links map[string]ResourceLink, name string, type_ SearchResultType, ) *SearchResultDto`

NewSearchResultDto instantiates a new SearchResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultDtoWithDefaults

`func NewSearchResultDtoWithDefaults() *SearchResultDto`

NewSearchResultDtoWithDefaults instantiates a new SearchResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SearchResultDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SearchResultDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SearchResultDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetName

`func (o *SearchResultDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchResultDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchResultDto) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SearchResultDto) GetType() SearchResultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResultDto) GetTypeOk() (*SearchResultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResultDto) SetType(v SearchResultType)`

SetType sets Type field to given value.


### GetLabel

`func (o *SearchResultDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SearchResultDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SearchResultDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SearchResultDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *SearchResultDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *SearchResultDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


