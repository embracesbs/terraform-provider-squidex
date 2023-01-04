# RuleEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The id of the event. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The time when the event has been created. | [optional] 
**Description** | **string** | The description. | 
**EventName** | **string** | The name of the event. | 
**LastDump** | Pointer to **string** | The last dump. | [optional] 
**NumCalls** | **int32** | The number of calls. | [optional] 
**NextAttempt** | Pointer to [**time.Time**](time.Time.md) | The next attempt. | [optional] 
**Result** | [**OneOfRuleResult**](oneOf&lt;RuleResult&gt;.md) | The result of the event. | [optional] 
**JobResult** | [**OneOfRuleJobResult**](oneOf&lt;RuleJobResult&gt;.md) | The result of the job. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


