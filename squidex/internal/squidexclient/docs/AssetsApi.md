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

> *os.File AssetContentGetAssetContent(ctx, id).Version(version).Cache(cache).Download(download).Width(width).Height(height).Quality(quality).Mode(mode).Bg(bg).FocusX(focusX).FocusY(focusY).Nofocus(nofocus).Auto(auto).Force(force).Format(format).Execute()

Get the asset content.

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
    id := "id_example" // string | The id of the asset.
    version := int64(789) // int64 | The optional version of the asset. (optional)
    cache := int64(789) // int64 | The cache duration in seconds. (optional)
    download := int32(56) // int32 | Set it to 0 to prevent download. (optional)
    width := int32(56) // int32 | The target width of the asset, if it is an image. (optional)
    height := int32(56) // int32 | The target height of the asset, if it is an image. (optional)
    quality := int32(56) // int32 | Optional image quality, it is is an jpeg image. (optional)
    mode := openapiclient.ResizeMode("Crop") // ResizeMode | The resize mode when the width and height is defined. (optional)
    bg := "bg_example" // string | Optional background color. (optional)
    focusX := float32(3.4) // float32 | Override the y focus point. (optional)
    focusY := float32(3.4) // float32 | Override the x focus point. (optional)
    nofocus := true // bool | True to ignore the asset focus point if any. (optional)
    auto := true // bool | True to use auto format. (optional)
    force := true // bool | True to force a new resize even if it already stored. (optional)
    format := openapiclient.ImageFormat("AVIF") // ImageFormat | True to force a new resize even if it already stored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetContentGetAssetContent(context.Background(), id).Version(version).Cache(cache).Download(download).Width(width).Height(height).Quality(quality).Mode(mode).Bg(bg).FocusX(focusX).FocusY(focusY).Nofocus(nofocus).Auto(auto).Force(force).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetContentGetAssetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetContentGetAssetContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetContentGetAssetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetContentGetAssetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int64** | The optional version of the asset. | 
 **cache** | **int64** | The cache duration in seconds. | 
 **download** | **int32** | Set it to 0 to prevent download. | 
 **width** | **int32** | The target width of the asset, if it is an image. | 
 **height** | **int32** | The target height of the asset, if it is an image. | 
 **quality** | **int32** | Optional image quality, it is is an jpeg image. | 
 **mode** | [**ResizeMode**](ResizeMode.md) | The resize mode when the width and height is defined. | 
 **bg** | **string** | Optional background color. | 
 **focusX** | **float32** | Override the y focus point. | 
 **focusY** | **float32** | Override the x focus point. | 
 **nofocus** | **bool** | True to ignore the asset focus point if any. | 
 **auto** | **bool** | True to use auto format. | 
 **force** | **bool** | True to force a new resize even if it already stored. | 
 **format** | [**ImageFormat**](ImageFormat.md) | True to force a new resize even if it already stored. | 

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

> *os.File AssetContentGetAssetContentBySlug(ctx, app, idOrSlug, more).Version(version).Cache(cache).Download(download).Width(width).Height(height).Quality(quality).Mode(mode).Bg(bg).FocusX(focusX).FocusY(focusY).Nofocus(nofocus).Auto(auto).Force(force).Format(format).Execute()

Get the asset content.

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
    idOrSlug := "idOrSlug_example" // string | The id or slug of the asset.
    more := "more_example" // string | Optional suffix that can be used to seo-optimize the link to the image Has not effect.
    version := int64(789) // int64 | The optional version of the asset. (optional)
    cache := int64(789) // int64 | The cache duration in seconds. (optional)
    download := int32(56) // int32 | Set it to 0 to prevent download. (optional)
    width := int32(56) // int32 | The target width of the asset, if it is an image. (optional)
    height := int32(56) // int32 | The target height of the asset, if it is an image. (optional)
    quality := int32(56) // int32 | Optional image quality, it is is an jpeg image. (optional)
    mode := openapiclient.ResizeMode("Crop") // ResizeMode | The resize mode when the width and height is defined. (optional)
    bg := "bg_example" // string | Optional background color. (optional)
    focusX := float32(3.4) // float32 | Override the y focus point. (optional)
    focusY := float32(3.4) // float32 | Override the x focus point. (optional)
    nofocus := true // bool | True to ignore the asset focus point if any. (optional)
    auto := true // bool | True to use auto format. (optional)
    force := true // bool | True to force a new resize even if it already stored. (optional)
    format := openapiclient.ImageFormat("AVIF") // ImageFormat | True to force a new resize even if it already stored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetContentGetAssetContentBySlug(context.Background(), app, idOrSlug, more).Version(version).Cache(cache).Download(download).Width(width).Height(height).Quality(quality).Mode(mode).Bg(bg).FocusX(focusX).FocusY(focusY).Nofocus(nofocus).Auto(auto).Force(force).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetContentGetAssetContentBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetContentGetAssetContentBySlug`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetContentGetAssetContentBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**idOrSlug** | **string** | The id or slug of the asset. | 
