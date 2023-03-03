# \UsersApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetUser**](UsersApi.md#UsersGetUser) | **Get** /api/users/{id} | Get user by id.
[**UsersGetUserPicture**](UsersApi.md#UsersGetUserPicture) | **Get** /api/users/{id}/picture | Get user picture by id.
[**UsersGetUserResources**](UsersApi.md#UsersGetUserResources) | **Get** /api | Get the user resources.
[**UsersGetUsers**](UsersApi.md#UsersGetUsers) | **Get** /api/users | Get users by query.



## UsersGetUser

> UserDto UsersGetUser(ctx, id)

Get user by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the user (GUID). | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUserPicture

> *os.File UsersGetUserPicture(ctx, id)

Get user picture by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the user (GUID). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUserResources

> ResourcesDto UsersGetUserResources(ctx, )

Get the user resources.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ResourcesDto**](ResourcesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUsers

> []UserDto UsersGetUsers(ctx, optional)

Get users by query.

Search the user by query that contains the email address or the part of the email address.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The query to search the user by email address. Case invariant. | 

### Return type

[**[]UserDto**](UserDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

