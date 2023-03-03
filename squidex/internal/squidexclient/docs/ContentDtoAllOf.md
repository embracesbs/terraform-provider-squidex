# ContentDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The if of the content item. | [optional] 
**CreatedBy** | **string** | The user that has created the content item. | 
**LastModifiedBy** | **string** | The user that has updated the content item. | 
**Data** | [**interface{}**](.md) | The data of the content item. | 
**ReferenceData** | [**map[string]map[string]interface{}**](map.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date and time when the content item has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the content item has been modified last. | [optional] 
**Status** | **string** | The status of the content. | [optional] 
**NewStatus** | Pointer to **string** | The new status of the content. | [optional] 
**StatusColor** | **string** | The color of the status. | [optional] 
**NewStatusColor** | Pointer to **string** | The color of the new status. | [optional] 
**EditToken** | Pointer to **string** | The UI token. | [optional] 
**ScheduleJob** | [**ScheduleJobDto**](ScheduleJobDto.md) |  | [optional] 
**SchemaId** | **string** | The id of the schema. | [optional] 
**SchemaName** | Pointer to **string** | The name of the schema. | [optional] 
**SchemaDisplayName** | Pointer to **string** | The display name of the schema. | [optional] 
**ReferenceFields** | Pointer to [**[]FieldDto**](FieldDto.md) | The reference fields. | [optional] 
**IsDeleted** | **bool** | Indicates whether the content is deleted. | [optional] 
**Version** | **int64** | The version of the content. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


