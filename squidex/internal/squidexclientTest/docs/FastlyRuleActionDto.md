# FastlyRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | The API key to grant access to Squidex. | 
**ServiceId** | **string** | The ID of the fastly service. | 

## Methods

### NewFastlyRuleActionDto

`func NewFastlyRuleActionDto(apiKey string, serviceId string, ) *FastlyRuleActionDto`

NewFastlyRuleActionDto instantiates a new FastlyRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastlyRuleActionDtoWithDefaults

`func NewFastlyRuleActionDtoWithDefaults() *FastlyRuleActionDto`

NewFastlyRuleActionDtoWithDefaults instantiates a new FastlyRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *FastlyRuleActionDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FastlyRuleActionDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FastlyRuleActionDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetServiceId

`func (o *FastlyRuleActionDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FastlyRuleActionDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FastlyRuleActionDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


