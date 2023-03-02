# CommentRuleActionDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The comment text. | 
**Client** | Pointer to **NullableString** | An optional client name. | [optional] 

## Methods

### NewCommentRuleActionDtoAllOf

`func NewCommentRuleActionDtoAllOf(text string, ) *CommentRuleActionDtoAllOf`

NewCommentRuleActionDtoAllOf instantiates a new CommentRuleActionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRuleActionDtoAllOfWithDefaults

`func NewCommentRuleActionDtoAllOfWithDefaults() *CommentRuleActionDtoAllOf`

NewCommentRuleActionDtoAllOfWithDefaults instantiates a new CommentRuleActionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CommentRuleActionDtoAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CommentRuleActionDtoAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CommentRuleActionDtoAllOf) SetText(v string)`

SetText sets Text field to given value.


### GetClient

`func (o *CommentRuleActionDtoAllOf) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CommentRuleActionDtoAllOf) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CommentRuleActionDtoAllOf) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *CommentRuleActionDtoAllOf) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *CommentRuleActionDtoAllOf) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *CommentRuleActionDtoAllOf) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


