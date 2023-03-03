# CreateSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**SchemaPropertiesDto**](SchemaPropertiesDto.md) |  | [optional] 
**Scripts** | [**SchemaScriptsDto**](SchemaScriptsDto.md) |  | [optional] 
**FieldsInReferences** | **[]string** |  | [optional] 
**FieldsInLists** | **[]string** |  | [optional] 
**Fields** | Pointer to [**[]UpsertSchemaFieldDto**](UpsertSchemaFieldDto.md) | Optional fields. | [optional] 
**PreviewUrls** | Pointer to **map[string]string** | The optional preview urls. | [optional] 
**FieldRules** | Pointer to [**[]FieldRuleDto**](FieldRuleDto.md) | The optional field Rules. | [optional] 
**Category** | Pointer to **string** | The category. | [optional] 
**IsPublished** | **bool** | Set it to true to autopublish the schema. | [optional] 
**Name** | **string** | The name of the schema. | 
**Type** | [**SchemaType**](SchemaType.md) |  | [optional] 
**IsSingleton** | **bool** | Set to true to allow a single content item only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


