# \RulesApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesDeleteEvent**](RulesApi.md#RulesDeleteEvent) | **Delete** /apps/{app}/rules/events/{id} | Cancels the event and retries.
[**RulesDeleteRule**](RulesApi.md#RulesDeleteRule) | **Delete** /apps/{app}/rules/{id} | Delete a rule.
[**RulesDisableRule**](RulesApi.md#RulesDisableRule) | **Put** /apps/{app}/rules/{id}/disable | Disable a rule.
[**RulesEnableRule**](RulesApi.md#RulesEnableRule) | **Put** /apps/{app}/rules/{id}/enable | Enable a rule.
[**RulesGetActions**](RulesApi.md#RulesGetActions) | **Get** /rules/actions | Get the supported rule actions.
[**RulesGetEvents**](RulesApi.md#RulesGetEvents) | **Get** /apps/{app}/rules/events | Get rule events.
[**RulesGetRules**](RulesApi.md#RulesGetRules) | **Get** /apps/{app}/rules | Get rules.
[**RulesPostRule**](RulesApi.md#RulesPostRule) | **Post** /apps/{app}/rules | Create a new rule.
[**RulesPutEvent**](RulesApi.md#RulesPutEvent) | **Put** /apps/{app}/rules/events/{id} | Retry the event immediately.
[**RulesPutRule**](RulesApi.md#RulesPutRule) | **Put** /apps/{app}/rules/{id} | Update a rule.
[**RulesTriggerRule**](RulesApi.md#RulesTriggerRule) | **Put** /apps/{app}/rules/{id}/trigger | Trigger a rule.



## RulesDeleteEvent

> RulesDeleteEvent(ctx, app, id)

Cancels the event and retries.

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

Get the supported rule actions.

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

All events for the specified schemas will be sent to the url. The timeout is 2 seconds.

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

