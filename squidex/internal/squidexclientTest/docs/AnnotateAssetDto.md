# AnnotateAssetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **NullableString** | The new name of the asset. | [optional] 
**Slug** | Pointer to **NullableString** | The new slug of the asset. | [optional] 
**IsProtected** | Pointer to **NullableBool** | True, when the asset is not public. | [optional] 
**Tags** | Pointer to **[]string** | The new asset tags. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The asset metadata. | [optional] 

## Methods

### NewAnnotateAssetDto

`func NewAnnotateAssetDto() *AnnotateAssetDto`

NewAnnotateAssetDto instantiates a new AnnotateAssetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotateAssetDtoWithDefaults

`func NewAnnotateAssetDtoWithDefaults() *AnnotateAssetDto`

NewAnnotateAssetDtoWithDefaults instantiates a new AnnotateAssetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *AnnotateAssetDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AnnotateAssetDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AnnotateAssetDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AnnotateAssetDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AnnotateAssetDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AnnotateAssetDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetSlug

`func (o *AnnotateAssetDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AnnotateAssetDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AnnotateAssetDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AnnotateAssetDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *AnnotateAssetDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *AnnotateAssetDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetIsProtected

`func (o *AnnotateAssetDto) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *AnnotateAssetDto) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *AnnotateAssetDto) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *AnnotateAssetDto) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *AnnotateAssetDto) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *AnnotateAssetDto) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetTags

`func (o *AnnotateAssetDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnnotateAssetDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnnotateAssetDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AnnotateAssetDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AnnotateAssetDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AnnotateAssetDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetMetadata

`func (o *AnnotateAssetDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AnnotateAssetDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AnnotateAssetDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AnnotateAssetDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AnnotateAssetDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AnnotateAssetDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


