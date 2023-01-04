# ArrayFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinItems** | Pointer to **NullableInt32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | The maximum allowed items for the field value. | [optional] 
**UniqueFields** | Pointer to **[]string** | The fields that must be unique. | [optional] 

## Methods

### NewArrayFieldPropertiesDto

`func NewArrayFieldPropertiesDto() *ArrayFieldPropertiesDto`

NewArrayFieldPropertiesDto instantiates a new ArrayFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayFieldPropertiesDtoWithDefaults

`func NewArrayFieldPropertiesDtoWithDefaults() *ArrayFieldPropertiesDto`

NewArrayFieldPropertiesDtoWithDefaults instantiates a new ArrayFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinItems

`func (o *ArrayFieldPropertiesDto) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *ArrayFieldPropertiesDto) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *ArrayFieldPropertiesDto) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *ArrayFieldPropertiesDto) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *ArrayFieldPropertiesDto) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *ArrayFieldPropertiesDto) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *ArrayFieldPropertiesDto) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *ArrayFieldPropertiesDto) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *ArrayFieldPropertiesDto) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *ArrayFieldPropertiesDto) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *ArrayFieldPropertiesDto) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *ArrayFieldPropertiesDto) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetUniqueFields

`func (o *ArrayFieldPropertiesDto) GetUniqueFields() []string`

GetUniqueFields returns the UniqueFields field if non-nil, zero value otherwise.

### GetUniqueFieldsOk

`func (o *ArrayFieldPropertiesDto) GetUniqueFieldsOk() (*[]string, bool)`

GetUniqueFieldsOk returns a tuple with the UniqueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueFields

`func (o *ArrayFieldPropertiesDto) SetUniqueFields(v []string)`

SetUniqueFields sets UniqueFields field to given value.

### HasUniqueFields

`func (o *ArrayFieldPropertiesDto) HasUniqueFields() bool`

HasUniqueFields returns a boolean if a field has been set.

### SetUniqueFieldsNil

`func (o *ArrayFieldPropertiesDto) SetUniqueFieldsNil(b bool)`

 SetUniqueFieldsNil sets the value for UniqueFields to be an explicit nil

### UnsetUniqueFields
`func (o *ArrayFieldPropertiesDto) UnsetUniqueFields()`

UnsetUniqueFields ensures that no value is present for UniqueFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


