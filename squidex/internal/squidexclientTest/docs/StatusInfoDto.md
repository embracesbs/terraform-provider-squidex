# StatusInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The name of the status. | 
**Color** | **string** | The color of the status. | 

## Methods

### NewStatusInfoDto

`func NewStatusInfoDto(status string, color string, ) *StatusInfoDto`

NewStatusInfoDto instantiates a new StatusInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusInfoDtoWithDefaults

`func NewStatusInfoDtoWithDefaults() *StatusInfoDto`

NewStatusInfoDtoWithDefaults instantiates a new StatusInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusInfoDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusInfoDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusInfoDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetColor

`func (o *StatusInfoDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *StatusInfoDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *StatusInfoDto) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


