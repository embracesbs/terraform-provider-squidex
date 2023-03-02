# AllContentsByPostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | The list of ids to query. | [optional] 
**ScheduledFrom** | Pointer to **NullableTime** | The start of the schedule. | [optional] 
**ScheduledTo** | Pointer to **NullableTime** | The end of the schedule. | [optional] 

## Methods

### NewAllContentsByPostDto

`func NewAllContentsByPostDto() *AllContentsByPostDto`

NewAllContentsByPostDto instantiates a new AllContentsByPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllContentsByPostDtoWithDefaults

`func NewAllContentsByPostDtoWithDefaults() *AllContentsByPostDto`

NewAllContentsByPostDtoWithDefaults instantiates a new AllContentsByPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *AllContentsByPostDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AllContentsByPostDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AllContentsByPostDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AllContentsByPostDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *AllContentsByPostDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *AllContentsByPostDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetScheduledFrom

`func (o *AllContentsByPostDto) GetScheduledFrom() time.Time`

GetScheduledFrom returns the ScheduledFrom field if non-nil, zero value otherwise.

### GetScheduledFromOk

`func (o *AllContentsByPostDto) GetScheduledFromOk() (*time.Time, bool)`

GetScheduledFromOk returns a tuple with the ScheduledFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFrom

`func (o *AllContentsByPostDto) SetScheduledFrom(v time.Time)`

SetScheduledFrom sets ScheduledFrom field to given value.

### HasScheduledFrom

`func (o *AllContentsByPostDto) HasScheduledFrom() bool`

HasScheduledFrom returns a boolean if a field has been set.

### SetScheduledFromNil

`func (o *AllContentsByPostDto) SetScheduledFromNil(b bool)`

 SetScheduledFromNil sets the value for ScheduledFrom to be an explicit nil

### UnsetScheduledFrom
`func (o *AllContentsByPostDto) UnsetScheduledFrom()`

UnsetScheduledFrom ensures that no value is present for ScheduledFrom, not even an explicit nil
### GetScheduledTo

`func (o *AllContentsByPostDto) GetScheduledTo() time.Time`

GetScheduledTo returns the ScheduledTo field if non-nil, zero value otherwise.

### GetScheduledToOk

`func (o *AllContentsByPostDto) GetScheduledToOk() (*time.Time, bool)`

GetScheduledToOk returns a tuple with the ScheduledTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTo

`func (o *AllContentsByPostDto) SetScheduledTo(v time.Time)`

SetScheduledTo sets ScheduledTo field to given value.

### HasScheduledTo

`func (o *AllContentsByPostDto) HasScheduledTo() bool`

HasScheduledTo returns a boolean if a field has been set.

### SetScheduledToNil

`func (o *AllContentsByPostDto) SetScheduledToNil(b bool)`

 SetScheduledToNil sets the value for ScheduledTo to be an explicit nil

### UnsetScheduledTo
`func (o *AllContentsByPostDto) UnsetScheduledTo()`

UnsetScheduledTo ensures that no value is present for ScheduledTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


