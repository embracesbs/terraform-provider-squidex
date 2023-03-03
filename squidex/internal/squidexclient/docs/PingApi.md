# \PingApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingGetAppPing**](PingApi.md#PingGetAppPing) | **Get** /api/ping/{app} | Get ping status.
[**PingGetInfo**](PingApi.md#PingGetInfo) | **Get** /api/info | Get API information.
[**PingGetPing**](PingApi.md#PingGetPing) | **Get** /api/ping | Get ping status of the API.



## PingGetAppPing

> PingGetAppPing(ctx, app)

Get ping status.

Can be used to test, if the Squidex API is alive and responding.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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


## PingGetInfo

> map[string]string PingGetInfo(ctx, )

Get API information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingGetPing

> PingGetPing(ctx, )

Get ping status of the API.

Can be used to test, if the Squidex API is alive and responding.

### Required Parameters

This endpoint does not need any parameter.

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

