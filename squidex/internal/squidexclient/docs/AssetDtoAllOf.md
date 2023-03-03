# AssetDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the asset. | [optional] 
**ParentId** | **string** | The id of the parent folder. Empty for files without parent. | [optional] 
**FileName** | **string** | The file name. | 
**FileHash** | Pointer to **string** | The file hash. | [optional] 
**IsProtected** | **bool** | True, when the asset is not public. | [optional] 
**Slug** | **string** | The slug. | 
**MimeType** | **string** | The mime type. | 
**FileType** | **string** | The file type. | 
**MetadataText** | **string** | The formatted text representation of the metadata. | 
**EditToken** | Pointer to **string** | The UI token. | [optional] 
**Metadata** | **map[string]interface{}** | The asset metadata. | 
**Tags** | **[]string** | The asset tags. | 
**FileSize** | **int64** | The size of the file in bytes. | [optional] 
**FileVersion** | **int64** | The version of the file. | [optional] 
**Type** | [**AssetType**](AssetType.md) |  | [optional] 
**CreatedBy** | **string** | The user that has created the schema. | 
**LastModifiedBy** | **string** | The user that has updated the asset. | 
**Created** | [**time.Time**](time.Time.md) | The date and time when the asset has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the asset has been modified last. | [optional] 
**Version** | **int64** | The version of the asset. | [optional] 
**Meta** | [**AssetMeta**](AssetMeta.md) |  | [optional] 
**IsImage** | **bool** | Determines of the created file is an image. | [optional] 
**PixelWidth** | Pointer to **int32** | The width of the image in pixels if the asset is an image. | [optional] 
**PixelHeight** | Pointer to **int32** | The height of the image in pixels if the asset is an image. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


