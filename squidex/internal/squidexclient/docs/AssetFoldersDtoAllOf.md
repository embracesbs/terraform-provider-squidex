# AssetFoldersDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | The total number of assets. | [optional] 
**Items** | [**[]AssetFolderDto**](AssetFolderDto.md) | The assets folders. | 
**Path** | [**[]AssetFolderDto**](AssetFolderDto.md) | The path to the current folder. | 

## Methods

### NewAssetFoldersDtoAllOf

`func NewAssetFoldersDtoAllOf(items []AssetFolderDto, path []AssetFolderDto, ) *AssetFoldersDtoAllOf`

NewAssetFoldersDtoAllOf instantiates a new AssetFoldersDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFoldersDtoAllOfWithDefaults

`func NewAssetFoldersDtoAllOfWithDefaults() *AssetFoldersDtoAllOf`

NewAssetFoldersDtoAllOfWithDefaults instantiates a new AssetFoldersDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AssetFoldersDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetFoldersDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetFoldersDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AssetFoldersDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *AssetFoldersDtoAllOf) GetItems() []AssetFolderDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetFoldersDtoAllOf) GetItemsOk() (*[]AssetFolderDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetFoldersDtoAllOf) SetItems(v []AssetFolderDto)`

SetItems sets Items field to given value.


### GetPath

`func (o *AssetFoldersDtoAllOf) GetPath() []AssetFolderDto`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AssetFoldersDtoAllOf) GetPathOk() (*[]AssetFolderDto, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AssetFoldersDtoAllOf) SetPath(v []AssetFolderDto)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


