# UpdateAppSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patterns** | [**[]PatternDto**](PatternDto.md) | The configured app patterns. | 
**Editors** | [**[]EditorDto**](EditorDto.md) | The configured UI editors. | 
**HideScheduler** | Pointer to **bool** | Hide the scheduler for content items. | [optional] 
**HideDateTimeModeButton** | Pointer to **bool** | Hide the datetime mode button. | [optional] 

## Methods

### NewUpdateAppSettingsDto

`func NewUpdateAppSettingsDto(patterns []PatternDto, editors []EditorDto, ) *UpdateAppSettingsDto`

NewUpdateAppSettingsDto instantiates a new UpdateAppSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppSettingsDtoWithDefaults

`func NewUpdateAppSettingsDtoWithDefaults() *UpdateAppSettingsDto`

NewUpdateAppSettingsDtoWithDefaults instantiates a new UpdateAppSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatterns

`func (o *UpdateAppSettingsDto) GetPatterns() []PatternDto`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *UpdateAppSettingsDto) GetPatternsOk() (*[]PatternDto, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *UpdateAppSettingsDto) SetPatterns(v []PatternDto)`

SetPatterns sets Patterns field to given value.


### GetEditors

`func (o *UpdateAppSettingsDto) GetEditors() []EditorDto`

GetEditors returns the Editors field if non-nil, zero value otherwise.

### GetEditorsOk

`func (o *UpdateAppSettingsDto) GetEditorsOk() (*[]EditorDto, bool)`

GetEditorsOk returns a tuple with the Editors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditors

`func (o *UpdateAppSettingsDto) SetEditors(v []EditorDto)`

SetEditors sets Editors field to given value.


### GetHideScheduler

`func (o *UpdateAppSettingsDto) GetHideScheduler() bool`

GetHideScheduler returns the HideScheduler field if non-nil, zero value otherwise.

### GetHideSchedulerOk

`func (o *UpdateAppSettingsDto) GetHideSchedulerOk() (*bool, bool)`

GetHideSchedulerOk returns a tuple with the HideScheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideScheduler

`func (o *UpdateAppSettingsDto) SetHideScheduler(v bool)`

SetHideScheduler sets HideScheduler field to given value.

### HasHideScheduler

`func (o *UpdateAppSettingsDto) HasHideScheduler() bool`

HasHideScheduler returns a boolean if a field has been set.

### GetHideDateTimeModeButton

`func (o *UpdateAppSettingsDto) GetHideDateTimeModeButton() bool`

GetHideDateTimeModeButton returns the HideDateTimeModeButton field if non-nil, zero value otherwise.

### GetHideDateTimeModeButtonOk

`func (o *UpdateAppSettingsDto) GetHideDateTimeModeButtonOk() (*bool, bool)`

GetHideDateTimeModeButtonOk returns a tuple with the HideDateTimeModeButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDateTimeModeButton

`func (o *UpdateAppSettingsDto) SetHideDateTimeModeButton(v bool)`

SetHideDateTimeModeButton sets HideDateTimeModeButton field to given value.

### HasHideDateTimeModeButton

`func (o *UpdateAppSettingsDto) HasHideDateTimeModeButton() bool`

HasHideDateTimeModeButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


