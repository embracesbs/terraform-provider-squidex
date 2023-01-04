# NumberFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to **map[string]float64** |  | [optional] 
**DefaultValue** | Pointer to **NullableFloat64** | The default value for the field value. | [optional] 
**MaxValue** | Pointer to **NullableFloat64** | The maximum allowed value for the field value. | [optional] 
**MinValue** | Pointer to **NullableFloat64** | The minimum allowed value for the field value. | [optional] 
**AllowedValues** | Pointer to **[]float64** | The allowed values for the field value. | [optional] 
**IsUnique** | Pointer to **bool** | Indicates if the field value must be unique. Ignored for nested fields and localized fields. | [optional] 
**InlineEditable** | Pointer to **bool** | Indicates that the inline editor is enabled for this field. | [optional] 
**Editor** | Pointer to [**NumberFieldEditor**](NumberFieldEditor.md) |  | [optional] 

## Methods

### NewNumberFieldPropertiesDto

`func NewNumberFieldPropertiesDto() *NumberFieldPropertiesDto`

NewNumberFieldPropertiesDto instantiates a new NumberFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFieldPropertiesDtoWithDefaults

`func NewNumberFieldPropertiesDtoWithDefaults() *NumberFieldPropertiesDto`

NewNumberFieldPropertiesDtoWithDefaults instantiates a new NumberFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *NumberFieldPropertiesDto) GetDefaultValues() map[string]float64`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *NumberFieldPropertiesDto) GetDefaultValuesOk() (*map[string]float64, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *NumberFieldPropertiesDto) SetDefaultValues(v map[string]float64)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *NumberFieldPropertiesDto) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *NumberFieldPropertiesDto) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NumberFieldPropertiesDto) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NumberFieldPropertiesDto) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NumberFieldPropertiesDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NumberFieldPropertiesDto) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NumberFieldPropertiesDto) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMaxValue

`func (o *NumberFieldPropertiesDto) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *NumberFieldPropertiesDto) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *NumberFieldPropertiesDto) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *NumberFieldPropertiesDto) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *NumberFieldPropertiesDto) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *NumberFieldPropertiesDto) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinValue

`func (o *NumberFieldPropertiesDto) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *NumberFieldPropertiesDto) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *NumberFieldPropertiesDto) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *NumberFieldPropertiesDto) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *NumberFieldPropertiesDto) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *NumberFieldPropertiesDto) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetAllowedValues

`func (o *NumberFieldPropertiesDto) GetAllowedValues() []float64`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *NumberFieldPropertiesDto) GetAllowedValuesOk() (*[]float64, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *NumberFieldPropertiesDto) SetAllowedValues(v []float64)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *NumberFieldPropertiesDto) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *NumberFieldPropertiesDto) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *NumberFieldPropertiesDto) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetIsUnique

`func (o *NumberFieldPropertiesDto) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *NumberFieldPropertiesDto) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *NumberFieldPropertiesDto) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *NumberFieldPropertiesDto) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetInlineEditable

`func (o *NumberFieldPropertiesDto) GetInlineEditable() bool`

GetInlineEditable returns the InlineEditable field if non-nil, zero value otherwise.

### GetInlineEditableOk

`func (o *NumberFieldPropertiesDto) GetInlineEditableOk() (*bool, bool)`

GetInlineEditableOk returns a tuple with the InlineEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEditable

`func (o *NumberFieldPropertiesDto) SetInlineEditable(v bool)`

SetInlineEditable sets InlineEditable field to given value.

### HasInlineEditable

`func (o *NumberFieldPropertiesDto) HasInlineEditable() bool`

HasInlineEditable returns a boolean if a field has been set.

### GetEditor

`func (o *NumberFieldPropertiesDto) GetEditor() NumberFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *NumberFieldPropertiesDto) GetEditorOk() (*NumberFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *NumberFieldPropertiesDto) SetEditor(v NumberFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *NumberFieldPropertiesDto) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


