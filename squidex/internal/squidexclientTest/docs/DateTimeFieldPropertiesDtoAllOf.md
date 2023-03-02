# DateTimeFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] 
**DefaultValue** | Pointer to **NullableTime** | The default value for the field value. | [optional] 
**MaxValue** | Pointer to **NullableTime** | The maximum allowed value for the field value. | [optional] 
**MinValue** | Pointer to **NullableTime** | The minimum allowed value for the field value. | [optional] 
**Format** | Pointer to **NullableString** | The format pattern when displayed in the UI. | [optional] 
**Editor** | Pointer to [**DateTimeFieldEditor**](DateTimeFieldEditor.md) |  | [optional] 
**CalculatedDefaultValue** | Pointer to [**DateTimeCalculatedDefaultValue**](DateTimeCalculatedDefaultValue.md) |  | [optional] 

## Methods

### NewDateTimeFieldPropertiesDtoAllOf

`func NewDateTimeFieldPropertiesDtoAllOf() *DateTimeFieldPropertiesDtoAllOf`

NewDateTimeFieldPropertiesDtoAllOf instantiates a new DateTimeFieldPropertiesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeFieldPropertiesDtoAllOfWithDefaults

`func NewDateTimeFieldPropertiesDtoAllOfWithDefaults() *DateTimeFieldPropertiesDtoAllOf`

NewDateTimeFieldPropertiesDtoAllOfWithDefaults instantiates a new DateTimeFieldPropertiesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValues() map[string]time.Time`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string]time.Time, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValues(v map[string]time.Time)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *DateTimeFieldPropertiesDtoAllOf) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValue() time.Time`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetDefaultValueOk() (*time.Time, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValue(v time.Time)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *DateTimeFieldPropertiesDtoAllOf) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *DateTimeFieldPropertiesDtoAllOf) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMaxValue

`func (o *DateTimeFieldPropertiesDtoAllOf) GetMaxValue() time.Time`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetMaxValueOk() (*time.Time, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *DateTimeFieldPropertiesDtoAllOf) SetMaxValue(v time.Time)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *DateTimeFieldPropertiesDtoAllOf) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *DateTimeFieldPropertiesDtoAllOf) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *DateTimeFieldPropertiesDtoAllOf) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinValue

`func (o *DateTimeFieldPropertiesDtoAllOf) GetMinValue() time.Time`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetMinValueOk() (*time.Time, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *DateTimeFieldPropertiesDtoAllOf) SetMinValue(v time.Time)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *DateTimeFieldPropertiesDtoAllOf) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *DateTimeFieldPropertiesDtoAllOf) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *DateTimeFieldPropertiesDtoAllOf) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetFormat

`func (o *DateTimeFieldPropertiesDtoAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DateTimeFieldPropertiesDtoAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DateTimeFieldPropertiesDtoAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *DateTimeFieldPropertiesDtoAllOf) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *DateTimeFieldPropertiesDtoAllOf) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetEditor

`func (o *DateTimeFieldPropertiesDtoAllOf) GetEditor() DateTimeFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetEditorOk() (*DateTimeFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *DateTimeFieldPropertiesDtoAllOf) SetEditor(v DateTimeFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *DateTimeFieldPropertiesDtoAllOf) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetCalculatedDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) GetCalculatedDefaultValue() DateTimeCalculatedDefaultValue`

GetCalculatedDefaultValue returns the CalculatedDefaultValue field if non-nil, zero value otherwise.

### GetCalculatedDefaultValueOk

`func (o *DateTimeFieldPropertiesDtoAllOf) GetCalculatedDefaultValueOk() (*DateTimeCalculatedDefaultValue, bool)`

GetCalculatedDefaultValueOk returns a tuple with the CalculatedDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) SetCalculatedDefaultValue(v DateTimeCalculatedDefaultValue)`

SetCalculatedDefaultValue sets CalculatedDefaultValue field to given value.

### HasCalculatedDefaultValue

`func (o *DateTimeFieldPropertiesDtoAllOf) HasCalculatedDefaultValue() bool`

HasCalculatedDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


