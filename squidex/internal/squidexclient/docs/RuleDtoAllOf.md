# RuleDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the rule. | [optional] 
**CreatedBy** | **string** | The user that has created the rule. | 
**LastModifiedBy** | **string** | The user that has updated the rule. | 
**Created** | [**time.Time**](time.Time.md) | The date and time when the rule has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The date and time when the rule has been modified last. | [optional] 
**Version** | **int64** | The version of the rule. | [optional] 
**IsEnabled** | **bool** | Determines if the rule is enabled. | [optional] 
**Name** | Pointer to **string** | Optional rule name. | [optional] 
**Trigger** | [**OneOfRuleTriggerDto**](oneOf&lt;RuleTriggerDto&gt;.md) | The trigger properties. | 
**Action** | [**OneOfRuleAction**](oneOf&lt;RuleAction&gt;.md) | The action properties. | 
**NumSucceeded** | **int32** | The number of completed executions. | [optional] 
**NumFailed** | **int32** | The number of failed executions. | [optional] 
**LastExecuted** | Pointer to [**time.Time**](time.Time.md) | The date and time when the rule was executed the last time. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