**more** | **string** | Optional suffix that can be used to seo-optimize the link to the image Has not effect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetContentGetAssetContentBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **int64** | The optional version of the asset. | 
 **cache** | **int64** | The cache duration in seconds. | 
 **download** | **int32** | Set it to 0 to prevent download. | 
 **width** | **int32** | The target width of the asset, if it is an image. | 
 **height** | **int32** | The target height of the asset, if it is an image. | 
 **quality** | **int32** | Optional image quality, it is is an jpeg image. | 
 **mode** | [**ResizeMode**](ResizeMode.md) | The resize mode when the width and height is defined. | 
 **bg** | **string** | Optional background color. | 
 **focusX** | **float32** | Override the y focus point. | 
 **focusY** | **float32** | Override the x focus point. | 
 **nofocus** | **bool** | True to ignore the asset focus point if any. | 
 **auto** | **bool** | True to use auto format. | 
 **force** | **bool** | True to force a new resize even if it already stored. | 
 **format** | [**ImageFormat**](ImageFormat.md) | True to force a new resize even if it already stored. | 

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

> AssetFoldersDeleteAssetFolder(ctx, app, id).Execute()

Delete an asset folder.

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
    id := "id_example" // string | The id of the asset folder to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetFoldersDeleteAssetFolder(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetFoldersDeleteAssetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetFoldersDeleteAssetFolderRequest struct via the builder pattern


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


## AssetFoldersGetAssetFolders

> AssetFoldersDto AssetFoldersGetAssetFolders(ctx, app).ParentId(parentId).Scope(scope).Execute()

Get asset folders.



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
    parentId := "parentId_example" // string | The optional parent folder id. (optional)
    scope := openapiclient.AssetFolderScope("PathAndItems") // AssetFolderScope | The scope of the query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetFoldersGetAssetFolders(context.Background(), app).ParentId(parentId).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetFoldersGetAssetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetFoldersGetAssetFolders`: AssetFoldersDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetFoldersGetAssetFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetFoldersGetAssetFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | The optional parent folder id. | 
 **scope** | [**AssetFolderScope**](AssetFolderScope.md) | The scope of the query. | 

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

> AssetFolderDto AssetFoldersPostAssetFolder(ctx, app).CreateAssetFolderDto(createAssetFolderDto).Execute()

Create an asset folder.

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
    createAssetFolderDto := *openapiclient.NewCreateAssetFolderDto("FolderName_example") // CreateAssetFolderDto | The asset folder object that needs to be added to the App.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetFoldersPostAssetFolder(context.Background(), app).CreateAssetFolderDto(createAssetFolderDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetFoldersPostAssetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetFoldersPostAssetFolder`: AssetFolderDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetFoldersPostAssetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetFoldersPostAssetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAssetFolderDto** | [**CreateAssetFolderDto**](CreateAssetFolderDto.md) | The asset folder object that needs to be added to the App. | 

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

> AssetFolderDto AssetFoldersPutAssetFolder(ctx, app, id).RenameAssetFolderDto(renameAssetFolderDto).Execute()

