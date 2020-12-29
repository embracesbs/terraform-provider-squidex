# \NewsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NewsGetLanguages**](NewsApi.md#NewsGetLanguages) | **Get** /news/features | Get features since version.



## NewsGetLanguages

> FeaturesDto NewsGetLanguages(ctx, optional)

Get features since version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NewsGetLanguagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NewsGetLanguagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **optional.Int32**| The latest received version. | [default to 0]

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

