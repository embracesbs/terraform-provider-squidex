# UpdateSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **NullableString** | Hints to describe the schema. | [optional] 
**ContentsSidebarUrl** | Pointer to **NullableString** | The url to a the sidebar plugin for content lists. | [optional] 
**ContentSidebarUrl** | Pointer to **NullableString** | The url to a the sidebar plugin for content items. | [optional] 
**ContentEditorUrl** | Pointer to **NullableString** | The url to the editor plugin. | [optional] 
**ValidateOnPublish** | Pointer to **bool** | True to validate the content items on publish. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 

## Methods

### NewUpdateSchemaDto

`func NewUpdateSchemaDto() *UpdateSchemaDto`

NewUpdateSchemaDto instantiates a new UpdateSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSchemaDtoWithDefaults

`func NewUpdateSchemaDtoWithDefaults() *UpdateSchemaDto`

NewUpdateSchemaDtoWithDefaults instantiates a new UpdateSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateSchemaDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateSchemaDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateSchemaDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateSchemaDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *UpdateSchemaDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *UpdateSchemaDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetHints

`func (o *UpdateSchemaDto) GetHints() string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *UpdateSchemaDto) GetHintsOk() (*string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *UpdateSchemaDto) SetHints(v string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *UpdateSchemaDto) HasHints() bool`

HasHints returns a boolean if a field has been set.

### SetHintsNil

`func (o *UpdateSchemaDto) SetHintsNil(b bool)`

 SetHintsNil sets the value for Hints to be an explicit nil

### UnsetHints
`func (o *UpdateSchemaDto) UnsetHints()`

UnsetHints ensures that no value is present for Hints, not even an explicit nil
### GetContentsSidebarUrl

`func (o *UpdateSchemaDto) GetContentsSidebarUrl() string`

GetContentsSidebarUrl returns the ContentsSidebarUrl field if non-nil, zero value otherwise.

### GetContentsSidebarUrlOk

`func (o *UpdateSchemaDto) GetContentsSidebarUrlOk() (*string, bool)`

GetContentsSidebarUrlOk returns a tuple with the ContentsSidebarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsSidebarUrl

`func (o *UpdateSchemaDto) SetContentsSidebarUrl(v string)`

SetContentsSidebarUrl sets ContentsSidebarUrl field to given value.

### HasContentsSidebarUrl

`func (o *UpdateSchemaDto) HasContentsSidebarUrl() bool`

HasContentsSidebarUrl returns a boolean if a field has been set.

### SetContentsSidebarUrlNil

`func (o *UpdateSchemaDto) SetContentsSidebarUrlNil(b bool)`

 SetContentsSidebarUrlNil sets the value for ContentsSidebarUrl to be an explicit nil

### UnsetContentsSidebarUrl
`func (o *UpdateSchemaDto) UnsetContentsSidebarUrl()`

UnsetContentsSidebarUrl ensures that no value is present for ContentsSidebarUrl, not even an explicit nil
### GetContentSidebarUrl

`func (o *UpdateSchemaDto) GetContentSidebarUrl() string`

GetContentSidebarUrl returns the ContentSidebarUrl field if non-nil, zero value otherwise.

### GetContentSidebarUrlOk

`func (o *UpdateSchemaDto) GetContentSidebarUrlOk() (*string, bool)`

GetContentSidebarUrlOk returns a tuple with the ContentSidebarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSidebarUrl

`func (o *UpdateSchemaDto) SetContentSidebarUrl(v string)`

SetContentSidebarUrl sets ContentSidebarUrl field to given value.

### HasContentSidebarUrl

`func (o *UpdateSchemaDto) HasContentSidebarUrl() bool`

HasContentSidebarUrl returns a boolean if a field has been set.

### SetContentSidebarUrlNil

`func (o *UpdateSchemaDto) SetContentSidebarUrlNil(b bool)`

 SetContentSidebarUrlNil sets the value for ContentSidebarUrl to be an explicit nil

### UnsetContentSidebarUrl
`func (o *UpdateSchemaDto) UnsetContentSidebarUrl()`

UnsetContentSidebarUrl ensures that no value is present for ContentSidebarUrl, not even an explicit nil
### GetContentEditorUrl

`func (o *UpdateSchemaDto) GetContentEditorUrl() string`

GetContentEditorUrl returns the ContentEditorUrl field if non-nil, zero value otherwise.

### GetContentEditorUrlOk

`func (o *UpdateSchemaDto) GetContentEditorUrlOk() (*string, bool)`

GetContentEditorUrlOk returns a tuple with the ContentEditorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEditorUrl

`func (o *UpdateSchemaDto) SetContentEditorUrl(v string)`

SetContentEditorUrl sets ContentEditorUrl field to given value.

### HasContentEditorUrl

`func (o *UpdateSchemaDto) HasContentEditorUrl() bool`

HasContentEditorUrl returns a boolean if a field has been set.

### SetContentEditorUrlNil

`func (o *UpdateSchemaDto) SetContentEditorUrlNil(b bool)`

 SetContentEditorUrlNil sets the value for ContentEditorUrl to be an explicit nil

### UnsetContentEditorUrl
`func (o *UpdateSchemaDto) UnsetContentEditorUrl()`

UnsetContentEditorUrl ensures that no value is present for ContentEditorUrl, not even an explicit nil
### GetValidateOnPublish

`func (o *UpdateSchemaDto) GetValidateOnPublish() bool`

GetValidateOnPublish returns the ValidateOnPublish field if non-nil, zero value otherwise.

### GetValidateOnPublishOk

`func (o *UpdateSchemaDto) GetValidateOnPublishOk() (*bool, bool)`

GetValidateOnPublishOk returns a tuple with the ValidateOnPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnPublish

`func (o *UpdateSchemaDto) SetValidateOnPublish(v bool)`

SetValidateOnPublish sets ValidateOnPublish field to given value.

### HasValidateOnPublish

`func (o *UpdateSchemaDto) HasValidateOnPublish() bool`

HasValidateOnPublish returns a boolean if a field has been set.

### GetTags

`func (o *UpdateSchemaDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSchemaDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSchemaDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSchemaDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateSchemaDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateSchemaDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


