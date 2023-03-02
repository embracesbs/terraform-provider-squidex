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

> AssetScriptsDto AppAssetsGetAssetScripts(ctx, app)

Get the app asset scripts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to get the asset scripts for. | 

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

> AssetScriptsDto AppAssetsPutAssetScripts(ctx, app, updateAssetScriptsDto)

Update the app asset scripts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to update. | 
**updateAssetScriptsDto** | [**UpdateAssetScriptsDto**](UpdateAssetScriptsDto.md)| The values to update. | 

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

> ClientsDto AppClientsDeleteClient(ctx, app, id)

Revoke an app client.

The application that uses this client credentials cannot access the API after it has been revoked.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the client that must be deleted. | 

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

> ClientsDto AppClientsGetClients(ctx, app)

Get app clients.

Gets all configured clients for the app with the specified name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> ClientsDto AppClientsPostClient(ctx, app, createClientDto)

Create a new app client.

Create a new client for the app with the specified name. The client secret is auto generated on the server and returned. The client does not expire, the access token is valid for 30 days.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createClientDto** | [**CreateClientDto**](CreateClientDto.md)| Client object that needs to be added to the app. | 

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

> ClientsDto AppClientsPutClient(ctx, app, id, updateClientDto)

Updates an app client.

Only the display name can be changed, create a new client if necessary.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the client that must be updated. | 
**updateClientDto** | [**UpdateClientDto**](UpdateClientDto.md)| Client object that needs to be updated. | 

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

> ContributorsDto AppContributorsDeleteContributor(ctx, app, id)

Remove contributor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the contributor. | 

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

> ContributorsDto AppContributorsDeleteMyself(ctx, app)

Remove yourself.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> ContributorsDto AppContributorsGetContributors(ctx, app)

Get app contributors.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> ContributorsDto AppContributorsPostContributor(ctx, app, assignContributorDto)

Assign contributor to app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**assignContributorDto** | [**AssignContributorDto**](AssignContributorDto.md)| Contributor object that needs to be added to the app. | 

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

> *os.File AppImageGetImage(ctx, app)

Get the app image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> AppLanguagesDto AppLanguagesDeleteLanguage(ctx, app, language)

Deletes an app language.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**language** | **string**| The language to delete from the app. | 

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

> AppLanguagesDto AppLanguagesGetLanguages(ctx, app)

Get app languages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> AppLanguagesDto AppLanguagesPostLanguage(ctx, app, addLanguageDto)

Attaches an app language.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**addLanguageDto** | [**AddLanguageDto**](AddLanguageDto.md)| The language to add to the app. | 

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

> AppLanguagesDto AppLanguagesPutLanguage(ctx, app, language, updateLanguageDto)

Updates an app language.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**language** | **string**| The language to update. | 
**updateLanguageDto** | [**UpdateLanguageDto**](UpdateLanguageDto.md)| The language object. | 

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

> RolesDto AppRolesDeleteRole(ctx, app, roleName)

Remove role from app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**roleName** | **string**| The name of the role. | 

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

> []string AppRolesGetPermissions(ctx, app)

Get app permissions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> RolesDto AppRolesGetRoles(ctx, app)

Get app roles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> RolesDto AppRolesPostRole(ctx, app, addRoleDto)

Add role to app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**addRoleDto** | [**AddRoleDto**](AddRoleDto.md)| Role object that needs to be added to the app. | 

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

> RolesDto AppRolesPutRole(ctx, app, roleName, updateRoleDto)

Update an app role.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**roleName** | **string**| The name of the role to be updated. | 
**updateRoleDto** | [**UpdateRoleDto**](UpdateRoleDto.md)| Role to be updated for the app. | 

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

> AppSettingsDto AppSettingsGetSettings(ctx, app)

Get the app settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to get the settings for. | 

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

> AppSettingsDto AppSettingsPutSettings(ctx, app, updateAppSettingsDto)

Update the app settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to update. | 
**updateAppSettingsDto** | [**UpdateAppSettingsDto**](UpdateAppSettingsDto.md)| The values to update. | 

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

> WorkflowsDto AppWorkflowsDeleteWorkflow(ctx, app, id)

Delete a workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the workflow to update. | 

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

> WorkflowsDto AppWorkflowsGetWorkflows(ctx, app)

Get app workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> WorkflowsDto AppWorkflowsPostWorkflow(ctx, app, addWorkflowDto)

Create a workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**addWorkflowDto** | [**AddWorkflowDto**](AddWorkflowDto.md)| The new workflow. | 

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

> WorkflowsDto AppWorkflowsPutWorkflow(ctx, app, id, updateWorkflowDto)

Update a workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the workflow to update. | 
**updateWorkflowDto** | [**UpdateWorkflowDto**](UpdateWorkflowDto.md)| The new workflow. | 

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

> AppsDeleteApp(ctx, app)

Delete the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to delete. | 

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

> AppDto AppsDeleteImage(ctx, app)

Remove the app image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to update. | 

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

> AppDto AppsGetApp(ctx, app)

Get an app by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> []AppDto AppsGetApps(ctx, )

Get your apps.

You can only retrieve the list of apps when you are authenticated as a user (OpenID implicit flow). You will retrieve all apps, where you are assigned as a contributor.

### Required Parameters

This endpoint does not need any parameter.

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

> AppDto AppsPostApp(ctx, createAppDto)

Create a new app.

You can only create an app when you are authenticated as a user (OpenID implicit flow). You will be assigned as owner of the new app automatically.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAppDto** | [**CreateAppDto**](CreateAppDto.md)| The app object that needs to be added to Squidex. | 

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

> AppDto AppsPutApp(ctx, app, updateAppDto)

Update the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to update. | 
**updateAppDto** | [**UpdateAppDto**](UpdateAppDto.md)| The values to update. | 

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

> AppDto AppsUploadImage(ctx, app, optional)

Upload the app image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to update. | 
 **optional** | ***AppsUploadImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppsUploadImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

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

