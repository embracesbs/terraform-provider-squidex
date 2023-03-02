# SignalRRuleActionDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | **string** | The connection string to the Azure SignalR. | 
**HubName** | **string** | The name of the hub. | 
**Action** | [**ActionTypeEnum**](ActionTypeEnum.md) |  | 
**MethodName** | Pointer to **NullableString** | Set the Name of the hub method received by the customer. | [optional] 
**Target** | Pointer to **NullableString** | Define target users or groups by id or name. One item per line. Not needed for Broadcast action. | [optional] 
**Payload** | Pointer to **NullableString** | Leave it empty to use the full event as body. | [optional] 

## Methods

### NewSignalRRuleActionDtoAllOf

`func NewSignalRRuleActionDtoAllOf(connectionString string, hubName string, action ActionTypeEnum, ) *SignalRRuleActionDtoAllOf`

NewSignalRRuleActionDtoAllOf instantiates a new SignalRRuleActionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalRRuleActionDtoAllOfWithDefaults

`func NewSignalRRuleActionDtoAllOfWithDefaults() *SignalRRuleActionDtoAllOf`

NewSignalRRuleActionDtoAllOfWithDefaults instantiates a new SignalRRuleActionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *SignalRRuleActionDtoAllOf) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *SignalRRuleActionDtoAllOf) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *SignalRRuleActionDtoAllOf) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetHubName

`func (o *SignalRRuleActionDtoAllOf) GetHubName() string`

GetHubName returns the HubName field if non-nil, zero value otherwise.

### GetHubNameOk

`func (o *SignalRRuleActionDtoAllOf) GetHubNameOk() (*string, bool)`

GetHubNameOk returns a tuple with the HubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubName

`func (o *SignalRRuleActionDtoAllOf) SetHubName(v string)`

SetHubName sets HubName field to given value.


### GetAction

`func (o *SignalRRuleActionDtoAllOf) GetAction() ActionTypeEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SignalRRuleActionDtoAllOf) GetActionOk() (*ActionTypeEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SignalRRuleActionDtoAllOf) SetAction(v ActionTypeEnum)`

SetAction sets Action field to given value.


### GetMethodName

`func (o *SignalRRuleActionDtoAllOf) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *SignalRRuleActionDtoAllOf) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *SignalRRuleActionDtoAllOf) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.

### HasMethodName

`func (o *SignalRRuleActionDtoAllOf) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.

### SetMethodNameNil

`func (o *SignalRRuleActionDtoAllOf) SetMethodNameNil(b bool)`

 SetMethodNameNil sets the value for MethodName to be an explicit nil

### UnsetMethodName
`func (o *SignalRRuleActionDtoAllOf) UnsetMethodName()`

UnsetMethodName ensures that no value is present for MethodName, not even an explicit nil
### GetTarget

`func (o *SignalRRuleActionDtoAllOf) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SignalRRuleActionDtoAllOf) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SignalRRuleActionDtoAllOf) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SignalRRuleActionDtoAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *SignalRRuleActionDtoAllOf) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *SignalRRuleActionDtoAllOf) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetPayload

`func (o *SignalRRuleActionDtoAllOf) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SignalRRuleActionDtoAllOf) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SignalRRuleActionDtoAllOf) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SignalRRuleActionDtoAllOf) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *SignalRRuleActionDtoAllOf) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *SignalRRuleActionDtoAllOf) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


