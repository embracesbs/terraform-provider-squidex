# ContributorsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]ContributorDto**](ContributorDto.md) | The contributors. | 
**MaxContributors** | Pointer to **int32** | The maximum number of allowed contributors. | [optional] 
**Meta** | Pointer to [**ContributorsMetadata**](ContributorsMetadata.md) |  | [optional] 

## Methods

### NewContributorsDto

`func NewContributorsDto(links map[string]ResourceLink, items []ContributorDto, ) *ContributorsDto`

NewContributorsDto instantiates a new ContributorsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorsDtoWithDefaults

`func NewContributorsDtoWithDefaults() *ContributorsDto`

NewContributorsDtoWithDefaults instantiates a new ContributorsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ContributorsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContributorsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContributorsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *ContributorsDto) GetItems() []ContributorDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContributorsDto) GetItemsOk() (*[]ContributorDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContributorsDto) SetItems(v []ContributorDto)`

SetItems sets Items field to given value.


### GetMaxContributors

`func (o *ContributorsDto) GetMaxContributors() int32`

GetMaxContributors returns the MaxContributors field if non-nil, zero value otherwise.

### GetMaxContributorsOk

`func (o *ContributorsDto) GetMaxContributorsOk() (*int32, bool)`

GetMaxContributorsOk returns a tuple with the MaxContributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContributors

`func (o *ContributorsDto) SetMaxContributors(v int32)`

SetMaxContributors sets MaxContributors field to given value.

### HasMaxContributors

`func (o *ContributorsDto) HasMaxContributors() bool`

HasMaxContributors returns a boolean if a field has been set.

### GetMeta

`func (o *ContributorsDto) GetMeta() ContributorsMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContributorsDto) GetMetaOk() (*ContributorsMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContributorsDto) SetMeta(v ContributorsMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ContributorsDto) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


