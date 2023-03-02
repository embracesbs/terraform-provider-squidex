# FieldDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **int64** | The id of the field. | [optional] 
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | Pointer to **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | Pointer to **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | Pointer to **bool** | Defines if the field is disabled. | [optional] 
**Partitioning** | **string** | Defines the partitioning of the field. | 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 
**Nested** | Pointer to [**[]NestedFieldDto**](NestedFieldDto.md) | The nested fields. | [optional] 

## Methods

### NewFieldDtoAllOf

`func NewFieldDtoAllOf(name string, partitioning string, properties FieldPropertiesDto, ) *FieldDtoAllOf`

NewFieldDtoAllOf instantiates a new FieldDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDtoAllOfWithDefaults

`func NewFieldDtoAllOfWithDefaults() *FieldDtoAllOf`

NewFieldDtoAllOfWithDefaults instantiates a new FieldDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *FieldDtoAllOf) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *FieldDtoAllOf) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *FieldDtoAllOf) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *FieldDtoAllOf) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetName

`func (o *FieldDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetIsHidden

`func (o *FieldDtoAllOf) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *FieldDtoAllOf) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *FieldDtoAllOf) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *FieldDtoAllOf) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLocked

`func (o *FieldDtoAllOf) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *FieldDtoAllOf) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *FieldDtoAllOf) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *FieldDtoAllOf) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDisabled

`func (o *FieldDtoAllOf) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *FieldDtoAllOf) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *FieldDtoAllOf) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *FieldDtoAllOf) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetPartitioning

`func (o *FieldDtoAllOf) GetPartitioning() string`

GetPartitioning returns the Partitioning field if non-nil, zero value otherwise.

### GetPartitioningOk

`func (o *FieldDtoAllOf) GetPartitioningOk() (*string, bool)`

GetPartitioningOk returns a tuple with the Partitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioning

`func (o *FieldDtoAllOf) SetPartitioning(v string)`

SetPartitioning sets Partitioning field to given value.


### GetProperties

`func (o *FieldDtoAllOf) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FieldDtoAllOf) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FieldDtoAllOf) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.


### GetNested

`func (o *FieldDtoAllOf) GetNested() []NestedFieldDto`

GetNested returns the Nested field if non-nil, zero value otherwise.

### GetNestedOk

`func (o *FieldDtoAllOf) GetNestedOk() (*[]NestedFieldDto, bool)`

GetNestedOk returns a tuple with the Nested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNested

`func (o *FieldDtoAllOf) SetNested(v []NestedFieldDto)`

SetNested sets Nested field to given value.

### HasNested

`func (o *FieldDtoAllOf) HasNested() bool`

HasNested returns a boolean if a field has been set.

### SetNestedNil

`func (o *FieldDtoAllOf) SetNestedNil(b bool)`

 SetNestedNil sets the value for Nested to be an explicit nil

### UnsetNested
`func (o *FieldDtoAllOf) UnsetNested()`

UnsetNested ensures that no value is present for Nested, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


