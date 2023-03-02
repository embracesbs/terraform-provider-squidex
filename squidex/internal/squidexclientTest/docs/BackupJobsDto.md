# BackupJobsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]BackupJobDto**](BackupJobDto.md) | The backups. | 

## Methods

### NewBackupJobsDto

`func NewBackupJobsDto(links map[string]ResourceLink, items []BackupJobDto, ) *BackupJobsDto`

NewBackupJobsDto instantiates a new BackupJobsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobsDtoWithDefaults

`func NewBackupJobsDtoWithDefaults() *BackupJobsDto`

NewBackupJobsDtoWithDefaults instantiates a new BackupJobsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BackupJobsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupJobsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupJobsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *BackupJobsDto) GetItems() []BackupJobDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupJobsDto) GetItemsOk() (*[]BackupJobDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupJobsDto) SetItems(v []BackupJobDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


