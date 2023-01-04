# \DiagnosticsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnosticsGetDump**](DiagnosticsApi.md#DiagnosticsGetDump) | **Get** /api/diagnostics/dump | Creates a dump and writes it into storage..
[**DiagnosticsGetGCDump**](DiagnosticsApi.md#DiagnosticsGetGCDump) | **Get** /api/diagnostics/gcdump | Creates a gc dump and writes it into storage.



## DiagnosticsGetDump

> DiagnosticsGetDump(ctx).Execute()

Creates a dump and writes it into storage..

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiagnosticsApi.DiagnosticsGetDump(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticsApi.DiagnosticsGetDump``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiagnosticsGetDumpRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiagnosticsGetGCDump

> DiagnosticsGetGCDump(ctx).Execute()

Creates a gc dump and writes it into storage.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiagnosticsApi.DiagnosticsGetGCDump(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticsApi.DiagnosticsGetGCDump``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiagnosticsGetGCDumpRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

