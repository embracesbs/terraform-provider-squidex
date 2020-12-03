# \StatisticsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsagesGetCurrentStorageSize**](StatisticsApi.md#UsagesGetCurrentStorageSize) | **Get** /apps/{app}/usages/storage/today | Get total asset size.
[**UsagesGetLog**](StatisticsApi.md#UsagesGetLog) | **Get** /apps/{app}/usages/log | Get api calls as log file.
[**UsagesGetMonthlyCalls**](StatisticsApi.md#UsagesGetMonthlyCalls) | **Get** /apps/{app}/usages/calls/month | Get api calls for this month.
[**UsagesGetStorageSizes**](StatisticsApi.md#UsagesGetStorageSizes) | **Get** /apps/{app}/usages/storage/{fromDate}/{toDate} | Get asset usage by date.
[**UsagesGetUsages**](StatisticsApi.md#UsagesGetUsages) | **Get** /apps/{app}/usages/calls/{fromDate}/{toDate} | Get api calls in date range.



## UsagesGetCurrentStorageSize

> CurrentStorageDto UsagesGetCurrentStorageSize(ctx, app)

Get total asset size.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**CurrentStorageDto**](CurrentStorageDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsagesGetLog

> LogDownloadDto UsagesGetLog(ctx, app)

Get api calls as log file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**LogDownloadDto**](LogDownloadDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsagesGetMonthlyCalls

> CurrentCallsDto UsagesGetMonthlyCalls(ctx, app)

Get api calls for this month.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**CurrentCallsDto**](CurrentCallsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsagesGetStorageSizes

> []StorageUsageDto UsagesGetStorageSizes(ctx, app, fromDate, toDate)

Get asset usage by date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**fromDate** | **string**| The from date. | 
**toDate** | **string**| The to date. | 

### Return type

[**[]StorageUsageDto**](StorageUsageDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsagesGetUsages

> map[string][]CallsUsageDto UsagesGetUsages(ctx, app, fromDate, toDate)

Get api calls in date range.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**fromDate** | **string**| The from date. | 
**toDate** | **string**| The to date. | 

### Return type

[**map[string][]CallsUsageDto**](array.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

