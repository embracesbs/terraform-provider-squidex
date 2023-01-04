# ComponentsFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinItems** | Pointer to **NullableInt32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | The maximum allowed items for the field value. | [optional] 
**SchemaIds** | Pointer to **[]string** | The id of the embedded schemas. | [optional] 
**UniqueFields** | Pointer to **[]string** | The fields that must be unique. | [optional] 

## Methods

### NewComponentsFieldPropertiesDto

`func NewComponentsFieldPropertiesDto() *ComponentsFieldPropertiesDto`

NewComponentsFieldPropertiesDto instantiates a new ComponentsFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentsFieldPropertiesDtoWithDefaults

`func NewComponentsFieldPropertiesDtoWithDefaults() *ComponentsFieldPropertiesDto`

NewComponentsFieldPropertiesDtoWithDefaults instantiates a new ComponentsFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinItems

`func (o *ComponentsFieldPropertiesDto) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *ComponentsFieldPropertiesDto) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *ComponentsFieldPropertiesDto) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *ComponentsFieldPropertiesDto) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *ComponentsFieldPropertiesDto) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *ComponentsFieldPropertiesDto) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *ComponentsFieldPropertiesDto) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *ComponentsFieldPropertiesDto) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *ComponentsFieldPropertiesDto) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *ComponentsFieldPropertiesDto) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *ComponentsFieldPropertiesDto) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *ComponentsFieldPropertiesDto) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetSchemaIds

`func (o *ComponentsFieldPropertiesDto) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *ComponentsFieldPropertiesDto) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *ComponentsFieldPropertiesDto) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *ComponentsFieldPropertiesDto) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *ComponentsFieldPropertiesDto) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *ComponentsFieldPropertiesDto) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil
### GetUniqueFields

`func (o *ComponentsFieldPropertiesDto) GetUniqueFields() []string`

GetUniqueFields returns the UniqueFields field if non-nil, zero value otherwise.

### GetUniqueFieldsOk

`func (o *ComponentsFieldPropertiesDto) GetUniqueFieldsOk() (*[]string, bool)`

GetUniqueFieldsOk returns a tuple with the UniqueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueFields

`func (o *ComponentsFieldPropertiesDto) SetUniqueFields(v []string)`

SetUniqueFields sets UniqueFields field to given value.

### HasUniqueFields

`func (o *ComponentsFieldPropertiesDto) HasUniqueFields() bool`

HasUniqueFields returns a boolean if a field has been set.

### SetUniqueFieldsNil

`func (o *ComponentsFieldPropertiesDto) SetUniqueFieldsNil(b bool)`

 SetUniqueFieldsNil sets the value for UniqueFields to be an explicit nil

### UnsetUniqueFields
`func (o *ComponentsFieldPropertiesDto) UnsetUniqueFields()`

UnsetUniqueFields ensures that no value is present for UniqueFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


