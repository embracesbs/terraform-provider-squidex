# MediumRuleActionDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The self issued access token. | 
**Title** | **string** | The title, used for the url. | 
**Content** | **string** | The content, either html or markdown. | 
**CanonicalUrl** | Pointer to **NullableString** | The original home of this content, if it was originally published elsewhere. | [optional] 
**Tags** | Pointer to **NullableString** | The optional comma separated list of tags. | [optional] 
**PublicationId** | Pointer to **NullableString** | Optional publication id. | [optional] 
**IsHtml** | Pointer to **bool** | Indicates whether the content is markdown or html. | [optional] 

## Methods

### NewMediumRuleActionDtoAllOf

`func NewMediumRuleActionDtoAllOf(accessToken string, title string, content string, ) *MediumRuleActionDtoAllOf`

NewMediumRuleActionDtoAllOf instantiates a new MediumRuleActionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediumRuleActionDtoAllOfWithDefaults

`func NewMediumRuleActionDtoAllOfWithDefaults() *MediumRuleActionDtoAllOf`

NewMediumRuleActionDtoAllOfWithDefaults instantiates a new MediumRuleActionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MediumRuleActionDtoAllOf) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MediumRuleActionDtoAllOf) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MediumRuleActionDtoAllOf) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTitle

`func (o *MediumRuleActionDtoAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MediumRuleActionDtoAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MediumRuleActionDtoAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *MediumRuleActionDtoAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MediumRuleActionDtoAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MediumRuleActionDtoAllOf) SetContent(v string)`

SetContent sets Content field to given value.


### GetCanonicalUrl

`func (o *MediumRuleActionDtoAllOf) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *MediumRuleActionDtoAllOf) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *MediumRuleActionDtoAllOf) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *MediumRuleActionDtoAllOf) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *MediumRuleActionDtoAllOf) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *MediumRuleActionDtoAllOf) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetTags

`func (o *MediumRuleActionDtoAllOf) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MediumRuleActionDtoAllOf) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MediumRuleActionDtoAllOf) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MediumRuleActionDtoAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MediumRuleActionDtoAllOf) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MediumRuleActionDtoAllOf) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPublicationId

`func (o *MediumRuleActionDtoAllOf) GetPublicationId() string`

GetPublicationId returns the PublicationId field if non-nil, zero value otherwise.

### GetPublicationIdOk

`func (o *MediumRuleActionDtoAllOf) GetPublicationIdOk() (*string, bool)`

GetPublicationIdOk returns a tuple with the PublicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationId

`func (o *MediumRuleActionDtoAllOf) SetPublicationId(v string)`

SetPublicationId sets PublicationId field to given value.

### HasPublicationId

`func (o *MediumRuleActionDtoAllOf) HasPublicationId() bool`

HasPublicationId returns a boolean if a field has been set.

### SetPublicationIdNil

`func (o *MediumRuleActionDtoAllOf) SetPublicationIdNil(b bool)`

 SetPublicationIdNil sets the value for PublicationId to be an explicit nil

### UnsetPublicationId
`func (o *MediumRuleActionDtoAllOf) UnsetPublicationId()`

UnsetPublicationId ensures that no value is present for PublicationId, not even an explicit nil
### GetIsHtml

`func (o *MediumRuleActionDtoAllOf) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *MediumRuleActionDtoAllOf) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *MediumRuleActionDtoAllOf) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.

### HasIsHtml

`func (o *MediumRuleActionDtoAllOf) HasIsHtml() bool`

HasIsHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


