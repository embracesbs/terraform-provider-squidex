# AppSettingsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patterns** | [**[]PatternDto**](PatternDto.md) | The configured app patterns. | 
**Editors** | [**[]EditorDto**](EditorDto.md) | The configured UI editors. | 
**HideScheduler** | Pointer to **bool** | Hide the scheduler for content items. | [optional] 
**HideDateTimeModeButton** | Pointer to **bool** | Hide the datetime mode button. | [optional] 
**Version** | Pointer to **int64** | The version of the app. | [optional] 

## Methods

### NewAppSettingsDtoAllOf

`func NewAppSettingsDtoAllOf(patterns []PatternDto, editors []EditorDto, ) *AppSettingsDtoAllOf`

NewAppSettingsDtoAllOf instantiates a new AppSettingsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSettingsDtoAllOfWithDefaults

`func NewAppSettingsDtoAllOfWithDefaults() *AppSettingsDtoAllOf`

NewAppSettingsDtoAllOfWithDefaults instantiates a new AppSettingsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatterns

`func (o *AppSettingsDtoAllOf) GetPatterns() []PatternDto`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *AppSettingsDtoAllOf) GetPatternsOk() (*[]PatternDto, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *AppSettingsDtoAllOf) SetPatterns(v []PatternDto)`

SetPatterns sets Patterns field to given value.


### GetEditors

`func (o *AppSettingsDtoAllOf) GetEditors() []EditorDto`

GetEditors returns the Editors field if non-nil, zero value otherwise.

### GetEditorsOk

`func (o *AppSettingsDtoAllOf) GetEditorsOk() (*[]EditorDto, bool)`

GetEditorsOk returns a tuple with the Editors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditors

`func (o *AppSettingsDtoAllOf) SetEditors(v []EditorDto)`

SetEditors sets Editors field to given value.


### GetHideScheduler

`func (o *AppSettingsDtoAllOf) GetHideScheduler() bool`

GetHideScheduler returns the HideScheduler field if non-nil, zero value otherwise.

### GetHideSchedulerOk

`func (o *AppSettingsDtoAllOf) GetHideSchedulerOk() (*bool, bool)`

GetHideSchedulerOk returns a tuple with the HideScheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideScheduler

`func (o *AppSettingsDtoAllOf) SetHideScheduler(v bool)`

SetHideScheduler sets HideScheduler field to given value.

### HasHideScheduler

`func (o *AppSettingsDtoAllOf) HasHideScheduler() bool`

HasHideScheduler returns a boolean if a field has been set.

### GetHideDateTimeModeButton

`func (o *AppSettingsDtoAllOf) GetHideDateTimeModeButton() bool`

GetHideDateTimeModeButton returns the HideDateTimeModeButton field if non-nil, zero value otherwise.

### GetHideDateTimeModeButtonOk

`func (o *AppSettingsDtoAllOf) GetHideDateTimeModeButtonOk() (*bool, bool)`

GetHideDateTimeModeButtonOk returns a tuple with the HideDateTimeModeButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDateTimeModeButton

`func (o *AppSettingsDtoAllOf) SetHideDateTimeModeButton(v bool)`

SetHideDateTimeModeButton sets HideDateTimeModeButton field to given value.

### HasHideDateTimeModeButton

`func (o *AppSettingsDtoAllOf) HasHideDateTimeModeButton() bool`

HasHideDateTimeModeButton returns a boolean if a field has been set.

### GetVersion

`func (o *AppSettingsDtoAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppSettingsDtoAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppSettingsDtoAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppSettingsDtoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


