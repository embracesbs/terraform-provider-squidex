# CommentRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The comment text. | 
**Client** | Pointer to **NullableString** | An optional client name. | [optional] 

## Methods

### NewCommentRuleActionDto

`func NewCommentRuleActionDto(text string, ) *CommentRuleActionDto`

NewCommentRuleActionDto instantiates a new CommentRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRuleActionDtoWithDefaults

`func NewCommentRuleActionDtoWithDefaults() *CommentRuleActionDto`

NewCommentRuleActionDtoWithDefaults instantiates a new CommentRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CommentRuleActionDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CommentRuleActionDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CommentRuleActionDto) SetText(v string)`

SetText sets Text field to given value.


### GetClient

`func (o *CommentRuleActionDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CommentRuleActionDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CommentRuleActionDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *CommentRuleActionDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *CommentRuleActionDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *CommentRuleActionDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


