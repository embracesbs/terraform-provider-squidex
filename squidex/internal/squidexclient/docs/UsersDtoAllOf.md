# UsersDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | The total number of users. | [optional] 
**Items** | [**[]UserDto**](UserDto.md) | The users. | 

## Methods

### NewUsersDtoAllOf

`func NewUsersDtoAllOf(items []UserDto, ) *UsersDtoAllOf`

NewUsersDtoAllOf instantiates a new UsersDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersDtoAllOfWithDefaults

`func NewUsersDtoAllOfWithDefaults() *UsersDtoAllOf`

NewUsersDtoAllOfWithDefaults instantiates a new UsersDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *UsersDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsersDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsersDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UsersDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *UsersDtoAllOf) GetItems() []UserDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsersDtoAllOf) GetItemsOk() (*[]UserDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsersDtoAllOf) SetItems(v []UserDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


