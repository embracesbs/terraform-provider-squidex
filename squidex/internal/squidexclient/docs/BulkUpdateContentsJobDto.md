# BulkUpdateContentsJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [**QueryJsonDto**](QueryJsonDto.md) |  | [optional] 
**Id** | Pointer to **string** | An optional id of the content to update. | [optional] 
**Data** | [**map[string]map[string]interface{}**](map.md) |  | [optional] 
**Status** | Pointer to **string** | The new status when the type is set to &#39;ChangeStatus&#39; or &#39;Upsert&#39;. | [optional] 
**DueTime** | Pointer to [**time.Time**](time.Time.md) | The due time. | [optional] 
**Type** | [**BulkUpdateContentType**](BulkUpdateContentType.md) |  | [optional] 
**Schema** | Pointer to **string** | The optional schema id or name. | [optional] 
**Patch** | **bool** | Makes the update as patch. | [optional] 
**Permanent** | **bool** | True to delete the content permanently. | [optional] 
**ExpectedCount** | **int64** | The number of expected items. Set it to a higher number to update multiple items when a query is defined. | [optional] 
**ExpectedVersion** | **int64** | The expected version. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


