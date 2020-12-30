# \PlansApi

All URIs are relative to *https://squidex.dev.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPlansGetPlans**](PlansApi.md#AppPlansGetPlans) | **Get** /apps/{app}/plans | Get app plan information.
[**AppPlansPutPlan**](PlansApi.md#AppPlansPutPlan) | **Put** /apps/{app}/plan | Change the app plan.



## AppPlansGetPlans

> AppPlansDto AppPlansGetPlans(ctx, app)

Get app plan information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**AppPlansDto**](AppPlansDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPlansPutPlan

> PlanChangedDto AppPlansPutPlan(ctx, app, changePlanDto)

Change the app plan.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**changePlanDto** | [**ChangePlanDto**](ChangePlanDto.md)| Plan object that needs to be changed. | 

### Return type

[**PlanChangedDto**](PlanChangedDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

