# RuleEventDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the event. | [optional] 
**Created** | Pointer to **time.Time** | The time when the event has been created. | [optional] 
**Description** | **string** | The description. | 
**EventName** | **string** | The name of the event. | 
**LastDump** | Pointer to **NullableString** | The last dump. | [optional] 
**NumCalls** | Pointer to **int32** | The number of calls. | [optional] 
**NextAttempt** | Pointer to **NullableTime** | The next attempt. | [optional] 
**Result** | Pointer to [**RuleResult**](RuleResult.md) |  | [optional] 
**JobResult** | Pointer to [**RuleJobResult**](RuleJobResult.md) |  | [optional] 

## Methods

### NewRuleEventDtoAllOf

`func NewRuleEventDtoAllOf(description string, eventName string, ) *RuleEventDtoAllOf`

NewRuleEventDtoAllOf instantiates a new RuleEventDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEventDtoAllOfWithDefaults

`func NewRuleEventDtoAllOfWithDefaults() *RuleEventDtoAllOf`

NewRuleEventDtoAllOfWithDefaults instantiates a new RuleEventDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleEventDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleEventDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleEventDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleEventDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *RuleEventDtoAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RuleEventDtoAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RuleEventDtoAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RuleEventDtoAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *RuleEventDtoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleEventDtoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleEventDtoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventName

`func (o *RuleEventDtoAllOf) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *RuleEventDtoAllOf) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *RuleEventDtoAllOf) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetLastDump

`func (o *RuleEventDtoAllOf) GetLastDump() string`

GetLastDump returns the LastDump field if non-nil, zero value otherwise.

### GetLastDumpOk

`func (o *RuleEventDtoAllOf) GetLastDumpOk() (*string, bool)`

GetLastDumpOk returns a tuple with the LastDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDump

`func (o *RuleEventDtoAllOf) SetLastDump(v string)`

SetLastDump sets LastDump field to given value.

### HasLastDump

`func (o *RuleEventDtoAllOf) HasLastDump() bool`

HasLastDump returns a boolean if a field has been set.

### SetLastDumpNil

`func (o *RuleEventDtoAllOf) SetLastDumpNil(b bool)`

 SetLastDumpNil sets the value for LastDump to be an explicit nil

### UnsetLastDump
`func (o *RuleEventDtoAllOf) UnsetLastDump()`

UnsetLastDump ensures that no value is present for LastDump, not even an explicit nil
### GetNumCalls

`func (o *RuleEventDtoAllOf) GetNumCalls() int32`

GetNumCalls returns the NumCalls field if non-nil, zero value otherwise.

### GetNumCallsOk

`func (o *RuleEventDtoAllOf) GetNumCallsOk() (*int32, bool)`

GetNumCallsOk returns a tuple with the NumCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCalls

`func (o *RuleEventDtoAllOf) SetNumCalls(v int32)`

SetNumCalls sets NumCalls field to given value.

### HasNumCalls

`func (o *RuleEventDtoAllOf) HasNumCalls() bool`

HasNumCalls returns a boolean if a field has been set.

### GetNextAttempt

`func (o *RuleEventDtoAllOf) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *RuleEventDtoAllOf) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *RuleEventDtoAllOf) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *RuleEventDtoAllOf) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### SetNextAttemptNil

`func (o *RuleEventDtoAllOf) SetNextAttemptNil(b bool)`

 SetNextAttemptNil sets the value for NextAttempt to be an explicit nil

### UnsetNextAttempt
`func (o *RuleEventDtoAllOf) UnsetNextAttempt()`

UnsetNextAttempt ensures that no value is present for NextAttempt, not even an explicit nil
### GetResult

`func (o *RuleEventDtoAllOf) GetResult() RuleResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RuleEventDtoAllOf) GetResultOk() (*RuleResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RuleEventDtoAllOf) SetResult(v RuleResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *RuleEventDtoAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetJobResult

`func (o *RuleEventDtoAllOf) GetJobResult() RuleJobResult`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *RuleEventDtoAllOf) GetJobResultOk() (*RuleJobResult, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *RuleEventDtoAllOf) SetJobResult(v RuleJobResult)`

SetJobResult sets JobResult field to given value.

### HasJobResult

`func (o *RuleEventDtoAllOf) HasJobResult() bool`

HasJobResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


