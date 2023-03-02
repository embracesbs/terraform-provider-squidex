# CurrentStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int64** | The size in bytes. | [optional] 
**MaxAllowed** | Pointer to **int64** | The maximum allowed asset size. | [optional] 

## Methods

### NewCurrentStorageDto

`func NewCurrentStorageDto() *CurrentStorageDto`

NewCurrentStorageDto instantiates a new CurrentStorageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentStorageDtoWithDefaults

`func NewCurrentStorageDtoWithDefaults() *CurrentStorageDto`

NewCurrentStorageDtoWithDefaults instantiates a new CurrentStorageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *CurrentStorageDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CurrentStorageDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CurrentStorageDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CurrentStorageDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaxAllowed

`func (o *CurrentStorageDto) GetMaxAllowed() int64`

GetMaxAllowed returns the MaxAllowed field if non-nil, zero value otherwise.

### GetMaxAllowedOk

`func (o *CurrentStorageDto) GetMaxAllowedOk() (*int64, bool)`

GetMaxAllowedOk returns a tuple with the MaxAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowed

`func (o *CurrentStorageDto) SetMaxAllowed(v int64)`

SetMaxAllowed sets MaxAllowed field to given value.

### HasMaxAllowed

`func (o *CurrentStorageDto) HasMaxAllowed() bool`

HasMaxAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