Update an asset folder.

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
    id := "id_example" // string | The id of the asset folder.
    renameAssetFolderDto := *openapiclient.NewRenameAssetFolderDto("FolderName_example") // RenameAssetFolderDto | The asset folder object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetFoldersPutAssetFolder(context.Background(), app, id).RenameAssetFolderDto(renameAssetFolderDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetFoldersPutAssetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetFoldersPutAssetFolder`: AssetFolderDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetFoldersPutAssetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetFoldersPutAssetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renameAssetFolderDto** | [**RenameAssetFolderDto**](RenameAssetFolderDto.md) | The asset folder object that needs to updated. | 

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

> AssetFolderDto AssetFoldersPutAssetFolderParent(ctx, app, id).MoveAssetFolderDto(moveAssetFolderDto).Execute()

Move an asset folder.

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
    id := "id_example" // string | The id of the asset folder.
    moveAssetFolderDto := *openapiclient.NewMoveAssetFolderDto() // MoveAssetFolderDto | The asset folder object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetFoldersPutAssetFolderParent(context.Background(), app, id).MoveAssetFolderDto(moveAssetFolderDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetFoldersPutAssetFolderParent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetFoldersPutAssetFolderParent`: AssetFolderDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetFoldersPutAssetFolderParent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetFoldersPutAssetFolderParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **moveAssetFolderDto** | [**MoveAssetFolderDto**](MoveAssetFolderDto.md) | The asset folder object that needs to updated. | 

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

> []BulkResultDto AssetsBulkUpdateAssets(ctx, app).BulkUpdateAssetsDto(bulkUpdateAssetsDto).Execute()

Bulk update assets.

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
    bulkUpdateAssetsDto := *openapiclient.NewBulkUpdateAssetsDto([]openapiclient.BulkUpdateAssetsJobDto{*openapiclient.NewBulkUpdateAssetsJobDto()}) // BulkUpdateAssetsDto | The bulk update request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsBulkUpdateAssets(context.Background(), app).BulkUpdateAssetsDto(bulkUpdateAssetsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsBulkUpdateAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsBulkUpdateAssets`: []BulkResultDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsBulkUpdateAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsBulkUpdateAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateAssetsDto** | [**BulkUpdateAssetsDto**](BulkUpdateAssetsDto.md) | The bulk update request. | 

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

> AssetsDeleteAsset(ctx, app, id).CheckReferrers(checkReferrers).Permanent(permanent).Execute()

Delete an asset.

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
    id := "id_example" // string | The id of the asset to delete.
    checkReferrers := true // bool | True to check referrers of this asset. (optional)
    permanent := true // bool | True to delete the asset permanently. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsDeleteAsset(context.Background(), app, id).CheckReferrers(checkReferrers).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsDeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsDeleteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checkReferrers** | **bool** | True to check referrers of this asset. | 
 **permanent** | **bool** | True to delete the asset permanently. | 

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

> AssetDto AssetsGetAsset(ctx, app, id).Execute()

Get an asset by id.

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
    id := "id_example" // string | The id of the asset to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsGetAsset(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsGetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsGetAsset`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsGetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> AssetsDto AssetsGetAssets(ctx, app).ParentId(parentId).Ids(ids).Q(q).Top(top).Skip(skip).Orderby(orderby).Filter(filter).Execute()

Get assets.



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
    parentId := "parentId_example" // string | The optional parent folder id. (optional)
    ids := "ids_example" // string | The optional asset ids. (optional)
    q := "q_example" // string | The optional json query. (optional)
    top := float32(8.14) // float32 | Optional number of contents to take. (optional)
    skip := float32(8.14) // float32 | Optional number of contents to skip. (optional)
    orderby := "orderby_example" // string | Optional OData order definition. (optional)
    filter := "filter_example" // string | Optional OData filter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsGetAssets(context.Background(), app).ParentId(parentId).Ids(ids).Q(q).Top(top).Skip(skip).Orderby(orderby).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsGetAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsGetAssets`: AssetsDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsGetAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | The optional parent folder id. | 
 **ids** | **string** | The optional asset ids. | 
 **q** | **string** | The optional json query. | 
 **top** | **float32** | Optional number of contents to take. | 
 **skip** | **float32** | Optional number of contents to skip. | 
 **orderby** | **string** | Optional OData order definition. | 
 **filter** | **string** | Optional OData filter. | 

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

