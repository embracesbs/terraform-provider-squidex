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

> SchemaDto SchemaFieldsDeleteField(ctx, app, schema, id).Execute()

Delete a schema field.

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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsDeleteField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsDeleteField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsDeleteField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsDeleteField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsDeleteFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsDeleteNestedField(ctx, app, schema, parentId, id).Execute()

Delete a nested field.

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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsDeleteNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsDeleteNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsDeleteNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsDeleteNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsDeleteNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemaDto SchemaFieldsDisableField(ctx, app, schema, id).Execute()

Disable a schema field.



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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsDisableField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsDisableField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsDisableField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsDisableField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsDisableFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsDisableNestedField(ctx, app, schema, parentId, id).Execute()

Disable a nested field.



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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsDisableNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsDisableNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsDisableNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsDisableNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsDisableNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemaDto SchemaFieldsEnableField(ctx, app, schema, id).Execute()

Enable a schema field.



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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to enable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsEnableField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsEnableField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsEnableField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsEnableField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsEnableFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsEnableNestedField(ctx, app, schema, parentId, id).Execute()

Enable a nested field.



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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to enable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsEnableNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsEnableNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsEnableNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsEnableNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsEnableNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemaDto SchemaFieldsHideField(ctx, app, schema, id).Execute()

Hide a schema field.



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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to hide.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsHideField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsHideField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsHideField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsHideField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to hide. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsHideFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsHideNestedField(ctx, app, schema, parentId, id).Execute()

Hide a nested field.



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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to hide.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsHideNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsHideNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsHideNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsHideNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to hide. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsHideNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemaDto SchemaFieldsLockField(ctx, app, schema, id).Execute()

Lock a schema field.



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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to lock.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsLockField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsLockField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsLockField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsLockField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to lock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsLockFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsLockNestedField(ctx, app, schema, parentId, id).Execute()

Lock a nested field.



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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to lock.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsLockNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsLockNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsLockNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsLockNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to lock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsLockNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemaDto SchemaFieldsPostField(ctx, app, schema).AddFieldDto(addFieldDto).Execute()

Add a schema field.

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
    schema := "schema_example" // string | The name of the schema.
    addFieldDto := *openapiclient.NewAddFieldDto("Name_example", *openapiclient.NewFieldPropertiesDto("FieldType_example")) // AddFieldDto | The field object that needs to be added to the schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPostField(context.Background(), app, schema).AddFieldDto(addFieldDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPostField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPostField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPostField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPostFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addFieldDto** | [**AddFieldDto**](AddFieldDto.md) | The field object that needs to be added to the schema. | 

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

> SchemaDto SchemaFieldsPostNestedField(ctx, app, schema, parentId).AddFieldDto(addFieldDto).Execute()

Add a nested field.

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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    addFieldDto := *openapiclient.NewAddFieldDto("Name_example", *openapiclient.NewFieldPropertiesDto("FieldType_example")) // AddFieldDto | The field object that needs to be added to the schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPostNestedField(context.Background(), app, schema, parentId).AddFieldDto(addFieldDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPostNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPostNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPostNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPostNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addFieldDto** | [**AddFieldDto**](AddFieldDto.md) | The field object that needs to be added to the schema. | 

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

> SchemaDto SchemaFieldsPutField(ctx, app, schema, id).UpdateFieldDto(updateFieldDto).Execute()

Update a schema field.

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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to update.
    updateFieldDto := *openapiclient.NewUpdateFieldDto(*openapiclient.NewFieldPropertiesDto("FieldType_example")) // UpdateFieldDto | The field object that needs to be added to the schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPutField(context.Background(), app, schema, id).UpdateFieldDto(updateFieldDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPutField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPutField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPutField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPutFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md) | The field object that needs to be added to the schema. | 

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

> SchemaDto SchemaFieldsPutNestedField(ctx, app, schema, parentId, id).UpdateFieldDto(updateFieldDto).Execute()

Update a nested field.

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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to update.
    updateFieldDto := *openapiclient.NewUpdateFieldDto(*openapiclient.NewFieldPropertiesDto("FieldType_example")) // UpdateFieldDto | The field object that needs to be added to the schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPutNestedField(context.Background(), app, schema, parentId, id).UpdateFieldDto(updateFieldDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPutNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPutNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPutNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPutNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateFieldDto** | [**UpdateFieldDto**](UpdateFieldDto.md) | The field object that needs to be added to the schema. | 

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

