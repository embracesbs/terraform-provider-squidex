# WorkflowStepDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transitions** | [**map[string]WorkflowTransitionDto**](WorkflowTransitionDto.md) | The transitions. | 
**Color** | Pointer to **string** | The optional color. | [optional] 
**NoUpdate** | **bool** | Indicates if updates should not be allowed. | [optional] 
**NoUpdateExpression** | Pointer to **string** | Optional expression that must evaluate to true when you want to prevent updates. | [optional] 
**NoUpdateRoles** | Pointer to **[]string** | Optional list of roles to restrict the updates for users with these roles. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


