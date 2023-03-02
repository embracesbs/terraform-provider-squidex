# ReferencesFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **string** | Hints to describe the field. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder to show when no value has been entered. | [optional] 
**IsRequired** | **bool** | Indicates if the field is required. | [optional] 
**IsRequiredOnPublish** | **bool** | Indicates if the field is required when publishing. | [optional] 
**IsHalfWidth** | **bool** | Indicates if the field should be rendered with half width only. | [optional] 
**EditorUrl** | Pointer to **string** | Optional url to the editor. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 
**FieldType** | **string** |  | 
**DefaultValues** | [**map[string][]string**](array.md) |  | [optional] 
**DefaultValue** | Pointer to **[]string** | The default value as a list of content ids. | [optional] 
**MinItems** | Pointer to **int32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **int32** | The maximum allowed items for the field value. | [optional] 
**AllowDuplicates** | **bool** | True, if duplicate values are allowed. | [optional] 
**ResolveReference** | **bool** | True to resolve references in the content list. | [optional] 
**MustBePublished** | **bool** | True when all references must be published. | [optional] 
**Editor** | [**ReferencesFieldEditor**](ReferencesFieldEditor.md) |  | [optional] 
**SchemaIds** | Pointer to **[]string** | The id of the referenced schemas. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


