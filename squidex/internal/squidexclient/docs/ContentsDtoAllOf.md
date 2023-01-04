# ContentsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | The total number of content items. | [optional] 
**Items** | [**[]ContentDto**](ContentDto.md) | The content items. | 
**Statuses** | [**[]StatusInfoDto**](StatusInfoDto.md) | The possible statuses. | 

## Methods

### NewContentsDtoAllOf

`func NewContentsDtoAllOf(items []ContentDto, statuses []StatusInfoDto, ) *ContentsDtoAllOf`

NewContentsDtoAllOf instantiates a new ContentsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentsDtoAllOfWithDefaults

`func NewContentsDtoAllOfWithDefaults() *ContentsDtoAllOf`

NewContentsDtoAllOfWithDefaults instantiates a new ContentsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ContentsDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContentsDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContentsDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContentsDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *ContentsDtoAllOf) GetItems() []ContentDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContentsDtoAllOf) GetItemsOk() (*[]ContentDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContentsDtoAllOf) SetItems(v []ContentDto)`

SetItems sets Items field to given value.


### GetStatuses

`func (o *ContentsDtoAllOf) GetStatuses() []StatusInfoDto`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ContentsDtoAllOf) GetStatusesOk() (*[]StatusInfoDto, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ContentsDtoAllOf) SetStatuses(v []StatusInfoDto)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


