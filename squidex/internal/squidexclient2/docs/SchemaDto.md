# SchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The id of the schema. | [optional] 
**Name** | **string** | The name of the schema. Unique within the app. | 
**Category** | Pointer to **string** | The name of the category. | [optional] 
**Properties** | [**OneOfSchemaPropertiesDto**](oneOf&lt;SchemaPropertiesDto&gt;.md) | The schema properties. | 
**IsSingleton** | **bool** | Indicates if the schema is a singleton. | [optional] 
**IsPublished** | **bool** | Indicates if the schema is published. | [optional] 
**CreatedBy** | **string** | The user that has created the schema. | 
**LastModifiedBy** | **string** | The user that has updated the schema. | 
**Created** | [**time.Time**](time.Time.md) | The date and time when the schema has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the schema has been modified last. | [optional] 
**Version** | **int64** | The version of the schema. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


