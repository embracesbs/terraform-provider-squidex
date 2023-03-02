# AddFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. Must be unique within the schema. | 
**Partitioning** | Pointer to **NullableString** | Determines the optional partitioning of the field. | [optional] 
**Properties** | [**FieldPropertiesDto**](FieldPropertiesDto.md) |  | 

## Methods

### NewAddFieldDto

`func NewAddFieldDto(name string, properties FieldPropertiesDto, ) *AddFieldDto`

NewAddFieldDto instantiates a new AddFieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFieldDtoWithDefaults

`func NewAddFieldDtoWithDefaults() *AddFieldDto`

NewAddFieldDtoWithDefaults instantiates a new AddFieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddFieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddFieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddFieldDto) SetName(v string)`

SetName sets Name field to given value.


### GetPartitioning

`func (o *AddFieldDto) GetPartitioning() string`

GetPartitioning returns the Partitioning field if non-nil, zero value otherwise.

### GetPartitioningOk

`func (o *AddFieldDto) GetPartitioningOk() (*string, bool)`

GetPartitioningOk returns a tuple with the Partitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioning

`func (o *AddFieldDto) SetPartitioning(v string)`

SetPartitioning sets Partitioning field to given value.

### HasPartitioning

`func (o *AddFieldDto) HasPartitioning() bool`

HasPartitioning returns a boolean if a field has been set.

### SetPartitioningNil

`func (o *AddFieldDto) SetPartitioningNil(b bool)`

 SetPartitioningNil sets the value for Partitioning to be an explicit nil

### UnsetPartitioning
`func (o *AddFieldDto) UnsetPartitioning()`

UnsetPartitioning ensures that no value is present for Partitioning, not even an explicit nil
### GetProperties

`func (o *AddFieldDto) GetProperties() FieldPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AddFieldDto) GetPropertiesOk() (*FieldPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AddFieldDto) SetProperties(v FieldPropertiesDto)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


