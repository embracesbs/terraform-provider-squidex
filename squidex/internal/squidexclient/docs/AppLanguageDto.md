# AppLanguageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Iso2Code** | **string** | The iso code of the language. | 
**EnglishName** | **string** | The english name of the language. | 
**Fallback** | **[]string** | The fallback languages. | 
**IsMaster** | Pointer to **bool** | Indicates if the language is the master language. | [optional] 
**IsOptional** | Pointer to **bool** | Indicates if the language is optional. | [optional] 

## Methods

### NewAppLanguageDto

`func NewAppLanguageDto(links map[string]ResourceLink, iso2Code string, englishName string, fallback []string, ) *AppLanguageDto`

NewAppLanguageDto instantiates a new AppLanguageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLanguageDtoWithDefaults

`func NewAppLanguageDtoWithDefaults() *AppLanguageDto`

NewAppLanguageDtoWithDefaults instantiates a new AppLanguageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppLanguageDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppLanguageDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppLanguageDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetIso2Code

`func (o *AppLanguageDto) GetIso2Code() string`

GetIso2Code returns the Iso2Code field if non-nil, zero value otherwise.

### GetIso2CodeOk

`func (o *AppLanguageDto) GetIso2CodeOk() (*string, bool)`

GetIso2CodeOk returns a tuple with the Iso2Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2Code

`func (o *AppLanguageDto) SetIso2Code(v string)`

SetIso2Code sets Iso2Code field to given value.


### GetEnglishName

`func (o *AppLanguageDto) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *AppLanguageDto) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *AppLanguageDto) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.


### GetFallback

`func (o *AppLanguageDto) GetFallback() []string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *AppLanguageDto) GetFallbackOk() (*[]string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *AppLanguageDto) SetFallback(v []string)`

SetFallback sets Fallback field to given value.


### GetIsMaster

`func (o *AppLanguageDto) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *AppLanguageDto) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *AppLanguageDto) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *AppLanguageDto) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetIsOptional

`func (o *AppLanguageDto) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *AppLanguageDto) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *AppLanguageDto) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *AppLanguageDto) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


