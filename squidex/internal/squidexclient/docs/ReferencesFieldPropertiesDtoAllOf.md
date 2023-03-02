# ReferencesFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


