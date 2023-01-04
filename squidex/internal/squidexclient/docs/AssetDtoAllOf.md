# AssetDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the asset. | [optional] 
**ParentId** | Pointer to **string** | The id of the parent folder. Empty for files without parent. | [optional] 
**FileName** | **string** | The file name. | 
**FileHash** | Pointer to **NullableString** | The file hash. | [optional] 
**IsProtected** | Pointer to **bool** | True, when the asset is not public. | [optional] 
**Slug** | **string** | The slug. | 
**MimeType** | **string** | The mime type. | 
**FileType** | **string** | The file type. | 
**MetadataText** | **string** | The formatted text representation of the metadata. | 
**EditToken** | Pointer to **NullableString** | The UI token. | [optional] 
**Metadata** | **map[string]interface{}** | The asset metadata. | 
**Tags** | **[]string** | The asset tags. | 
**FileSize** | Pointer to **int64** | The size of the file in bytes. | [optional] 
**FileVersion** | Pointer to **int64** | The version of the file. | [optional] 
**Type** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**CreatedBy** | **string** | The user that has created the schema. | 
**LastModifiedBy** | **string** | The user that has updated the asset. | 
**Created** | Pointer to **time.Time** | The date and time when the asset has been created. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the asset has been modified last. | [optional] 
**Version** | Pointer to **int64** | The version of the asset. | [optional] 
**Meta** | Pointer to [**AssetMeta**](AssetMeta.md) |  | [optional] 
**IsImage** | Pointer to **bool** | Determines of the created file is an image. | [optional] 
**PixelWidth** | Pointer to **NullableInt32** | The width of the image in pixels if the asset is an image. | [optional] 
**PixelHeight** | Pointer to **NullableInt32** | The height of the image in pixels if the asset is an image. | [optional] 

## Methods

### NewAssetDtoAllOf

`func NewAssetDtoAllOf(fileName string, slug string, mimeType string, fileType string, metadataText string, metadata map[string]interface{}, tags []string, createdBy string, lastModifiedBy string, ) *AssetDtoAllOf`

NewAssetDtoAllOf instantiates a new AssetDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDtoAllOfWithDefaults

`func NewAssetDtoAllOfWithDefaults() *AssetDtoAllOf`

NewAssetDtoAllOfWithDefaults instantiates a new AssetDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *AssetDtoAllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AssetDtoAllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AssetDtoAllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AssetDtoAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFileName

`func (o *AssetDtoAllOf) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AssetDtoAllOf) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AssetDtoAllOf) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileHash

`func (o *AssetDtoAllOf) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *AssetDtoAllOf) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *AssetDtoAllOf) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *AssetDtoAllOf) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHashNil

`func (o *AssetDtoAllOf) SetFileHashNil(b bool)`

 SetFileHashNil sets the value for FileHash to be an explicit nil

### UnsetFileHash
`func (o *AssetDtoAllOf) UnsetFileHash()`

UnsetFileHash ensures that no value is present for FileHash, not even an explicit nil
### GetIsProtected

`func (o *AssetDtoAllOf) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *AssetDtoAllOf) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *AssetDtoAllOf) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *AssetDtoAllOf) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetSlug

`func (o *AssetDtoAllOf) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AssetDtoAllOf) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AssetDtoAllOf) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetMimeType

`func (o *AssetDtoAllOf) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AssetDtoAllOf) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AssetDtoAllOf) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetFileType

`func (o *AssetDtoAllOf) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *AssetDtoAllOf) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *AssetDtoAllOf) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetMetadataText

`func (o *AssetDtoAllOf) GetMetadataText() string`

GetMetadataText returns the MetadataText field if non-nil, zero value otherwise.

### GetMetadataTextOk

`func (o *AssetDtoAllOf) GetMetadataTextOk() (*string, bool)`

GetMetadataTextOk returns a tuple with the MetadataText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataText

`func (o *AssetDtoAllOf) SetMetadataText(v string)`

SetMetadataText sets MetadataText field to given value.


### GetEditToken

`func (o *AssetDtoAllOf) GetEditToken() string`

GetEditToken returns the EditToken field if non-nil, zero value otherwise.

### GetEditTokenOk

`func (o *AssetDtoAllOf) GetEditTokenOk() (*string, bool)`

GetEditTokenOk returns a tuple with the EditToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditToken

`func (o *AssetDtoAllOf) SetEditToken(v string)`

SetEditToken sets EditToken field to given value.

### HasEditToken

`func (o *AssetDtoAllOf) HasEditToken() bool`

HasEditToken returns a boolean if a field has been set.

### SetEditTokenNil

`func (o *AssetDtoAllOf) SetEditTokenNil(b bool)`

 SetEditTokenNil sets the value for EditToken to be an explicit nil

