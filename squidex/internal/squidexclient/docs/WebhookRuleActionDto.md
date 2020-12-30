# WebhookRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** |  | 
**Url** | **string** | The url to the webhook. | 
**Method** | [**OneOfWebhookMethod**](oneOf&lt;WebhookMethod&gt;.md) | The type of the request. | 
**Payload** | Pointer to **string** | Leave it empty to use the full event as body. | [optional] 
**PayloadType** | Pointer to **string** | The mime type of the payload. | [optional] 
**Headers** | Pointer to **string** | The message headers in the format &#39;[Key]&#x3D;[Value]&#39;, one entry per line. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret that is used to calculate the payload signature. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


