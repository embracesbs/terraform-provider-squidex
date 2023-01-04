# UpdateLanguageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMaster** | Pointer to **NullableBool** | Set the value to true to make the language the master. | [optional] 
**IsOptional** | Pointer to **bool** | Set the value to true to make the language optional. | [optional] 
**Fallback** | Pointer to **[]string** | Optional fallback languages. | [optional] 

## Methods

### NewUpdateLanguageDto

`func NewUpdateLanguageDto() *UpdateLanguageDto`

NewUpdateLanguageDto instantiates a new UpdateLanguageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLanguageDtoWithDefaults

`func NewUpdateLanguageDtoWithDefaults() *UpdateLanguageDto`

NewUpdateLanguageDtoWithDefaults instantiates a new UpdateLanguageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMaster

`func (o *UpdateLanguageDto) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *UpdateLanguageDto) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *UpdateLanguageDto) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *UpdateLanguageDto) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### SetIsMasterNil

`func (o *UpdateLanguageDto) SetIsMasterNil(b bool)`

 SetIsMasterNil sets the value for IsMaster to be an explicit nil

### UnsetIsMaster
`func (o *UpdateLanguageDto) UnsetIsMaster()`

UnsetIsMaster ensures that no value is present for IsMaster, not even an explicit nil
### GetIsOptional

`func (o *UpdateLanguageDto) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *UpdateLanguageDto) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *UpdateLanguageDto) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *UpdateLanguageDto) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetFallback

`func (o *UpdateLanguageDto) GetFallback() []string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *UpdateLanguageDto) GetFallbackOk() (*[]string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *UpdateLanguageDto) SetFallback(v []string)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *UpdateLanguageDto) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### SetFallbackNil

`func (o *UpdateLanguageDto) SetFallbackNil(b bool)`

 SetFallbackNil sets the value for Fallback to be an explicit nil

### UnsetFallback
`func (o *UpdateLanguageDto) UnsetFallback()`

UnsetFallback ensures that no value is present for Fallback, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


