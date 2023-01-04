# MediumRuleActionDto

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

### NewMediumRuleActionDto

`func NewMediumRuleActionDto(accessToken string, title string, content string, ) *MediumRuleActionDto`

NewMediumRuleActionDto instantiates a new MediumRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediumRuleActionDtoWithDefaults

`func NewMediumRuleActionDtoWithDefaults() *MediumRuleActionDto`

NewMediumRuleActionDtoWithDefaults instantiates a new MediumRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MediumRuleActionDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MediumRuleActionDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MediumRuleActionDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTitle

`func (o *MediumRuleActionDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MediumRuleActionDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MediumRuleActionDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *MediumRuleActionDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MediumRuleActionDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MediumRuleActionDto) SetContent(v string)`

SetContent sets Content field to given value.


### GetCanonicalUrl

`func (o *MediumRuleActionDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *MediumRuleActionDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *MediumRuleActionDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *MediumRuleActionDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *MediumRuleActionDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *MediumRuleActionDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetTags

`func (o *MediumRuleActionDto) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MediumRuleActionDto) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MediumRuleActionDto) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MediumRuleActionDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MediumRuleActionDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MediumRuleActionDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPublicationId

`func (o *MediumRuleActionDto) GetPublicationId() string`

GetPublicationId returns the PublicationId field if non-nil, zero value otherwise.

### GetPublicationIdOk

`func (o *MediumRuleActionDto) GetPublicationIdOk() (*string, bool)`

GetPublicationIdOk returns a tuple with the PublicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationId

`func (o *MediumRuleActionDto) SetPublicationId(v string)`

SetPublicationId sets PublicationId field to given value.

### HasPublicationId

`func (o *MediumRuleActionDto) HasPublicationId() bool`

HasPublicationId returns a boolean if a field has been set.

### SetPublicationIdNil

`func (o *MediumRuleActionDto) SetPublicationIdNil(b bool)`

 SetPublicationIdNil sets the value for PublicationId to be an explicit nil

### UnsetPublicationId
`func (o *MediumRuleActionDto) UnsetPublicationId()`

UnsetPublicationId ensures that no value is present for PublicationId, not even an explicit nil
### GetIsHtml

`func (o *MediumRuleActionDto) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *MediumRuleActionDto) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *MediumRuleActionDto) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.

### HasIsHtml

`func (o *MediumRuleActionDto) HasIsHtml() bool`

HasIsHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


