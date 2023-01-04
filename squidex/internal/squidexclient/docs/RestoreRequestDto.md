# RestoreRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the app. | [optional] 
**Url** | **string** | The url to the restore file. | 

## Methods

### NewRestoreRequestDto

`func NewRestoreRequestDto(url string, ) *RestoreRequestDto`

NewRestoreRequestDto instantiates a new RestoreRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRequestDtoWithDefaults

`func NewRestoreRequestDtoWithDefaults() *RestoreRequestDto`

NewRestoreRequestDtoWithDefaults instantiates a new RestoreRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestoreRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestoreRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RestoreRequestDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RestoreRequestDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *RestoreRequestDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RestoreRequestDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RestoreRequestDto) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


