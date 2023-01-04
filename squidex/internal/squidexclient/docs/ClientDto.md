# ClientDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | **string** | The client id. | 
**Secret** | **string** | The client secret. | 
**Name** | **string** | The client name. | 
**Role** | Pointer to **NullableString** | The role of the client. | [optional] 
**ApiCallsLimit** | Pointer to **int64** | The number of allowed api calls per month for this client. | [optional] 
**ApiTrafficLimit** | Pointer to **int64** | The number of allowed api traffic bytes per month for this client. | [optional] 
**AllowAnonymous** | Pointer to **bool** | True to allow anonymous access without an access token for this client. | [optional] 

## Methods

### NewClientDto

`func NewClientDto(links map[string]ResourceLink, id string, secret string, name string, ) *ClientDto`

NewClientDto instantiates a new ClientDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDtoWithDefaults

`func NewClientDtoWithDefaults() *ClientDto`

NewClientDtoWithDefaults instantiates a new ClientDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ClientDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *ClientDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientDto) SetId(v string)`

SetId sets Id field to given value.


### GetSecret

`func (o *ClientDto) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClientDto) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClientDto) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetName

`func (o *ClientDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientDto) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ClientDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClientDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClientDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ClientDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ClientDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ClientDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetApiCallsLimit

`func (o *ClientDto) GetApiCallsLimit() int64`

GetApiCallsLimit returns the ApiCallsLimit field if non-nil, zero value otherwise.

### GetApiCallsLimitOk

`func (o *ClientDto) GetApiCallsLimitOk() (*int64, bool)`

GetApiCallsLimitOk returns a tuple with the ApiCallsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCallsLimit

`func (o *ClientDto) SetApiCallsLimit(v int64)`

SetApiCallsLimit sets ApiCallsLimit field to given value.

### HasApiCallsLimit

`func (o *ClientDto) HasApiCallsLimit() bool`

HasApiCallsLimit returns a boolean if a field has been set.

### GetApiTrafficLimit

`func (o *ClientDto) GetApiTrafficLimit() int64`

GetApiTrafficLimit returns the ApiTrafficLimit field if non-nil, zero value otherwise.

### GetApiTrafficLimitOk

`func (o *ClientDto) GetApiTrafficLimitOk() (*int64, bool)`

GetApiTrafficLimitOk returns a tuple with the ApiTrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTrafficLimit

`func (o *ClientDto) SetApiTrafficLimit(v int64)`

SetApiTrafficLimit sets ApiTrafficLimit field to given value.

### HasApiTrafficLimit

`func (o *ClientDto) HasApiTrafficLimit() bool`

HasApiTrafficLimit returns a boolean if a field has been set.

### GetAllowAnonymous

`func (o *ClientDto) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *ClientDto) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *ClientDto) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *ClientDto) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


