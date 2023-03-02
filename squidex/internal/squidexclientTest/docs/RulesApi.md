# \RulesApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesDeleteEvent**](RulesApi.md#RulesDeleteEvent) | **Delete** /api/apps/{app}/rules/events/{id} | Cancels an event.
[**RulesDeleteEvents**](RulesApi.md#RulesDeleteEvents) | **Delete** /api/apps/{app}/rules/events | Cancels all events.
[**RulesDeleteRule**](RulesApi.md#RulesDeleteRule) | **Delete** /api/apps/{app}/rules/{id} | Delete a rule.
[**RulesDeleteRuleEvents**](RulesApi.md#RulesDeleteRuleEvents) | **Delete** /api/apps/{app}/rules/{id}/events | Cancels all rule events.
[**RulesDeleteRuleRun**](RulesApi.md#RulesDeleteRuleRun) | **Delete** /api/apps/{app}/rules/run | Cancel the current run.
[**RulesDisableRule**](RulesApi.md#RulesDisableRule) | **Put** /api/apps/{app}/rules/{id}/disable | Disable a rule.
[**RulesEnableRule**](RulesApi.md#RulesEnableRule) | **Put** /api/apps/{app}/rules/{id}/enable | Enable a rule.
[**RulesGetActions**](RulesApi.md#RulesGetActions) | **Get** /api/rules/actions | Get supported rule actions.
[**RulesGetEventSchema**](RulesApi.md#RulesGetEventSchema) | **Get** /api/rules/eventtypes/{type} | Provide the json schema for the event with the specified name.
[**RulesGetEventTypes**](RulesApi.md#RulesGetEventTypes) | **Get** /api/rules/eventtypes | Provide a list of all event types that are used in rules.
[**RulesGetEvents**](RulesApi.md#RulesGetEvents) | **Get** /api/apps/{app}/rules/events | Get rule events.
[**RulesGetRules**](RulesApi.md#RulesGetRules) | **Get** /api/apps/{app}/rules | Get rules.
[**RulesPostRule**](RulesApi.md#RulesPostRule) | **Post** /api/apps/{app}/rules | Create a new rule.
[**RulesPutEvent**](RulesApi.md#RulesPutEvent) | **Put** /api/apps/{app}/rules/events/{id} | Retry the event immediately.
[**RulesPutRule**](RulesApi.md#RulesPutRule) | **Put** /api/apps/{app}/rules/{id} | Update a rule.
[**RulesPutRuleRun**](RulesApi.md#RulesPutRuleRun) | **Put** /api/apps/{app}/rules/{id}/run | Run a rule.
[**RulesSimulateGET**](RulesApi.md#RulesSimulateGET) | **Get** /api/apps/{app}/rules/{id}/simulate | Simulate a rule.
[**RulesSimulatePOST**](RulesApi.md#RulesSimulatePOST) | **Post** /api/apps/{app}/rules/simulate | Simulate a rule.
[**RulesTriggerRule**](RulesApi.md#RulesTriggerRule) | **Put** /api/apps/{app}/rules/{id}/trigger | Trigger a rule.



## RulesDeleteEvent

> RulesDeleteEvent(ctx, app, id).Execute()

Cancels an event.

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
    id := "id_example" // string | The event to enqueue.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesDeleteEvent(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDeleteEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The event to enqueue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDeleteEventRequest struct via the builder pattern


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


## RulesDeleteEvents

> RulesDeleteEvents(ctx, app).Execute()

Cancels all events.

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
    resp, r, err := apiClient.RulesApi.RulesDeleteEvents(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDeleteEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDeleteEventsRequest struct via the builder pattern


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


## RulesDeleteRule

> RulesDeleteRule(ctx, app, id).Execute()

Delete a rule.

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
    id := "id_example" // string | The id of the rule to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesDeleteRule(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDeleteRuleRequest struct via the builder pattern


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


## RulesDeleteRuleEvents

> RulesDeleteRuleEvents(ctx, app, id).Execute()

Cancels all rule events.

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
    id := "id_example" // string | The id of the rule to cancel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesDeleteRuleEvents(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDeleteRuleEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDeleteRuleEventsRequest struct via the builder pattern


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


## RulesDeleteRuleRun

> RulesDeleteRuleRun(ctx, app).Execute()

Cancel the current run.

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
    resp, r, err := apiClient.RulesApi.RulesDeleteRuleRun(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDeleteRuleRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDeleteRuleRunRequest struct via the builder pattern


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


## RulesDisableRule

> RuleDto RulesDisableRule(ctx, app, id).Execute()

Disable a rule.

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
    id := "id_example" // string | The id of the rule to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesDisableRule(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesDisableRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesDisableRule`: RuleDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesDisableRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesDisableRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RuleDto**](RuleDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesEnableRule

> RuleDto RulesEnableRule(ctx, app, id).Execute()

Enable a rule.

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
    id := "id_example" // string | The id of the rule to enable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesEnableRule(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesEnableRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesEnableRule`: RuleDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesEnableRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesEnableRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RuleDto**](RuleDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesGetActions

> map[string]RuleElementDto RulesGetActions(ctx).Execute()

Get supported rule actions.

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
    resp, r, err := apiClient.RulesApi.RulesGetActions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesGetActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesGetActions`: map[string]RuleElementDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesGetActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRulesGetActionsRequest struct via the builder pattern


### Return type

[**map[string]RuleElementDto**](RuleElementDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesGetEventSchema

> interface{} RulesGetEventSchema(ctx, type_).Execute()

Provide the json schema for the event with the specified name.

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
    type_ := "type__example" // string | The type name of the event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesGetEventSchema(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesGetEventSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesGetEventSchema`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesGetEventSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type name of the event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesGetEventSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesGetEventTypes

> []string RulesGetEventTypes(ctx).Execute()

Provide a list of all event types that are used in rules.

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
    resp, r, err := apiClient.RulesApi.RulesGetEventTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesGetEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesGetEventTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesGetEventTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRulesGetEventTypesRequest struct via the builder pattern


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


## RulesGetEvents

> RuleEventsDto RulesGetEvents(ctx, app).RuleId(ruleId).Skip(skip).Take(take).Execute()

Get rule events.

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
    ruleId := "ruleId_example" // string | The optional rule id to filter to events. (optional)
    skip := int32(56) // int32 | The number of events to skip. (optional) (default to 0)
    take := int32(56) // int32 | The number of events to take. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesGetEvents(context.Background(), app).RuleId(ruleId).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesGetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesGetEvents`: RuleEventsDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesGetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleId** | **string** | The optional rule id to filter to events. | 
 **skip** | **int32** | The number of events to skip. | [default to 0]
 **take** | **int32** | The number of events to take. | [default to 20]

### Return type

[**RuleEventsDto**](RuleEventsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesGetRules

> RulesDto RulesGetRules(ctx, app).Execute()

Get rules.

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
    resp, r, err := apiClient.RulesApi.RulesGetRules(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesGetRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesGetRules`: RulesDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesGetRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RulesDto**](RulesDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesPostRule

> RuleDto RulesPostRule(ctx, app).CreateRuleDto(createRuleDto).Execute()

Create a new rule.

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
    createRuleDto := *openapiclient.NewCreateRuleDto(*openapiclient.NewRuleTriggerDto("TriggerType_example"), *openapiclient.NewRuleAction("ActionType_example")) // CreateRuleDto | The rule object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesPostRule(context.Background(), app).CreateRuleDto(createRuleDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesPostRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesPostRule`: RuleDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesPostRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesPostRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRuleDto** | [**CreateRuleDto**](CreateRuleDto.md) | The rule object that needs to be added to the app. | 

### Return type

[**RuleDto**](RuleDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesPutEvent

> RulesPutEvent(ctx, app, id).Execute()

Retry the event immediately.

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
    id := "id_example" // string | The event to enqueue.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesPutEvent(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesPutEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The event to enqueue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesPutEventRequest struct via the builder pattern


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


## RulesPutRule

> RuleDto RulesPutRule(ctx, app, id).UpdateRuleDto(updateRuleDto).Execute()

Update a rule.

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
    id := "id_example" // string | The id of the rule to update.
    updateRuleDto := *openapiclient.NewUpdateRuleDto() // UpdateRuleDto | The rule object that needs to be added to the app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesPutRule(context.Background(), app, id).UpdateRuleDto(updateRuleDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesPutRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesPutRule`: RuleDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesPutRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesPutRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRuleDto** | [**UpdateRuleDto**](UpdateRuleDto.md) | The rule object that needs to be added to the app. | 

### Return type

[**RuleDto**](RuleDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesPutRuleRun

> RulesPutRuleRun(ctx, app, id).FromSnapshots(fromSnapshots).Execute()

Run a rule.

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
    id := "id_example" // string | The id of the rule to run.
    fromSnapshots := true // bool | Runs the rule from snapeshots if possible. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesPutRuleRun(context.Background(), app, id).FromSnapshots(fromSnapshots).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesPutRuleRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesPutRuleRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromSnapshots** | **bool** | Runs the rule from snapeshots if possible. | [default to false]

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


## RulesSimulateGET

> SimulatedRuleEventsDto RulesSimulateGET(ctx, app, id).Execute()

Simulate a rule.

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
    id := "id_example" // string | The id of the rule to simulate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesSimulateGET(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesSimulateGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesSimulateGET`: SimulatedRuleEventsDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesSimulateGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to simulate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesSimulateGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SimulatedRuleEventsDto**](SimulatedRuleEventsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesSimulatePOST

> SimulatedRuleEventsDto RulesSimulatePOST(ctx, app).CreateRuleDto(createRuleDto).Execute()

Simulate a rule.

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
    createRuleDto := *openapiclient.NewCreateRuleDto(*openapiclient.NewRuleTriggerDto("TriggerType_example"), *openapiclient.NewRuleAction("ActionType_example")) // CreateRuleDto | The rule to simulate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesSimulatePOST(context.Background(), app).CreateRuleDto(createRuleDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesSimulatePOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesSimulatePOST`: SimulatedRuleEventsDto
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.RulesSimulatePOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesSimulatePOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRuleDto** | [**CreateRuleDto**](CreateRuleDto.md) | The rule to simulate. | 

### Return type

[**SimulatedRuleEventsDto**](SimulatedRuleEventsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesTriggerRule

> RulesTriggerRule(ctx, app, id).Execute()

Trigger a rule.

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
    id := "id_example" // string | The id of the rule to disable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.RulesTriggerRule(context.Background(), app, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.RulesTriggerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 
**id** | **string** | The id of the rule to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesTriggerRuleRequest struct via the builder pattern


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

