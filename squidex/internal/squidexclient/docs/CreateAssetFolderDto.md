# CreateAssetFolderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderName** | **string** | The name of the folder. | 
**ParentId** | Pointer to **string** | The id of the parent folder. | [optional] 

## Methods

### NewCreateAssetFolderDto

`func NewCreateAssetFolderDto(folderName string, ) *CreateAssetFolderDto`

NewCreateAssetFolderDto instantiates a new CreateAssetFolderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetFolderDtoWithDefaults

`func NewCreateAssetFolderDtoWithDefaults() *CreateAssetFolderDto`

NewCreateAssetFolderDtoWithDefaults instantiates a new CreateAssetFolderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderName

`func (o *CreateAssetFolderDto) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *CreateAssetFolderDto) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *CreateAssetFolderDto) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetParentId

`func (o *CreateAssetFolderDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateAssetFolderDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateAssetFolderDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateAssetFolderDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


