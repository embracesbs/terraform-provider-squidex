# ScheduleJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the schedule job. | [optional] 
**Status** | Pointer to **string** | The new status. | [optional] 
**DueTime** | Pointer to **time.Time** | The target date and time when the content should be scheduled. | [optional] 
**Color** | Pointer to **string** | The color of the scheduled status. | [optional] 
**ScheduledBy** | **string** | The user who schedule the content. | 

## Methods

### NewScheduleJobDto

`func NewScheduleJobDto(scheduledBy string, ) *ScheduleJobDto`

NewScheduleJobDto instantiates a new ScheduleJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleJobDtoWithDefaults

`func NewScheduleJobDtoWithDefaults() *ScheduleJobDto`

NewScheduleJobDtoWithDefaults instantiates a new ScheduleJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleJobDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleJobDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleJobDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleJobDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduleJobDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleJobDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleJobDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduleJobDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDueTime

`func (o *ScheduleJobDto) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *ScheduleJobDto) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *ScheduleJobDto) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *ScheduleJobDto) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetColor

`func (o *ScheduleJobDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ScheduleJobDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ScheduleJobDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ScheduleJobDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetScheduledBy

`func (o *ScheduleJobDto) GetScheduledBy() string`

GetScheduledBy returns the ScheduledBy field if non-nil, zero value otherwise.

### GetScheduledByOk

`func (o *ScheduleJobDto) GetScheduledByOk() (*string, bool)`

GetScheduledByOk returns a tuple with the ScheduledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBy

`func (o *ScheduleJobDto) SetScheduledBy(v string)`

SetScheduledBy sets ScheduledBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


