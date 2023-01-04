# SimulatedRuleEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The unique event id. | 
**EventName** | **string** | The name of the event. | 
**Event** | **interface{}** | The source event. | 
**EnrichedEvent** | Pointer to **interface{}** | The enriched event. | [optional] 
**ActionName** | Pointer to **NullableString** | The data for the action. | [optional] 
**ActionData** | Pointer to **NullableString** | The name of the action. | [optional] 
**Error** | Pointer to **NullableString** | The name of the event. | [optional] 
**SkipReasons** | [**[]SkipReason**](SkipReason.md) | The reason why the event has been skipped. | 

## Methods

### NewSimulatedRuleEventDto

`func NewSimulatedRuleEventDto(eventId string, eventName string, event interface{}, skipReasons []SkipReason, ) *SimulatedRuleEventDto`

NewSimulatedRuleEventDto instantiates a new SimulatedRuleEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatedRuleEventDtoWithDefaults

`func NewSimulatedRuleEventDtoWithDefaults() *SimulatedRuleEventDto`

NewSimulatedRuleEventDtoWithDefaults instantiates a new SimulatedRuleEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SimulatedRuleEventDto) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SimulatedRuleEventDto) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SimulatedRuleEventDto) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventName

`func (o *SimulatedRuleEventDto) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *SimulatedRuleEventDto) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *SimulatedRuleEventDto) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEvent

`func (o *SimulatedRuleEventDto) GetEvent() interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SimulatedRuleEventDto) GetEventOk() (*interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SimulatedRuleEventDto) SetEvent(v interface{})`

SetEvent sets Event field to given value.


### SetEventNil

`func (o *SimulatedRuleEventDto) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *SimulatedRuleEventDto) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetEnrichedEvent

`func (o *SimulatedRuleEventDto) GetEnrichedEvent() interface{}`

GetEnrichedEvent returns the EnrichedEvent field if non-nil, zero value otherwise.

### GetEnrichedEventOk

`func (o *SimulatedRuleEventDto) GetEnrichedEventOk() (*interface{}, bool)`

GetEnrichedEventOk returns a tuple with the EnrichedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrichedEvent

`func (o *SimulatedRuleEventDto) SetEnrichedEvent(v interface{})`

SetEnrichedEvent sets EnrichedEvent field to given value.

### HasEnrichedEvent

`func (o *SimulatedRuleEventDto) HasEnrichedEvent() bool`

HasEnrichedEvent returns a boolean if a field has been set.

### SetEnrichedEventNil

`func (o *SimulatedRuleEventDto) SetEnrichedEventNil(b bool)`

 SetEnrichedEventNil sets the value for EnrichedEvent to be an explicit nil

### UnsetEnrichedEvent
`func (o *SimulatedRuleEventDto) UnsetEnrichedEvent()`

UnsetEnrichedEvent ensures that no value is present for EnrichedEvent, not even an explicit nil
### GetActionName

`func (o *SimulatedRuleEventDto) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *SimulatedRuleEventDto) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *SimulatedRuleEventDto) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *SimulatedRuleEventDto) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionNameNil

`func (o *SimulatedRuleEventDto) SetActionNameNil(b bool)`

 SetActionNameNil sets the value for ActionName to be an explicit nil

### UnsetActionName
`func (o *SimulatedRuleEventDto) UnsetActionName()`

UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
### GetActionData

`func (o *SimulatedRuleEventDto) GetActionData() string`

GetActionData returns the ActionData field if non-nil, zero value otherwise.

### GetActionDataOk

`func (o *SimulatedRuleEventDto) GetActionDataOk() (*string, bool)`

GetActionDataOk returns a tuple with the ActionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionData

`func (o *SimulatedRuleEventDto) SetActionData(v string)`

SetActionData sets ActionData field to given value.

### HasActionData

`func (o *SimulatedRuleEventDto) HasActionData() bool`

HasActionData returns a boolean if a field has been set.

### SetActionDataNil

`func (o *SimulatedRuleEventDto) SetActionDataNil(b bool)`

 SetActionDataNil sets the value for ActionData to be an explicit nil

### UnsetActionData
`func (o *SimulatedRuleEventDto) UnsetActionData()`

UnsetActionData ensures that no value is present for ActionData, not even an explicit nil
### GetError

`func (o *SimulatedRuleEventDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SimulatedRuleEventDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SimulatedRuleEventDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SimulatedRuleEventDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SimulatedRuleEventDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SimulatedRuleEventDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetSkipReasons

`func (o *SimulatedRuleEventDto) GetSkipReasons() []SkipReason`

GetSkipReasons returns the SkipReasons field if non-nil, zero value otherwise.

### GetSkipReasonsOk

`func (o *SimulatedRuleEventDto) GetSkipReasonsOk() (*[]SkipReason, bool)`

GetSkipReasonsOk returns a tuple with the SkipReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReasons

`func (o *SimulatedRuleEventDto) SetSkipReasons(v []SkipReason)`

SetSkipReasons sets SkipReasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


