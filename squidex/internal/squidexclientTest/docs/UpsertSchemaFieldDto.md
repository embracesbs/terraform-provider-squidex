# UpsertSchemaFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | Pointer to **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | Pointer to **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | Pointer to **bool** | Defines if the field is disabled. | [optional] 
**Partitioning** | Pointer to **NullableString** | Determines the optional partitioning of the field. | [optional] 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 
**Nested** | Pointer to [**[]UpsertSchemaNestedFieldDto**](UpsertSchemaNestedFieldDto.md) | The nested fields. | [optional] 

## Methods

### NewUpsertSchemaFieldDto

`func NewUpsertSchemaFieldDto(name string, properties FieldPropertiesDto, ) *UpsertSchemaFieldDto`

NewUpsertSchemaFieldDto instantiates a new UpsertSchemaFieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertSchemaFieldDtoWithDefaults

`func NewUpsertSchemaFieldDtoWithDefaults() *UpsertSchemaFieldDto`

NewUpsertSchemaFieldDtoWithDefaults instantiates a new UpsertSchemaFieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpsertSchemaFieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertSchemaFieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertSchemaFieldDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsHidden

`func (o *UpsertSchemaFieldDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UpsertSchemaFieldDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UpsertSchemaFieldDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UpsertSchemaFieldDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLocked

`func (o *UpsertSchemaFieldDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UpsertSchemaFieldDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UpsertSchemaFieldDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UpsertSchemaFieldDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UpsertSchemaFieldDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UpsertSchemaFieldDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UpsertSchemaFieldDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UpsertSchemaFieldDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetPartitioning

`func (o *UpsertSchemaFieldDto) GetPartitioning() string`

GetPartitioning returns the Partitioning field if non-nil, zero value otherwise.

### GetPartitioningOk

`func (o *UpsertSchemaFieldDto) GetPartitioningOk() (*string, bool)`

GetPartitioningOk returns a tuple with the Partitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioning

`func (o *UpsertSchemaFieldDto) SetPartitioning(v string)`

SetPartitioning sets Partitioning field to given value.

### HasPartitioning

`func (o *UpsertSchemaFieldDto) HasPartitioning() bool`

HasPartitioning returns a boolean if a field has been set.

### SetPartitioningNil

`func (o *UpsertSchemaFieldDto) SetPartitioningNil(b bool)`

 SetPartitioningNil sets the value for Partitioning to be an explicit nil

### UnsetPartitioning
`func (o *UpsertSchemaFieldDto) UnsetPartitioning()`

UnsetPartitioning ensures that no value is present for Partitioning, not even an explicit nil
### GetProperties

`func (o *UpsertSchemaFieldDto) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpsertSchemaFieldDto) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpsertSchemaFieldDto) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.


### GetNested

`func (o *UpsertSchemaFieldDto) GetNested() []UpsertSchemaNestedFieldDto`

GetNested returns the Nested field if non-nil, zero value otherwise.

### GetNestedOk

`func (o *UpsertSchemaFieldDto) GetNestedOk() (*[]UpsertSchemaNestedFieldDto, bool)`

GetNestedOk returns a tuple with the Nested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNested

`func (o *UpsertSchemaFieldDto) SetNested(v []UpsertSchemaNestedFieldDto)`

SetNested sets Nested field to given value.

### HasNested

`func (o *UpsertSchemaFieldDto) HasNested() bool`

HasNested returns a boolean if a field has been set.

### SetNestedNil

`func (o *UpsertSchemaFieldDto) SetNestedNil(b bool)`

 SetNestedNil sets the value for Nested to be an explicit nil

### UnsetNested
`func (o *UpsertSchemaFieldDto) UnsetNested()`

UnsetNested ensures that no value is present for Nested, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


