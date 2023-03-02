# AppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Name** | **string** | The name of the app. | 
**Label** | Pointer to **NullableString** | The optional label of the app. | [optional] 
**Description** | Pointer to **NullableString** | The optional description of the app. | [optional] 
**Version** | Pointer to **int64** | The version of the app. | [optional] 
**Id** | Pointer to **string** | The id of the app. | [optional] 
**Created** | Pointer to **time.Time** | The timestamp when the app has been created. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp when the app has been modified last. | [optional] 
**Permissions** | Pointer to **[]string** | The permission level of the user. | [optional] 
**CanAccessApi** | Pointer to **bool** | Indicates if the user can access the api. | [optional] 
**CanAccessContent** | Pointer to **bool** | Indicates if the user can access at least one content. | [optional] 
**RoleName** | Pointer to **NullableString** | The role name of the user. | [optional] 
**RoleProperties** | **interface{}** | The properties from the role. | 

## Methods

### NewAppDto

`func NewAppDto(links map[string]ResourceLink, name string, roleProperties interface{}, ) *AppDto`

NewAppDto instantiates a new AppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDtoWithDefaults

`func NewAppDtoWithDefaults() *AppDto`

NewAppDtoWithDefaults instantiates a new AppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetName

`func (o *AppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDto) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *AppDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *AppDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *AppDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDescription

`func (o *AppDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *AppDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *AppDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *AppDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *AppDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AppDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AppDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AppDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetPermissions

`func (o *AppDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AppDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AppDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AppDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCanAccessApi

`func (o *AppDto) GetCanAccessApi() bool`

GetCanAccessApi returns the CanAccessApi field if non-nil, zero value otherwise.

### GetCanAccessApiOk

`func (o *AppDto) GetCanAccessApiOk() (*bool, bool)`

GetCanAccessApiOk returns a tuple with the CanAccessApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessApi

`func (o *AppDto) SetCanAccessApi(v bool)`

SetCanAccessApi sets CanAccessApi field to given value.

### HasCanAccessApi

`func (o *AppDto) HasCanAccessApi() bool`

HasCanAccessApi returns a boolean if a field has been set.

### GetCanAccessContent

`func (o *AppDto) GetCanAccessContent() bool`

GetCanAccessContent returns the CanAccessContent field if non-nil, zero value otherwise.

### GetCanAccessContentOk

`func (o *AppDto) GetCanAccessContentOk() (*bool, bool)`

GetCanAccessContentOk returns a tuple with the CanAccessContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessContent

`func (o *AppDto) SetCanAccessContent(v bool)`

SetCanAccessContent sets CanAccessContent field to given value.

### HasCanAccessContent

`func (o *AppDto) HasCanAccessContent() bool`

HasCanAccessContent returns a boolean if a field has been set.

### GetRoleName

`func (o *AppDto) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AppDto) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AppDto) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AppDto) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *AppDto) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *AppDto) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRoleProperties

`func (o *AppDto) GetRoleProperties() interface{}`

GetRoleProperties returns the RoleProperties field if non-nil, zero value otherwise.

### GetRolePropertiesOk

`func (o *AppDto) GetRolePropertiesOk() (*interface{}, bool)`

GetRolePropertiesOk returns a tuple with the RoleProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProperties

`func (o *AppDto) SetRoleProperties(v interface{})`

SetRoleProperties sets RoleProperties field to given value.


### SetRolePropertiesNil

`func (o *AppDto) SetRolePropertiesNil(b bool)`

 SetRolePropertiesNil sets the value for RoleProperties to be an explicit nil

### UnsetRoleProperties
`func (o *AppDto) UnsetRoleProperties()`

UnsetRoleProperties ensures that no value is present for RoleProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


