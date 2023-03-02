# FieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** | Optional label for the editor. | [optional] 
**Hints** | Pointer to **NullableString** | Hints to describe the field. | [optional] 
**Placeholder** | Pointer to **NullableString** | Placeholder to show when no value has been entered. | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the field is required. | [optional] 
**IsRequiredOnPublish** | Pointer to **bool** | Indicates if the field is required when publishing. | [optional] 
**IsHalfWidth** | Pointer to **bool** | Indicates if the field should be rendered with half width only. | [optional] 
**EditorUrl** | Pointer to **NullableString** | Optional url to the editor. | [optional] 
**Tags** | Pointer to **[]string** | Tags for automation processes. | [optional] 
**FieldType** | **string** |  | 

## Methods

### NewFieldPropertiesDto

`func NewFieldPropertiesDto(fieldType string, ) *FieldPropertiesDto`

NewFieldPropertiesDto instantiates a new FieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldPropertiesDtoWithDefaults

`func NewFieldPropertiesDtoWithDefaults() *FieldPropertiesDto`

NewFieldPropertiesDtoWithDefaults instantiates a new FieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FieldPropertiesDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldPropertiesDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldPropertiesDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldPropertiesDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FieldPropertiesDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FieldPropertiesDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetHints

`func (o *FieldPropertiesDto) GetHints() string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *FieldPropertiesDto) GetHintsOk() (*string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *FieldPropertiesDto) SetHints(v string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *FieldPropertiesDto) HasHints() bool`

HasHints returns a boolean if a field has been set.

### SetHintsNil

`func (o *FieldPropertiesDto) SetHintsNil(b bool)`

 SetHintsNil sets the value for Hints to be an explicit nil

### UnsetHints
`func (o *FieldPropertiesDto) UnsetHints()`

UnsetHints ensures that no value is present for Hints, not even an explicit nil
### GetPlaceholder

`func (o *FieldPropertiesDto) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *FieldPropertiesDto) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *FieldPropertiesDto) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *FieldPropertiesDto) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *FieldPropertiesDto) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *FieldPropertiesDto) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetIsRequired

`func (o *FieldPropertiesDto) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FieldPropertiesDto) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FieldPropertiesDto) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FieldPropertiesDto) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsRequiredOnPublish

`func (o *FieldPropertiesDto) GetIsRequiredOnPublish() bool`

GetIsRequiredOnPublish returns the IsRequiredOnPublish field if non-nil, zero value otherwise.

### GetIsRequiredOnPublishOk

`func (o *FieldPropertiesDto) GetIsRequiredOnPublishOk() (*bool, bool)`

GetIsRequiredOnPublishOk returns a tuple with the IsRequiredOnPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequiredOnPublish

`func (o *FieldPropertiesDto) SetIsRequiredOnPublish(v bool)`

SetIsRequiredOnPublish sets IsRequiredOnPublish field to given value.

### HasIsRequiredOnPublish

`func (o *FieldPropertiesDto) HasIsRequiredOnPublish() bool`

HasIsRequiredOnPublish returns a boolean if a field has been set.

### GetIsHalfWidth

`func (o *FieldPropertiesDto) GetIsHalfWidth() bool`

GetIsHalfWidth returns the IsHalfWidth field if non-nil, zero value otherwise.

### GetIsHalfWidthOk

`func (o *FieldPropertiesDto) GetIsHalfWidthOk() (*bool, bool)`

GetIsHalfWidthOk returns a tuple with the IsHalfWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHalfWidth

`func (o *FieldPropertiesDto) SetIsHalfWidth(v bool)`

SetIsHalfWidth sets IsHalfWidth field to given value.

### HasIsHalfWidth

`func (o *FieldPropertiesDto) HasIsHalfWidth() bool`

HasIsHalfWidth returns a boolean if a field has been set.

### GetEditorUrl

`func (o *FieldPropertiesDto) GetEditorUrl() string`

GetEditorUrl returns the EditorUrl field if non-nil, zero value otherwise.

### GetEditorUrlOk

`func (o *FieldPropertiesDto) GetEditorUrlOk() (*string, bool)`

GetEditorUrlOk returns a tuple with the EditorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUrl

`func (o *FieldPropertiesDto) SetEditorUrl(v string)`

SetEditorUrl sets EditorUrl field to given value.

### HasEditorUrl

`func (o *FieldPropertiesDto) HasEditorUrl() bool`

HasEditorUrl returns a boolean if a field has been set.

### SetEditorUrlNil

`func (o *FieldPropertiesDto) SetEditorUrlNil(b bool)`

 SetEditorUrlNil sets the value for EditorUrl to be an explicit nil

### UnsetEditorUrl
`func (o *FieldPropertiesDto) UnsetEditorUrl()`

UnsetEditorUrl ensures that no value is present for EditorUrl, not even an explicit nil
### GetTags

`func (o *FieldPropertiesDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FieldPropertiesDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FieldPropertiesDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FieldPropertiesDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FieldPropertiesDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FieldPropertiesDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetFieldType

`func (o *FieldPropertiesDto) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *FieldPropertiesDto) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *FieldPropertiesDto) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


