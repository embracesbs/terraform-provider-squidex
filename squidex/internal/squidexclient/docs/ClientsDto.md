# ClientsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]ClientDto**](ClientDto.md) | The clients. | 

## Methods

### NewClientsDto

`func NewClientsDto(links map[string]ResourceLink, items []ClientDto, ) *ClientsDto`

NewClientsDto instantiates a new ClientsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientsDtoWithDefaults

`func NewClientsDtoWithDefaults() *ClientsDto`

NewClientsDtoWithDefaults instantiates a new ClientsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ClientsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *ClientsDto) GetItems() []ClientDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClientsDto) GetItemsOk() (*[]ClientDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClientsDto) SetItems(v []ClientDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


