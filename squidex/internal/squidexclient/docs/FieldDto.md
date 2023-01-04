# FieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**FieldId** | Pointer to **int64** | The id of the field. | [optional] 
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | Pointer to **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | Pointer to **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | Pointer to **bool** | Defines if the field is disabled. | [optional] 
**Partitioning** | **string** | Defines the partitioning of the field. | 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 
**Nested** | Pointer to [**[]NestedFieldDto**](NestedFieldDto.md) | The nested fields. | [optional] 

## Methods

### NewFieldDto

`func NewFieldDto(links map[string]ResourceLink, name string, partitioning string, properties FieldPropertiesDto, ) *FieldDto`

NewFieldDto instantiates a new FieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDtoWithDefaults

`func NewFieldDtoWithDefaults() *FieldDto`

NewFieldDtoWithDefaults instantiates a new FieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FieldDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FieldDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FieldDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetFieldId

`func (o *FieldDto) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *FieldDto) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *FieldDto) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *FieldDto) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetName

`func (o *FieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsHidden

`func (o *FieldDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *FieldDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *FieldDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *FieldDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLocked

`func (o *FieldDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *FieldDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *FieldDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *FieldDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDisabled

`func (o *FieldDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *FieldDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *FieldDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *FieldDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetPartitioning

`func (o *FieldDto) GetPartitioning() string`

GetPartitioning returns the Partitioning field if non-nil, zero value otherwise.

### GetPartitioningOk

`func (o *FieldDto) GetPartitioningOk() (*string, bool)`

GetPartitioningOk returns a tuple with the Partitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioning

`func (o *FieldDto) SetPartitioning(v string)`

SetPartitioning sets Partitioning field to given value.


### GetProperties

`func (o *FieldDto) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FieldDto) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FieldDto) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.


### GetNested

`func (o *FieldDto) GetNested() []NestedFieldDto`

GetNested returns the Nested field if non-nil, zero value otherwise.

### GetNestedOk

`func (o *FieldDto) GetNestedOk() (*[]NestedFieldDto, bool)`

GetNestedOk returns a tuple with the Nested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNested

`func (o *FieldDto) SetNested(v []NestedFieldDto)`

SetNested sets Nested field to given value.

### HasNested

`func (o *FieldDto) HasNested() bool`

HasNested returns a boolean if a field has been set.

### SetNestedNil

`func (o *FieldDto) SetNestedNil(b bool)`

 SetNestedNil sets the value for Nested to be an explicit nil

### UnsetNested
`func (o *FieldDto) UnsetNested()`

UnsetNested ensures that no value is present for Nested, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


