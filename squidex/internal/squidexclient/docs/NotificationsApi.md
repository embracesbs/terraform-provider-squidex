# \NotificationsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserNotificationsDeleteComment**](NotificationsApi.md#UserNotificationsDeleteComment) | **Delete** /api/users/{userId}/notifications/{commentId} | Delete a notification.
[**UserNotificationsGetNotifications**](NotificationsApi.md#UserNotificationsGetNotifications) | **Get** /api/users/{userId}/notifications | Get all notifications.



## UserNotificationsDeleteComment

> UserNotificationsDeleteComment(ctx, userId, commentId).Execute()

Delete a notification.

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
    userId := "userId_example" // string | The user id.
    commentId := "commentId_example" // string | The id of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UserNotificationsDeleteComment(context.Background(), userId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UserNotificationsDeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 
**commentId** | **string** | The id of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserNotificationsDeleteCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> CommentsDto UserNotificationsGetNotifications(ctx, userId).Version(version).Execute()

Get all notifications.



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
    userId := "userId_example" // string | The user id.
    version := int64(789) // int64 | The current version. (optional) (default to -2)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.UserNotificationsGetNotifications(context.Background(), userId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UserNotificationsGetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserNotificationsGetNotifications`: CommentsDto
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UserNotificationsGetNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserNotificationsGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int64** | The current version. | [default to -2]

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

