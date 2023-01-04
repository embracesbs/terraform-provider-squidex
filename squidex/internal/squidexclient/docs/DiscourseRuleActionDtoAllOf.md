# DiscourseRuleActionDtoAllOf

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

### NewDiscourseRuleActionDtoAllOf

`func NewDiscourseRuleActionDtoAllOf(url string, apiKey string, apiUsername string, text string, ) *DiscourseRuleActionDtoAllOf`

NewDiscourseRuleActionDtoAllOf instantiates a new DiscourseRuleActionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscourseRuleActionDtoAllOfWithDefaults

`func NewDiscourseRuleActionDtoAllOfWithDefaults() *DiscourseRuleActionDtoAllOf`

NewDiscourseRuleActionDtoAllOfWithDefaults instantiates a new DiscourseRuleActionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DiscourseRuleActionDtoAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DiscourseRuleActionDtoAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DiscourseRuleActionDtoAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetApiKey

`func (o *DiscourseRuleActionDtoAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DiscourseRuleActionDtoAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DiscourseRuleActionDtoAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiUsername

`func (o *DiscourseRuleActionDtoAllOf) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *DiscourseRuleActionDtoAllOf) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *DiscourseRuleActionDtoAllOf) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.


### GetText

`func (o *DiscourseRuleActionDtoAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DiscourseRuleActionDtoAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DiscourseRuleActionDtoAllOf) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *DiscourseRuleActionDtoAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DiscourseRuleActionDtoAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DiscourseRuleActionDtoAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DiscourseRuleActionDtoAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DiscourseRuleActionDtoAllOf) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DiscourseRuleActionDtoAllOf) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTopic

`func (o *DiscourseRuleActionDtoAllOf) GetTopic() int32`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *DiscourseRuleActionDtoAllOf) GetTopicOk() (*int32, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *DiscourseRuleActionDtoAllOf) SetTopic(v int32)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *DiscourseRuleActionDtoAllOf) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *DiscourseRuleActionDtoAllOf) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *DiscourseRuleActionDtoAllOf) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCategory

`func (o *DiscourseRuleActionDtoAllOf) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DiscourseRuleActionDtoAllOf) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DiscourseRuleActionDtoAllOf) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DiscourseRuleActionDtoAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *DiscourseRuleActionDtoAllOf) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *DiscourseRuleActionDtoAllOf) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


