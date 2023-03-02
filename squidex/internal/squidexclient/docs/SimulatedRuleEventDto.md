# SimulatedRuleEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The unique event id. | 
**EventName** | **string** | The name of the event. | 
**Event** | [**interface{}**](.md) | The source event. | 
**EnrichedEvent** | Pointer to [**interface{}**](.md) | The enriched event. | [optional] 
**ActionName** | Pointer to **string** | The data for the action. | [optional] 
**ActionData** | Pointer to **string** | The name of the action. | [optional] 
**Error** | Pointer to **string** | The name of the event. | [optional] 
**SkipReasons** | [**[]SkipReason**](SkipReason.md) | The reason why the event has been skipped. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


