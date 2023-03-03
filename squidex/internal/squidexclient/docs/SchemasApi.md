# \SchemasApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaFieldsDeleteField**](SchemasApi.md#SchemaFieldsDeleteField) | **Delete** /api/apps/{app}/schemas/{schema}/fields/{id} | Delete a schema field.
[**SchemaFieldsDeleteNestedField**](SchemasApi.md#SchemaFieldsDeleteNestedField) | **Delete** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id} | Delete a nested field.
[**SchemaFieldsDisableField**](SchemasApi.md#SchemaFieldsDisableField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id}/disable | Disable a schema field.
[**SchemaFieldsDisableNestedField**](SchemasApi.md#SchemaFieldsDisableNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id}/disable | Disable a nested field.
[**SchemaFieldsEnableField**](SchemasApi.md#SchemaFieldsEnableField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id}/enable | Enable a schema field.
[**SchemaFieldsEnableNestedField**](SchemasApi.md#SchemaFieldsEnableNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id}/enable | Enable a nested field.
[**SchemaFieldsHideField**](SchemasApi.md#SchemaFieldsHideField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id}/hide | Hide a schema field.
[**SchemaFieldsHideNestedField**](SchemasApi.md#SchemaFieldsHideNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id}/hide | Hide a nested field.
[**SchemaFieldsLockField**](SchemasApi.md#SchemaFieldsLockField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id}/lock | Lock a schema field.
[**SchemaFieldsLockNestedField**](SchemasApi.md#SchemaFieldsLockNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id}/lock | Lock a nested field.
[**SchemaFieldsPostField**](SchemasApi.md#SchemaFieldsPostField) | **Post** /api/apps/{app}/schemas/{schema}/fields | Add a schema field.
[**SchemaFieldsPostNestedField**](SchemasApi.md#SchemaFieldsPostNestedField) | **Post** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested | Add a nested field.
[**SchemaFieldsPutField**](SchemasApi.md#SchemaFieldsPutField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id} | Update a schema field.
[**SchemaFieldsPutNestedField**](SchemasApi.md#SchemaFieldsPutNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id} | Update a nested field.
[**SchemaFieldsPutNestedFieldOrdering**](SchemasApi.md#SchemaFieldsPutNestedFieldOrdering) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/ordering | Reorder all nested fields.
[**SchemaFieldsPutSchemaFieldOrdering**](SchemasApi.md#SchemaFieldsPutSchemaFieldOrdering) | **Put** /api/apps/{app}/schemas/{schema}/fields/ordering | Reorder all fields.
[**SchemaFieldsPutSchemaUIFields**](SchemasApi.md#SchemaFieldsPutSchemaUIFields) | **Put** /api/apps/{app}/schemas/{schema}/fields/ui | Configure UI fields.
[**SchemaFieldsShowField**](SchemasApi.md#SchemaFieldsShowField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{id}/show | Show a schema field.
[**SchemaFieldsShowNestedField**](SchemasApi.md#SchemaFieldsShowNestedField) | **Put** /api/apps/{app}/schemas/{schema}/fields/{parentId}/nested/{id}/show | Show a nested field.
[**SchemasDeleteSchema**](SchemasApi.md#SchemasDeleteSchema) | **Delete** /api/apps/{app}/schemas/{schema} | Delete a schema.
[**SchemasGetSchema**](SchemasApi.md#SchemasGetSchema) | **Get** /api/apps/{app}/schemas/{schema} | Get a schema by name.
[**SchemasGetSchemas**](SchemasApi.md#SchemasGetSchemas) | **Get** /api/apps/{app}/schemas | Get schemas.
[**SchemasPostSchema**](SchemasApi.md#SchemasPostSchema) | **Post** /api/apps/{app}/schemas | Create a new schema.
[**SchemasPublishSchema**](SchemasApi.md#SchemasPublishSchema) | **Put** /api/apps/{app}/schemas/{schema}/publish | Publish a schema.
[**SchemasPutCategory**](SchemasApi.md#SchemasPutCategory) | **Put** /api/apps/{app}/schemas/{schema}/category | Update a schema category.
[**SchemasPutPreviewUrls**](SchemasApi.md#SchemasPutPreviewUrls) | **Put** /api/apps/{app}/schemas/{schema}/preview-urls | Update the preview urls.
[**SchemasPutRules**](SchemasApi.md#SchemasPutRules) | **Put** /api/apps/{app}/schemas/{schema}/rules | Update the rules.
[**SchemasPutSchema**](SchemasApi.md#SchemasPutSchema) | **Put** /api/apps/{app}/schemas/{schema} | Update a schema.
[**SchemasPutSchemaSync**](SchemasApi.md#SchemasPutSchemaSync) | **Put** /api/apps/{app}/schemas/{schema}/sync | Synchronize a schema.
[**SchemasPutScripts**](SchemasApi.md#SchemasPutScripts) | **Put** /api/apps/{app}/schemas/{schema}/scripts | Update the scripts.
[**SchemasUnpublishSchema**](SchemasApi.md#SchemasUnpublishSchema) | **Put** /api/apps/{app}/schemas/{schema}/unpublish | Unpublish a schema.



## SchemaFieldsDeleteField

> SchemaDto SchemaFieldsDeleteField(ctx, app, schema, id)

Delete a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDeleteNestedField

> SchemaDto SchemaFieldsDeleteNestedField(ctx, app, schema, parentId, id)

Delete a nested field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDisableField

> SchemaDto SchemaFieldsDisableField(ctx, app, schema, id)

Disable a schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsDisableNestedField

> SchemaDto SchemaFieldsDisableNestedField(ctx, app, schema, parentId, id)

Disable a nested field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to disable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsEnableField

> SchemaDto SchemaFieldsEnableField(ctx, app, schema, id)

Enable a schema field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to enable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsEnableNestedField

> SchemaDto SchemaFieldsEnableNestedField(ctx, app, schema, parentId, id)

Enable a nested field.

A disabled field cannot not be edited in the squidex portal anymore, but will be part of the API response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to enable. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsHideField

> SchemaDto SchemaFieldsHideField(ctx, app, schema, id)

Hide a schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to hide. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsHideNestedField

> SchemaDto SchemaFieldsHideNestedField(ctx, app, schema, parentId, id)

Hide a nested field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to hide. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsLockField

> SchemaDto SchemaFieldsLockField(ctx, app, schema, id)

Lock a schema field.

A locked field cannot be updated or deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to lock. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsLockNestedField

> SchemaDto SchemaFieldsLockNestedField(ctx, app, schema, parentId, id)

Lock a nested field.

A locked field cannot be edited or deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to lock. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPostField

> SchemaDto SchemaFieldsPostField(ctx, app, schema, addFieldDto)

Add a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**addFieldDto** | [**AddFieldDto**](AddFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPostNestedField

> SchemaDto SchemaFieldsPostNestedField(ctx, app, schema, parentId, addFieldDto)

Add a nested field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**addFieldDto** | [**AddFieldDto**](AddFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutField

> SchemaDto SchemaFieldsPutField(ctx, app, schema, id, updateFieldDto)

Update a schema field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to update. | 
**updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutNestedField

> SchemaDto SchemaFieldsPutNestedField(ctx, app, schema, parentId, id, updateFieldDto)

Update a nested field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to update. | 
**updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md)| The field object that needs to be added to the schema. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutNestedFieldOrdering

> SchemaDto SchemaFieldsPutNestedFieldOrdering(ctx, app, schema, parentId, reorderFieldsDto)

Reorder all nested fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md)| The request that contains the field ids. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutSchemaFieldOrdering

> SchemaDto SchemaFieldsPutSchemaFieldOrdering(ctx, app, schema, reorderFieldsDto)

Reorder all fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md)| The request that contains the field ids. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsPutSchemaUIFields

> SchemaDto SchemaFieldsPutSchemaUIFields(ctx, app, schema, configureUiFieldsDto)

Configure UI fields.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**configureUiFieldsDto** | [**ConfigureUiFieldsDto**](ConfigureUiFieldsDto.md)| The request that contains the field names. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsShowField

> SchemaDto SchemaFieldsShowField(ctx, app, schema, id)

Show a schema field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**id** | **int64**| The id of the field to show. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaFieldsShowNestedField

> SchemaDto SchemaFieldsShowNestedField(ctx, app, schema, parentId, id)

Show a nested field.

A hidden field is not part of the API response, but can still be edited in the portal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**parentId** | **int64**| The parent field id. | 
**id** | **int64**| The id of the field to show. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasDeleteSchema

> SchemasDeleteSchema(ctx, app, schema)

Delete a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema to delete. | 

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

> SchemaDto SchemasGetSchema(ctx, app, schema)

Get a schema by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema to retrieve. | 

### Return type

[**SchemaDto**](SchemaDto.md)

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

> SchemaDto SchemasPostSchema(ctx, app, createSchemaDto)

Create a new schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createSchemaDto** | [**CreateSchemaDto**](CreateSchemaDto.md)| The schema object that needs to be added to the app. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPublishSchema

> SchemaDto SchemasPublishSchema(ctx, app, schema)

Publish a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema to publish. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutCategory

> SchemaDto SchemasPutCategory(ctx, app, schema, changeCategoryDto)

Update a schema category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**changeCategoryDto** | [**ChangeCategoryDto**](ChangeCategoryDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutPreviewUrls

> SchemaDto SchemasPutPreviewUrls(ctx, app, schema, requestBody)

Update the preview urls.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**requestBody** | [**map[string]string**](string.md)| The preview urls for the schema. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutRules

> SchemaDto SchemasPutRules(ctx, app, schema, configureFieldRulesDto)

Update the rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**configureFieldRulesDto** | [**ConfigureFieldRulesDto**](ConfigureFieldRulesDto.md)| The schema rules object that needs to updated. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutSchema

> SchemaDto SchemasPutSchema(ctx, app, schema, updateSchemaDto)

Update a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**updateSchemaDto** | [**UpdateSchemaDto**](UpdateSchemaDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutSchemaSync

> SchemaDto SchemasPutSchemaSync(ctx, app, schema, synchronizeSchemaDto)

Synchronize a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**synchronizeSchemaDto** | [**SynchronizeSchemaDto**](SynchronizeSchemaDto.md)| The schema object that needs to updated. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPutScripts

> SchemaDto SchemasPutScripts(ctx, app, schema, schemaScriptsDto)

Update the scripts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema. | 
**schemaScriptsDto** | [**SchemaScriptsDto**](SchemaScriptsDto.md)| The schema scripts object that needs to updated. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasUnpublishSchema

> SchemaDto SchemasUnpublishSchema(ctx, app, schema)

Unpublish a schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**schema** | **string**| The name of the schema to unpublish. | 

### Return type

[**SchemaDto**](SchemaDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

