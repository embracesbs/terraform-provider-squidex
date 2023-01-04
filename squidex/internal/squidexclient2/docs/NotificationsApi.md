# \NotificationsApi

All URIs are relative to *https://squidex.dev.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserNotificationsDeleteComment**](NotificationsApi.md#UserNotificationsDeleteComment) | **Delete** /users/{userId}/notifications/{commentId} | Deletes the notification.
[**UserNotificationsGetNotifications**](NotificationsApi.md#UserNotificationsGetNotifications) | **Get** /users/{userId}/notifications | Get all notifications.



## UserNotificationsDeleteComment

> UserNotificationsDeleteComment(ctx, userId, commentId)

Deletes the notification.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user id. | 
**commentId** | **string**| The id of the comment. | 

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


## UserNotificationsGetNotifications

> CommentsDto UserNotificationsGetNotifications(ctx, userId, optional)

Get all notifications.

When passing in a version you can retrieve all updates since then.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user id. | 
 **optional** | ***UserNotificationsGetNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserNotificationsGetNotificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int64**| The current version. | [default to -2]

### Return type

[**CommentsDto**](CommentsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

