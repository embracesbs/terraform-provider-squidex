# NumberFieldPropertiesDto

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
**DefaultValues** | **map[string]float64** |  | [optional] 
**DefaultValue** | Pointer to **float64** | The default value for the field value. | [optional] 
**MaxValue** | Pointer to **float64** | The maximum allowed value for the field value. | [optional] 
**MinValue** | Pointer to **float64** | The minimum allowed value for the field value. | [optional] 
**AllowedValues** | Pointer to **[]float64** | The allowed values for the field value. | [optional] 
**IsUnique** | **bool** | Indicates if the field value must be unique. Ignored for nested fields and localized fields. | [optional] 
**InlineEditable** | **bool** | Indicates that the inline editor is enabled for this field. | [optional] 
**Editor** | [**NumberFieldEditor**](NumberFieldEditor.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


