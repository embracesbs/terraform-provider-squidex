# NestedFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**FieldId** | Pointer to **int64** | The id of the field. | [optional] 
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**IsHidden** | Pointer to **bool** | Defines if the field is hidden. | [optional] 
**IsLocked** | Pointer to **bool** | Defines if the field is locked. | [optional] 
**IsDisabled** | Pointer to **bool** | Defines if the field is disabled. | [optional] 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 

## Methods

### NewNestedFieldDto

`func NewNestedFieldDto(links map[string]ResourceLink, name string, properties FieldPropertiesDto, ) *NestedFieldDto`

NewNestedFieldDto instantiates a new NestedFieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedFieldDtoWithDefaults

`func NewNestedFieldDtoWithDefaults() *NestedFieldDto`

NewNestedFieldDtoWithDefaults instantiates a new NestedFieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NestedFieldDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NestedFieldDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NestedFieldDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetFieldId

`func (o *NestedFieldDto) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *NestedFieldDto) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *NestedFieldDto) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *NestedFieldDto) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetName

`func (o *NestedFieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedFieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedFieldDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsHidden

`func (o *NestedFieldDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *NestedFieldDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *NestedFieldDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *NestedFieldDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLocked

`func (o *NestedFieldDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *NestedFieldDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *NestedFieldDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *NestedFieldDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDisabled

`func (o *NestedFieldDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *NestedFieldDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *NestedFieldDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *NestedFieldDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetProperties

`func (o *NestedFieldDto) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NestedFieldDto) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NestedFieldDto) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


