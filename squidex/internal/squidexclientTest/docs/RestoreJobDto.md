# RestoreJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The uri to load from. | 
**Log** | **[]string** | The status log. | 
**Started** | Pointer to **time.Time** | The time when the job has been started. | [optional] 
**Stopped** | Pointer to **NullableTime** | The time when the job has been stopped. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 

## Methods

### NewRestoreJobDto

`func NewRestoreJobDto(url string, log []string, ) *RestoreJobDto`

NewRestoreJobDto instantiates a new RestoreJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobDtoWithDefaults

`func NewRestoreJobDtoWithDefaults() *RestoreJobDto`

NewRestoreJobDtoWithDefaults instantiates a new RestoreJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RestoreJobDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RestoreJobDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RestoreJobDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLog

`func (o *RestoreJobDto) GetLog() []string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *RestoreJobDto) GetLogOk() (*[]string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *RestoreJobDto) SetLog(v []string)`

SetLog sets Log field to given value.


### GetStarted

`func (o *RestoreJobDto) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RestoreJobDto) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RestoreJobDto) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RestoreJobDto) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStopped

`func (o *RestoreJobDto) GetStopped() time.Time`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *RestoreJobDto) GetStoppedOk() (*time.Time, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *RestoreJobDto) SetStopped(v time.Time)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *RestoreJobDto) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### SetStoppedNil

`func (o *RestoreJobDto) SetStoppedNil(b bool)`

 SetStoppedNil sets the value for Stopped to be an explicit nil

### UnsetStopped
`func (o *RestoreJobDto) UnsetStopped()`

UnsetStopped ensures that no value is present for Stopped, not even an explicit nil
### GetStatus

`func (o *RestoreJobDto) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreJobDto) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreJobDto) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreJobDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


