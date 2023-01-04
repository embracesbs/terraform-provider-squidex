# UsersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Total** | Pointer to **int64** | The total number of users. | [optional] 
**Items** | [**[]UserDto**](UserDto.md) | The users. | 

## Methods

### NewUsersDto

`func NewUsersDto(links map[string]ResourceLink, items []UserDto, ) *UsersDto`

NewUsersDto instantiates a new UsersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersDtoWithDefaults

`func NewUsersDtoWithDefaults() *UsersDto`

NewUsersDtoWithDefaults instantiates a new UsersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UsersDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UsersDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UsersDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetTotal

`func (o *UsersDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsersDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsersDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UsersDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *UsersDto) GetItems() []UserDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsersDto) GetItemsOk() (*[]UserDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsersDto) SetItems(v []UserDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


