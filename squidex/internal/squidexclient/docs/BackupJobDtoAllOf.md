# BackupJobDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the backup job. | [optional] 
**Started** | Pointer to **time.Time** | The time when the job has been started. | [optional] 
**Stopped** | Pointer to **NullableTime** | The time when the job has been stopped. | [optional] 
**HandledEvents** | Pointer to **int32** | The number of handled events. | [optional] 
**HandledAssets** | Pointer to **int32** | The number of handled assets. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 

## Methods

### NewBackupJobDtoAllOf

`func NewBackupJobDtoAllOf() *BackupJobDtoAllOf`

NewBackupJobDtoAllOf instantiates a new BackupJobDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobDtoAllOfWithDefaults

`func NewBackupJobDtoAllOfWithDefaults() *BackupJobDtoAllOf`

NewBackupJobDtoAllOfWithDefaults instantiates a new BackupJobDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupJobDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupJobDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupJobDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupJobDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStarted

`func (o *BackupJobDtoAllOf) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *BackupJobDtoAllOf) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *BackupJobDtoAllOf) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *BackupJobDtoAllOf) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStopped

`func (o *BackupJobDtoAllOf) GetStopped() time.Time`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *BackupJobDtoAllOf) GetStoppedOk() (*time.Time, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *BackupJobDtoAllOf) SetStopped(v time.Time)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *BackupJobDtoAllOf) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### SetStoppedNil

`func (o *BackupJobDtoAllOf) SetStoppedNil(b bool)`

 SetStoppedNil sets the value for Stopped to be an explicit nil

### UnsetStopped
`func (o *BackupJobDtoAllOf) UnsetStopped()`

UnsetStopped ensures that no value is present for Stopped, not even an explicit nil
### GetHandledEvents

`func (o *BackupJobDtoAllOf) GetHandledEvents() int32`

GetHandledEvents returns the HandledEvents field if non-nil, zero value otherwise.

### GetHandledEventsOk

`func (o *BackupJobDtoAllOf) GetHandledEventsOk() (*int32, bool)`

GetHandledEventsOk returns a tuple with the HandledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledEvents

`func (o *BackupJobDtoAllOf) SetHandledEvents(v int32)`

SetHandledEvents sets HandledEvents field to given value.

### HasHandledEvents

`func (o *BackupJobDtoAllOf) HasHandledEvents() bool`

HasHandledEvents returns a boolean if a field has been set.

### GetHandledAssets

`func (o *BackupJobDtoAllOf) GetHandledAssets() int32`

GetHandledAssets returns the HandledAssets field if non-nil, zero value otherwise.

### GetHandledAssetsOk

`func (o *BackupJobDtoAllOf) GetHandledAssetsOk() (*int32, bool)`

GetHandledAssetsOk returns a tuple with the HandledAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAssets

`func (o *BackupJobDtoAllOf) SetHandledAssets(v int32)`

SetHandledAssets sets HandledAssets field to given value.

### HasHandledAssets

`func (o *BackupJobDtoAllOf) HasHandledAssets() bool`

HasHandledAssets returns a boolean if a field has been set.

### GetStatus

`func (o *BackupJobDtoAllOf) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupJobDtoAllOf) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupJobDtoAllOf) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupJobDtoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


