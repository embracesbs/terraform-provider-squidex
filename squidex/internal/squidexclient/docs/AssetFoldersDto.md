# AssetFoldersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Total** | Pointer to **int64** | The total number of assets. | [optional] 
**Items** | [**[]AssetFolderDto**](AssetFolderDto.md) | The assets folders. | 
**Path** | [**[]AssetFolderDto**](AssetFolderDto.md) | The path to the current folder. | 

## Methods

### NewAssetFoldersDto

`func NewAssetFoldersDto(links map[string]ResourceLink, items []AssetFolderDto, path []AssetFolderDto, ) *AssetFoldersDto`

NewAssetFoldersDto instantiates a new AssetFoldersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFoldersDtoWithDefaults

`func NewAssetFoldersDtoWithDefaults() *AssetFoldersDto`

NewAssetFoldersDtoWithDefaults instantiates a new AssetFoldersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AssetFoldersDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssetFoldersDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssetFoldersDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetTotal

`func (o *AssetFoldersDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetFoldersDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetFoldersDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AssetFoldersDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *AssetFoldersDto) GetItems() []AssetFolderDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetFoldersDto) GetItemsOk() (*[]AssetFolderDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetFoldersDto) SetItems(v []AssetFolderDto)`

SetItems sets Items field to given value.


### GetPath

`func (o *AssetFoldersDto) GetPath() []AssetFolderDto`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AssetFoldersDto) GetPathOk() (*[]AssetFolderDto, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AssetFoldersDto) SetPath(v []AssetFolderDto)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


