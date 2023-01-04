# ContributorsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ContributorDto**](ContributorDto.md) | The contributors. | 
**MaxContributors** | Pointer to **int32** | The maximum number of allowed contributors. | [optional] 
**Meta** | Pointer to [**ContributorsMetadata**](ContributorsMetadata.md) |  | [optional] 

## Methods

### NewContributorsDtoAllOf

`func NewContributorsDtoAllOf(items []ContributorDto, ) *ContributorsDtoAllOf`

NewContributorsDtoAllOf instantiates a new ContributorsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorsDtoAllOfWithDefaults

`func NewContributorsDtoAllOfWithDefaults() *ContributorsDtoAllOf`

NewContributorsDtoAllOfWithDefaults instantiates a new ContributorsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContributorsDtoAllOf) GetItems() []ContributorDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContributorsDtoAllOf) GetItemsOk() (*[]ContributorDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContributorsDtoAllOf) SetItems(v []ContributorDto)`

SetItems sets Items field to given value.


### GetMaxContributors

`func (o *ContributorsDtoAllOf) GetMaxContributors() int32`

GetMaxContributors returns the MaxContributors field if non-nil, zero value otherwise.

### GetMaxContributorsOk

`func (o *ContributorsDtoAllOf) GetMaxContributorsOk() (*int32, bool)`

GetMaxContributorsOk returns a tuple with the MaxContributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContributors

`func (o *ContributorsDtoAllOf) SetMaxContributors(v int32)`

SetMaxContributors sets MaxContributors field to given value.

### HasMaxContributors

`func (o *ContributorsDtoAllOf) HasMaxContributors() bool`

HasMaxContributors returns a boolean if a field has been set.

### GetMeta

`func (o *ContributorsDtoAllOf) GetMeta() ContributorsMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContributorsDtoAllOf) GetMetaOk() (*ContributorsMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContributorsDtoAllOf) SetMeta(v ContributorsMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ContributorsDtoAllOf) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


