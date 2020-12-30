# \SearchApi

All URIs are relative to *https://squidex.dev.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGetSearchResults**](SearchApi.md#SearchGetSearchResults) | **Get** /apps/{app}/search | Get search results.



## SearchGetSearchResults

> []SearchResultDto SearchGetSearchResults(ctx, app, optional)

Get search results.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***SearchGetSearchResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchGetSearchResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The search query. | 

### Return type

[**[]SearchResultDto**](SearchResultDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

