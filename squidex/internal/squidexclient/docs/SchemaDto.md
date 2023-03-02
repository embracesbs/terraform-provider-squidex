# SchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The id of the schema. | [optional] 
**CreatedBy** | **string** | The user that has created the schema. | 
**LastModifiedBy** | **string** | The user that has updated the schema. | 
**Name** | **string** | The name of the schema. Unique within the app. | 
**Type** | [**SchemaType**](SchemaType.md) |  | [optional] 
**Category** | Pointer to **string** | The name of the category. | [optional] 
**Properties** | [**SchemaPropertiesDto**](SchemaPropertiesDto.md) |  | 
**IsSingleton** | **bool** | Indicates if the schema is a singleton. | [optional] 
**IsPublished** | **bool** | Indicates if the schema is published. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date and time when the schema has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the schema has been modified last. | [optional] 
**Version** | **int64** | The version of the schema. | [optional] 
**Scripts** | [**SchemaScriptsDto**](SchemaScriptsDto.md) |  | 
**PreviewUrls** | **map[string]string** | The preview Urls. | 
**FieldsInLists** | **[]string** |  | 
**FieldsInReferences** | **[]string** |  | 
**FieldRules** | [**[]FieldRuleDto**](FieldRuleDto.md) | The field rules. | [optional] 
**Fields** | [**[]FieldDto**](FieldDto.md) | The list of fields. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


