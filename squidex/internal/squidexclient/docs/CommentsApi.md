# \CommentsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentsDeleteComment**](CommentsApi.md#CommentsDeleteComment) | **Delete** /apps/{app}/comments/{commentsId}/{commentId} | Deletes the comment.
[**CommentsGetComments**](CommentsApi.md#CommentsGetComments) | **Get** /apps/{app}/comments/{commentsId} | Get all comments.
[**CommentsPostComment**](CommentsApi.md#CommentsPostComment) | **Post** /apps/{app}/comments/{commentsId} | Create a new comment.
[**CommentsPutComment**](CommentsApi.md#CommentsPutComment) | **Put** /apps/{app}/comments/{commentsId}/{commentId} | Updates the comment.



## CommentsDeleteComment

> CommentsDeleteComment(ctx, app, commentsId, commentId)

Deletes the comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**commentsId** | **string**| The id of the comments. | 
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


## CommentsGetComments

> CommentsDto CommentsGetComments(ctx, app, commentsId, optional)

Get all comments.

When passing in a version you can retrieve all updates since then.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**commentsId** | **string**| The id of the comments. | 
 **optional** | ***CommentsGetCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommentsGetCommentsOpts struct


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


## CommentsPostComment

> CommentDto CommentsPostComment(ctx, app, commentsId, upsertCommentDto)

Create a new comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**commentsId** | **string**| The id of the comments. | 
**upsertCommentDto** | [**UpsertCommentDto**](UpsertCommentDto.md)| The comment object that needs to created. | 

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

> CommentsPutComment(ctx, app, commentsId, commentId, upsertCommentDto)

Updates the comment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**commentsId** | **string**| The id of the comments. | 
**commentId** | **string**| The id of the comment. | 
**upsertCommentDto** | [**UpsertCommentDto**](UpsertCommentDto.md)| The comment object that needs to updated. | 

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

