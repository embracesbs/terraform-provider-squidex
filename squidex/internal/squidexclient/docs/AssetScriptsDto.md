# AssetScriptsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Create** | Pointer to **NullableString** | The script that is executed when creating an asset. | [optional] 
**Update** | Pointer to **NullableString** | The script that is executed when updating a content. | [optional] 
**Annotate** | Pointer to **NullableString** | The script that is executed when annotating a content. | [optional] 
**Move** | Pointer to **NullableString** | The script that is executed when moving a content. | [optional] 
**Delete** | Pointer to **NullableString** | The script that is executed when deleting a content. | [optional] 
**Version** | Pointer to **int64** | The version of the app. | [optional] 

## Methods

### NewAssetScriptsDto

`func NewAssetScriptsDto(links map[string]ResourceLink, ) *AssetScriptsDto`

NewAssetScriptsDto instantiates a new AssetScriptsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetScriptsDtoWithDefaults

`func NewAssetScriptsDtoWithDefaults() *AssetScriptsDto`

NewAssetScriptsDtoWithDefaults instantiates a new AssetScriptsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AssetScriptsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssetScriptsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssetScriptsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetCreate

`func (o *AssetScriptsDto) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AssetScriptsDto) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AssetScriptsDto) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AssetScriptsDto) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreateNil

`func (o *AssetScriptsDto) SetCreateNil(b bool)`

 SetCreateNil sets the value for Create to be an explicit nil

### UnsetCreate
`func (o *AssetScriptsDto) UnsetCreate()`

UnsetCreate ensures that no value is present for Create, not even an explicit nil
### GetUpdate

`func (o *AssetScriptsDto) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AssetScriptsDto) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AssetScriptsDto) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AssetScriptsDto) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *AssetScriptsDto) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *AssetScriptsDto) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetAnnotate

`func (o *AssetScriptsDto) GetAnnotate() string`

GetAnnotate returns the Annotate field if non-nil, zero value otherwise.

### GetAnnotateOk

`func (o *AssetScriptsDto) GetAnnotateOk() (*string, bool)`

GetAnnotateOk returns a tuple with the Annotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotate

`func (o *AssetScriptsDto) SetAnnotate(v string)`

SetAnnotate sets Annotate field to given value.

### HasAnnotate

`func (o *AssetScriptsDto) HasAnnotate() bool`

HasAnnotate returns a boolean if a field has been set.

### SetAnnotateNil

`func (o *AssetScriptsDto) SetAnnotateNil(b bool)`

 SetAnnotateNil sets the value for Annotate to be an explicit nil

### UnsetAnnotate
`func (o *AssetScriptsDto) UnsetAnnotate()`

UnsetAnnotate ensures that no value is present for Annotate, not even an explicit nil
### GetMove

`func (o *AssetScriptsDto) GetMove() string`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *AssetScriptsDto) GetMoveOk() (*string, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *AssetScriptsDto) SetMove(v string)`

SetMove sets Move field to given value.

### HasMove

`func (o *AssetScriptsDto) HasMove() bool`

HasMove returns a boolean if a field has been set.

### SetMoveNil

`func (o *AssetScriptsDto) SetMoveNil(b bool)`

 SetMoveNil sets the value for Move to be an explicit nil

### UnsetMove
`func (o *AssetScriptsDto) UnsetMove()`

UnsetMove ensures that no value is present for Move, not even an explicit nil
### GetDelete

`func (o *AssetScriptsDto) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AssetScriptsDto) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AssetScriptsDto) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AssetScriptsDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *AssetScriptsDto) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *AssetScriptsDto) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetVersion

`func (o *AssetScriptsDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AssetScriptsDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AssetScriptsDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AssetScriptsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


