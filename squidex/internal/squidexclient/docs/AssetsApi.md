# \AssetsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetContentGetAssetContentBySlug**](AssetsApi.md#AssetContentGetAssetContentBySlug) | **Get** /assets/{app}/{idOrSlug}/{more} | Get the asset content.
[**AssetFoldersDeleteAssetFolder**](AssetsApi.md#AssetFoldersDeleteAssetFolder) | **Delete** /apps/{app}/assets/folders/{id} | Delete an asset folder.
[**AssetFoldersGetAssetFolders**](AssetsApi.md#AssetFoldersGetAssetFolders) | **Get** /apps/{app}/assets/folders | Get asset folders.
[**AssetFoldersPostAssetFolder**](AssetsApi.md#AssetFoldersPostAssetFolder) | **Post** /apps/{app}/assets/folders | Upload a new asset.
[**AssetFoldersPutAssetFolder**](AssetsApi.md#AssetFoldersPutAssetFolder) | **Put** /apps/{app}/assets/folders/{id} | Updates the asset folder.
[**AssetFoldersPutAssetFolderParent**](AssetsApi.md#AssetFoldersPutAssetFolderParent) | **Put** /apps/{app}/assets/folders/{id}/parent | Moves the asset folder.
[**AssetsDeleteAsset**](AssetsApi.md#AssetsDeleteAsset) | **Delete** /apps/{app}/assets/{id} | Delete an asset.
[**AssetsGetAsset**](AssetsApi.md#AssetsGetAsset) | **Get** /apps/{app}/assets/{id} | Get an asset by id.
[**AssetsGetTags**](AssetsApi.md#AssetsGetTags) | **Get** /apps/{app}/assets/tags | Get assets tags.
[**AssetsPutAsset**](AssetsApi.md#AssetsPutAsset) | **Put** /apps/{app}/assets/{id} | Updates the asset.
[**AssetsPutAssetParent**](AssetsApi.md#AssetsPutAssetParent) | **Put** /apps/{app}/assets/{id}/parent | Moves the asset.



## AssetContentGetAssetContentBySlug

> *os.File AssetContentGetAssetContentBySlug(ctx, app, idOrSlug, more, optional)

Get the asset content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**idOrSlug** | **string**| The id or slug of the asset. | 
**more** | **string**| Optional suffix that can be used to seo-optimize the link to the image Has not effect. | 
 **optional** | ***AssetContentGetAssetContentBySlugOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetContentGetAssetContentBySlugOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **optional.Int64**| The optional version of the asset. | 
 **cache** | **optional.Int64**| The cache duration in seconds. | 
 **download** | **optional.Int32**| Set it to 0 to prevent download. | 
 **width** | **optional.Int32**| The target width of the asset, if it is an image. | 
 **height** | **optional.Int32**| The target height of the asset, if it is an image. | 
 **quality** | **optional.Int32**| Optional image quality, it is is an jpeg image. | 
 **mode** | [**optional.Interface of OneOfResizeMode**](.md)| The resize mode when the width and height is defined. | 
 **focusX** | **optional.Float32**| Override the y focus point. | 
 **focusY** | **optional.Float32**| Override the x focus point. | 
 **nofocus** | **optional.Bool**| True to ignore the asset focus point if any. | 

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


## AssetFoldersDeleteAssetFolder

> AssetFoldersDeleteAssetFolder(ctx, app, id)

Delete an asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset folder to delete. | 

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


## AssetFoldersGetAssetFolders

> AssetsDto AssetFoldersGetAssetFolders(ctx, app, optional)

Get asset folders.

Get all asset folders for the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***AssetFoldersGetAssetFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetFoldersGetAssetFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **optional.String**| The optional parent folder id. | 

### Return type

[**AssetsDto**](AssetsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPostAssetFolder

> AssetDto AssetFoldersPostAssetFolder(ctx, app, createAssetFolderDto)

Upload a new asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createAssetFolderDto** | [**CreateAssetFolderDto**](CreateAssetFolderDto.md)| The asset folder object that needs to be added to the App. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPutAssetFolder

> AssetDto AssetFoldersPutAssetFolder(ctx, app, id, renameAssetFolderDto)

Updates the asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset folder. | 
**renameAssetFolderDto** | [**RenameAssetFolderDto**](RenameAssetFolderDto.md)| The asset folder object that needs to updated. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPutAssetFolderParent

> AssetDto AssetFoldersPutAssetFolderParent(ctx, app, id, moveAssetItemDto)

Moves the asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset folder. | 
**moveAssetItemDto** | [**MoveAssetItemDto**](MoveAssetItemDto.md)| The asset folder object that needs to updated. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsDeleteAsset

> AssetsDeleteAsset(ctx, app, id)

Delete an asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset to delete. | 

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


## AssetsGetAsset

> AssetDto AssetsGetAsset(ctx, app, id)

Get an asset by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset to retrieve. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetTags

> map[string]int32 AssetsGetTags(ctx, app)

Get assets tags.

Get all tags for assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

**map[string]int32**

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPutAsset

> AssetDto AssetsPutAsset(ctx, app, id, annotateAssetDto)

Updates the asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset. | 
**annotateAssetDto** | [**AnnotateAssetDto**](AnnotateAssetDto.md)| The asset object that needs to updated. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPutAssetParent

> AssetDto AssetsPutAssetParent(ctx, app, id, moveAssetItemDto)

Moves the asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset. | 
**moveAssetItemDto** | [**MoveAssetItemDto**](MoveAssetItemDto.md)| The asset object that needs to updated. | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

