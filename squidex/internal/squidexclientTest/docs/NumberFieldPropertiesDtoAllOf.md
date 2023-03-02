# NumberFieldPropertiesDtoAllOf

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

### NewNumberFieldPropertiesDtoAllOf

`func NewNumberFieldPropertiesDtoAllOf() *NumberFieldPropertiesDtoAllOf`

NewNumberFieldPropertiesDtoAllOf instantiates a new NumberFieldPropertiesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFieldPropertiesDtoAllOfWithDefaults

`func NewNumberFieldPropertiesDtoAllOfWithDefaults() *NumberFieldPropertiesDtoAllOf`

NewNumberFieldPropertiesDtoAllOfWithDefaults instantiates a new NumberFieldPropertiesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *NumberFieldPropertiesDtoAllOf) GetDefaultValues() map[string]float64`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *NumberFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string]float64, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *NumberFieldPropertiesDtoAllOf) SetDefaultValues(v map[string]float64)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *NumberFieldPropertiesDtoAllOf) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *NumberFieldPropertiesDtoAllOf) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NumberFieldPropertiesDtoAllOf) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NumberFieldPropertiesDtoAllOf) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NumberFieldPropertiesDtoAllOf) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NumberFieldPropertiesDtoAllOf) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NumberFieldPropertiesDtoAllOf) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMaxValue

`func (o *NumberFieldPropertiesDtoAllOf) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *NumberFieldPropertiesDtoAllOf) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *NumberFieldPropertiesDtoAllOf) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *NumberFieldPropertiesDtoAllOf) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *NumberFieldPropertiesDtoAllOf) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *NumberFieldPropertiesDtoAllOf) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinValue

`func (o *NumberFieldPropertiesDtoAllOf) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *NumberFieldPropertiesDtoAllOf) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *NumberFieldPropertiesDtoAllOf) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *NumberFieldPropertiesDtoAllOf) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *NumberFieldPropertiesDtoAllOf) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *NumberFieldPropertiesDtoAllOf) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetAllowedValues

`func (o *NumberFieldPropertiesDtoAllOf) GetAllowedValues() []float64`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *NumberFieldPropertiesDtoAllOf) GetAllowedValuesOk() (*[]float64, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *NumberFieldPropertiesDtoAllOf) SetAllowedValues(v []float64)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *NumberFieldPropertiesDtoAllOf) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *NumberFieldPropertiesDtoAllOf) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *NumberFieldPropertiesDtoAllOf) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetIsUnique

`func (o *NumberFieldPropertiesDtoAllOf) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *NumberFieldPropertiesDtoAllOf) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *NumberFieldPropertiesDtoAllOf) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *NumberFieldPropertiesDtoAllOf) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetInlineEditable

`func (o *NumberFieldPropertiesDtoAllOf) GetInlineEditable() bool`

GetInlineEditable returns the InlineEditable field if non-nil, zero value otherwise.

### GetInlineEditableOk

`func (o *NumberFieldPropertiesDtoAllOf) GetInlineEditableOk() (*bool, bool)`

GetInlineEditableOk returns a tuple with the InlineEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEditable

`func (o *NumberFieldPropertiesDtoAllOf) SetInlineEditable(v bool)`

SetInlineEditable sets InlineEditable field to given value.

### HasInlineEditable

`func (o *NumberFieldPropertiesDtoAllOf) HasInlineEditable() bool`

HasInlineEditable returns a boolean if a field has been set.

### GetEditor

`func (o *NumberFieldPropertiesDtoAllOf) GetEditor() NumberFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *NumberFieldPropertiesDtoAllOf) GetEditorOk() (*NumberFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *NumberFieldPropertiesDtoAllOf) SetEditor(v NumberFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *NumberFieldPropertiesDtoAllOf) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


