# \HistoryApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HistoryGetHistory**](HistoryApi.md#HistoryGetHistory) | **Get** /api/apps/{app}/history | Get historical events.



## HistoryGetHistory

> []HistoryEventDto HistoryGetHistory(ctx, app, optional)

Get historical events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***HistoryGetHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HistoryGetHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **optional.String**| The name of the channel. | 

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

