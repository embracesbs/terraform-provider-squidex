# UpsertCommentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The comment text. | 
**Url** | Pointer to **NullableString** | The url where the comment is created. | [optional] 

## Methods

### NewUpsertCommentDto

`func NewUpsertCommentDto(text string, ) *UpsertCommentDto`

NewUpsertCommentDto instantiates a new UpsertCommentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertCommentDtoWithDefaults

`func NewUpsertCommentDtoWithDefaults() *UpsertCommentDto`

NewUpsertCommentDtoWithDefaults instantiates a new UpsertCommentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *UpsertCommentDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UpsertCommentDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UpsertCommentDto) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *UpsertCommentDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpsertCommentDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpsertCommentDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpsertCommentDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpsertCommentDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpsertCommentDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


