# AssignContributorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContributorId** | **string** | The id or email of the user to add to the app. | 
**Role** | Pointer to **NullableString** | The role of the contributor. | [optional] 
**Invite** | Pointer to **bool** | Set to true to invite the user if he does not exist. | [optional] 

## Methods

### NewAssignContributorDto

`func NewAssignContributorDto(contributorId string, ) *AssignContributorDto`

NewAssignContributorDto instantiates a new AssignContributorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignContributorDtoWithDefaults

`func NewAssignContributorDtoWithDefaults() *AssignContributorDto`

NewAssignContributorDtoWithDefaults instantiates a new AssignContributorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributorId

`func (o *AssignContributorDto) GetContributorId() string`

GetContributorId returns the ContributorId field if non-nil, zero value otherwise.

### GetContributorIdOk

`func (o *AssignContributorDto) GetContributorIdOk() (*string, bool)`

GetContributorIdOk returns a tuple with the ContributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorId

`func (o *AssignContributorDto) SetContributorId(v string)`

SetContributorId sets ContributorId field to given value.


### GetRole

`func (o *AssignContributorDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssignContributorDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssignContributorDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssignContributorDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AssignContributorDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AssignContributorDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetInvite

`func (o *AssignContributorDto) GetInvite() bool`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *AssignContributorDto) GetInviteOk() (*bool, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *AssignContributorDto) SetInvite(v bool)`

SetInvite sets Invite field to given value.

### HasInvite

`func (o *AssignContributorDto) HasInvite() bool`

HasInvite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


