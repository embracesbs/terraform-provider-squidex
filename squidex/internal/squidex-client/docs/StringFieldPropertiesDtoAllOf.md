# StringFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | The default value for the field value. | [optional] 
**Pattern** | Pointer to **string** | The pattern to enforce a specific format for the field value. | [optional] 
**PatternMessage** | Pointer to **string** | The validation message for the pattern. | [optional] 
**MinLength** | Pointer to **int32** | The minimum allowed length for the field value. | [optional] 
**MaxLength** | Pointer to **int32** | The maximum allowed length for the field value. | [optional] 
**AllowedValues** | Pointer to **[]string** | The allowed values for the field value. | [optional] 
**IsUnique** | **bool** | Indicates if the field value must be unique. Ignored for nested fields and localized fields. | [optional] 
**InlineEditable** | **bool** | Indicates that the inline editor is enabled for this field. | [optional] 
**Editor** | [**OneOfStringFieldEditor**](oneOf&lt;StringFieldEditor&gt;.md) | The editor that is used to manage this field. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


