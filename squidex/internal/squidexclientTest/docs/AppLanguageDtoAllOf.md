# AppLanguageDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso2Code** | **string** | The iso code of the language. | 
**EnglishName** | **string** | The english name of the language. | 
**Fallback** | **[]string** | The fallback languages. | 
**IsMaster** | Pointer to **bool** | Indicates if the language is the master language. | [optional] 
**IsOptional** | Pointer to **bool** | Indicates if the language is optional. | [optional] 

## Methods

### NewAppLanguageDtoAllOf

`func NewAppLanguageDtoAllOf(iso2Code string, englishName string, fallback []string, ) *AppLanguageDtoAllOf`

NewAppLanguageDtoAllOf instantiates a new AppLanguageDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLanguageDtoAllOfWithDefaults

`func NewAppLanguageDtoAllOfWithDefaults() *AppLanguageDtoAllOf`

NewAppLanguageDtoAllOfWithDefaults instantiates a new AppLanguageDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso2Code

`func (o *AppLanguageDtoAllOf) GetIso2Code() string`

GetIso2Code returns the Iso2Code field if non-nil, zero value otherwise.

### GetIso2CodeOk

`func (o *AppLanguageDtoAllOf) GetIso2CodeOk() (*string, bool)`

GetIso2CodeOk returns a tuple with the Iso2Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2Code

`func (o *AppLanguageDtoAllOf) SetIso2Code(v string)`

SetIso2Code sets Iso2Code field to given value.


### GetEnglishName

`func (o *AppLanguageDtoAllOf) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *AppLanguageDtoAllOf) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *AppLanguageDtoAllOf) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.


### GetFallback

`func (o *AppLanguageDtoAllOf) GetFallback() []string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *AppLanguageDtoAllOf) GetFallbackOk() (*[]string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *AppLanguageDtoAllOf) SetFallback(v []string)`

SetFallback sets Fallback field to given value.


### GetIsMaster

`func (o *AppLanguageDtoAllOf) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *AppLanguageDtoAllOf) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *AppLanguageDtoAllOf) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *AppLanguageDtoAllOf) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetIsOptional

`func (o *AppLanguageDtoAllOf) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *AppLanguageDtoAllOf) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *AppLanguageDtoAllOf) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *AppLanguageDtoAllOf) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


