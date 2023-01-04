# UpdateAssetScriptsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **NullableString** | The script that is executed when creating an asset. | [optional] 
**Update** | Pointer to **NullableString** | The script that is executed when updating a content. | [optional] 
**Annotate** | Pointer to **NullableString** | The script that is executed when annotating a content. | [optional] 
**Move** | Pointer to **NullableString** | The script that is executed when moving a content. | [optional] 
**Delete** | Pointer to **NullableString** | The script that is executed when deleting a content. | [optional] 

## Methods

### NewUpdateAssetScriptsDto

`func NewUpdateAssetScriptsDto() *UpdateAssetScriptsDto`

NewUpdateAssetScriptsDto instantiates a new UpdateAssetScriptsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAssetScriptsDtoWithDefaults

`func NewUpdateAssetScriptsDtoWithDefaults() *UpdateAssetScriptsDto`

NewUpdateAssetScriptsDtoWithDefaults instantiates a new UpdateAssetScriptsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *UpdateAssetScriptsDto) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UpdateAssetScriptsDto) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UpdateAssetScriptsDto) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UpdateAssetScriptsDto) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreateNil

`func (o *UpdateAssetScriptsDto) SetCreateNil(b bool)`

 SetCreateNil sets the value for Create to be an explicit nil

### UnsetCreate
`func (o *UpdateAssetScriptsDto) UnsetCreate()`

UnsetCreate ensures that no value is present for Create, not even an explicit nil
### GetUpdate

`func (o *UpdateAssetScriptsDto) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UpdateAssetScriptsDto) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UpdateAssetScriptsDto) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UpdateAssetScriptsDto) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *UpdateAssetScriptsDto) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *UpdateAssetScriptsDto) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetAnnotate

`func (o *UpdateAssetScriptsDto) GetAnnotate() string`

GetAnnotate returns the Annotate field if non-nil, zero value otherwise.

### GetAnnotateOk

`func (o *UpdateAssetScriptsDto) GetAnnotateOk() (*string, bool)`

GetAnnotateOk returns a tuple with the Annotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotate

`func (o *UpdateAssetScriptsDto) SetAnnotate(v string)`

SetAnnotate sets Annotate field to given value.

### HasAnnotate

`func (o *UpdateAssetScriptsDto) HasAnnotate() bool`

HasAnnotate returns a boolean if a field has been set.

### SetAnnotateNil

`func (o *UpdateAssetScriptsDto) SetAnnotateNil(b bool)`

 SetAnnotateNil sets the value for Annotate to be an explicit nil

### UnsetAnnotate
`func (o *UpdateAssetScriptsDto) UnsetAnnotate()`

UnsetAnnotate ensures that no value is present for Annotate, not even an explicit nil
### GetMove

`func (o *UpdateAssetScriptsDto) GetMove() string`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *UpdateAssetScriptsDto) GetMoveOk() (*string, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *UpdateAssetScriptsDto) SetMove(v string)`

SetMove sets Move field to given value.

### HasMove

`func (o *UpdateAssetScriptsDto) HasMove() bool`

HasMove returns a boolean if a field has been set.

### SetMoveNil

`func (o *UpdateAssetScriptsDto) SetMoveNil(b bool)`

 SetMoveNil sets the value for Move to be an explicit nil

### UnsetMove
`func (o *UpdateAssetScriptsDto) UnsetMove()`

UnsetMove ensures that no value is present for Move, not even an explicit nil
### GetDelete

`func (o *UpdateAssetScriptsDto) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateAssetScriptsDto) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateAssetScriptsDto) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateAssetScriptsDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *UpdateAssetScriptsDto) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *UpdateAssetScriptsDto) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


