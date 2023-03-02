# \TemplatesApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesGetTemplate**](TemplatesApi.md#TemplatesGetTemplate) | **Get** /api/templates/{name} | Get template details.
[**TemplatesGetTemplates**](TemplatesApi.md#TemplatesGetTemplates) | **Get** /api/templates | Get all templates.



## TemplatesGetTemplate

> TemplateDetailsDto TemplatesGetTemplate(ctx, name).Execute()

Get template details.

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
    name := "name_example" // string | The name of the template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetTemplate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetTemplate`: TemplateDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateDetailsDto**](TemplateDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesGetTemplates

> TemplatesDto TemplatesGetTemplates(ctx).Execute()

Get all templates.

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
    resp, r, err := apiClient.TemplatesApi.TemplatesGetTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetTemplates`: TemplatesDto
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetTemplatesRequest struct via the builder pattern


### Return type

[**TemplatesDto**](TemplatesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