> SchemaDto SchemaFieldsPutNestedFieldOrdering(ctx, app, schema, parentId).ReorderFieldsDto(reorderFieldsDto).Execute()

Reorder all nested fields.

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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    reorderFieldsDto := *openapiclient.NewReorderFieldsDto([]int64{int64(123)}) // ReorderFieldsDto | The request that contains the field ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPutNestedFieldOrdering(context.Background(), app, schema, parentId).ReorderFieldsDto(reorderFieldsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPutNestedFieldOrdering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPutNestedFieldOrdering`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPutNestedFieldOrdering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPutNestedFieldOrderingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md) | The request that contains the field ids. | 

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

> SchemaDto SchemaFieldsPutSchemaFieldOrdering(ctx, app, schema).ReorderFieldsDto(reorderFieldsDto).Execute()

Reorder all fields.

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
    schema := "schema_example" // string | The name of the schema.
    reorderFieldsDto := *openapiclient.NewReorderFieldsDto([]int64{int64(123)}) // ReorderFieldsDto | The request that contains the field ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPutSchemaFieldOrdering(context.Background(), app, schema).ReorderFieldsDto(reorderFieldsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPutSchemaFieldOrdering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPutSchemaFieldOrdering`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPutSchemaFieldOrdering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPutSchemaFieldOrderingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reorderFieldsDto** | [**ReorderFieldsDto**](ReorderFieldsDto.md) | The request that contains the field ids. | 

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

> SchemaDto SchemaFieldsPutSchemaUIFields(ctx, app, schema).ConfigureUIFieldsDto(configureUIFieldsDto).Execute()

Configure UI fields.

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
    schema := "schema_example" // string | The name of the schema.
    configureUIFieldsDto := *openapiclient.NewConfigureUIFieldsDto() // ConfigureUIFieldsDto | The request that contains the field names.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsPutSchemaUIFields(context.Background(), app, schema).ConfigureUIFieldsDto(configureUIFieldsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsPutSchemaUIFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsPutSchemaUIFields`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsPutSchemaUIFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsPutSchemaUIFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configureUIFieldsDto** | [**ConfigureUIFieldsDto**](ConfigureUIFieldsDto.md) | The request that contains the field names. | 

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

> SchemaDto SchemaFieldsShowField(ctx, app, schema, id).Execute()

Show a schema field.



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
    schema := "schema_example" // string | The name of the schema.
    id := int64(789) // int64 | The id of the field to show.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsShowField(context.Background(), app, schema, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsShowField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsShowField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsShowField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**id** | **int64** | The id of the field to show. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsShowFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> SchemaDto SchemaFieldsShowNestedField(ctx, app, schema, parentId, id).Execute()

Show a nested field.



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
    schema := "schema_example" // string | The name of the schema.
    parentId := int64(789) // int64 | The parent field id.
    id := int64(789) // int64 | The id of the field to show.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemaFieldsShowNestedField(context.Background(), app, schema, parentId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemaFieldsShowNestedField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaFieldsShowNestedField`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemaFieldsShowNestedField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 
**parentId** | **int64** | The parent field id. | 
**id** | **int64** | The id of the field to show. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaFieldsShowNestedFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> SchemasDeleteSchema(ctx, app, schema).Execute()

Delete a schema.

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
    schema := "schema_example" // string | The name of the schema to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasDeleteSchema(context.Background(), app, schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasDeleteSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasDeleteSchemaRequest struct via the builder pattern


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


## SchemasGetSchema

> SchemaDto SchemasGetSchema(ctx, app, schema).Execute()

Get a schema by name.

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
    schema := "schema_example" // string | The name of the schema to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasGetSchema(context.Background(), app, schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasGetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasGetSchema`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasGetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> SchemasDto SchemasGetSchemas(ctx, app).Execute()

