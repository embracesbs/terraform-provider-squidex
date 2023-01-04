# TranslationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**TranslationResultCode**](TranslationResultCode.md) |  | [optional] 
**Text** | Pointer to **NullableString** | The translated text. | [optional] 

## Methods

### NewTranslationDto

`func NewTranslationDto() *TranslationDto`

NewTranslationDto instantiates a new TranslationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationDtoWithDefaults

`func NewTranslationDtoWithDefaults() *TranslationDto`

NewTranslationDtoWithDefaults instantiates a new TranslationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TranslationDto) GetResult() TranslationResultCode`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TranslationDto) GetResultOk() (*TranslationResultCode, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TranslationDto) SetResult(v TranslationResultCode)`

SetResult sets Result field to given value.

### HasResult

`func (o *TranslationDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetText

`func (o *TranslationDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranslationDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranslationDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TranslationDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *TranslationDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *TranslationDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


