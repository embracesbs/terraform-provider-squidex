# CommentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the comment. | [optional] 
**Time** | **time.Time** | The time when the comment was created or updated last. | 
**User** | **string** | The user who created or updated the comment. | 
**Text** | **string** | The text of the comment. | 
**Url** | Pointer to **NullableString** | The url where the comment is created. | [optional] 

## Methods

### NewCommentDto

`func NewCommentDto(time time.Time, user string, text string, ) *CommentDto`

NewCommentDto instantiates a new CommentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentDtoWithDefaults

`func NewCommentDtoWithDefaults() *CommentDto`

NewCommentDtoWithDefaults instantiates a new CommentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *CommentDto) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CommentDto) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CommentDto) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetUser

`func (o *CommentDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommentDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommentDto) SetUser(v string)`

SetUser sets User field to given value.


### GetText

`func (o *CommentDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CommentDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CommentDto) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *CommentDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommentDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommentDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CommentDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CommentDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CommentDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


