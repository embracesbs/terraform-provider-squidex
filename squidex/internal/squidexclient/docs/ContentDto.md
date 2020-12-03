# ContentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The if of the content item. | [optional] 
**CreatedBy** | **string** | The user that has created the content item. | 
**LastModifiedBy** | **string** | The user that has updated the content item. | 
**Data** | [**interface{}**](.md) | The data of the content item. | 
**DataDraft** | Pointer to [**interface{}**](.md) | The pending changes of the content item. | [optional] 
**ReferenceData** | Pointer to [**OneOfNamedContentData**](oneOf&lt;NamedContentData&gt;.md) | The reference data for the frontend UI. | [optional] 
**IsPending** | **bool** | Indicates if the draft data is pending. | [optional] 
**ScheduleJob** | Pointer to [**OneOfScheduleJobDto**](oneOf&lt;ScheduleJobDto&gt;.md) | The scheduled status. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date and time when the content item has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the content item has been modified last. | [optional] 
**Status** | **string** | The status of the content. | [optional] 
**StatusColor** | **string** | The color of the status. | [optional] 
**SchemaName** | Pointer to **string** | The name of the schema. | [optional] 
**SchemaDisplayName** | Pointer to **string** | The display name of the schema. | [optional] 
**ReferenceFields** | Pointer to [**[]FieldDto**](FieldDto.md) | The reference fields. | [optional] 
**Version** | **int64** | The version of the content. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


