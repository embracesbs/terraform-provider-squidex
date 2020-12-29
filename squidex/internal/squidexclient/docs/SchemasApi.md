# \SchemasApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaFieldsDeleteField**](SchemasApi.md#SchemaFieldsDeleteField) | **Delete** /apps/{app}/schemas/{name}/fields/{id} | Delete a schema field.
[**SchemaFieldsDeleteNestedField**](SchemasApi.md#SchemaFieldsDeleteNestedField) | **Delete** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id} | Delete a nested schema field.
[**SchemaFieldsDisableField**](SchemasApi.md#SchemaFieldsDisableField) | **Put** /apps/{app}/schemas/{name}/fields/{id}/disable | Disable a schema field.
[**SchemaFieldsDisableNestedField**](SchemasApi.md#SchemaFieldsDisableNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id}/disable | Disable nested a schema field.
[**SchemaFieldsEnableField**](SchemasApi.md#SchemaFieldsEnableField) | **Put** /apps/{app}/schemas/{name}/fields/{id}/enable | Enable a schema field.
[**SchemaFieldsEnableNestedField**](SchemasApi.md#SchemaFieldsEnableNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id}/enable | Enable a nested schema field.
[**SchemaFieldsHideField**](SchemasApi.md#SchemaFieldsHideField) | **Put** /apps/{app}/schemas/{name}/fields/{id}/hide | Hide a schema field.
[**SchemaFieldsHideNestedField**](SchemasApi.md#SchemaFieldsHideNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id}/hide | Hide a nested schema field.
[**SchemaFieldsLockField**](SchemasApi.md#SchemaFieldsLockField) | **Put** /apps/{app}/schemas/{name}/fields/{id}/lock | Lock a schema field.
[**SchemaFieldsLockNestedField**](SchemasApi.md#SchemaFieldsLockNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id}/lock | Lock a nested schema field.
[**SchemaFieldsPostField**](SchemasApi.md#SchemaFieldsPostField) | **Post** /apps/{app}/schemas/{name}/fields | Add a schema field.
[**SchemaFieldsPostNestedField**](SchemasApi.md#SchemaFieldsPostNestedField) | **Post** /apps/{app}/schemas/{name}/fields/{parentId}/nested | Add a nested schema field.
[**SchemaFieldsPutField**](SchemasApi.md#SchemaFieldsPutField) | **Put** /apps/{app}/schemas/{name}/fields/{id} | Update a schema field.
[**SchemaFieldsPutNestedField**](SchemasApi.md#SchemaFieldsPutNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id} | Update a nested schema field.
[**SchemaFieldsPutNestedFieldOrdering**](SchemasApi.md#SchemaFieldsPutNestedFieldOrdering) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/ordering | Reorders the nested fields.
[**SchemaFieldsPutSchemaFieldOrdering**](SchemasApi.md#SchemaFieldsPutSchemaFieldOrdering) | **Put** /apps/{app}/schemas/{name}/fields/ordering | Reorders the fields.
[**SchemaFieldsPutSchemaUIFields**](SchemasApi.md#SchemaFieldsPutSchemaUIFields) | **Put** /apps/{app}/schemas/{name}/fields/ui | Configure UI fields.
[**SchemaFieldsShowField**](SchemasApi.md#SchemaFieldsShowField) | **Put** /apps/{app}/schemas/{name}/fields/{id}/show | Show a schema field.
[**SchemaFieldsShowNestedField**](SchemasApi.md#SchemaFieldsShowNestedField) | **Put** /apps/{app}/schemas/{name}/fields/{parentId}/nested/{id}/show | Show a nested schema field.
[**SchemasDeleteSchema**](SchemasApi.md#SchemasDeleteSchema) | **Delete** /apps/{app}/schemas/{name} | Delete a schema.
[**SchemasGetSchema**](SchemasApi.md#SchemasGetSchema) | **Get** /apps/{app}/schemas/{name} | Get a schema by name.
[**SchemasGetSchemas**](SchemasApi.md#SchemasGetSchemas) | **Get** /apps/{app}/schemas | Get schemas.
[**SchemasPostSchema**](SchemasApi.md#SchemasPostSchema) | **Post** /apps/{app}/schemas | Create a new schema.
[**SchemasPublishSchema**](SchemasApi.md#SchemasPublishSchema) | **Put** /apps/{app}/schemas/{name}/publish | Publish a schema.
[**SchemasPutCategory**](SchemasApi.md#SchemasPutCategory) | **Put** /apps/{app}/schemas/{name}/category | Update a schema category.
[**SchemasPutPreviewUrls**](SchemasApi.md#SchemasPutPreviewUrls) | **Put** /apps/{app}/schemas/{name}/preview-urls | Update the preview urls.
[**SchemasPutSchema**](SchemasApi.md#SchemasPutSchema) | **Put** /apps/{app}/schemas/{name} | Update a schema.
[**SchemasPutSchemaSync**](SchemasApi.md#SchemasPutSchemaSync) | **Put** /apps/{app}/schemas/{name}/sync | Synchronize a schema.
[**SchemasPutScripts**](SchemasApi.md#SchemasPutScripts) | **Put** /apps/{app}/schemas/{name}/scripts | Update the scripts.
[**SchemasUnpublishSchema**](SchemasApi.md#SchemasUnpublishSchema) | **Put** /apps/{app}/schemas/{name}/unpublish | Unpublish a schema.



## SchemaFieldsDeleteField

> SchemaDetailsDto SchemaFieldsDeleteField(ctx, app, name, id)

Delete a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDeleteNestedField

> SchemaDetailsDto SchemaFieldsDeleteNestedField(ctx, app, name, parentId, id)

Delete a nested schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDisableField

> SchemaDetailsDto SchemaFieldsDisableField(ctx, app, name, id)

Disable a schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDisableNestedField

> SchemaDetailsDto SchemaFieldsDisableNestedField(ctx, app, name, parentId, id)

Disable nested a schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsEnableField

> SchemaDetailsDto SchemaFieldsEnableField(ctx, app, name, id)

Enable a schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to enable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsEnableNestedField

> SchemaDetailsDto SchemaFieldsEnableNestedField(ctx, app, name, parentId, id)

Enable a nested schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to enable. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsHideField

> SchemaDetailsDto SchemaFieldsHideField(ctx, app, name, id)

Hide a schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to hide. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsHideNestedField

> SchemaDetailsDto SchemaFieldsHideNestedField(ctx, app, name, parentId, id)

Hide a nested schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to hide. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsLockField

> SchemaDetailsDto SchemaFieldsLockField(ctx, app, name, id)

Lock a schema field.

A locked field cannot be updated or deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to lock. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsLockNestedField

> SchemaDetailsDto SchemaFieldsLockNestedField(ctx, app, name, parentId, id)

Lock a nested schema field.

A locked field cannot be edited or deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to lock. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPostField

> SchemaDetailsDto SchemaFieldsPostField(ctx, app, name, addFieldDto)

Add a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**addFieldDto** | [**AddFieldDto**](AddFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPostNestedField

> SchemaDetailsDto SchemaFieldsPostNestedField(ctx, app, name, parentId, addFieldDto)

Add a nested schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**addFieldDto** | [**AddFieldDto**](AddFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutField

> SchemaDetailsDto SchemaFieldsPutField(ctx, app, name, id, updateFieldDto)

Update a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to update. | 
**updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutNestedField

> SchemaDetailsDto SchemaFieldsPutNestedField(ctx, app, name, parentId, id, updateFieldDto)

Update a nested schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to update. | 
**updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutNestedFieldOrdering

> SchemaDetailsDto SchemaFieldsPutNestedFieldOrdering(ctx, app, name, parentId, reorderFieldsDto)

Reorders the nested fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md)| The request that contains the field ids. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutSchemaFieldOrdering

> SchemaDetailsDto SchemaFieldsPutSchemaFieldOrdering(ctx, app, name, reorderFieldsDto)

Reorders the fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md)| The request that contains the field ids. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutSchemaUIFields

> SchemaDetailsDto SchemaFieldsPutSchemaUIFields(ctx, app, name, configureUiFieldsDto)

Configure UI fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**configureUiFieldsDto** | [**ConfigureUiFieldsDto**](ConfigureUiFieldsDto.md)| The request that contains the field names. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsShowField

> SchemaDetailsDto SchemaFieldsShowField(ctx, app, name, id)

Show a schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to show. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsShowNestedField

> SchemaDetailsDto SchemaFieldsShowNestedField(ctx, app, name, parentId, id)

Show a nested schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to show. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasDeleteSchema

> SchemasDeleteSchema(ctx, app, name)

Delete a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema to delete. | 

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


## SchemasGetSchema

> SchemaDetailsDto SchemasGetSchema(ctx, app, name)

Get a schema by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema to retrieve. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasGetSchemas

> SchemasDto SchemasGetSchemas(ctx, app)

Get schemas.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**SchemasDto**](SchemasDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPostSchema

> SchemaDetailsDto SchemasPostSchema(ctx, app, createSchemaDto)

Create a new schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createSchemaDto** | [**CreateSchemaDto**](CreateSchemaDto.md)| The schema object that needs to be added to the app. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPublishSchema

> SchemaDetailsDto SchemasPublishSchema(ctx, app, name)

Publish a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema to publish. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutCategory

> SchemaDetailsDto SchemasPutCategory(ctx, app, name, changeCategoryDto)

Update a schema category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**changeCategoryDto** | [**ChangeCategoryDto**](ChangeCategoryDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutPreviewUrls

> SchemaDetailsDto SchemasPutPreviewUrls(ctx, app, name, requestBody)

Update the preview urls.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**requestBody** | [**map[string]string**](string.md)| The preview urls for the schema. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutSchema

> SchemaDetailsDto SchemasPutSchema(ctx, app, name, updateSchemaDto)

Update a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**updateSchemaDto** | [**UpdateSchemaDto**](UpdateSchemaDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutSchemaSync

> SchemaDetailsDto SchemasPutSchemaSync(ctx, app, name, synchronizeSchemaDto)

Synchronize a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**synchronizeSchemaDto** | [**SynchronizeSchemaDto**](SynchronizeSchemaDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutScripts

> SchemaDetailsDto SchemasPutScripts(ctx, app, name, schemaScriptsDto)

Update the scripts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema. | 
**schemaScriptsDto** | [**SchemaScriptsDto**](SchemaScriptsDto.md)| The schema scripts object that needs to updated. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasUnpublishSchema

> SchemaDetailsDto SchemasUnpublishSchema(ctx, app, name)

Unpublish a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**name** | **string**| The name of the schema to unpublish. | 

### Return type

[**SchemaDetailsDto**](SchemaDetailsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

