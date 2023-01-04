# ChangeStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The new status. | 
**DueTime** | Pointer to **NullableTime** | The due time. | [optional] 
**CheckReferrers** | Pointer to **bool** | True to check referrers of this content. | [optional] 

## Methods

### NewChangeStatusDto

`func NewChangeStatusDto(status string, ) *ChangeStatusDto`

NewChangeStatusDto instantiates a new ChangeStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeStatusDtoWithDefaults

`func NewChangeStatusDtoWithDefaults() *ChangeStatusDto`

NewChangeStatusDtoWithDefaults instantiates a new ChangeStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChangeStatusDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeStatusDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeStatusDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDueTime

`func (o *ChangeStatusDto) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *ChangeStatusDto) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *ChangeStatusDto) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *ChangeStatusDto) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *ChangeStatusDto) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *ChangeStatusDto) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil
### GetCheckReferrers

`func (o *ChangeStatusDto) GetCheckReferrers() bool`

GetCheckReferrers returns the CheckReferrers field if non-nil, zero value otherwise.

### GetCheckReferrersOk

`func (o *ChangeStatusDto) GetCheckReferrersOk() (*bool, bool)`

GetCheckReferrersOk returns a tuple with the CheckReferrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckReferrers

`func (o *ChangeStatusDto) SetCheckReferrers(v bool)`

SetCheckReferrers sets CheckReferrers field to given value.

### HasCheckReferrers

`func (o *ChangeStatusDto) HasCheckReferrers() bool`

HasCheckReferrers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


