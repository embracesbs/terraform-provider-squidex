# StringFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | **map[string]string** |  | [optional] 
**DefaultValue** | Pointer to **string** | The default value for the field value. | [optional] 
**Pattern** | Pointer to **string** | The pattern to enforce a specific format for the field value. | [optional] 
**PatternMessage** | Pointer to **string** | The validation message for the pattern. | [optional] 
**FolderId** | Pointer to **string** | The initial id to the folder when the control supports file uploads. | [optional] 
**MinLength** | Pointer to **int32** | The minimum allowed length for the field value. | [optional] 
**MaxLength** | Pointer to **int32** | The maximum allowed length for the field value. | [optional] 
**MinCharacters** | Pointer to **int32** | The minimum allowed of normal characters for the field value. | [optional] 
**MaxCharacters** | Pointer to **int32** | The maximum allowed of normal characters for the field value. | [optional] 
**MinWords** | Pointer to **int32** | The minimum allowed number of words for the field value. | [optional] 
**MaxWords** | Pointer to **int32** | The maximum allowed number of words for the field value. | [optional] 
**AllowedValues** | Pointer to **[]string** | The allowed values for the field value. | [optional] 
**SchemaIds** | Pointer to **[]string** | The allowed schema ids that can be embedded. | [optional] 
**IsUnique** | **bool** | Indicates if the field value must be unique. Ignored for nested fields and localized fields. | [optional] 
**IsEmbeddable** | **bool** | Indicates that other content items or references are embedded. | [optional] 
**InlineEditable** | **bool** | Indicates that the inline editor is enabled for this field. | [optional] 
**CreateEnum** | **bool** | Indicates whether GraphQL Enum should be created. | [optional] 
**ContentType** | [**StringContentType**](StringContentType.md) |  | [optional] 
**Editor** | [**StringFieldEditor**](StringFieldEditor.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


