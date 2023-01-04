# \PlansApi

All URIs are relative to *http://squidex.localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPlansGetPlans**](PlansApi.md#AppPlansGetPlans) | **Get** /api/apps/{app}/plans | Get app plan information.
[**AppPlansPutPlan**](PlansApi.md#AppPlansPutPlan) | **Put** /api/apps/{app}/plan | Change the app plan.



## AppPlansGetPlans

> AppPlansDto AppPlansGetPlans(ctx, app).Execute()

Get app plan information.

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
    resp, r, err := apiClient.PlansApi.AppPlansGetPlans(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.AppPlansGetPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPlansGetPlans`: AppPlansDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.AppPlansGetPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPlansGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PlanChangedDto AppPlansPutPlan(ctx, app).ChangePlanDto(changePlanDto).Execute()

Change the app plan.

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
    changePlanDto := *openapiclient.NewChangePlanDto("PlanId_example") // ChangePlanDto | Plan object that needs to be changed.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.AppPlansPutPlan(context.Background(), app).ChangePlanDto(changePlanDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.AppPlansPutPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPlansPutPlan`: PlanChangedDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.AppPlansPutPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | The name of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPlansPutPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePlanDto** | [**ChangePlanDto**](ChangePlanDto.md) | Plan object that needs to be changed. | 

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

