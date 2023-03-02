# UpdateClientDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The new display name of the client. | [optional] 
**Role** | Pointer to **NullableString** | The role of the client. | [optional] 
**AllowAnonymous** | Pointer to **NullableBool** | True to allow anonymous access without an access token for this client. | [optional] 
**ApiCallsLimit** | Pointer to **NullableInt64** | The number of allowed api calls per month for this client. | [optional] 
**ApiTrafficLimit** | Pointer to **NullableInt64** | The number of allowed api traffic bytes per month for this client. | [optional] 

## Methods

### NewUpdateClientDto

`func NewUpdateClientDto() *UpdateClientDto`

NewUpdateClientDto instantiates a new UpdateClientDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClientDtoWithDefaults

`func NewUpdateClientDtoWithDefaults() *UpdateClientDto`

NewUpdateClientDtoWithDefaults instantiates a new UpdateClientDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClientDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClientDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClientDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClientDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateClientDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateClientDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *UpdateClientDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateClientDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateClientDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateClientDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *UpdateClientDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *UpdateClientDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetAllowAnonymous

`func (o *UpdateClientDto) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *UpdateClientDto) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *UpdateClientDto) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *UpdateClientDto) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.

### SetAllowAnonymousNil

`func (o *UpdateClientDto) SetAllowAnonymousNil(b bool)`

 SetAllowAnonymousNil sets the value for AllowAnonymous to be an explicit nil

### UnsetAllowAnonymous
`func (o *UpdateClientDto) UnsetAllowAnonymous()`

UnsetAllowAnonymous ensures that no value is present for AllowAnonymous, not even an explicit nil
### GetApiCallsLimit

`func (o *UpdateClientDto) GetApiCallsLimit() int64`

GetApiCallsLimit returns the ApiCallsLimit field if non-nil, zero value otherwise.

### GetApiCallsLimitOk

`func (o *UpdateClientDto) GetApiCallsLimitOk() (*int64, bool)`

GetApiCallsLimitOk returns a tuple with the ApiCallsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCallsLimit

`func (o *UpdateClientDto) SetApiCallsLimit(v int64)`

SetApiCallsLimit sets ApiCallsLimit field to given value.

### HasApiCallsLimit

`func (o *UpdateClientDto) HasApiCallsLimit() bool`

HasApiCallsLimit returns a boolean if a field has been set.

### SetApiCallsLimitNil

`func (o *UpdateClientDto) SetApiCallsLimitNil(b bool)`

 SetApiCallsLimitNil sets the value for ApiCallsLimit to be an explicit nil

### UnsetApiCallsLimit
`func (o *UpdateClientDto) UnsetApiCallsLimit()`

UnsetApiCallsLimit ensures that no value is present for ApiCallsLimit, not even an explicit nil
### GetApiTrafficLimit

`func (o *UpdateClientDto) GetApiTrafficLimit() int64`

GetApiTrafficLimit returns the ApiTrafficLimit field if non-nil, zero value otherwise.

### GetApiTrafficLimitOk

`func (o *UpdateClientDto) GetApiTrafficLimitOk() (*int64, bool)`

GetApiTrafficLimitOk returns a tuple with the ApiTrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTrafficLimit

`func (o *UpdateClientDto) SetApiTrafficLimit(v int64)`

SetApiTrafficLimit sets ApiTrafficLimit field to given value.

### HasApiTrafficLimit

`func (o *UpdateClientDto) HasApiTrafficLimit() bool`

HasApiTrafficLimit returns a boolean if a field has been set.

### SetApiTrafficLimitNil

`func (o *UpdateClientDto) SetApiTrafficLimitNil(b bool)`

 SetApiTrafficLimitNil sets the value for ApiTrafficLimit to be an explicit nil

### UnsetApiTrafficLimit
`func (o *UpdateClientDto) UnsetApiTrafficLimit()`

UnsetApiTrafficLimit ensures that no value is present for ApiTrafficLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


