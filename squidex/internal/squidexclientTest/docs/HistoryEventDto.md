# HistoryEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The message for the event. | 
**EventType** | **string** | The type of the original event. | 
**Actor** | **string** | The user who called the action. | 
**EventId** | Pointer to **string** | Gets a unique id for the event. | [optional] 
**Created** | Pointer to **time.Time** | The time when the event happened. | [optional] 
**Version** | Pointer to **int64** | The version identifier. | [optional] 

## Methods

### NewHistoryEventDto

`func NewHistoryEventDto(message string, eventType string, actor string, ) *HistoryEventDto`

NewHistoryEventDto instantiates a new HistoryEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEventDtoWithDefaults

`func NewHistoryEventDtoWithDefaults() *HistoryEventDto`

NewHistoryEventDtoWithDefaults instantiates a new HistoryEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *HistoryEventDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HistoryEventDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HistoryEventDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEventType

`func (o *HistoryEventDto) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *HistoryEventDto) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *HistoryEventDto) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetActor

`func (o *HistoryEventDto) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *HistoryEventDto) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *HistoryEventDto) SetActor(v string)`

SetActor sets Actor field to given value.


### GetEventId

`func (o *HistoryEventDto) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *HistoryEventDto) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *HistoryEventDto) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *HistoryEventDto) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetCreated

`func (o *HistoryEventDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HistoryEventDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HistoryEventDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *HistoryEventDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *HistoryEventDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HistoryEventDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HistoryEventDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HistoryEventDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


