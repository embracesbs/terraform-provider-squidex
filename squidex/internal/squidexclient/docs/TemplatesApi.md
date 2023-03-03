# \TemplatesApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesGetTemplate**](TemplatesApi.md#TemplatesGetTemplate) | **Get** /api/templates/{name} | Get template details.
[**TemplatesGetTemplates**](TemplatesApi.md#TemplatesGetTemplates) | **Get** /api/templates | Get all templates.



## TemplatesGetTemplate

> TemplateDetailsDto TemplatesGetTemplate(ctx, name)

Get template details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the template. | 

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

> TemplatesDto TemplatesGetTemplates(ctx, )

Get all templates.

### Required Parameters

This endpoint does not need any parameter.

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