### UnsetEditToken
`func (o *AssetDtoAllOf) UnsetEditToken()`

UnsetEditToken ensures that no value is present for EditToken, not even an explicit nil
### GetMetadata

`func (o *AssetDtoAllOf) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssetDtoAllOf) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssetDtoAllOf) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTags

`func (o *AssetDtoAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetDtoAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetDtoAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetFileSize

`func (o *AssetDtoAllOf) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AssetDtoAllOf) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AssetDtoAllOf) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AssetDtoAllOf) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileVersion

`func (o *AssetDtoAllOf) GetFileVersion() int64`

GetFileVersion returns the FileVersion field if non-nil, zero value otherwise.

### GetFileVersionOk

`func (o *AssetDtoAllOf) GetFileVersionOk() (*int64, bool)`

GetFileVersionOk returns a tuple with the FileVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVersion

`func (o *AssetDtoAllOf) SetFileVersion(v int64)`

SetFileVersion sets FileVersion field to given value.

### HasFileVersion

`func (o *AssetDtoAllOf) HasFileVersion() bool`

HasFileVersion returns a boolean if a field has been set.

### GetType

`func (o *AssetDtoAllOf) GetType() AssetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetDtoAllOf) GetTypeOk() (*AssetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetDtoAllOf) SetType(v AssetType)`

SetType sets Type field to given value.

### HasType

`func (o *AssetDtoAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AssetDtoAllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AssetDtoAllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AssetDtoAllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModifiedBy

`func (o *AssetDtoAllOf) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AssetDtoAllOf) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AssetDtoAllOf) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetCreated

`func (o *AssetDtoAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AssetDtoAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AssetDtoAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AssetDtoAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *AssetDtoAllOf) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AssetDtoAllOf) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AssetDtoAllOf) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AssetDtoAllOf) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersion

`func (o *AssetDtoAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AssetDtoAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AssetDtoAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AssetDtoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMeta

`func (o *AssetDtoAllOf) GetMeta() AssetMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AssetDtoAllOf) GetMetaOk() (*AssetMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AssetDtoAllOf) SetMeta(v AssetMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AssetDtoAllOf) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetIsImage

`func (o *AssetDtoAllOf) GetIsImage() bool`

GetIsImage returns the IsImage field if non-nil, zero value otherwise.

### GetIsImageOk

`func (o *AssetDtoAllOf) GetIsImageOk() (*bool, bool)`

GetIsImageOk returns a tuple with the IsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImage

`func (o *AssetDtoAllOf) SetIsImage(v bool)`

SetIsImage sets IsImage field to given value.

### HasIsImage

`func (o *AssetDtoAllOf) HasIsImage() bool`

HasIsImage returns a boolean if a field has been set.

### GetPixelWidth

`func (o *AssetDtoAllOf) GetPixelWidth() int32`

GetPixelWidth returns the PixelWidth field if non-nil, zero value otherwise.

### GetPixelWidthOk

`func (o *AssetDtoAllOf) GetPixelWidthOk() (*int32, bool)`

GetPixelWidthOk returns a tuple with the PixelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelWidth

`func (o *AssetDtoAllOf) SetPixelWidth(v int32)`

SetPixelWidth sets PixelWidth field to given value.

### HasPixelWidth

`func (o *AssetDtoAllOf) HasPixelWidth() bool`

HasPixelWidth returns a boolean if a field has been set.

### SetPixelWidthNil

`func (o *AssetDtoAllOf) SetPixelWidthNil(b bool)`

 SetPixelWidthNil sets the value for PixelWidth to be an explicit nil

### UnsetPixelWidth
`func (o *AssetDtoAllOf) UnsetPixelWidth()`

UnsetPixelWidth ensures that no value is present for PixelWidth, not even an explicit nil
### GetPixelHeight

`func (o *AssetDtoAllOf) GetPixelHeight() int32`

GetPixelHeight returns the PixelHeight field if non-nil, zero value otherwise.

### GetPixelHeightOk

`func (o *AssetDtoAllOf) GetPixelHeightOk() (*int32, bool)`

GetPixelHeightOk returns a tuple with the PixelHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelHeight

`func (o *AssetDtoAllOf) SetPixelHeight(v int32)`

SetPixelHeight sets PixelHeight field to given value.

### HasPixelHeight

`func (o *AssetDtoAllOf) HasPixelHeight() bool`

HasPixelHeight returns a boolean if a field has been set.

### SetPixelHeightNil

`func (o *AssetDtoAllOf) SetPixelHeightNil(b bool)`

 SetPixelHeightNil sets the value for PixelHeight to be an explicit nil

### UnsetPixelHeight
`func (o *AssetDtoAllOf) UnsetPixelHeight()`

UnsetPixelHeight ensures that no value is present for PixelHeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


