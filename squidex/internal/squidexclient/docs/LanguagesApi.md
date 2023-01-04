# \LanguagesApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LanguagesGetLanguages**](LanguagesApi.md#LanguagesGetLanguages) | **Get** /api/languages | Get supported languages.



## LanguagesGetLanguages

> []LanguageDto LanguagesGetLanguages(ctx).Execute()

Get supported languages.



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
    resp, r, err := apiClient.LanguagesApi.LanguagesGetLanguages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.LanguagesGetLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LanguagesGetLanguages`: []LanguageDto
    fmt.Fprintf(os.Stdout, "Response from `LanguagesApi.LanguagesGetLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLanguagesGetLanguagesRequest struct via the builder pattern


### Return type

[**[]LanguageDto**](LanguageDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

