# \NewsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NewsGetNews**](NewsApi.md#NewsGetNews) | **Get** /api/news/features | Get features since version.



## NewsGetNews

> FeaturesDto NewsGetNews(ctx).Version(version).Execute()

Get features since version.

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
    version := int32(56) // int32 | The latest received version. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NewsApi.NewsGetNews(context.Background()).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewsApi.NewsGetNews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewsGetNews`: FeaturesDto
    fmt.Fprintf(os.Stdout, "Response from `NewsApi.NewsGetNews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewsGetNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **int32** | The latest received version. | [default to 0]

### Return type

[**FeaturesDto**](FeaturesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

