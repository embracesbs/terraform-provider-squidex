# TagsFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to **map[string][]string** |  | [optional] 
**DefaultValue** | Pointer to **[]string** | The default value. | [optional] 
**MinItems** | Pointer to **NullableInt32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | The maximum allowed items for the field value. | [optional] 
**AllowedValues** | Pointer to **[]string** | The allowed values for the field value. | [optional] 
**CreateEnum** | Pointer to **bool** | Indicates whether GraphQL Enum should be created. | [optional] 
**Editor** | Pointer to [**TagsFieldEditor**](TagsFieldEditor.md) |  | [optional] 

## Methods

### NewTagsFieldPropertiesDtoAllOf

`func NewTagsFieldPropertiesDtoAllOf() *TagsFieldPropertiesDtoAllOf`

NewTagsFieldPropertiesDtoAllOf instantiates a new TagsFieldPropertiesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsFieldPropertiesDtoAllOfWithDefaults

`func NewTagsFieldPropertiesDtoAllOfWithDefaults() *TagsFieldPropertiesDtoAllOf`

NewTagsFieldPropertiesDtoAllOfWithDefaults instantiates a new TagsFieldPropertiesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *TagsFieldPropertiesDtoAllOf) GetDefaultValues() map[string][]string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *TagsFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string][]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *TagsFieldPropertiesDtoAllOf) SetDefaultValues(v map[string][]string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *TagsFieldPropertiesDtoAllOf) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *TagsFieldPropertiesDtoAllOf) GetDefaultValue() []string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *TagsFieldPropertiesDtoAllOf) GetDefaultValueOk() (*[]string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *TagsFieldPropertiesDtoAllOf) SetDefaultValue(v []string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *TagsFieldPropertiesDtoAllOf) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *TagsFieldPropertiesDtoAllOf) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *TagsFieldPropertiesDtoAllOf) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMinItems

`func (o *TagsFieldPropertiesDtoAllOf) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *TagsFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *TagsFieldPropertiesDtoAllOf) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *TagsFieldPropertiesDtoAllOf) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *TagsFieldPropertiesDtoAllOf) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *TagsFieldPropertiesDtoAllOf) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *TagsFieldPropertiesDtoAllOf) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *TagsFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *TagsFieldPropertiesDtoAllOf) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *TagsFieldPropertiesDtoAllOf) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *TagsFieldPropertiesDtoAllOf) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *TagsFieldPropertiesDtoAllOf) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetAllowedValues

`func (o *TagsFieldPropertiesDtoAllOf) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *TagsFieldPropertiesDtoAllOf) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *TagsFieldPropertiesDtoAllOf) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *TagsFieldPropertiesDtoAllOf) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *TagsFieldPropertiesDtoAllOf) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *TagsFieldPropertiesDtoAllOf) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetCreateEnum

`func (o *TagsFieldPropertiesDtoAllOf) GetCreateEnum() bool`

GetCreateEnum returns the CreateEnum field if non-nil, zero value otherwise.

### GetCreateEnumOk

`func (o *TagsFieldPropertiesDtoAllOf) GetCreateEnumOk() (*bool, bool)`

GetCreateEnumOk returns a tuple with the CreateEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnum

`func (o *TagsFieldPropertiesDtoAllOf) SetCreateEnum(v bool)`

SetCreateEnum sets CreateEnum field to given value.

### HasCreateEnum

`func (o *TagsFieldPropertiesDtoAllOf) HasCreateEnum() bool`

HasCreateEnum returns a boolean if a field has been set.

### GetEditor

`func (o *TagsFieldPropertiesDtoAllOf) GetEditor() TagsFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *TagsFieldPropertiesDtoAllOf) GetEditorOk() (*TagsFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *TagsFieldPropertiesDtoAllOf) SetEditor(v TagsFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *TagsFieldPropertiesDtoAllOf) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


