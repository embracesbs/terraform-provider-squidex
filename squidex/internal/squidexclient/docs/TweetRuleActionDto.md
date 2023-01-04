# TweetRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  The generated access token. | 
**AccessSecret** | **string** |  The generated access secret. | 
**Text** | **string** | The text that is sent as tweet to twitter. | 

## Methods

### NewTweetRuleActionDto

`func NewTweetRuleActionDto(accessToken string, accessSecret string, text string, ) *TweetRuleActionDto`

NewTweetRuleActionDto instantiates a new TweetRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetRuleActionDtoWithDefaults

`func NewTweetRuleActionDtoWithDefaults() *TweetRuleActionDto`

NewTweetRuleActionDtoWithDefaults instantiates a new TweetRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TweetRuleActionDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TweetRuleActionDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TweetRuleActionDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccessSecret

`func (o *TweetRuleActionDto) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *TweetRuleActionDto) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *TweetRuleActionDto) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.


### GetText

`func (o *TweetRuleActionDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TweetRuleActionDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TweetRuleActionDto) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


