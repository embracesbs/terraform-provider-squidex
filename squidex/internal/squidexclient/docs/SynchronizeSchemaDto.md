# SynchronizeSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**OneOfSchemaPropertiesDto**](oneOf&lt;SchemaPropertiesDto&gt;.md) | The optional properties. | [optional] 
**Scripts** | Pointer to [**OneOfSchemaScriptsDto**](oneOf&lt;SchemaScriptsDto&gt;.md) | The optional scripts. | [optional] 
**FieldsInReferences** | Pointer to [**OneOfarray**](oneOf&lt;array&gt;.md) | The names of the fields that should be used in references. | [optional] 
**FieldsInLists** | Pointer to [**OneOfarray**](oneOf&lt;array&gt;.md) | The names of the fields that should be shown in lists, including meta fields. | [optional] 
**Fields** | Pointer to [**[]UpsertSchemaFieldDto**](UpsertSchemaFieldDto.md) | Optional fields. | [optional] 
**PreviewUrls** | Pointer to **map[string]string** | The optional preview urls. | [optional] 
**Category** | Pointer to **string** | The category. | [optional] 
**IsPublished** | **bool** | Set it to true to autopublish the schema. | [optional] 
**NoFieldDeletion** | **bool** | True, when fields should not be deleted. | [optional] 
**NoFieldRecreation** | **bool** | True, when fields with different types should not be recreated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


