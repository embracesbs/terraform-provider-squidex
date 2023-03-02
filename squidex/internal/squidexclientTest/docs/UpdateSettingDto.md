# UpdateSettingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | The value for the setting. | [optional] 

## Methods

### NewUpdateSettingDto

`func NewUpdateSettingDto() *UpdateSettingDto`

NewUpdateSettingDto instantiates a new UpdateSettingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingDtoWithDefaults

`func NewUpdateSettingDtoWithDefaults() *UpdateSettingDto`

NewUpdateSettingDtoWithDefaults instantiates a new UpdateSettingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateSettingDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateSettingDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateSettingDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateSettingDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


