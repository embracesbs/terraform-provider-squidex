# UserDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the user. | 
**Email** | **string** | The email of the user. Unique value. | 
**DisplayName** | **string** | The display name (usually first name and last name) of the user. | 
**IsLocked** | **bool** | Determines if the user is locked. | 
**Permissions** | **[]string** | Additional permissions for the user. | 

## Methods

### NewUserDtoAllOf

`func NewUserDtoAllOf(id string, email string, displayName string, isLocked bool, permissions []string, ) *UserDtoAllOf`

NewUserDtoAllOf instantiates a new UserDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoAllOfWithDefaults

`func NewUserDtoAllOfWithDefaults() *UserDtoAllOf`

NewUserDtoAllOfWithDefaults instantiates a new UserDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDtoAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserDtoAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDtoAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDtoAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *UserDtoAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDtoAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDtoAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsLocked

`func (o *UserDtoAllOf) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserDtoAllOf) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserDtoAllOf) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetPermissions

`func (o *UserDtoAllOf) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserDtoAllOf) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserDtoAllOf) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


