# ClientDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The client id. | 
**Secret** | **string** | The client secret. | 
**Name** | **string** | The client name. | 
**Role** | Pointer to **NullableString** | The role of the client. | [optional] 
**ApiCallsLimit** | Pointer to **int64** | The number of allowed api calls per month for this client. | [optional] 
**ApiTrafficLimit** | Pointer to **int64** | The number of allowed api traffic bytes per month for this client. | [optional] 
**AllowAnonymous** | Pointer to **bool** | True to allow anonymous access without an access token for this client. | [optional] 

## Methods

### NewClientDtoAllOf

`func NewClientDtoAllOf(id string, secret string, name string, ) *ClientDtoAllOf`

NewClientDtoAllOf instantiates a new ClientDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDtoAllOfWithDefaults

`func NewClientDtoAllOfWithDefaults() *ClientDtoAllOf`

NewClientDtoAllOfWithDefaults instantiates a new ClientDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientDtoAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetSecret

`func (o *ClientDtoAllOf) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClientDtoAllOf) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClientDtoAllOf) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetName

`func (o *ClientDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ClientDtoAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClientDtoAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClientDtoAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ClientDtoAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ClientDtoAllOf) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ClientDtoAllOf) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetApiCallsLimit

`func (o *ClientDtoAllOf) GetApiCallsLimit() int64`

GetApiCallsLimit returns the ApiCallsLimit field if non-nil, zero value otherwise.

### GetApiCallsLimitOk

`func (o *ClientDtoAllOf) GetApiCallsLimitOk() (*int64, bool)`

GetApiCallsLimitOk returns a tuple with the ApiCallsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCallsLimit

`func (o *ClientDtoAllOf) SetApiCallsLimit(v int64)`

SetApiCallsLimit sets ApiCallsLimit field to given value.

### HasApiCallsLimit

`func (o *ClientDtoAllOf) HasApiCallsLimit() bool`

HasApiCallsLimit returns a boolean if a field has been set.

### GetApiTrafficLimit

`func (o *ClientDtoAllOf) GetApiTrafficLimit() int64`

GetApiTrafficLimit returns the ApiTrafficLimit field if non-nil, zero value otherwise.

### GetApiTrafficLimitOk

`func (o *ClientDtoAllOf) GetApiTrafficLimitOk() (*int64, bool)`

GetApiTrafficLimitOk returns a tuple with the ApiTrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTrafficLimit

`func (o *ClientDtoAllOf) SetApiTrafficLimit(v int64)`

SetApiTrafficLimit sets ApiTrafficLimit field to given value.

### HasApiTrafficLimit

`func (o *ClientDtoAllOf) HasApiTrafficLimit() bool`

HasApiTrafficLimit returns a boolean if a field has been set.

### GetAllowAnonymous

`func (o *ClientDtoAllOf) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *ClientDtoAllOf) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *ClientDtoAllOf) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *ClientDtoAllOf) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


