# \HistoryApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HistoryGetHistory**](HistoryApi.md#HistoryGetHistory) | **Get** /api/apps/{app}/history | Get historical events.



## HistoryGetHistory

> []HistoryEventDto HistoryGetHistory(ctx, app).Channel(channel).Execute()

Get historical events.

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
    channel := "channel_example" // string | The name of the channel. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.HistoryGetHistory(context.Background(), app).Channel(channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.HistoryGetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HistoryGetHistory`: []HistoryEventDto
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.HistoryGetHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoryGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** | The name of the channel. | 

### Return type

[**[]HistoryEventDto**](HistoryEventDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

