# CallsUsageDtoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCalls** | **int64** | The total number of API calls. | [optional] 
**TotalBytes** | **int64** | The total number of bytes transferred. | [optional] 
**MonthCalls** | **int64** | The total number of API calls this month. | [optional] 
**MonthBytes** | **int64** | The total number of bytes transferred this month. | [optional] 
**BlockingApiCalls** | **int64** | The amount of calls that will block the app. | [optional] 
**AllowedBytes** | **int64** | The included API traffic. | [optional] 
**AllowedCalls** | **int64** | The included API calls. | [optional] 
**AverageElapsedMs** | **float64** | The average duration in milliseconds. | [optional] 
**Details** | [**map[string][]CallsUsagePerDateDto**](array.md) | The statistics by date and group. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


