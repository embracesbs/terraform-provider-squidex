# \AppsApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppAssetsGetAssetScripts**](AppsApi.md#AppAssetsGetAssetScripts) | **Get** /api/apps/{app}/assets/scripts | Get the app asset scripts.
[**AppAssetsPutAssetScripts**](AppsApi.md#AppAssetsPutAssetScripts) | **Put** /api/apps/{app}/assets/scripts | Update the app asset scripts.
[**AppClientsDeleteClient**](AppsApi.md#AppClientsDeleteClient) | **Delete** /api/apps/{app}/clients/{id} | Revoke an app client.
[**AppClientsGetClients**](AppsApi.md#AppClientsGetClients) | **Get** /api/apps/{app}/clients | Get app clients.
[**AppClientsPostClient**](AppsApi.md#AppClientsPostClient) | **Post** /api/apps/{app}/clients | Create a new app client.
[**AppClientsPutClient**](AppsApi.md#AppClientsPutClient) | **Put** /api/apps/{app}/clients/{id} | Updates an app client.
[**AppContributorsDeleteContributor**](AppsApi.md#AppContributorsDeleteContributor) | **Delete** /api/apps/{app}/contributors/{id} | Remove contributor.
[**AppContributorsDeleteMyself**](AppsApi.md#AppContributorsDeleteMyself) | **Delete** /api/apps/{app}/contributors/me | Remove yourself.
[**AppContributorsGetContributors**](AppsApi.md#AppContributorsGetContributors) | **Get** /api/apps/{app}/contributors | Get app contributors.
[**AppContributorsPostContributor**](AppsApi.md#AppContributorsPostContributor) | **Post** /api/apps/{app}/contributors | Assign contributor to app.
[**AppImageGetImage**](AppsApi.md#AppImageGetImage) | **Get** /api/apps/{app}/image | Get the app image.
[**AppLanguagesDeleteLanguage**](AppsApi.md#AppLanguagesDeleteLanguage) | **Delete** /api/apps/{app}/languages/{language} | Deletes an app language.
[**AppLanguagesGetLanguages**](AppsApi.md#AppLanguagesGetLanguages) | **Get** /api/apps/{app}/languages | Get app languages.
[**AppLanguagesPostLanguage**](AppsApi.md#AppLanguagesPostLanguage) | **Post** /api/apps/{app}/languages | Attaches an app language.
[**AppLanguagesPutLanguage**](AppsApi.md#AppLanguagesPutLanguage) | **Put** /api/apps/{app}/languages/{language} | Updates an app language.
[**AppRolesDeleteRole**](AppsApi.md#AppRolesDeleteRole) | **Delete** /api/apps/{app}/roles/{roleName} | Remove role from app.
[**AppRolesGetPermissions**](AppsApi.md#AppRolesGetPermissions) | **Get** /api/apps/{app}/roles/permissions | Get app permissions.
[**AppRolesGetRoles**](AppsApi.md#AppRolesGetRoles) | **Get** /api/apps/{app}/roles | Get app roles.
[**AppRolesPostRole**](AppsApi.md#AppRolesPostRole) | **Post** /api/apps/{app}/roles | Add role to app.
[**AppRolesPutRole**](AppsApi.md#AppRolesPutRole) | **Put** /api/apps/{app}/roles/{roleName} | Update an app role.
[**AppSettingsGetSettings**](AppsApi.md#AppSettingsGetSettings) | **Get** /api/apps/{app}/settings | Get the app settings.
[**AppSettingsPutSettings**](AppsApi.md#AppSettingsPutSettings) | **Put** /api/apps/{app}/settings | Update the app settings.
[**AppWorkflowsDeleteWorkflow**](AppsApi.md#AppWorkflowsDeleteWorkflow) | **Delete** /api/apps/{app}/workflows/{id} | Delete a workflow.
[**AppWorkflowsGetWorkflows**](AppsApi.md#AppWorkflowsGetWorkflows) | **Get** /api/apps/{app}/workflows | Get app workflow.
[**AppWorkflowsPostWorkflow**](AppsApi.md#AppWorkflowsPostWorkflow) | **Post** /api/apps/{app}/workflows | Create a workflow.
[**AppWorkflowsPutWorkflow**](AppsApi.md#AppWorkflowsPutWorkflow) | **Put** /api/apps/{app}/workflows/{id} | Update a workflow.
[**AppsDeleteApp**](AppsApi.md#AppsDeleteApp) | **Delete** /api/apps/{app} | Delete the app.
[**AppsDeleteImage**](AppsApi.md#AppsDeleteImage) | **Delete** /api/apps/{app}/image | Remove the app image.
[**AppsGetApp**](AppsApi.md#AppsGetApp) | **Get** /api/apps/{app} | Get an app by name.
[**AppsGetApps**](AppsApi.md#AppsGetApps) | **Get** /api/apps | Get your apps.
[**AppsPostApp**](AppsApi.md#AppsPostApp) | **Post** /api/apps | Create a new app.
[**AppsPutApp**](AppsApi.md#AppsPutApp) | **Put** /api/apps/{app} | Update the app.
[**AppsUploadImage**](AppsApi.md#AppsUploadImage) | **Post** /api/apps/{app}/image | Upload the app image.



## AppAssetsGetAssetScripts

> AssetScriptsDto AppAssetsGetAssetScripts(ctx, app).Execute()

Get the app asset scripts.

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
    app := "app_example" // string | The name of the app to get the asset scripts for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppAssetsGetAssetScripts(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppAssetsGetAssetScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAssetsGetAssetScripts`: AssetScriptsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppAssetsGetAssetScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to get the asset scripts for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAssetsGetAssetScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetScriptsDto**](AssetScriptsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAssetsPutAssetScripts

> AssetScriptsDto AppAssetsPutAssetScripts(ctx, app).UpdateAssetScriptsDto(updateAssetScriptsDto).Execute()

Update the app asset scripts.

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
    app := "app_example" // string | The name of the app to update.
    updateAssetScriptsDto := *openapiclient.NewUpdateAssetScriptsDto() // UpdateAssetScriptsDto | The values to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppAssetsPutAssetScripts(context.Background(), app).UpdateAssetScriptsDto(updateAssetScriptsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppAssetsPutAssetScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAssetsPutAssetScripts`: AssetScriptsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppAssetsPutAssetScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAssetsPutAssetScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAssetScriptsDto** | [**UpdateAssetScriptsDto**](UpdateAssetScriptsDto.md) | The values to update. | 

### Return type

[**AssetScriptsDto**](AssetScriptsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClientsDeleteClient

> ClientsDto AppClientsDeleteClient(ctx, app, id).Execute()

Revoke an app client.



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
    id := "id_example" // string | The id of the client that must be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppClientsDeleteClient(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppClientsDeleteClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClientsDeleteClient`: ClientsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppClientsDeleteClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the client that must be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClientsDeleteClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientsDto**](ClientsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClientsGetClients

> ClientsDto AppClientsGetClients(ctx, app).Execute()

Get app clients.



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
    resp, r, err := apiClient.AppsApi.AppClientsGetClients(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppClientsGetClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClientsGetClients`: ClientsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppClientsGetClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClientsGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientsDto**](ClientsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClientsPostClient

> ClientsDto AppClientsPostClient(ctx, app).CreateClientDto(createClientDto).Execute()

Create a new app client.



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
    createClientDto := *openapiclient.NewCreateClientDto("Id_example") // CreateClientDto | Client object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppClientsPostClient(context.Background(), app).CreateClientDto(createClientDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppClientsPostClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClientsPostClient`: ClientsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppClientsPostClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClientsPostClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createClientDto** | [**CreateClientDto**](CreateClientDto.md) | Client object that needs to be added to the app. | 

### Return type

[**ClientsDto**](ClientsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClientsPutClient

> ClientsDto AppClientsPutClient(ctx, app, id).UpdateClientDto(updateClientDto).Execute()

Updates an app client.



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
    id := "id_example" // string | The id of the client that must be updated.
    updateClientDto := *openapiclient.NewUpdateClientDto() // UpdateClientDto | Client object that needs to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppClientsPutClient(context.Background(), app, id).UpdateClientDto(updateClientDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppClientsPutClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClientsPutClient`: ClientsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppClientsPutClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the client that must be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClientsPutClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateClientDto** | [**UpdateClientDto**](UpdateClientDto.md) | Client object that needs to be updated. | 

### Return type

[**ClientsDto**](ClientsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppContributorsDeleteContributor

> ContributorsDto AppContributorsDeleteContributor(ctx, app, id).Execute()

Remove contributor.

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
    id := "id_example" // string | The id of the contributor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppContributorsDeleteContributor(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppContributorsDeleteContributor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppContributorsDeleteContributor`: ContributorsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppContributorsDeleteContributor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the contributor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppContributorsDeleteContributorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContributorsDto**](ContributorsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppContributorsDeleteMyself

> ContributorsDto AppContributorsDeleteMyself(ctx, app).Execute()

Remove yourself.

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
    resp, r, err := apiClient.AppsApi.AppContributorsDeleteMyself(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppContributorsDeleteMyself``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppContributorsDeleteMyself`: ContributorsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppContributorsDeleteMyself`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppContributorsDeleteMyselfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContributorsDto**](ContributorsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppContributorsGetContributors

> ContributorsDto AppContributorsGetContributors(ctx, app).Execute()

Get app contributors.

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
    resp, r, err := apiClient.AppsApi.AppContributorsGetContributors(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppContributorsGetContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppContributorsGetContributors`: ContributorsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppContributorsGetContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppContributorsGetContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContributorsDto**](ContributorsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppContributorsPostContributor

> ContributorsDto AppContributorsPostContributor(ctx, app).AssignContributorDto(assignContributorDto).Execute()

Assign contributor to app.

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
    assignContributorDto := *openapiclient.NewAssignContributorDto("ContributorId_example") // AssignContributorDto | Contributor object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppContributorsPostContributor(context.Background(), app).AssignContributorDto(assignContributorDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppContributorsPostContributor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppContributorsPostContributor`: ContributorsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppContributorsPostContributor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppContributorsPostContributorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignContributorDto** | [**AssignContributorDto**](AssignContributorDto.md) | Contributor object that needs to be added to the app. | 

### Return type

[**ContributorsDto**](ContributorsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppImageGetImage

> *os.File AppImageGetImage(ctx, app).Execute()

Get the app image.

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
    resp, r, err := apiClient.AppsApi.AppImageGetImage(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppImageGetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppImageGetImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppImageGetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppImageGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AppLanguagesDeleteLanguage

> AppLanguagesDto AppLanguagesDeleteLanguage(ctx, app, language).Execute()

Deletes an app language.

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
    language := "language_example" // string | The language to delete from the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppLanguagesDeleteLanguage(context.Background(), app, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppLanguagesDeleteLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLanguagesDeleteLanguage`: AppLanguagesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppLanguagesDeleteLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**language** | **string** | The language to delete from the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLanguagesDeleteLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLanguagesDto**](AppLanguagesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLanguagesGetLanguages

> AppLanguagesDto AppLanguagesGetLanguages(ctx, app).Execute()

Get app languages.

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
    resp, r, err := apiClient.AppsApi.AppLanguagesGetLanguages(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppLanguagesGetLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLanguagesGetLanguages`: AppLanguagesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppLanguagesGetLanguages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLanguagesGetLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppLanguagesDto**](AppLanguagesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLanguagesPostLanguage

> AppLanguagesDto AppLanguagesPostLanguage(ctx, app).AddLanguageDto(addLanguageDto).Execute()

Attaches an app language.

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
    addLanguageDto := *openapiclient.NewAddLanguageDto("Language_example") // AddLanguageDto | The language to add to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppLanguagesPostLanguage(context.Background(), app).AddLanguageDto(addLanguageDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppLanguagesPostLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLanguagesPostLanguage`: AppLanguagesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppLanguagesPostLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLanguagesPostLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLanguageDto** | [**AddLanguageDto**](AddLanguageDto.md) | The language to add to the app. | 

### Return type

[**AppLanguagesDto**](AppLanguagesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLanguagesPutLanguage

> AppLanguagesDto AppLanguagesPutLanguage(ctx, app, language).UpdateLanguageDto(updateLanguageDto).Execute()

Updates an app language.

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
    language := "language_example" // string | The language to update.
    updateLanguageDto := *openapiclient.NewUpdateLanguageDto() // UpdateLanguageDto | The language object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppLanguagesPutLanguage(context.Background(), app, language).UpdateLanguageDto(updateLanguageDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppLanguagesPutLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLanguagesPutLanguage`: AppLanguagesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppLanguagesPutLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**language** | **string** | The language to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLanguagesPutLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLanguageDto** | [**UpdateLanguageDto**](UpdateLanguageDto.md) | The language object. | 

### Return type

[**AppLanguagesDto**](AppLanguagesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesDeleteRole

> RolesDto AppRolesDeleteRole(ctx, app, roleName).Execute()

Remove role from app.

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
    roleName := "roleName_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppRolesDeleteRole(context.Background(), app, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppRolesDeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesDeleteRole`: RolesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppRolesDeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**roleName** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RolesDto**](RolesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesGetPermissions

> []string AppRolesGetPermissions(ctx, app).Execute()

Get app permissions.

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
    resp, r, err := apiClient.AppsApi.AppRolesGetPermissions(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppRolesGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesGetPermissions`: []string
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppRolesGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesGetPermissionsRequest struct via the builder pattern


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


## AppRolesGetRoles

> RolesDto AppRolesGetRoles(ctx, app).Execute()

Get app roles.

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
    resp, r, err := apiClient.AppsApi.AppRolesGetRoles(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppRolesGetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesGetRoles`: RolesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppRolesGetRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RolesDto**](RolesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesPostRole

> RolesDto AppRolesPostRole(ctx, app).AddRoleDto(addRoleDto).Execute()

Add role to app.

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
    addRoleDto := *openapiclient.NewAddRoleDto("Name_example") // AddRoleDto | Role object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppRolesPostRole(context.Background(), app).AddRoleDto(addRoleDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppRolesPostRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesPostRole`: RolesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppRolesPostRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesPostRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addRoleDto** | [**AddRoleDto**](AddRoleDto.md) | Role object that needs to be added to the app. | 

### Return type

[**RolesDto**](RolesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesPutRole

> RolesDto AppRolesPutRole(ctx, app, roleName).UpdateRoleDto(updateRoleDto).Execute()

Update an app role.

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
    roleName := "roleName_example" // string | The name of the role to be updated.
    updateRoleDto := *openapiclient.NewUpdateRoleDto([]string{"Permissions_example"}) // UpdateRoleDto | Role to be updated for the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppRolesPutRole(context.Background(), app, roleName).UpdateRoleDto(updateRoleDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppRolesPutRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesPutRole`: RolesDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppRolesPutRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**roleName** | **string** | The name of the role to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesPutRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleDto** | [**UpdateRoleDto**](UpdateRoleDto.md) | Role to be updated for the app. | 

### Return type

[**RolesDto**](RolesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingsGetSettings

> AppSettingsDto AppSettingsGetSettings(ctx, app).Execute()

Get the app settings.

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
    app := "app_example" // string | The name of the app to get the settings for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppSettingsGetSettings(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppSettingsGetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingsGetSettings`: AppSettingsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppSettingsGetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to get the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingsGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppSettingsDto**](AppSettingsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingsPutSettings

> AppSettingsDto AppSettingsPutSettings(ctx, app).UpdateAppSettingsDto(updateAppSettingsDto).Execute()

Update the app settings.

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
    app := "app_example" // string | The name of the app to update.
    updateAppSettingsDto := *openapiclient.NewUpdateAppSettingsDto([]openapiclient.PatternDto{*openapiclient.NewPatternDto("Name_example", "Regex_example")}, []openapiclient.EditorDto{*openapiclient.NewEditorDto("Name_example", "Url_example")}) // UpdateAppSettingsDto | The values to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppSettingsPutSettings(context.Background(), app).UpdateAppSettingsDto(updateAppSettingsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppSettingsPutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingsPutSettings`: AppSettingsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppSettingsPutSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingsPutSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppSettingsDto** | [**UpdateAppSettingsDto**](UpdateAppSettingsDto.md) | The values to update. | 

### Return type

[**AppSettingsDto**](AppSettingsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppWorkflowsDeleteWorkflow

> WorkflowsDto AppWorkflowsDeleteWorkflow(ctx, app, id).Execute()

Delete a workflow.

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
    id := "id_example" // string | The id of the workflow to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppWorkflowsDeleteWorkflow(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppWorkflowsDeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppWorkflowsDeleteWorkflow`: WorkflowsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppWorkflowsDeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the workflow to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppWorkflowsDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowsDto**](WorkflowsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppWorkflowsGetWorkflows

> WorkflowsDto AppWorkflowsGetWorkflows(ctx, app).Execute()

Get app workflow.

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
    resp, r, err := apiClient.AppsApi.AppWorkflowsGetWorkflows(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppWorkflowsGetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppWorkflowsGetWorkflows`: WorkflowsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppWorkflowsGetWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppWorkflowsGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowsDto**](WorkflowsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppWorkflowsPostWorkflow

> WorkflowsDto AppWorkflowsPostWorkflow(ctx, app).AddWorkflowDto(addWorkflowDto).Execute()

Create a workflow.

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
    addWorkflowDto := *openapiclient.NewAddWorkflowDto("Name_example") // AddWorkflowDto | The new workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppWorkflowsPostWorkflow(context.Background(), app).AddWorkflowDto(addWorkflowDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppWorkflowsPostWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppWorkflowsPostWorkflow`: WorkflowsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppWorkflowsPostWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppWorkflowsPostWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addWorkflowDto** | [**AddWorkflowDto**](AddWorkflowDto.md) | The new workflow. | 

### Return type

[**WorkflowsDto**](WorkflowsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppWorkflowsPutWorkflow

> WorkflowsDto AppWorkflowsPutWorkflow(ctx, app, id).UpdateWorkflowDto(updateWorkflowDto).Execute()

Update a workflow.

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
    id := "id_example" // string | The id of the workflow to update.
    updateWorkflowDto := *openapiclient.NewUpdateWorkflowDto(map[string]WorkflowStepDto{"key": *openapiclient.NewWorkflowStepDto(map[string]WorkflowTransitionDto{"key": *openapiclient.NewWorkflowTransitionDto()})}, "Initial_example") // UpdateWorkflowDto | The new workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppWorkflowsPutWorkflow(context.Background(), app, id).UpdateWorkflowDto(updateWorkflowDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppWorkflowsPutWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppWorkflowsPutWorkflow`: WorkflowsDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppWorkflowsPutWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the workflow to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppWorkflowsPutWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWorkflowDto** | [**UpdateWorkflowDto**](UpdateWorkflowDto.md) | The new workflow. | 

### Return type

[**WorkflowsDto**](WorkflowsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDeleteApp

> AppsDeleteApp(ctx, app).Execute()

Delete the app.

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
    app := "app_example" // string | The name of the app to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsDeleteApp(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsDeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDeleteAppRequest struct via the builder pattern


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


## AppsDeleteImage

> AppDto AppsDeleteImage(ctx, app).Execute()

Remove the app image.

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
    app := "app_example" // string | The name of the app to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsDeleteImage(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsDeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsDeleteImage`: AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsDeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetApp

> AppDto AppsGetApp(ctx, app).Execute()

Get an app by name.

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
    resp, r, err := apiClient.AppsApi.AppsGetApp(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetApp`: AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetApps

> []AppDto AppsGetApps(ctx).Execute()

Get your apps.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetApps`: []AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetAppsRequest struct via the builder pattern


### Return type

[**[]AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPostApp

> AppDto AppsPostApp(ctx).CreateAppDto(createAppDto).Execute()

Create a new app.



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
    createAppDto := *openapiclient.NewCreateAppDto("Name_example") // CreateAppDto | The app object that needs to be added to Squidex.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsPostApp(context.Background()).CreateAppDto(createAppDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPostApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPostApp`: AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPostApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsPostAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAppDto** | [**CreateAppDto**](CreateAppDto.md) | The app object that needs to be added to Squidex. | 

### Return type

[**AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPutApp

> AppDto AppsPutApp(ctx, app).UpdateAppDto(updateAppDto).Execute()

Update the app.

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
    app := "app_example" // string | The name of the app to update.
    updateAppDto := *openapiclient.NewUpdateAppDto() // UpdateAppDto | The values to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsPutApp(context.Background(), app).UpdateAppDto(updateAppDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPutApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPutApp`: AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPutApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPutAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppDto** | [**UpdateAppDto**](UpdateAppDto.md) | The values to update. | 

### Return type

[**AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsUploadImage

> AppDto AppsUploadImage(ctx, app).File(file).Execute()

Upload the app image.

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
    app := "app_example" // string | The name of the app to update.
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsUploadImage(context.Background(), app).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsUploadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsUploadImage`: AppDto
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsUploadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**AppDto**](AppDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

