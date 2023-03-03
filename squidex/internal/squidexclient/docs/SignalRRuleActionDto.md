# SignalRRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** |  | 
**ConnectionString** | **string** | The connection string to the Azure SignalR. | 
**HubName** | **string** | The name of the hub. | 
**Action** | [**ActionTypeEnum**](ActionTypeEnum.md) |  | 
**MethodName** | Pointer to **string** | Set the Name of the hub method received by the customer. | [optional] 
**Target** | Pointer to **string** | Define target users or groups by id or name. One item per line. Not needed for Broadcast action. | [optional] 
**Payload** | Pointer to **string** | Leave it empty to use the full event as body. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


