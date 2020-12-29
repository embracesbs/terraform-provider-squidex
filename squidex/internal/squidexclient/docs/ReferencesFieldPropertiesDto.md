# ReferencesFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **string** | Hints to describe the schema. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder to show when no value has been entered. | [optional] 
**IsRequired** | **bool** | Indicates if the field is required. | [optional] 
**EditorUrl** | Pointer to **string** | Optional url to the editor. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 
**FieldType** | **string** |  | 
**MinItems** | Pointer to **int32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **int32** | The maximum allowed items for the field value. | [optional] 
**AllowDuplicates** | **bool** | True, if duplicate values are allowed. | [optional] 
**ResolveReference** | **bool** | True to resolve references in the content list. | [optional] 
**Editor** | [**OneOfReferencesFieldEditor**](oneOf&lt;ReferencesFieldEditor&gt;.md) | The editor that is used to manage this field. | [optional] 
**SchemaIds** | Pointer to **[]string** | The id of the referenced schemas. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


