# ContributorDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContributorId** | **string** | The id of the user that contributes to the app. | 
**ContributorName** | **string** | The display name. | 
**ContributorEmail** | **string** | The email address. | 
**Role** | Pointer to **NullableString** | The role of the contributor. | [optional] 

## Methods

### NewContributorDtoAllOf

`func NewContributorDtoAllOf(contributorId string, contributorName string, contributorEmail string, ) *ContributorDtoAllOf`

NewContributorDtoAllOf instantiates a new ContributorDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorDtoAllOfWithDefaults

`func NewContributorDtoAllOfWithDefaults() *ContributorDtoAllOf`

NewContributorDtoAllOfWithDefaults instantiates a new ContributorDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributorId

`func (o *ContributorDtoAllOf) GetContributorId() string`

GetContributorId returns the ContributorId field if non-nil, zero value otherwise.

### GetContributorIdOk

`func (o *ContributorDtoAllOf) GetContributorIdOk() (*string, bool)`

GetContributorIdOk returns a tuple with the ContributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorId

`func (o *ContributorDtoAllOf) SetContributorId(v string)`

SetContributorId sets ContributorId field to given value.


### GetContributorName

`func (o *ContributorDtoAllOf) GetContributorName() string`

GetContributorName returns the ContributorName field if non-nil, zero value otherwise.

### GetContributorNameOk

`func (o *ContributorDtoAllOf) GetContributorNameOk() (*string, bool)`

GetContributorNameOk returns a tuple with the ContributorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorName

`func (o *ContributorDtoAllOf) SetContributorName(v string)`

SetContributorName sets ContributorName field to given value.


### GetContributorEmail

`func (o *ContributorDtoAllOf) GetContributorEmail() string`

GetContributorEmail returns the ContributorEmail field if non-nil, zero value otherwise.

### GetContributorEmailOk

`func (o *ContributorDtoAllOf) GetContributorEmailOk() (*string, bool)`

GetContributorEmailOk returns a tuple with the ContributorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorEmail

`func (o *ContributorDtoAllOf) SetContributorEmail(v string)`

SetContributorEmail sets ContributorEmail field to given value.


### GetRole

`func (o *ContributorDtoAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContributorDtoAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContributorDtoAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ContributorDtoAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ContributorDtoAllOf) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ContributorDtoAllOf) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