> AssetsDto AssetsGetAssetsPost(ctx, app).QueryDto(queryDto).Execute()

Get assets.



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
    queryDto := *openapiclient.NewQueryDto() // QueryDto | The required query object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsGetAssetsPost(context.Background(), app).QueryDto(queryDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsGetAssetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsGetAssetsPost`: AssetsDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsGetAssetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetAssetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryDto** | [**QueryDto**](QueryDto.md) | The required query object. | 

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

> map[string]int32 AssetsGetTags(ctx, app).Execute()

Get assets tags.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsGetTags(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsGetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsGetTags`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsGetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> AssetDto AssetsPostAsset(ctx, app).ParentId(parentId).Id(id).Duplicate(duplicate).File(file).Execute()

Upload a new asset.



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
    parentId := "parentId_example" // string | The optional parent folder id. (optional)
    id := "id_example" // string | The optional custom asset id. (optional)
    duplicate := true // bool | True to duplicate the asset, event if the file has been uploaded. (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPostAsset(context.Background(), app).ParentId(parentId).Id(id).Duplicate(duplicate).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPostAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPostAsset`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPostAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPostAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | The optional parent folder id. | 
 **id** | **string** | The optional custom asset id. | 
 **duplicate** | **bool** | True to duplicate the asset, event if the file has been uploaded. | 
 **file** | ***os.File** |  | 

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

> AssetDto AssetsPostUpsertAsset(ctx, app, id).ParentId(parentId).Duplicate(duplicate).File(file).Execute()

Upsert an asset.



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
    id := "id_example" // string | The optional custom asset id.
    parentId := "parentId_example" // string | The optional parent folder id. (optional)
    duplicate := true // bool | True to duplicate the asset, event if the file has been uploaded. (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPostUpsertAsset(context.Background(), app, id).ParentId(parentId).Duplicate(duplicate).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPostUpsertAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPostUpsertAsset`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPostUpsertAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The optional custom asset id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPostUpsertAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentId** | **string** | The optional parent folder id. | 
 **duplicate** | **bool** | True to duplicate the asset, event if the file has been uploaded. | 
 **file** | ***os.File** |  | 

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

> AssetDto AssetsPutAsset(ctx, app, id).AnnotateAssetDto(annotateAssetDto).Execute()

Update an asset.

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
    id := "id_example" // string | The id of the asset.
    annotateAssetDto := *openapiclient.NewAnnotateAssetDto() // AnnotateAssetDto | The asset object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPutAsset(context.Background(), app, id).AnnotateAssetDto(annotateAssetDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPutAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPutAsset`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPutAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPutAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **annotateAssetDto** | [**AnnotateAssetDto**](AnnotateAssetDto.md) | The asset object that needs to updated. | 

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

> AssetDto AssetsPutAssetContent(ctx, app, id).File(file).Execute()

Replace asset content.



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
    id := "id_example" // string | The id of the asset.
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPutAssetContent(context.Background(), app, id).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPutAssetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPutAssetContent`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPutAssetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPutAssetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

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

> AssetDto AssetsPutAssetParent(ctx, app, id).MoveAssetDto(moveAssetDto).Execute()

Moves the asset.

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
    id := "id_example" // string | The id of the asset.
    moveAssetDto := *openapiclient.NewMoveAssetDto() // MoveAssetDto | The asset object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPutAssetParent(context.Background(), app, id).MoveAssetDto(moveAssetDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPutAssetParent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPutAssetParent`: AssetDto
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPutAssetParent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPutAssetParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **moveAssetDto** | [**MoveAssetDto**](MoveAssetDto.md) | The asset object that needs to updated. | 

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

> map[string]int32 AssetsPutTag(ctx, app, name).RenameTagDto(renameTagDto).Execute()

Rename an asset tag.

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
    name := "name_example" // string | The tag to return.
    renameTagDto := *openapiclient.NewRenameTagDto("TagName_example") // RenameTagDto | The required request object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.AssetsPutTag(context.Background(), app, name).RenameTagDto(renameTagDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsPutTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsPutTag`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsPutTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**name** | **string** | The tag to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPutTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renameTagDto** | [**RenameTagDto**](RenameTagDto.md) | The required request object. | 

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

