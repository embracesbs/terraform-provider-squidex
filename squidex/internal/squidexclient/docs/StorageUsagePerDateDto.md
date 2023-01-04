# StorageUsagePerDateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date when the usage was tracked. | [optional] 
**TotalCount** | Pointer to **int64** | The number of assets. | [optional] 
**TotalSize** | Pointer to **int64** | The size in bytes. | [optional] 

## Methods

### NewStorageUsagePerDateDto

`func NewStorageUsagePerDateDto() *StorageUsagePerDateDto`

NewStorageUsagePerDateDto instantiates a new StorageUsagePerDateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageUsagePerDateDtoWithDefaults

`func NewStorageUsagePerDateDtoWithDefaults() *StorageUsagePerDateDto`

NewStorageUsagePerDateDtoWithDefaults instantiates a new StorageUsagePerDateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *StorageUsagePerDateDto) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StorageUsagePerDateDto) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StorageUsagePerDateDto) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StorageUsagePerDateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTotalCount

`func (o *StorageUsagePerDateDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *StorageUsagePerDateDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *StorageUsagePerDateDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *StorageUsagePerDateDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalSize

`func (o *StorageUsagePerDateDto) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *StorageUsagePerDateDto) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *StorageUsagePerDateDto) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *StorageUsagePerDateDto) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


