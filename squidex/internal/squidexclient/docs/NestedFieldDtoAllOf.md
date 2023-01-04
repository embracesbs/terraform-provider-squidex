# NestedFieldDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **int64** | The id of the field. | [optional] 
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | Pointer to **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | Pointer to **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | Pointer to **bool** | Defines if the field is disabled. | [optional] 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 

## Methods

### NewNestedFieldDtoAllOf

`func NewNestedFieldDtoAllOf(name string, properties FieldPropertiesDto, ) *NestedFieldDtoAllOf`

NewNestedFieldDtoAllOf instantiates a new NestedFieldDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedFieldDtoAllOfWithDefaults

`func NewNestedFieldDtoAllOfWithDefaults() *NestedFieldDtoAllOf`

NewNestedFieldDtoAllOfWithDefaults instantiates a new NestedFieldDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *NestedFieldDtoAllOf) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *NestedFieldDtoAllOf) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *NestedFieldDtoAllOf) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *NestedFieldDtoAllOf) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetName

`func (o *NestedFieldDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedFieldDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedFieldDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetIsHidden

`func (o *NestedFieldDtoAllOf) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *NestedFieldDtoAllOf) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *NestedFieldDtoAllOf) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *NestedFieldDtoAllOf) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLocked

`func (o *NestedFieldDtoAllOf) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *NestedFieldDtoAllOf) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *NestedFieldDtoAllOf) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *NestedFieldDtoAllOf) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDisabled

`func (o *NestedFieldDtoAllOf) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *NestedFieldDtoAllOf) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *NestedFieldDtoAllOf) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *NestedFieldDtoAllOf) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetProperties

`func (o *NestedFieldDtoAllOf) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NestedFieldDtoAllOf) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NestedFieldDtoAllOf) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


