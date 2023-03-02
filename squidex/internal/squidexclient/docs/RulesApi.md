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

> RulesDeleteEvent(ctx, app, id)

Cancels an event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The event to enqueue. | 

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

> RulesDeleteEvents(ctx, app)

Cancels all events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> RulesDeleteRule(ctx, app, id)

Delete a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to delete. | 

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

> RulesDeleteRuleEvents(ctx, app, id)

Cancels all rule events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to cancel. | 

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

> RulesDeleteRuleRun(ctx, app)

Cancel the current run.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> RuleDto RulesDisableRule(ctx, app, id)

Disable a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to disable. | 

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

> RuleDto RulesEnableRule(ctx, app, id)

Enable a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to enable. | 

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

> map[string]RuleElementDto RulesGetActions(ctx, )

Get supported rule actions.

### Required Parameters

This endpoint does not need any parameter.

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

> interface{} RulesGetEventSchema(ctx, type_)

Provide the json schema for the event with the specified name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The type name of the event. | 

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

> []string RulesGetEventTypes(ctx, )

Provide a list of all event types that are used in rules.

### Required Parameters

This endpoint does not need any parameter.

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

> RuleEventsDto RulesGetEvents(ctx, app, optional)

Get rule events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
 **optional** | ***RulesGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RulesGetEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleId** | **optional.String**| The optional rule id to filter to events. | 
 **skip** | **optional.Int32**| The number of events to skip. | [default to 0]
 **take** | **optional.Int32**| The number of events to take. | [default to 20]

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

> RulesDto RulesGetRules(ctx, app)

Get rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

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

> RuleDto RulesPostRule(ctx, app, createRuleDto)

Create a new rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createRuleDto** | [**CreateRuleDto**](CreateRuleDto.md)| The rule object that needs to be added to the app. | 

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

> RulesPutEvent(ctx, app, id)

Retry the event immediately.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The event to enqueue. | 

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

> RuleDto RulesPutRule(ctx, app, id, updateRuleDto)

Update a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to update. | 
**updateRuleDto** | [**UpdateRuleDto**](UpdateRuleDto.md)| The rule object that needs to be added to the app. | 

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

> RulesPutRuleRun(ctx, app, id, optional)

Run a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to run. | 
 **optional** | ***RulesPutRuleRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RulesPutRuleRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromSnapshots** | **optional.Bool**| Runs the rule from snapeshots if possible. | [default to false]

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

> SimulatedRuleEventsDto RulesSimulateGET(ctx, app, id)

Simulate a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to simulate. | 

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

> SimulatedRuleEventsDto RulesSimulatePOST(ctx, app, createRuleDto)

Simulate a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**createRuleDto** | [**CreateRuleDto**](CreateRuleDto.md)| The rule to simulate. | 

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

> RulesTriggerRule(ctx, app, id)

Trigger a rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the rule to disable. | 

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

