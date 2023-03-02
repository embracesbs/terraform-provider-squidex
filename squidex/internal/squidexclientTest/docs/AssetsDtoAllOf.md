# AssetsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | The total number of assets. | [optional] 
**Items** | [**[]AssetDto**](AssetDto.md) | The assets. | 

## Methods

### NewAssetsDtoAllOf

`func NewAssetsDtoAllOf(items []AssetDto, ) *AssetsDtoAllOf`

NewAssetsDtoAllOf instantiates a new AssetsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsDtoAllOfWithDefaults

`func NewAssetsDtoAllOfWithDefaults() *AssetsDtoAllOf`

NewAssetsDtoAllOfWithDefaults instantiates a new AssetsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AssetsDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetsDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetsDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AssetsDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *AssetsDtoAllOf) GetItems() []AssetDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetsDtoAllOf) GetItemsOk() (*[]AssetDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetsDtoAllOf) SetItems(v []AssetDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


