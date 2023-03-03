# BulkUpdateAssetsJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An optional id of the asset to update. | [optional] 
**Type** | [**BulkUpdateAssetType**](BulkUpdateAssetType.md) |  | [optional] 
**ParentId** | **string** | The parent folder id. | [optional] 
**FileName** | Pointer to **string** | The new name of the asset. | [optional] 
**Slug** | Pointer to **string** | The new slug of the asset. | [optional] 
**IsProtected** | Pointer to **bool** | True, when the asset is not public. | [optional] 
**Tags** | Pointer to **[]string** | The new asset tags. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The asset metadata. | [optional] 
**Permanent** | **bool** | True to delete the asset permanently. | [optional] 
**ExpectedVersion** | **int64** | The expected version. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


