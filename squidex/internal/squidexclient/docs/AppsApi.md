# \AppsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppClientsDeleteClient**](AppsApi.md#AppClientsDeleteClient) | **Delete** /apps/{app}/clients/{id} | Revoke an app client.
[**AppClientsGetClients**](AppsApi.md#AppClientsGetClients) | **Get** /apps/{app}/clients | Get app clients.
[**AppClientsPostClient**](AppsApi.md#AppClientsPostClient) | **Post** /apps/{app}/clients | Create a new app client.
[**AppClientsPutClient**](AppsApi.md#AppClientsPutClient) | **Put** /apps/{app}/clients/{id} | Updates an app client.
[**AppContributorsDeleteContributor**](AppsApi.md#AppContributorsDeleteContributor) | **Delete** /apps/{app}/contributors/{id} | Remove contributor from app.
[**AppContributorsGetContributors**](AppsApi.md#AppContributorsGetContributors) | **Get** /apps/{app}/contributors | Get app contributors.
[**AppContributorsPostContributor**](AppsApi.md#AppContributorsPostContributor) | **Post** /apps/{app}/contributors | Assign contributor to app.
[**AppLanguagesDeleteLanguage**](AppsApi.md#AppLanguagesDeleteLanguage) | **Delete** /apps/{app}/languages/{language} | Deletes an app language.
[**AppLanguagesGetLanguages**](AppsApi.md#AppLanguagesGetLanguages) | **Get** /apps/{app}/languages | Get app languages.
[**AppLanguagesPostLanguage**](AppsApi.md#AppLanguagesPostLanguage) | **Post** /apps/{app}/languages | Attaches an app language.
[**AppLanguagesPutLanguage**](AppsApi.md#AppLanguagesPutLanguage) | **Put** /apps/{app}/languages/{language} | Updates an app language.
[**AppPatternsDeletePattern**](AppsApi.md#AppPatternsDeletePattern) | **Delete** /apps/{app}/patterns/{id} | Delete an existing app pattern.
[**AppPatternsGetPatterns**](AppsApi.md#AppPatternsGetPatterns) | **Get** /apps/{app}/patterns | Get app patterns.
[**AppPatternsPostPattern**](AppsApi.md#AppPatternsPostPattern) | **Post** /apps/{app}/patterns | Create a new app pattern.
[**AppPatternsPutPattern**](AppsApi.md#AppPatternsPutPattern) | **Put** /apps/{app}/patterns/{id} | Update an existing app pattern.
[**AppRolesDeleteRole**](AppsApi.md#AppRolesDeleteRole) | **Delete** /apps/{app}/roles/{name} | Remove role from app.
[**AppRolesGetPermissions**](AppsApi.md#AppRolesGetPermissions) | **Get** /apps/{app}/roles/permissions | Get app permissions.
[**AppRolesGetRoles**](AppsApi.md#AppRolesGetRoles) | **Get** /apps/{app}/roles | Get app roles.
[**AppRolesPostRole**](AppsApi.md#AppRolesPostRole) | **Post** /apps/{app}/roles | Add role to app.
[**AppRolesPutRole**](AppsApi.md#AppRolesPutRole) | **Put** /apps/{app}/roles/{name} | Update an existing app role.
[**AppWorkflowsDeleteWorkflow**](AppsApi.md#AppWorkflowsDeleteWorkflow) | **Delete** /apps/{app}/workflows/{id} | Delete a workflow.
[**AppWorkflowsGetWorkflows**](AppsApi.md#AppWorkflowsGetWorkflows) | **Get** /apps/{app}/workflows | Get app workflow.
[**AppWorkflowsPostWorkflow**](AppsApi.md#AppWorkflowsPostWorkflow) | **Post** /apps/{app}/workflows | Create a workflow.
[**AppWorkflowsPutWorkflow**](AppsApi.md#AppWorkflowsPutWorkflow) | **Put** /apps/{app}/workflows/{id} | Update a workflow.
[**AppsDeleteApp**](AppsApi.md#AppsDeleteApp) | **Delete** /apps/{app} | Archive the app.
[**AppsGetApp**](AppsApi.md#AppsGetApp) | **Get** /apps/{app} | Get an app by name.
[**AppsGetApps**](AppsApi.md#AppsGetApps) | **Get** /apps | Get your apps.
[**AppsPostApp**](AppsApi.md#AppsPostApp) | **Post** /apps | Create a new app.
[**AppsUpdateApp**](AppsApi.md#AppsUpdateApp) | **Put** /apps/{app} | Update the app.



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

Create a new client for the app with the specified name. The client secret is auto generated on the server and returned. The client does not exire, the access token is valid for 30 days.

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

Remove contributor from app.

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


## AppPatternsDeletePattern

> PatternsDto AppPatternsDeletePattern(ctx, app, id)

Delete an existing app pattern.

Schemas using this pattern will still function using the same Regular Expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the pattern to be deleted. | 

### Return type

[**PatternsDto**](PatternsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPatternsGetPatterns

> PatternsDto AppPatternsGetPatterns(ctx, app)

Get app patterns.

Gets all configured regex patterns for the app with the specified name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**PatternsDto**](PatternsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPatternsPostPattern

> PatternsDto AppPatternsPostPattern(ctx, app, updatePatternDto)

Create a new app pattern.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**updatePatternDto** | [**UpdatePatternDto**](UpdatePatternDto.md)| Pattern to be added to the app. | 

### Return type

[**PatternsDto**](PatternsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPatternsPutPattern

> PatternsDto AppPatternsPutPattern(ctx, app, id, updatePatternDto)

Update an existing app pattern.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the pattern to be updated. | 
**updatePatternDto** | [**UpdatePatternDto**](UpdatePatternDto.md)| Pattern to be updated for the app. | 

### Return type

[**PatternsDto**](PatternsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesDeleteRole

> RolesDto AppRolesDeleteRole(ctx, app, name)

Remove role from app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the role. | 

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

> RolesDto AppRolesPutRole(ctx, app, name, updateRoleDto)

Update an existing app role.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the role to be updated. | 
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

Archive the app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app to archive. | 

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


## AppsUpdateApp

> AppDto AppsUpdateApp(ctx, app, updateAppDto)

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

