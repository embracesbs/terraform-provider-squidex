# BulkUpdateJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**OneOfQueryOfIJsonValue**](oneOf&lt;QueryOfIJsonValue&gt;.md) | An optional query to identify the content to update. | [optional] 
**Id** | Pointer to **string** | An optional id of the content to update. | [optional] 
**Data** | Pointer to [**OneOfNamedContentData**](oneOf&lt;NamedContentData&gt;.md) | The data of the content when type is set to &#39;Upsert&#39;. | [optional] 
**Status** | Pointer to **string** | The new status when the type is set to &#39;ChangeStatus&#39;. | [optional] 
**Type** | [**OneOfBulkUpdateType**](oneOf&lt;BulkUpdateType&gt;.md) | The update type. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


