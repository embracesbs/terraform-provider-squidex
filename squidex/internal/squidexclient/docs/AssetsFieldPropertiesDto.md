# AssetsFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **string** | Hints to describe the field. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder to show when no value has been entered. | [optional] 
**IsRequired** | **bool** | Indicates if the field is required. | [optional] 
**IsRequiredOnPublish** | **bool** | Indicates if the field is required when publishing. | [optional] 
**IsHalfWidth** | **bool** | Indicates if the field should be rendered with half width only. | [optional] 
**EditorUrl** | Pointer to **string** | Optional url to the editor. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 
**FieldType** | **string** |  | 
**PreviewMode** | [**AssetPreviewMode**](AssetPreviewMode.md) |  | [optional] 
**DefaultValues** | [**map[string][]string**](array.md) |  | [optional] 
**DefaultValue** | Pointer to **[]string** | The default value as a list of asset ids. | [optional] 
**FolderId** | Pointer to **string** | The initial id to the folder. | [optional] 
**MinItems** | Pointer to **int32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **int32** | The maximum allowed items for the field value. | [optional] 
**MinSize** | Pointer to **int32** | The minimum file size in bytes. | [optional] 
**MaxSize** | Pointer to **int32** | The maximum file size in bytes. | [optional] 
**MinWidth** | Pointer to **int32** | The minimum image width in pixels. | [optional] 
**MaxWidth** | Pointer to **int32** | The maximum image width in pixels. | [optional] 
**MinHeight** | Pointer to **int32** | The minimum image height in pixels. | [optional] 
**MaxHeight** | Pointer to **int32** | The maximum image height in pixels. | [optional] 
**AspectWidth** | Pointer to **int32** | The image aspect width in pixels. | [optional] 
**AspectHeight** | Pointer to **int32** | The image aspect height in pixels. | [optional] 
**ExpectedType** | [**AssetType**](AssetType.md) |  | [optional] 
**ResolveFirst** | **bool** | True to resolve first asset in the content list. | [optional] 
**MustBeImage** | **bool** | True to resolve first image in the content list. | [optional] 
**ResolveImage** | **bool** | True to resolve first image in the content list. | [optional] 
**AllowedExtensions** | Pointer to **[]string** | The allowed file extensions. | [optional] 
**AllowDuplicates** | **bool** | True, if duplicate values are allowed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


