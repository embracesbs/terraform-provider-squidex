# \StatisticsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsagesGetCurrentStorageSize**](StatisticsApi.md#UsagesGetCurrentStorageSize) | **Get** /api/apps/{app}/usages/storage/today | Get total asset size.
[**UsagesGetLog**](StatisticsApi.md#UsagesGetLog) | **Get** /api/apps/{app}/usages/log | Get api calls as log file.
[**UsagesGetStorageSizes**](StatisticsApi.md#UsagesGetStorageSizes) | **Get** /api/apps/{app}/usages/storage/{fromDate}/{toDate} | Get asset usage by date.
[**UsagesGetUsages**](StatisticsApi.md#UsagesGetUsages) | **Get** /api/apps/{app}/usages/calls/{fromDate}/{toDate} | Get api calls in date range.



## UsagesGetCurrentStorageSize

> CurrentStorageDto UsagesGetCurrentStorageSize(ctx, app).Execute()

Get total asset size.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    app := "app_example" // string | The name of the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticsApi.UsagesGetCurrentStorageSize(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.UsagesGetCurrentStorageSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsagesGetCurrentStorageSize`: CurrentStorageDto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.UsagesGetCurrentStorageSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsagesGetCurrentStorageSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> LogDownloadDto UsagesGetLog(ctx, app).Execute()

Get api calls as log file.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    app := "app_example" // string | The name of the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticsApi.UsagesGetLog(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.UsagesGetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsagesGetLog`: LogDownloadDto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.UsagesGetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsagesGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UsagesGetStorageSizes

> []StorageUsagePerDateDto UsagesGetStorageSizes(ctx, app, fromDate, toDate).Execute()

Get asset usage by date.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    app := "app_example" // string | The name of the app.
    fromDate := "fromDate_example" // string | The from date.
    toDate := "toDate_example" // string | The to date.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticsApi.UsagesGetStorageSizes(context.Background(), app, fromDate, toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.UsagesGetStorageSizes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsagesGetStorageSizes`: []StorageUsagePerDateDto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.UsagesGetStorageSizes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**fromDate** | **string** | The from date. | 
**toDate** | **string** | The to date. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsagesGetStorageSizesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]StorageUsagePerDateDto**](StorageUsagePerDateDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsagesGetUsages

> CallsUsageDtoDto UsagesGetUsages(ctx, app, fromDate, toDate).Execute()

Get api calls in date range.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    app := "app_example" // string | The name of the app.
    fromDate := "fromDate_example" // string | The from date.
    toDate := "toDate_example" // string | The to date.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticsApi.UsagesGetUsages(context.Background(), app, fromDate, toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.UsagesGetUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsagesGetUsages`: CallsUsageDtoDto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.UsagesGetUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**fromDate** | **string** | The from date. | 
**toDate** | **string** | The to date. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsagesGetUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CallsUsageDtoDto**](CallsUsageDtoDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

