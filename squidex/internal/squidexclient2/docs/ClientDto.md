# ClientDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The client id. | 
**Secret** | **string** | The client secret. | 
**Name** | **string** | The client name. | 
**Role** | Pointer to **string** | The role of the client. | [optional] 
**ApiCallsLimit** | **int64** | The number of allowed api calls per month for this client. | [optional] 
**ApiTrafficLimit** | **int64** | The number of allowed api traffic bytes per month for this client. | [optional] 
**AllowAnonymous** | **bool** | True to allow anonymous access without an access token for this client. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


