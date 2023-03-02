# TranslateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The text to translate. | 
**TargetLanguage** | **string** | The target language. | 
**SourceLanguage** | Pointer to **string** | The optional source language. | [optional] 

## Methods

### NewTranslateDto

`func NewTranslateDto(text string, targetLanguage string, ) *TranslateDto`

NewTranslateDto instantiates a new TranslateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslateDtoWithDefaults

`func NewTranslateDtoWithDefaults() *TranslateDto`

NewTranslateDtoWithDefaults instantiates a new TranslateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TranslateDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranslateDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranslateDto) SetText(v string)`

SetText sets Text field to given value.


### GetTargetLanguage

`func (o *TranslateDto) GetTargetLanguage() string`

GetTargetLanguage returns the TargetLanguage field if non-nil, zero value otherwise.

### GetTargetLanguageOk

`func (o *TranslateDto) GetTargetLanguageOk() (*string, bool)`

GetTargetLanguageOk returns a tuple with the TargetLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguage

`func (o *TranslateDto) SetTargetLanguage(v string)`

SetTargetLanguage sets TargetLanguage field to given value.


### GetSourceLanguage

`func (o *TranslateDto) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *TranslateDto) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *TranslateDto) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *TranslateDto) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


