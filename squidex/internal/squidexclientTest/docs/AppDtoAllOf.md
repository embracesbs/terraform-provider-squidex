# AppDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewAppDtoAllOf

`func NewAppDtoAllOf(name string, roleProperties interface{}, ) *AppDtoAllOf`

NewAppDtoAllOf instantiates a new AppDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDtoAllOfWithDefaults

`func NewAppDtoAllOfWithDefaults() *AppDtoAllOf`

NewAppDtoAllOfWithDefaults instantiates a new AppDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *AppDtoAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppDtoAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppDtoAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppDtoAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *AppDtoAllOf) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *AppDtoAllOf) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDescription

`func (o *AppDtoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDtoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDtoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDtoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppDtoAllOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppDtoAllOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *AppDtoAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppDtoAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppDtoAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppDtoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *AppDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *AppDtoAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppDtoAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppDtoAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppDtoAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *AppDtoAllOf) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AppDtoAllOf) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AppDtoAllOf) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AppDtoAllOf) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetPermissions

`func (o *AppDtoAllOf) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AppDtoAllOf) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AppDtoAllOf) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AppDtoAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCanAccessApi

`func (o *AppDtoAllOf) GetCanAccessApi() bool`

GetCanAccessApi returns the CanAccessApi field if non-nil, zero value otherwise.

### GetCanAccessApiOk

`func (o *AppDtoAllOf) GetCanAccessApiOk() (*bool, bool)`

GetCanAccessApiOk returns a tuple with the CanAccessApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessApi

`func (o *AppDtoAllOf) SetCanAccessApi(v bool)`

SetCanAccessApi sets CanAccessApi field to given value.

### HasCanAccessApi

`func (o *AppDtoAllOf) HasCanAccessApi() bool`

HasCanAccessApi returns a boolean if a field has been set.

### GetCanAccessContent

`func (o *AppDtoAllOf) GetCanAccessContent() bool`

GetCanAccessContent returns the CanAccessContent field if non-nil, zero value otherwise.

### GetCanAccessContentOk

`func (o *AppDtoAllOf) GetCanAccessContentOk() (*bool, bool)`

GetCanAccessContentOk returns a tuple with the CanAccessContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessContent

`func (o *AppDtoAllOf) SetCanAccessContent(v bool)`

SetCanAccessContent sets CanAccessContent field to given value.

### HasCanAccessContent

`func (o *AppDtoAllOf) HasCanAccessContent() bool`

HasCanAccessContent returns a boolean if a field has been set.

### GetRoleName

`func (o *AppDtoAllOf) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AppDtoAllOf) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AppDtoAllOf) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AppDtoAllOf) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *AppDtoAllOf) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *AppDtoAllOf) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRoleProperties

`func (o *AppDtoAllOf) GetRoleProperties() interface{}`

GetRoleProperties returns the RoleProperties field if non-nil, zero value otherwise.

### GetRolePropertiesOk

`func (o *AppDtoAllOf) GetRolePropertiesOk() (*interface{}, bool)`

GetRolePropertiesOk returns a tuple with the RoleProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProperties

`func (o *AppDtoAllOf) SetRoleProperties(v interface{})`

SetRoleProperties sets RoleProperties field to given value.


### SetRolePropertiesNil

`func (o *AppDtoAllOf) SetRolePropertiesNil(b bool)`

 SetRolePropertiesNil sets the value for RoleProperties to be an explicit nil

### UnsetRoleProperties
`func (o *AppDtoAllOf) UnsetRoleProperties()`

UnsetRoleProperties ensures that no value is present for RoleProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


