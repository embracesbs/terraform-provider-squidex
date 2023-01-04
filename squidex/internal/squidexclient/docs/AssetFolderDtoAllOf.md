# AssetFolderDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the asset. | [optional] 
**ParentId** | Pointer to **string** | The id of the parent folder. Empty for files without parent. | [optional] 
**FolderName** | **string** | The folder name. | 
**Version** | Pointer to **int64** | The version of the asset folder. | [optional] 

## Methods

### NewAssetFolderDtoAllOf

`func NewAssetFolderDtoAllOf(folderName string, ) *AssetFolderDtoAllOf`

NewAssetFolderDtoAllOf instantiates a new AssetFolderDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFolderDtoAllOfWithDefaults

`func NewAssetFolderDtoAllOfWithDefaults() *AssetFolderDtoAllOf`

NewAssetFolderDtoAllOfWithDefaults instantiates a new AssetFolderDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetFolderDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetFolderDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetFolderDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetFolderDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *AssetFolderDtoAllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AssetFolderDtoAllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AssetFolderDtoAllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AssetFolderDtoAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFolderName

`func (o *AssetFolderDtoAllOf) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AssetFolderDtoAllOf) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AssetFolderDtoAllOf) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetVersion

`func (o *AssetFolderDtoAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AssetFolderDtoAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AssetFolderDtoAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AssetFolderDtoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


