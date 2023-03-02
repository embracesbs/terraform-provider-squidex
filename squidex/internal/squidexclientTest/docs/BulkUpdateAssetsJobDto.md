# BulkUpdateAssetsJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional id of the asset to update. | [optional] 
**Type** | Pointer to [**BulkUpdateAssetType**](BulkUpdateAssetType.md) |  | [optional] 
**ParentId** | Pointer to **string** | The parent folder id. | [optional] 
**FileName** | Pointer to **NullableString** | The new name of the asset. | [optional] 
**Slug** | Pointer to **NullableString** | The new slug of the asset. | [optional] 
**IsProtected** | Pointer to **NullableBool** | True, when the asset is not public. | [optional] 
**Tags** | Pointer to **[]string** | The new asset tags. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The asset metadata. | [optional] 
**Permanent** | Pointer to **bool** | True to delete the asset permanently. | [optional] 
**ExpectedVersion** | Pointer to **int64** | The expected version. | [optional] 

## Methods

### NewBulkUpdateAssetsJobDto

`func NewBulkUpdateAssetsJobDto() *BulkUpdateAssetsJobDto`

NewBulkUpdateAssetsJobDto instantiates a new BulkUpdateAssetsJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAssetsJobDtoWithDefaults

`func NewBulkUpdateAssetsJobDtoWithDefaults() *BulkUpdateAssetsJobDto`

NewBulkUpdateAssetsJobDtoWithDefaults instantiates a new BulkUpdateAssetsJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkUpdateAssetsJobDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateAssetsJobDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateAssetsJobDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUpdateAssetsJobDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BulkUpdateAssetsJobDto) GetType() BulkUpdateAssetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkUpdateAssetsJobDto) GetTypeOk() (*BulkUpdateAssetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkUpdateAssetsJobDto) SetType(v BulkUpdateAssetType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkUpdateAssetsJobDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParentId

`func (o *BulkUpdateAssetsJobDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BulkUpdateAssetsJobDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BulkUpdateAssetsJobDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BulkUpdateAssetsJobDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFileName

`func (o *BulkUpdateAssetsJobDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BulkUpdateAssetsJobDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BulkUpdateAssetsJobDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BulkUpdateAssetsJobDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *BulkUpdateAssetsJobDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *BulkUpdateAssetsJobDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetSlug

`func (o *BulkUpdateAssetsJobDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkUpdateAssetsJobDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkUpdateAssetsJobDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BulkUpdateAssetsJobDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *BulkUpdateAssetsJobDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *BulkUpdateAssetsJobDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetIsProtected

`func (o *BulkUpdateAssetsJobDto) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *BulkUpdateAssetsJobDto) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *BulkUpdateAssetsJobDto) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *BulkUpdateAssetsJobDto) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *BulkUpdateAssetsJobDto) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *BulkUpdateAssetsJobDto) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetTags

`func (o *BulkUpdateAssetsJobDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkUpdateAssetsJobDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkUpdateAssetsJobDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkUpdateAssetsJobDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BulkUpdateAssetsJobDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BulkUpdateAssetsJobDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetMetadata

`func (o *BulkUpdateAssetsJobDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BulkUpdateAssetsJobDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BulkUpdateAssetsJobDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BulkUpdateAssetsJobDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BulkUpdateAssetsJobDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BulkUpdateAssetsJobDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPermanent

`func (o *BulkUpdateAssetsJobDto) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *BulkUpdateAssetsJobDto) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *BulkUpdateAssetsJobDto) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *BulkUpdateAssetsJobDto) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *BulkUpdateAssetsJobDto) GetExpectedVersion() int64`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *BulkUpdateAssetsJobDto) GetExpectedVersionOk() (*int64, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *BulkUpdateAssetsJobDto) SetExpectedVersion(v int64)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *BulkUpdateAssetsJobDto) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


