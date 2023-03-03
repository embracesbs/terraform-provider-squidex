# DateTimeFieldPropertiesDto

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
**DefaultValues** | [**map[string]time.Time**](time.Time.md) |  | [optional] 
**DefaultValue** | Pointer to [**time.Time**](time.Time.md) | The default value for the field value. | [optional] 
**MaxValue** | Pointer to [**time.Time**](time.Time.md) | The maximum allowed value for the field value. | [optional] 
**MinValue** | Pointer to [**time.Time**](time.Time.md) | The minimum allowed value for the field value. | [optional] 
**Format** | Pointer to **string** | The format pattern when displayed in the UI. | [optional] 
**Editor** | [**DateTimeFieldEditor**](DateTimeFieldEditor.md) |  | [optional] 
**CalculatedDefaultValue** | [**DateTimeCalculatedDefaultValue**](DateTimeCalculatedDefaultValue.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


