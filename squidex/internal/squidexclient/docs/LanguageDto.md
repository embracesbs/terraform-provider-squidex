# LanguageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso2Code** | **string** | The iso code of the language. | 
**EnglishName** | **string** | The english name of the language. | 
**NativeName** | **string** | The native name of the language. | 

## Methods

### NewLanguageDto

`func NewLanguageDto(iso2Code string, englishName string, nativeName string, ) *LanguageDto`

NewLanguageDto instantiates a new LanguageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageDtoWithDefaults

`func NewLanguageDtoWithDefaults() *LanguageDto`

NewLanguageDtoWithDefaults instantiates a new LanguageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso2Code

`func (o *LanguageDto) GetIso2Code() string`

GetIso2Code returns the Iso2Code field if non-nil, zero value otherwise.

### GetIso2CodeOk

`func (o *LanguageDto) GetIso2CodeOk() (*string, bool)`

GetIso2CodeOk returns a tuple with the Iso2Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2Code

`func (o *LanguageDto) SetIso2Code(v string)`

SetIso2Code sets Iso2Code field to given value.


### GetEnglishName

`func (o *LanguageDto) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *LanguageDto) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *LanguageDto) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.


### GetNativeName

`func (o *LanguageDto) GetNativeName() string`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *LanguageDto) GetNativeNameOk() (*string, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *LanguageDto) SetNativeName(v string)`

SetNativeName sets NativeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


