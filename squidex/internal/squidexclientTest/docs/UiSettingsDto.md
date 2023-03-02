# UISettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCreateApps** | Pointer to **bool** | True when the user can create apps. | [optional] 

## Methods

### NewUISettingsDto

`func NewUISettingsDto() *UISettingsDto`

NewUISettingsDto instantiates a new UISettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUISettingsDtoWithDefaults

`func NewUISettingsDtoWithDefaults() *UISettingsDto`

NewUISettingsDtoWithDefaults instantiates a new UISettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCreateApps

`func (o *UISettingsDto) GetCanCreateApps() bool`

GetCanCreateApps returns the CanCreateApps field if non-nil, zero value otherwise.

### GetCanCreateAppsOk

`func (o *UISettingsDto) GetCanCreateAppsOk() (*bool, bool)`

GetCanCreateAppsOk returns a tuple with the CanCreateApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateApps

`func (o *UISettingsDto) SetCanCreateApps(v bool)`

SetCanCreateApps sets CanCreateApps field to given value.

### HasCanCreateApps

`func (o *UISettingsDto) HasCanCreateApps() bool`

HasCanCreateApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


