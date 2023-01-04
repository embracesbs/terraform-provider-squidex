# AzureQueueRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | **string** | The connection string to the storage account. | 
**Queue** | **string** | The name of the queue. | 
**Payload** | Pointer to **NullableString** | Leave it empty to use the full event as body. | [optional] 

## Methods

### NewAzureQueueRuleActionDto

`func NewAzureQueueRuleActionDto(connectionString string, queue string, ) *AzureQueueRuleActionDto`

NewAzureQueueRuleActionDto instantiates a new AzureQueueRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureQueueRuleActionDtoWithDefaults

`func NewAzureQueueRuleActionDtoWithDefaults() *AzureQueueRuleActionDto`

NewAzureQueueRuleActionDtoWithDefaults instantiates a new AzureQueueRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *AzureQueueRuleActionDto) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *AzureQueueRuleActionDto) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *AzureQueueRuleActionDto) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetQueue

`func (o *AzureQueueRuleActionDto) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AzureQueueRuleActionDto) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AzureQueueRuleActionDto) SetQueue(v string)`

SetQueue sets Queue field to given value.


### GetPayload

`func (o *AzureQueueRuleActionDto) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AzureQueueRuleActionDto) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AzureQueueRuleActionDto) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AzureQueueRuleActionDto) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *AzureQueueRuleActionDto) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *AzureQueueRuleActionDto) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


