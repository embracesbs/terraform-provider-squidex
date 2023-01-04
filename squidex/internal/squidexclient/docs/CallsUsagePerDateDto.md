# CallsUsagePerDateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date when the usage was tracked. | [optional] 
**TotalCalls** | Pointer to **int64** | The total number of API calls. | [optional] 
**TotalBytes** | Pointer to **int64** | The total number of bytes transferred. | [optional] 
**AverageElapsedMs** | Pointer to **float64** | The average duration in milliseconds. | [optional] 

## Methods

### NewCallsUsagePerDateDto

`func NewCallsUsagePerDateDto() *CallsUsagePerDateDto`

NewCallsUsagePerDateDto instantiates a new CallsUsagePerDateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallsUsagePerDateDtoWithDefaults

`func NewCallsUsagePerDateDtoWithDefaults() *CallsUsagePerDateDto`

NewCallsUsagePerDateDtoWithDefaults instantiates a new CallsUsagePerDateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CallsUsagePerDateDto) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CallsUsagePerDateDto) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CallsUsagePerDateDto) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CallsUsagePerDateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTotalCalls

`func (o *CallsUsagePerDateDto) GetTotalCalls() int64`

GetTotalCalls returns the TotalCalls field if non-nil, zero value otherwise.

### GetTotalCallsOk

`func (o *CallsUsagePerDateDto) GetTotalCallsOk() (*int64, bool)`

GetTotalCallsOk returns a tuple with the TotalCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalls

`func (o *CallsUsagePerDateDto) SetTotalCalls(v int64)`

SetTotalCalls sets TotalCalls field to given value.

### HasTotalCalls

`func (o *CallsUsagePerDateDto) HasTotalCalls() bool`

HasTotalCalls returns a boolean if a field has been set.

### GetTotalBytes

`func (o *CallsUsagePerDateDto) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *CallsUsagePerDateDto) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *CallsUsagePerDateDto) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *CallsUsagePerDateDto) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetAverageElapsedMs

`func (o *CallsUsagePerDateDto) GetAverageElapsedMs() float64`

GetAverageElapsedMs returns the AverageElapsedMs field if non-nil, zero value otherwise.

### GetAverageElapsedMsOk

`func (o *CallsUsagePerDateDto) GetAverageElapsedMsOk() (*float64, bool)`

GetAverageElapsedMsOk returns a tuple with the AverageElapsedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageElapsedMs

`func (o *CallsUsagePerDateDto) SetAverageElapsedMs(v float64)`

SetAverageElapsedMs sets AverageElapsedMs field to given value.

### HasAverageElapsedMs

`func (o *CallsUsagePerDateDto) HasAverageElapsedMs() bool`

HasAverageElapsedMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


