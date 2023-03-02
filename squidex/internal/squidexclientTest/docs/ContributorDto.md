# ContributorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**ContributorId** | **string** | The id of the user that contributes to the app. | 
**ContributorName** | **string** | The display name. | 
**ContributorEmail** | **string** | The email address. | 
**Role** | Pointer to **NullableString** | The role of the contributor. | [optional] 

## Methods

### NewContributorDto

`func NewContributorDto(links map[string]ResourceLink, contributorId string, contributorName string, contributorEmail string, ) *ContributorDto`

NewContributorDto instantiates a new ContributorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorDtoWithDefaults

`func NewContributorDtoWithDefaults() *ContributorDto`

NewContributorDtoWithDefaults instantiates a new ContributorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ContributorDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContributorDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContributorDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetContributorId

`func (o *ContributorDto) GetContributorId() string`

GetContributorId returns the ContributorId field if non-nil, zero value otherwise.

### GetContributorIdOk

`func (o *ContributorDto) GetContributorIdOk() (*string, bool)`

GetContributorIdOk returns a tuple with the ContributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorId

`func (o *ContributorDto) SetContributorId(v string)`

SetContributorId sets ContributorId field to given value.


### GetContributorName

`func (o *ContributorDto) GetContributorName() string`

GetContributorName returns the ContributorName field if non-nil, zero value otherwise.

### GetContributorNameOk

`func (o *ContributorDto) GetContributorNameOk() (*string, bool)`

GetContributorNameOk returns a tuple with the ContributorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorName

`func (o *ContributorDto) SetContributorName(v string)`

SetContributorName sets ContributorName field to given value.


### GetContributorEmail

`func (o *ContributorDto) GetContributorEmail() string`

GetContributorEmail returns the ContributorEmail field if non-nil, zero value otherwise.

### GetContributorEmailOk

`func (o *ContributorDto) GetContributorEmailOk() (*string, bool)`

GetContributorEmailOk returns a tuple with the ContributorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorEmail

`func (o *ContributorDto) SetContributorEmail(v string)`

SetContributorEmail sets ContributorEmail field to given value.


### GetRole

`func (o *ContributorDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContributorDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContributorDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ContributorDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ContributorDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ContributorDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


