# RuleEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
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

### NewRuleEventDto

`func NewRuleEventDto(links map[string]ResourceLink, description string, eventName string, ) *RuleEventDto`

NewRuleEventDto instantiates a new RuleEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEventDtoWithDefaults

`func NewRuleEventDtoWithDefaults() *RuleEventDto`

NewRuleEventDtoWithDefaults instantiates a new RuleEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RuleEventDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RuleEventDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RuleEventDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *RuleEventDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleEventDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleEventDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleEventDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *RuleEventDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RuleEventDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RuleEventDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RuleEventDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *RuleEventDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleEventDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleEventDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventName

`func (o *RuleEventDto) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *RuleEventDto) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *RuleEventDto) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetLastDump

`func (o *RuleEventDto) GetLastDump() string`

GetLastDump returns the LastDump field if non-nil, zero value otherwise.

### GetLastDumpOk

`func (o *RuleEventDto) GetLastDumpOk() (*string, bool)`

GetLastDumpOk returns a tuple with the LastDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDump

`func (o *RuleEventDto) SetLastDump(v string)`

SetLastDump sets LastDump field to given value.

### HasLastDump

`func (o *RuleEventDto) HasLastDump() bool`

HasLastDump returns a boolean if a field has been set.

### SetLastDumpNil

`func (o *RuleEventDto) SetLastDumpNil(b bool)`

 SetLastDumpNil sets the value for LastDump to be an explicit nil

### UnsetLastDump
`func (o *RuleEventDto) UnsetLastDump()`

UnsetLastDump ensures that no value is present for LastDump, not even an explicit nil
### GetNumCalls

`func (o *RuleEventDto) GetNumCalls() int32`

GetNumCalls returns the NumCalls field if non-nil, zero value otherwise.

### GetNumCallsOk

`func (o *RuleEventDto) GetNumCallsOk() (*int32, bool)`

GetNumCallsOk returns a tuple with the NumCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCalls

`func (o *RuleEventDto) SetNumCalls(v int32)`

SetNumCalls sets NumCalls field to given value.

### HasNumCalls

`func (o *RuleEventDto) HasNumCalls() bool`

HasNumCalls returns a boolean if a field has been set.

### GetNextAttempt

`func (o *RuleEventDto) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *RuleEventDto) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *RuleEventDto) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *RuleEventDto) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### SetNextAttemptNil

`func (o *RuleEventDto) SetNextAttemptNil(b bool)`

 SetNextAttemptNil sets the value for NextAttempt to be an explicit nil

### UnsetNextAttempt
`func (o *RuleEventDto) UnsetNextAttempt()`

UnsetNextAttempt ensures that no value is present for NextAttempt, not even an explicit nil
### GetResult

`func (o *RuleEventDto) GetResult() RuleResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RuleEventDto) GetResultOk() (*RuleResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RuleEventDto) SetResult(v RuleResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *RuleEventDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetJobResult

`func (o *RuleEventDto) GetJobResult() RuleJobResult`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *RuleEventDto) GetJobResultOk() (*RuleJobResult, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *RuleEventDto) SetJobResult(v RuleJobResult)`

SetJobResult sets JobResult field to given value.

### HasJobResult

`func (o *RuleEventDto) HasJobResult() bool`

HasJobResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


