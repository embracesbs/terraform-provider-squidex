# CallsUsageDtoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCalls** | Pointer to **int64** | The total number of API calls. | [optional] 
**TotalBytes** | Pointer to **int64** | The total number of bytes transferred. | [optional] 
**MonthCalls** | Pointer to **int64** | The total number of API calls this month. | [optional] 
**MonthBytes** | Pointer to **int64** | The total number of bytes transferred this month. | [optional] 
**BlockingApiCalls** | Pointer to **int64** | The amount of calls that will block the app. | [optional] 
**AllowedBytes** | Pointer to **int64** | The included API traffic. | [optional] 
**AllowedCalls** | Pointer to **int64** | The included API calls. | [optional] 
**AverageElapsedMs** | Pointer to **float64** | The average duration in milliseconds. | [optional] 
**Details** | [**map[string][]CallsUsagePerDateDto**](array.md) | The statistics by date and group. | 

## Methods

### NewCallsUsageDtoDto

`func NewCallsUsageDtoDto(details map[string][]CallsUsagePerDateDto, ) *CallsUsageDtoDto`

NewCallsUsageDtoDto instantiates a new CallsUsageDtoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallsUsageDtoDtoWithDefaults

`func NewCallsUsageDtoDtoWithDefaults() *CallsUsageDtoDto`

NewCallsUsageDtoDtoWithDefaults instantiates a new CallsUsageDtoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCalls

`func (o *CallsUsageDtoDto) GetTotalCalls() int64`

GetTotalCalls returns the TotalCalls field if non-nil, zero value otherwise.

### GetTotalCallsOk

`func (o *CallsUsageDtoDto) GetTotalCallsOk() (*int64, bool)`

GetTotalCallsOk returns a tuple with the TotalCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalls

`func (o *CallsUsageDtoDto) SetTotalCalls(v int64)`

SetTotalCalls sets TotalCalls field to given value.

### HasTotalCalls

`func (o *CallsUsageDtoDto) HasTotalCalls() bool`

HasTotalCalls returns a boolean if a field has been set.

### GetTotalBytes

`func (o *CallsUsageDtoDto) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *CallsUsageDtoDto) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *CallsUsageDtoDto) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *CallsUsageDtoDto) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetMonthCalls

`func (o *CallsUsageDtoDto) GetMonthCalls() int64`

GetMonthCalls returns the MonthCalls field if non-nil, zero value otherwise.

### GetMonthCallsOk

`func (o *CallsUsageDtoDto) GetMonthCallsOk() (*int64, bool)`

GetMonthCallsOk returns a tuple with the MonthCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthCalls

`func (o *CallsUsageDtoDto) SetMonthCalls(v int64)`

SetMonthCalls sets MonthCalls field to given value.

### HasMonthCalls

`func (o *CallsUsageDtoDto) HasMonthCalls() bool`

HasMonthCalls returns a boolean if a field has been set.

### GetMonthBytes

`func (o *CallsUsageDtoDto) GetMonthBytes() int64`

GetMonthBytes returns the MonthBytes field if non-nil, zero value otherwise.

### GetMonthBytesOk

`func (o *CallsUsageDtoDto) GetMonthBytesOk() (*int64, bool)`

GetMonthBytesOk returns a tuple with the MonthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthBytes

`func (o *CallsUsageDtoDto) SetMonthBytes(v int64)`

SetMonthBytes sets MonthBytes field to given value.

### HasMonthBytes

`func (o *CallsUsageDtoDto) HasMonthBytes() bool`

HasMonthBytes returns a boolean if a field has been set.

### GetBlockingApiCalls

`func (o *CallsUsageDtoDto) GetBlockingApiCalls() int64`

GetBlockingApiCalls returns the BlockingApiCalls field if non-nil, zero value otherwise.

### GetBlockingApiCallsOk

`func (o *CallsUsageDtoDto) GetBlockingApiCallsOk() (*int64, bool)`

GetBlockingApiCallsOk returns a tuple with the BlockingApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingApiCalls

`func (o *CallsUsageDtoDto) SetBlockingApiCalls(v int64)`

SetBlockingApiCalls sets BlockingApiCalls field to given value.

### HasBlockingApiCalls

`func (o *CallsUsageDtoDto) HasBlockingApiCalls() bool`

HasBlockingApiCalls returns a boolean if a field has been set.

### GetAllowedBytes

`func (o *CallsUsageDtoDto) GetAllowedBytes() int64`

GetAllowedBytes returns the AllowedBytes field if non-nil, zero value otherwise.

### GetAllowedBytesOk

`func (o *CallsUsageDtoDto) GetAllowedBytesOk() (*int64, bool)`

GetAllowedBytesOk returns a tuple with the AllowedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBytes

`func (o *CallsUsageDtoDto) SetAllowedBytes(v int64)`

SetAllowedBytes sets AllowedBytes field to given value.

### HasAllowedBytes

`func (o *CallsUsageDtoDto) HasAllowedBytes() bool`

HasAllowedBytes returns a boolean if a field has been set.

### GetAllowedCalls

`func (o *CallsUsageDtoDto) GetAllowedCalls() int64`

GetAllowedCalls returns the AllowedCalls field if non-nil, zero value otherwise.

### GetAllowedCallsOk

`func (o *CallsUsageDtoDto) GetAllowedCallsOk() (*int64, bool)`

GetAllowedCallsOk returns a tuple with the AllowedCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCalls

`func (o *CallsUsageDtoDto) SetAllowedCalls(v int64)`

SetAllowedCalls sets AllowedCalls field to given value.

### HasAllowedCalls

`func (o *CallsUsageDtoDto) HasAllowedCalls() bool`

HasAllowedCalls returns a boolean if a field has been set.

### GetAverageElapsedMs

`func (o *CallsUsageDtoDto) GetAverageElapsedMs() float64`

GetAverageElapsedMs returns the AverageElapsedMs field if non-nil, zero value otherwise.

### GetAverageElapsedMsOk

`func (o *CallsUsageDtoDto) GetAverageElapsedMsOk() (*float64, bool)`

GetAverageElapsedMsOk returns a tuple with the AverageElapsedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageElapsedMs

`func (o *CallsUsageDtoDto) SetAverageElapsedMs(v float64)`

SetAverageElapsedMs sets AverageElapsedMs field to given value.

### HasAverageElapsedMs

`func (o *CallsUsageDtoDto) HasAverageElapsedMs() bool`

HasAverageElapsedMs returns a boolean if a field has been set.

### GetDetails

`func (o *CallsUsageDtoDto) GetDetails() map[string][]CallsUsagePerDateDto`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CallsUsageDtoDto) GetDetailsOk() (*map[string][]CallsUsagePerDateDto, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CallsUsageDtoDto) SetDetails(v map[string][]CallsUsagePerDateDto)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


