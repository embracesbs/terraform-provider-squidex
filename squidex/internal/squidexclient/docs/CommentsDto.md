# CommentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedComments** | Pointer to [**[]CommentDto**](CommentDto.md) | The created comments including the updates. | [optional] 
**UpdatedComments** | Pointer to [**[]CommentDto**](CommentDto.md) | The updates comments since the last version. | [optional] 
**DeletedComments** | Pointer to **[]string** | The deleted comments since the last version. | [optional] 
**Version** | Pointer to **int64** | The current version. | [optional] 

## Methods

### NewCommentsDto

`func NewCommentsDto() *CommentsDto`

NewCommentsDto instantiates a new CommentsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsDtoWithDefaults

`func NewCommentsDtoWithDefaults() *CommentsDto`

NewCommentsDtoWithDefaults instantiates a new CommentsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedComments

`func (o *CommentsDto) GetCreatedComments() []CommentDto`

GetCreatedComments returns the CreatedComments field if non-nil, zero value otherwise.

### GetCreatedCommentsOk

`func (o *CommentsDto) GetCreatedCommentsOk() (*[]CommentDto, bool)`

GetCreatedCommentsOk returns a tuple with the CreatedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedComments

`func (o *CommentsDto) SetCreatedComments(v []CommentDto)`

SetCreatedComments sets CreatedComments field to given value.

### HasCreatedComments

`func (o *CommentsDto) HasCreatedComments() bool`

HasCreatedComments returns a boolean if a field has been set.

### SetCreatedCommentsNil

`func (o *CommentsDto) SetCreatedCommentsNil(b bool)`

 SetCreatedCommentsNil sets the value for CreatedComments to be an explicit nil

### UnsetCreatedComments
`func (o *CommentsDto) UnsetCreatedComments()`

UnsetCreatedComments ensures that no value is present for CreatedComments, not even an explicit nil
### GetUpdatedComments

`func (o *CommentsDto) GetUpdatedComments() []CommentDto`

GetUpdatedComments returns the UpdatedComments field if non-nil, zero value otherwise.

### GetUpdatedCommentsOk

`func (o *CommentsDto) GetUpdatedCommentsOk() (*[]CommentDto, bool)`

GetUpdatedCommentsOk returns a tuple with the UpdatedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedComments

`func (o *CommentsDto) SetUpdatedComments(v []CommentDto)`

SetUpdatedComments sets UpdatedComments field to given value.

### HasUpdatedComments

`func (o *CommentsDto) HasUpdatedComments() bool`

HasUpdatedComments returns a boolean if a field has been set.

### SetUpdatedCommentsNil

`func (o *CommentsDto) SetUpdatedCommentsNil(b bool)`

 SetUpdatedCommentsNil sets the value for UpdatedComments to be an explicit nil

### UnsetUpdatedComments
`func (o *CommentsDto) UnsetUpdatedComments()`

UnsetUpdatedComments ensures that no value is present for UpdatedComments, not even an explicit nil
### GetDeletedComments

`func (o *CommentsDto) GetDeletedComments() []string`

GetDeletedComments returns the DeletedComments field if non-nil, zero value otherwise.

### GetDeletedCommentsOk

`func (o *CommentsDto) GetDeletedCommentsOk() (*[]string, bool)`

GetDeletedCommentsOk returns a tuple with the DeletedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedComments

`func (o *CommentsDto) SetDeletedComments(v []string)`

SetDeletedComments sets DeletedComments field to given value.

### HasDeletedComments

`func (o *CommentsDto) HasDeletedComments() bool`

HasDeletedComments returns a boolean if a field has been set.

### SetDeletedCommentsNil

`func (o *CommentsDto) SetDeletedCommentsNil(b bool)`

 SetDeletedCommentsNil sets the value for DeletedComments to be an explicit nil

### UnsetDeletedComments
`func (o *CommentsDto) UnsetDeletedComments()`

UnsetDeletedComments ensures that no value is present for DeletedComments, not even an explicit nil
### GetVersion

`func (o *CommentsDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CommentsDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CommentsDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CommentsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


