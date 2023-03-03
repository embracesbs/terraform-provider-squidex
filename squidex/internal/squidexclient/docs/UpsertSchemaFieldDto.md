# UpsertSchemaFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | **bool** | Defines if the field is disabled. | [optional] 
**Partitioning** | Pointer to **string** | Determines the optional partitioning of the field. | [optional] 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 
**Nested** | Pointer to [**[]UpsertSchemaNestedFieldDto**](UpsertSchemaNestedFieldDto.md) | The nested fields. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


