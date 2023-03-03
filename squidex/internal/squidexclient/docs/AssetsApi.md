# \AssetsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetContentGetAssetContent**](AssetsApi.md#AssetContentGetAssetContent) | **Get** /api/assets/{id} | Get the asset content.
[**AssetContentGetAssetContentBySlug**](AssetsApi.md#AssetContentGetAssetContentBySlug) | **Get** /api/assets/{app}/{idOrSlug}/{more} | Get the asset content.
[**AssetFoldersDeleteAssetFolder**](AssetsApi.md#AssetFoldersDeleteAssetFolder) | **Delete** /api/apps/{app}/assets/folders/{id} | Delete an asset folder.
[**AssetFoldersGetAssetFolders**](AssetsApi.md#AssetFoldersGetAssetFolders) | **Get** /api/apps/{app}/assets/folders | Get asset folders.
[**AssetFoldersPostAssetFolder**](AssetsApi.md#AssetFoldersPostAssetFolder) | **Post** /api/apps/{app}/assets/folders | Create an asset folder.
[**AssetFoldersPutAssetFolder**](AssetsApi.md#AssetFoldersPutAssetFolder) | **Put** /api/apps/{app}/assets/folders/{id} | Update an asset folder.
[**AssetFoldersPutAssetFolderParent**](AssetsApi.md#AssetFoldersPutAssetFolderParent) | **Put** /api/apps/{app}/assets/folders/{id}/parent | Move an asset folder.
[**AssetsBulkUpdateAssets**](AssetsApi.md#AssetsBulkUpdateAssets) | **Post** /api/apps/{app}/assets/bulk | Bulk update assets.
[**AssetsDeleteAsset**](AssetsApi.md#AssetsDeleteAsset) | **Delete** /api/apps/{app}/assets/{id} | Delete an asset.
[**AssetsGetAsset**](AssetsApi.md#AssetsGetAsset) | **Get** /api/apps/{app}/assets/{id} | Get an asset by id.
[**AssetsGetAssets**](AssetsApi.md#AssetsGetAssets) | **Get** /api/apps/{app}/assets | Get assets.
[**AssetsGetAssetsPost**](AssetsApi.md#AssetsGetAssetsPost) | **Post** /api/apps/{app}/assets/query | Get assets.
[**AssetsGetTags**](AssetsApi.md#AssetsGetTags) | **Get** /api/apps/{app}/assets/tags | Get assets tags.
[**AssetsPostAsset**](AssetsApi.md#AssetsPostAsset) | **Post** /api/apps/{app}/assets | Upload a new asset.
[**AssetsPostUpsertAsset**](AssetsApi.md#AssetsPostUpsertAsset) | **Post** /api/apps/{app}/assets/{id} | Upsert an asset.
[**AssetsPutAsset**](AssetsApi.md#AssetsPutAsset) | **Put** /api/apps/{app}/assets/{id} | Update an asset.
[**AssetsPutAssetContent**](AssetsApi.md#AssetsPutAssetContent) | **Put** /api/apps/{app}/assets/{id}/content | Replace asset content.
[**AssetsPutAssetParent**](AssetsApi.md#AssetsPutAssetParent) | **Put** /api/apps/{app}/assets/{id}/parent | Moves the asset.
[**AssetsPutTag**](AssetsApi.md#AssetsPutTag) | **Put** /api/apps/{app}/assets/tags/{name} | Rename an asset tag.



## AssetContentGetAssetContent

> *os.File AssetContentGetAssetContent(ctx, id, optional)

Get the asset content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the asset. | 
 **optional** | ***AssetContentGetAssetContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetContentGetAssetContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int64**| The optional version of the asset. | 
 **cache** | **optional.Int64**| The cache duration in seconds. | 
 **download** | **optional.Int32**| Set it to 0 to prevent download. | 
 **width** | **optional.Int32**| The target width of the asset, if it is an image. | 
 **height** | **optional.Int32**| The target height of the asset, if it is an image. | 
 **quality** | **optional.Int32**| Optional image quality, it is is an jpeg image. | 
 **mode** | [**optional.Interface of ResizeMode**](.md)| The resize mode when the width and height is defined. | 
 **bg** | **optional.String**| Optional background color. | 
 **focusX** | **optional.Float32**| Override the y focus point. | 
 **focusY** | **optional.Float32**| Override the x focus point. | 
 **nofocus** | **optional.Bool**| True to ignore the asset focus point if any. | 
 **auto** | **optional.Bool**| True to use auto format. | 
 **force** | **optional.Bool**| True to force a new resize even if it already stored. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| True to force a new resize even if it already stored. | 

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
 **mode** | [**optional.Interface of ResizeMode**](.md)| The resize mode when the width and height is defined. | 
 **bg** | **optional.String**| Optional background color. | 
 **focusX** | **optional.Float32**| Override the y focus point. | 
 **focusY** | **optional.Float32**| Override the x focus point. | 
 **nofocus** | **optional.Bool**| True to ignore the asset focus point if any. | 
 **auto** | **optional.Bool**| True to use auto format. | 
 **force** | **optional.Bool**| True to force a new resize even if it already stored. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| True to force a new resize even if it already stored. | 

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

> AssetFoldersDto AssetFoldersGetAssetFolders(ctx, app, optional)

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
 **scope** | [**optional.Interface of AssetFolderScope**](.md)| The scope of the query. | 

### Return type

[**AssetFoldersDto**](AssetFoldersDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPostAssetFolder

> AssetFolderDto AssetFoldersPostAssetFolder(ctx, app, createAssetFolderDto)

Create an asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createAssetFolderDto** | [**CreateAssetFolderDto**](CreateAssetFolderDto.md)| The asset folder object that needs to be added to the App. | 

### Return type

[**AssetFolderDto**](AssetFolderDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPutAssetFolder

> AssetFolderDto AssetFoldersPutAssetFolder(ctx, app, id, renameAssetFolderDto)

Update an asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset folder. | 
**renameAssetFolderDto** | [**RenameAssetFolderDto**](RenameAssetFolderDto.md)| The asset folder object that needs to updated. | 

### Return type

[**AssetFolderDto**](AssetFolderDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetFoldersPutAssetFolderParent

> AssetFolderDto AssetFoldersPutAssetFolderParent(ctx, app, id, moveAssetFolderDto)

Move an asset folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset folder. | 
**moveAssetFolderDto** | [**MoveAssetFolderDto**](MoveAssetFolderDto.md)| The asset folder object that needs to updated. | 

### Return type

[**AssetFolderDto**](AssetFolderDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsBulkUpdateAssets

> []BulkResultDto AssetsBulkUpdateAssets(ctx, app, bulkUpdateAssetsDto)

Bulk update assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**bulkUpdateAssetsDto** | [**BulkUpdateAssetsDto**](BulkUpdateAssetsDto.md)| The bulk update request. | 

### Return type

[**[]BulkResultDto**](BulkResultDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsDeleteAsset

> AssetsDeleteAsset(ctx, app, id, optional)

Delete an asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset to delete. | 
 **optional** | ***AssetsDeleteAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetsDeleteAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checkReferrers** | **optional.Bool**| True to check referrers of this asset. | 
 **permanent** | **optional.Bool**| True to delete the asset permanently. | 

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


## AssetsGetAssets

> AssetsDto AssetsGetAssets(ctx, app, optional)

Get assets.

Get all assets for the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***AssetsGetAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetsGetAssetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **optional.String**| The optional parent folder id. | 
 **ids** | **optional.String**| The optional asset ids. | 
 **q** | **optional.String**| The optional json query. | 
 **top** | **optional.Float32**| Optional number of contents to take. | 
 **skip** | **optional.Float32**| Optional number of contents to skip. | 
 **orderby** | **optional.String**| Optional OData order definition. | 
 **filter** | **optional.String**| Optional OData filter. | 

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


## AssetsGetAssetsPost

> AssetsDto AssetsGetAssetsPost(ctx, app, queryDto)

Get assets.

Get all assets for the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**queryDto** | [**QueryDto**](QueryDto.md)| The required query object. | 

### Return type

[**AssetsDto**](AssetsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
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


## AssetsPostAsset

> AssetDto AssetsPostAsset(ctx, app, optional)

Upload a new asset.

You can only upload one file at a time. The mime type of the file is not calculated by Squidex and is required correctly.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***AssetsPostAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetsPostAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **optional.String**| The optional parent folder id. | 
 **id** | **optional.String**| The optional custom asset id. | 
 **duplicate** | **optional.Bool**| True to duplicate the asset, event if the file has been uploaded. | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPostUpsertAsset

> AssetDto AssetsPostUpsertAsset(ctx, app, id, optional)

Upsert an asset.

You can only upload one file at a time. The mime type of the file is not calculated by Squidex and is required correctly.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The optional custom asset id. | 
 **optional** | ***AssetsPostUpsertAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetsPostUpsertAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentId** | **optional.String**| The optional parent folder id. | 
 **duplicate** | **optional.Bool**| True to duplicate the asset, event if the file has been uploaded. | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPutAsset

> AssetDto AssetsPutAsset(ctx, app, id, annotateAssetDto)

Update an asset.

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


## AssetsPutAssetContent

> AssetDto AssetsPutAssetContent(ctx, app, id, optional)

Replace asset content.

Use multipart request to upload an asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset. | 
 **optional** | ***AssetsPutAssetContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetsPutAssetContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**AssetDto**](AssetDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPutAssetParent

> AssetDto AssetsPutAssetParent(ctx, app, id, moveAssetDto)

Moves the asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset. | 
**moveAssetDto** | [**MoveAssetDto**](MoveAssetDto.md)| The asset object that needs to updated. | 

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


## AssetsPutTag

> map[string]int32 AssetsPutTag(ctx, app, name, renameTagDto)

Rename an asset tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The tag to return. | 
**renameTagDto** | [**RenameTagDto**](RenameTagDto.md)| The required request object. | 

### Return type

**map[string]int32**

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

