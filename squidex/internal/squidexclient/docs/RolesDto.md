# RolesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]RoleDto**](RoleDto.md) | The roles. | 

## Methods

### NewRolesDto

`func NewRolesDto(links map[string]ResourceLink, items []RoleDto, ) *RolesDto`

NewRolesDto instantiates a new RolesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesDtoWithDefaults

`func NewRolesDtoWithDefaults() *RolesDto`

NewRolesDtoWithDefaults instantiates a new RolesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RolesDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RolesDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RolesDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *RolesDto) GetItems() []RoleDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RolesDto) GetItemsOk() (*[]RoleDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RolesDto) SetItems(v []RoleDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


