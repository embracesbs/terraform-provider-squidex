# DiscourseRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The url to the discourse server. | 
**ApiKey** | **string** | The api key to authenticate to your discourse server. | 
**ApiUsername** | **string** | The api username to authenticate to your discourse server. | 
**Text** | **string** | The text as markdown. | 
**Title** | Pointer to **NullableString** | The optional title when creating new topics. | [optional] 
**Topic** | Pointer to **NullableInt32** | The optional topic id. | [optional] 
**Category** | Pointer to **NullableInt32** | The optional category id. | [optional] 

## Methods

### NewDiscourseRuleActionDto

`func NewDiscourseRuleActionDto(url string, apiKey string, apiUsername string, text string, ) *DiscourseRuleActionDto`

NewDiscourseRuleActionDto instantiates a new DiscourseRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscourseRuleActionDtoWithDefaults

`func NewDiscourseRuleActionDtoWithDefaults() *DiscourseRuleActionDto`

NewDiscourseRuleActionDtoWithDefaults instantiates a new DiscourseRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DiscourseRuleActionDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DiscourseRuleActionDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DiscourseRuleActionDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetApiKey

`func (o *DiscourseRuleActionDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DiscourseRuleActionDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DiscourseRuleActionDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiUsername

`func (o *DiscourseRuleActionDto) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *DiscourseRuleActionDto) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *DiscourseRuleActionDto) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.


### GetText

`func (o *DiscourseRuleActionDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DiscourseRuleActionDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DiscourseRuleActionDto) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *DiscourseRuleActionDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DiscourseRuleActionDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DiscourseRuleActionDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DiscourseRuleActionDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DiscourseRuleActionDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DiscourseRuleActionDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTopic

`func (o *DiscourseRuleActionDto) GetTopic() int32`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *DiscourseRuleActionDto) GetTopicOk() (*int32, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *DiscourseRuleActionDto) SetTopic(v int32)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *DiscourseRuleActionDto) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *DiscourseRuleActionDto) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *DiscourseRuleActionDto) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCategory

`func (o *DiscourseRuleActionDto) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DiscourseRuleActionDto) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DiscourseRuleActionDto) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DiscourseRuleActionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *DiscourseRuleActionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *DiscourseRuleActionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


