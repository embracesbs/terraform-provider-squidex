# UpdateRoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | Associated list of permissions. | 
**Properties** | Pointer to **map[string]interface{}** | Associated list of UI properties. | [optional] 

## Methods

### NewUpdateRoleDto

`func NewUpdateRoleDto(permissions []string, ) *UpdateRoleDto`

NewUpdateRoleDto instantiates a new UpdateRoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleDtoWithDefaults

`func NewUpdateRoleDtoWithDefaults() *UpdateRoleDto`

NewUpdateRoleDtoWithDefaults instantiates a new UpdateRoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UpdateRoleDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateRoleDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateRoleDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetProperties

`func (o *UpdateRoleDto) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdateRoleDto) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdateRoleDto) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UpdateRoleDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


