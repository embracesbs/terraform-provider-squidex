# \TranslationsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslationsGetLanguages**](TranslationsApi.md#TranslationsGetLanguages) | **Post** /apps/{app}/translations | Translate a text.



## TranslationsGetLanguages

> TranslationDto TranslationsGetLanguages(ctx, app, translateDto)

Translate a text.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**translateDto** | [**TranslateDto**](TranslateDto.md)| The translation request. | 

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

