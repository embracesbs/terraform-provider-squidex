# CreateUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user. Unique value. | 
**DisplayName** | **string** | The display name (usually first name and last name) of the user. | 
**Password** | **string** | The password of the user. | 
**Permissions** | **[]string** | Additional permissions for the user. | 

## Methods

### NewCreateUserDto

`func NewCreateUserDto(email string, displayName string, password string, permissions []string, ) *CreateUserDto`

NewCreateUserDto instantiates a new CreateUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserDtoWithDefaults

`func NewCreateUserDtoWithDefaults() *CreateUserDto`

NewCreateUserDtoWithDefaults instantiates a new CreateUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *CreateUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPassword

`func (o *CreateUserDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserDto) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPermissions

`func (o *CreateUserDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateUserDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateUserDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