Get schemas.

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
    resp, r, err := apiClient.SchemasApi.SchemasGetSchemas(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasGetSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasGetSchemas`: SchemasDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasGetSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasGetSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SchemaDto SchemasPostSchema(ctx, app).CreateSchemaDto(createSchemaDto).Execute()

Create a new schema.

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
    createSchemaDto := *openapiclient.NewCreateSchemaDto("Name_example") // CreateSchemaDto | The schema object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPostSchema(context.Background(), app).CreateSchemaDto(createSchemaDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPostSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPostSchema`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPostSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPostSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSchemaDto** | [**CreateSchemaDto**](CreateSchemaDto.md) | The schema object that needs to be added to the app. | 

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

> SchemaDto SchemasPublishSchema(ctx, app, schema).Execute()

Publish a schema.

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
    schema := "schema_example" // string | The name of the schema to publish.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPublishSchema(context.Background(), app, schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPublishSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPublishSchema`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPublishSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPublishSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> SchemaDto SchemasPutCategory(ctx, app, schema).ChangeCategoryDto(changeCategoryDto).Execute()

Update a schema category.

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
    schema := "schema_example" // string | The name of the schema.
    changeCategoryDto := *openapiclient.NewChangeCategoryDto() // ChangeCategoryDto | The schema object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutCategory(context.Background(), app, schema).ChangeCategoryDto(changeCategoryDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutCategory`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeCategoryDto** | [**ChangeCategoryDto**](ChangeCategoryDto.md) | The schema object that needs to updated. | 

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

> SchemaDto SchemasPutPreviewUrls(ctx, app, schema).RequestBody(requestBody).Execute()

Update the preview urls.

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
    schema := "schema_example" // string | The name of the schema.
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string | The preview urls for the schema.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutPreviewUrls(context.Background(), app, schema).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutPreviewUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutPreviewUrls`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutPreviewUrls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutPreviewUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]string** | The preview urls for the schema. | 

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

> SchemaDto SchemasPutRules(ctx, app, schema).ConfigureFieldRulesDto(configureFieldRulesDto).Execute()

Update the rules.

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
    schema := "schema_example" // string | The name of the schema.
    configureFieldRulesDto := *openapiclient.NewConfigureFieldRulesDto() // ConfigureFieldRulesDto | The schema rules object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutRules(context.Background(), app, schema).ConfigureFieldRulesDto(configureFieldRulesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutRules`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configureFieldRulesDto** | [**ConfigureFieldRulesDto**](ConfigureFieldRulesDto.md) | The schema rules object that needs to updated. | 

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

> SchemaDto SchemasPutSchema(ctx, app, schema).UpdateSchemaDto(updateSchemaDto).Execute()

Update a schema.

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
    schema := "schema_example" // string | The name of the schema.
    updateSchemaDto := *openapiclient.NewUpdateSchemaDto() // UpdateSchemaDto | The schema object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutSchema(context.Background(), app, schema).UpdateSchemaDto(updateSchemaDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutSchema`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSchemaDto** | [**UpdateSchemaDto**](UpdateSchemaDto.md) | The schema object that needs to updated. | 

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

> SchemaDto SchemasPutSchemaSync(ctx, app, schema).SynchronizeSchemaDto(synchronizeSchemaDto).Execute()

Synchronize a schema.

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
    schema := "schema_example" // string | The name of the schema.
    synchronizeSchemaDto := *openapiclient.NewSynchronizeSchemaDto() // SynchronizeSchemaDto | The schema object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutSchemaSync(context.Background(), app, schema).SynchronizeSchemaDto(synchronizeSchemaDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutSchemaSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutSchemaSync`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutSchemaSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutSchemaSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **synchronizeSchemaDto** | [**SynchronizeSchemaDto**](SynchronizeSchemaDto.md) | The schema object that needs to updated. | 

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

> SchemaDto SchemasPutScripts(ctx, app, schema).SchemaScriptsDto(schemaScriptsDto).Execute()

Update the scripts.

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
    schema := "schema_example" // string | The name of the schema.
    schemaScriptsDto := *openapiclient.NewSchemaScriptsDto() // SchemaScriptsDto | The schema scripts object that needs to updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasPutScripts(context.Background(), app, schema).SchemaScriptsDto(schemaScriptsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasPutScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasPutScripts`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasPutScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPutScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schemaScriptsDto** | [**SchemaScriptsDto**](SchemaScriptsDto.md) | The schema scripts object that needs to updated. | 

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

> SchemaDto SchemasUnpublishSchema(ctx, app, schema).Execute()

Unpublish a schema.

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
    schema := "schema_example" // string | The name of the schema to unpublish.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.SchemasUnpublishSchema(context.Background(), app, schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.SchemasUnpublishSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemasUnpublishSchema`: SchemaDto
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.SchemasUnpublishSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**schema** | **string** | The name of the schema to unpublish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasUnpublishSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

