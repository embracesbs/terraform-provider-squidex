# DateTimeFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **string** | Hints to describe the schema. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder to show when no value has been entered. | [optional] 
**IsRequired** | **bool** | Indicates if the field is required. | [optional] 
**IsHalfWidth** | **bool** | Indicates if the field should be rendered with half width only. | [optional] 
**EditorUrl** | Pointer to **string** | Optional url to the editor. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 
**FieldType** | **string** |  | 
**DefaultValue** | Pointer to [**time.Time**](time.Time.md) | The default value for the field value. | [optional] 
**MaxValue** | Pointer to [**time.Time**](time.Time.md) | The maximum allowed value for the field value. | [optional] 
**MinValue** | Pointer to [**time.Time**](time.Time.md) | The minimum allowed value for the field value. | [optional] 
**Editor** | [**OneOfDateTimeFieldEditor**](oneOf&lt;DateTimeFieldEditor&gt;.md) | The editor that is used to manage this field. | [optional] 
**CalculatedDefaultValue** | Pointer to [**OneOfDateTimeCalculatedDefaultValue**](oneOf&lt;DateTimeCalculatedDefaultValue&gt;.md) | The calculated default value for the field value. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


