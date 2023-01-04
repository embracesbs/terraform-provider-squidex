# BackupJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | Pointer to **string** | The id of the backup job. | [optional] 
**Started** | Pointer to **time.Time** | The time when the job has been started. | [optional] 
**Stopped** | Pointer to **NullableTime** | The time when the job has been stopped. | [optional] 
**HandledEvents** | Pointer to **int32** | The number of handled events. | [optional] 
**HandledAssets** | Pointer to **int32** | The number of handled assets. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 

## Methods

### NewBackupJobDto

`func NewBackupJobDto(links map[string]ResourceLink, ) *BackupJobDto`

NewBackupJobDto instantiates a new BackupJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobDtoWithDefaults

`func NewBackupJobDtoWithDefaults() *BackupJobDto`

NewBackupJobDtoWithDefaults instantiates a new BackupJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BackupJobDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupJobDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupJobDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *BackupJobDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupJobDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupJobDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupJobDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStarted

`func (o *BackupJobDto) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *BackupJobDto) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *BackupJobDto) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *BackupJobDto) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStopped

`func (o *BackupJobDto) GetStopped() time.Time`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *BackupJobDto) GetStoppedOk() (*time.Time, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *BackupJobDto) SetStopped(v time.Time)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *BackupJobDto) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### SetStoppedNil

`func (o *BackupJobDto) SetStoppedNil(b bool)`

 SetStoppedNil sets the value for Stopped to be an explicit nil

### UnsetStopped
`func (o *BackupJobDto) UnsetStopped()`

UnsetStopped ensures that no value is present for Stopped, not even an explicit nil
### GetHandledEvents

`func (o *BackupJobDto) GetHandledEvents() int32`

GetHandledEvents returns the HandledEvents field if non-nil, zero value otherwise.

### GetHandledEventsOk

`func (o *BackupJobDto) GetHandledEventsOk() (*int32, bool)`

GetHandledEventsOk returns a tuple with the HandledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledEvents

`func (o *BackupJobDto) SetHandledEvents(v int32)`

SetHandledEvents sets HandledEvents field to given value.

### HasHandledEvents

`func (o *BackupJobDto) HasHandledEvents() bool`

HasHandledEvents returns a boolean if a field has been set.

### GetHandledAssets

`func (o *BackupJobDto) GetHandledAssets() int32`

GetHandledAssets returns the HandledAssets field if non-nil, zero value otherwise.

### GetHandledAssetsOk

`func (o *BackupJobDto) GetHandledAssetsOk() (*int32, bool)`

GetHandledAssetsOk returns a tuple with the HandledAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAssets

`func (o *BackupJobDto) SetHandledAssets(v int32)`

SetHandledAssets sets HandledAssets field to given value.

### HasHandledAssets

`func (o *BackupJobDto) HasHandledAssets() bool`

HasHandledAssets returns a boolean if a field has been set.

### GetStatus

`func (o *BackupJobDto) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupJobDto) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupJobDto) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupJobDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


