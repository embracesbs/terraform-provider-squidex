# AssetsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Total** | Pointer to **int64** | The total number of assets. | [optional] 
**Items** | [**[]AssetDto**](AssetDto.md) | The assets. | 

## Methods

### NewAssetsDto

`func NewAssetsDto(links map[string]ResourceLink, items []AssetDto, ) *AssetsDto`

NewAssetsDto instantiates a new AssetsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsDtoWithDefaults

`func NewAssetsDtoWithDefaults() *AssetsDto`

NewAssetsDtoWithDefaults instantiates a new AssetsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AssetsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssetsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssetsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetTotal

`func (o *AssetsDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetsDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetsDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AssetsDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *AssetsDto) GetItems() []AssetDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetsDto) GetItemsOk() (*[]AssetDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetsDto) SetItems(v []AssetDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


