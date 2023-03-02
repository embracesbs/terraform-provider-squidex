# RoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Name** | **string** | The role name. | 
**NumClients** | Pointer to **int32** | The number of clients with this role. | [optional] 
**NumContributors** | Pointer to **int32** | The number of contributors with this role. | [optional] 
**IsDefaultRole** | Pointer to **bool** | Indicates if the role is an builtin default role. | [optional] 
**Permissions** | **[]string** | Associated list of permissions. | 
**Properties** | **map[string]interface{}** | Associated list of UI properties. | 

## Methods

### NewRoleDto

`func NewRoleDto(links map[string]ResourceLink, name string, permissions []string, properties map[string]interface{}, ) *RoleDto`

NewRoleDto instantiates a new RoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDtoWithDefaults

`func NewRoleDtoWithDefaults() *RoleDto`

NewRoleDtoWithDefaults instantiates a new RoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RoleDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetName

`func (o *RoleDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDto) SetName(v string)`

SetName sets Name field to given value.


### GetNumClients

`func (o *RoleDto) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *RoleDto) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *RoleDto) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *RoleDto) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetNumContributors

`func (o *RoleDto) GetNumContributors() int32`

GetNumContributors returns the NumContributors field if non-nil, zero value otherwise.

### GetNumContributorsOk

`func (o *RoleDto) GetNumContributorsOk() (*int32, bool)`

GetNumContributorsOk returns a tuple with the NumContributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumContributors

`func (o *RoleDto) SetNumContributors(v int32)`

SetNumContributors sets NumContributors field to given value.

### HasNumContributors

`func (o *RoleDto) HasNumContributors() bool`

HasNumContributors returns a boolean if a field has been set.

### GetIsDefaultRole

`func (o *RoleDto) GetIsDefaultRole() bool`

GetIsDefaultRole returns the IsDefaultRole field if non-nil, zero value otherwise.

### GetIsDefaultRoleOk

`func (o *RoleDto) GetIsDefaultRoleOk() (*bool, bool)`

GetIsDefaultRoleOk returns a tuple with the IsDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRole

`func (o *RoleDto) SetIsDefaultRole(v bool)`

SetIsDefaultRole sets IsDefaultRole field to given value.

### HasIsDefaultRole

`func (o *RoleDto) HasIsDefaultRole() bool`

HasIsDefaultRole returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetProperties

`func (o *RoleDto) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RoleDto) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RoleDto) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

