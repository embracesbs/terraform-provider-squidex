# \CommentsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentsDeleteComment**](CommentsApi.md#CommentsDeleteComment) | **Delete** /api/apps/{app}/comments/{commentsId}/{commentId} | Delete a comment.
[**CommentsGetComments**](CommentsApi.md#CommentsGetComments) | **Get** /api/apps/{app}/comments/{commentsId} | Get all comments.
[**CommentsGetWatchingUsers**](CommentsApi.md#CommentsGetWatchingUsers) | **Get** /api/apps/{app}/watching/{resource} | Get all watching users..
[**CommentsPostComment**](CommentsApi.md#CommentsPostComment) | **Post** /api/apps/{app}/comments/{commentsId} | Create a new comment.
[**CommentsPutComment**](CommentsApi.md#CommentsPutComment) | **Put** /api/apps/{app}/comments/{commentsId}/{commentId} | Update a comment.



## CommentsDeleteComment

> CommentsDeleteComment(ctx, app, commentsId, commentId).Execute()

Delete a comment.

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
    app := "app_example" // string | The name of the app.
    commentsId := "commentsId_example" // string | The id of the comments.
    commentId := "commentId_example" // string | The id of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentsApi.CommentsDeleteComment(context.Background(), app, commentsId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.CommentsDeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**commentsId** | **string** | The id of the comments. | 
**commentId** | **string** | The id of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsDeleteCommentRequest struct via the builder pattern


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


## CommentsGetComments

> CommentsDto CommentsGetComments(ctx, app, commentsId).Version(version).Execute()

Get all comments.



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
    app := "app_example" // string | The name of the app.
    commentsId := "commentsId_example" // string | The id of the comments.
    version := int64(789) // int64 | The current version. (optional) (default to -2)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentsApi.CommentsGetComments(context.Background(), app, commentsId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.CommentsGetComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommentsGetComments`: CommentsDto
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.CommentsGetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**commentsId** | **string** | The id of the comments. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsGetCommentsRequest struct via the builder pattern


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


## CommentsGetWatchingUsers

> []string CommentsGetWatchingUsers(ctx, app, resource).Execute()

Get all watching users..

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
    app := "app_example" // string | The name of the app.
    resource := "resource_example" // string | The path to the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentsApi.CommentsGetWatchingUsers(context.Background(), app, resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.CommentsGetWatchingUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommentsGetWatchingUsers`: []string
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.CommentsGetWatchingUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**resource** | **string** | The path to the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsGetWatchingUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsPostComment

> CommentDto CommentsPostComment(ctx, app, commentsId).UpsertCommentDto(upsertCommentDto).Execute()

Create a new comment.

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
    app := "app_example" // string | The name of the app.
    commentsId := "commentsId_example" // string | The id of the comments.
    upsertCommentDto := *openapiclient.NewUpsertCommentDto("Text_example") // UpsertCommentDto | The comment object that needs to created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentsApi.CommentsPostComment(context.Background(), app, commentsId).UpsertCommentDto(upsertCommentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.CommentsPostComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommentsPostComment`: CommentDto
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.CommentsPostComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**commentsId** | **string** | The id of the comments. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsPostCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upsertCommentDto** | [**UpsertCommentDto**](UpsertCommentDto.md) | The comment object that needs to created. | 

### Return type

[**CommentDto**](CommentDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsPutComment

> CommentsPutComment(ctx, app, commentsId, commentId).UpsertCommentDto(upsertCommentDto).Execute()

Update a comment.

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
    app := "app_example" // string | The name of the app.
    commentsId := "commentsId_example" // string | The id of the comments.
    commentId := "commentId_example" // string | The id of the comment.
    upsertCommentDto := *openapiclient.NewUpsertCommentDto("Text_example") // UpsertCommentDto | The comment object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentsApi.CommentsPutComment(context.Background(), app, commentsId, commentId).UpsertCommentDto(upsertCommentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.CommentsPutComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**commentsId** | **string** | The id of the comments. | 
**commentId** | **string** | The id of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsPutCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upsertCommentDto** | [**UpsertCommentDto**](UpsertCommentDto.md) | The comment object that needs to updated. | 

### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

