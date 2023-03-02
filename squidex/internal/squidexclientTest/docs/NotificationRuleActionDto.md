# NotificationRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** | The user id or email. | 
**Text** | **string** | The text to send. | 
**Url** | Pointer to **NullableString** | The optional url to attach to the notification. | [optional] 
**Client** | Pointer to **NullableString** | An optional client name. | [optional] 

## Methods

### NewNotificationRuleActionDto

`func NewNotificationRuleActionDto(user string, text string, ) *NotificationRuleActionDto`

NewNotificationRuleActionDto instantiates a new NotificationRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleActionDtoWithDefaults

`func NewNotificationRuleActionDtoWithDefaults() *NotificationRuleActionDto`

NewNotificationRuleActionDtoWithDefaults instantiates a new NotificationRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *NotificationRuleActionDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationRuleActionDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationRuleActionDto) SetUser(v string)`

SetUser sets User field to given value.


### GetText

`func (o *NotificationRuleActionDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NotificationRuleActionDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NotificationRuleActionDto) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *NotificationRuleActionDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationRuleActionDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationRuleActionDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationRuleActionDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NotificationRuleActionDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NotificationRuleActionDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetClient

`func (o *NotificationRuleActionDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *NotificationRuleActionDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *NotificationRuleActionDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *NotificationRuleActionDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *NotificationRuleActionDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *NotificationRuleActionDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


