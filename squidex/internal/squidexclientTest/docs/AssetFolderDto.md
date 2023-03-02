# AssetFolderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | Pointer to **string** | The id of the asset. | [optional] 
**ParentId** | Pointer to **string** | The id of the parent folder. Empty for files without parent. | [optional] 
**FolderName** | **string** | The folder name. | 
**Version** | Pointer to **int64** | The version of the asset folder. | [optional] 

## Methods

### NewAssetFolderDto

`func NewAssetFolderDto(links map[string]ResourceLink, folderName string, ) *AssetFolderDto`

NewAssetFolderDto instantiates a new AssetFolderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFolderDtoWithDefaults

`func NewAssetFolderDtoWithDefaults() *AssetFolderDto`

NewAssetFolderDtoWithDefaults instantiates a new AssetFolderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AssetFolderDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssetFolderDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssetFolderDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *AssetFolderDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetFolderDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetFolderDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetFolderDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *AssetFolderDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AssetFolderDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AssetFolderDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AssetFolderDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFolderName

`func (o *AssetFolderDto) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AssetFolderDto) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AssetFolderDto) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetVersion

`func (o *AssetFolderDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AssetFolderDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AssetFolderDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AssetFolderDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


