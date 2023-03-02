# ContentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Total** | Pointer to **int64** | The total number of content items. | [optional] 
**Items** | [**[]ContentDto**](ContentDto.md) | The content items. | 
**Statuses** | [**[]StatusInfoDto**](StatusInfoDto.md) | The possible statuses. | 

## Methods

### NewContentsDto

`func NewContentsDto(links map[string]ResourceLink, items []ContentDto, statuses []StatusInfoDto, ) *ContentsDto`

NewContentsDto instantiates a new ContentsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentsDtoWithDefaults

`func NewContentsDtoWithDefaults() *ContentsDto`

NewContentsDtoWithDefaults instantiates a new ContentsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ContentsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetTotal

`func (o *ContentsDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContentsDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContentsDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContentsDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *ContentsDto) GetItems() []ContentDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContentsDto) GetItemsOk() (*[]ContentDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContentsDto) SetItems(v []ContentDto)`

SetItems sets Items field to given value.


### GetStatuses

`func (o *ContentsDto) GetStatuses() []StatusInfoDto`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ContentsDto) GetStatusesOk() (*[]StatusInfoDto, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ContentsDto) SetStatuses(v []StatusInfoDto)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


