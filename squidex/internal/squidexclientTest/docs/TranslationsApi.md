# \TranslationsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslationsPostTranslation**](TranslationsApi.md#TranslationsPostTranslation) | **Post** /api/apps/{app}/translations | Translate a text.



## TranslationsPostTranslation

> TranslationDto TranslationsPostTranslation(ctx, app).TranslateDto(translateDto).Execute()

Translate a text.

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
    translateDto := *openapiclient.NewTranslateDto("Text_example", "TargetLanguage_example") // TranslateDto | The translation request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationsApi.TranslationsPostTranslation(context.Background(), app).TranslateDto(translateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.TranslationsPostTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslationsPostTranslation`: TranslationDto
    fmt.Fprintf(os.Stdout, "Response from `TranslationsApi.TranslationsPostTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslationsPostTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **translateDto** | [**TranslateDto**](TranslateDto.md) | The translation request. | 

### Return type

[**TranslationDto**](TranslationDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

