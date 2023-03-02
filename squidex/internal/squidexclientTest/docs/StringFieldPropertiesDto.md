# StringFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to **map[string]string** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** | The default value for the field value. | [optional] 
**Pattern** | Pointer to **NullableString** | The pattern to enforce a specific format for the field value. | [optional] 
**PatternMessage** | Pointer to **NullableString** | The validation message for the pattern. | [optional] 
**FolderId** | Pointer to **NullableString** | The initial id to the folder when the control supports file uploads. | [optional] 
**MinLength** | Pointer to **NullableInt32** | The minimum allowed length for the field value. | [optional] 
**MaxLength** | Pointer to **NullableInt32** | The maximum allowed length for the field value. | [optional] 
**MinCharacters** | Pointer to **NullableInt32** | The minimum allowed of normal characters for the field value. | [optional] 
**MaxCharacters** | Pointer to **NullableInt32** | The maximum allowed of normal characters for the field value. | [optional] 
**MinWords** | Pointer to **NullableInt32** | The minimum allowed number of words for the field value. | [optional] 
**MaxWords** | Pointer to **NullableInt32** | The maximum allowed number of words for the field value. | [optional] 
**AllowedValues** | Pointer to **[]string** | The allowed values for the field value. | [optional] 
**SchemaIds** | Pointer to **[]string** | The allowed schema ids that can be embedded. | [optional] 
**IsUnique** | Pointer to **bool** | Indicates if the field value must be unique. Ignored for nested fields and localized fields. | [optional] 
**IsEmbeddable** | Pointer to **bool** | Indicates that other content items or references are embedded. | [optional] 
**InlineEditable** | Pointer to **bool** | Indicates that the inline editor is enabled for this field. | [optional] 
**CreateEnum** | Pointer to **bool** | Indicates whether GraphQL Enum should be created. | [optional] 
**ContentType** | Pointer to [**StringContentType**](StringContentType.md) |  | [optional] 
**Editor** | Pointer to [**StringFieldEditor**](StringFieldEditor.md) |  | [optional] 

## Methods

### NewStringFieldPropertiesDto

`func NewStringFieldPropertiesDto() *StringFieldPropertiesDto`

NewStringFieldPropertiesDto instantiates a new StringFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringFieldPropertiesDtoWithDefaults

`func NewStringFieldPropertiesDtoWithDefaults() *StringFieldPropertiesDto`

NewStringFieldPropertiesDtoWithDefaults instantiates a new StringFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *StringFieldPropertiesDto) GetDefaultValues() map[string]string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *StringFieldPropertiesDto) GetDefaultValuesOk() (*map[string]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *StringFieldPropertiesDto) SetDefaultValues(v map[string]string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *StringFieldPropertiesDto) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *StringFieldPropertiesDto) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *StringFieldPropertiesDto) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *StringFieldPropertiesDto) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *StringFieldPropertiesDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *StringFieldPropertiesDto) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *StringFieldPropertiesDto) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetPattern

`func (o *StringFieldPropertiesDto) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *StringFieldPropertiesDto) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *StringFieldPropertiesDto) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *StringFieldPropertiesDto) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *StringFieldPropertiesDto) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *StringFieldPropertiesDto) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetPatternMessage

`func (o *StringFieldPropertiesDto) GetPatternMessage() string`

GetPatternMessage returns the PatternMessage field if non-nil, zero value otherwise.

### GetPatternMessageOk

`func (o *StringFieldPropertiesDto) GetPatternMessageOk() (*string, bool)`

GetPatternMessageOk returns a tuple with the PatternMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternMessage

`func (o *StringFieldPropertiesDto) SetPatternMessage(v string)`

SetPatternMessage sets PatternMessage field to given value.

### HasPatternMessage

`func (o *StringFieldPropertiesDto) HasPatternMessage() bool`

HasPatternMessage returns a boolean if a field has been set.

### SetPatternMessageNil

`func (o *StringFieldPropertiesDto) SetPatternMessageNil(b bool)`

 SetPatternMessageNil sets the value for PatternMessage to be an explicit nil

### UnsetPatternMessage
`func (o *StringFieldPropertiesDto) UnsetPatternMessage()`

UnsetPatternMessage ensures that no value is present for PatternMessage, not even an explicit nil
### GetFolderId

`func (o *StringFieldPropertiesDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *StringFieldPropertiesDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *StringFieldPropertiesDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *StringFieldPropertiesDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *StringFieldPropertiesDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *StringFieldPropertiesDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetMinLength

`func (o *StringFieldPropertiesDto) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *StringFieldPropertiesDto) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *StringFieldPropertiesDto) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *StringFieldPropertiesDto) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *StringFieldPropertiesDto) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *StringFieldPropertiesDto) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *StringFieldPropertiesDto) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *StringFieldPropertiesDto) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *StringFieldPropertiesDto) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *StringFieldPropertiesDto) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *StringFieldPropertiesDto) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *StringFieldPropertiesDto) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetMinCharacters

`func (o *StringFieldPropertiesDto) GetMinCharacters() int32`

GetMinCharacters returns the MinCharacters field if non-nil, zero value otherwise.

### GetMinCharactersOk

`func (o *StringFieldPropertiesDto) GetMinCharactersOk() (*int32, bool)`

GetMinCharactersOk returns a tuple with the MinCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCharacters

`func (o *StringFieldPropertiesDto) SetMinCharacters(v int32)`

SetMinCharacters sets MinCharacters field to given value.

### HasMinCharacters

`func (o *StringFieldPropertiesDto) HasMinCharacters() bool`

HasMinCharacters returns a boolean if a field has been set.

### SetMinCharactersNil

`func (o *StringFieldPropertiesDto) SetMinCharactersNil(b bool)`

 SetMinCharactersNil sets the value for MinCharacters to be an explicit nil

### UnsetMinCharacters
`func (o *StringFieldPropertiesDto) UnsetMinCharacters()`

UnsetMinCharacters ensures that no value is present for MinCharacters, not even an explicit nil
### GetMaxCharacters

`func (o *StringFieldPropertiesDto) GetMaxCharacters() int32`

GetMaxCharacters returns the MaxCharacters field if non-nil, zero value otherwise.

### GetMaxCharactersOk

`func (o *StringFieldPropertiesDto) GetMaxCharactersOk() (*int32, bool)`

GetMaxCharactersOk returns a tuple with the MaxCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacters

`func (o *StringFieldPropertiesDto) SetMaxCharacters(v int32)`

SetMaxCharacters sets MaxCharacters field to given value.

### HasMaxCharacters

`func (o *StringFieldPropertiesDto) HasMaxCharacters() bool`

HasMaxCharacters returns a boolean if a field has been set.

### SetMaxCharactersNil

`func (o *StringFieldPropertiesDto) SetMaxCharactersNil(b bool)`

 SetMaxCharactersNil sets the value for MaxCharacters to be an explicit nil

### UnsetMaxCharacters
`func (o *StringFieldPropertiesDto) UnsetMaxCharacters()`

UnsetMaxCharacters ensures that no value is present for MaxCharacters, not even an explicit nil
### GetMinWords

`func (o *StringFieldPropertiesDto) GetMinWords() int32`

GetMinWords returns the MinWords field if non-nil, zero value otherwise.

### GetMinWordsOk

`func (o *StringFieldPropertiesDto) GetMinWordsOk() (*int32, bool)`

GetMinWordsOk returns a tuple with the MinWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWords

`func (o *StringFieldPropertiesDto) SetMinWords(v int32)`

SetMinWords sets MinWords field to given value.

### HasMinWords

`func (o *StringFieldPropertiesDto) HasMinWords() bool`

HasMinWords returns a boolean if a field has been set.

### SetMinWordsNil

`func (o *StringFieldPropertiesDto) SetMinWordsNil(b bool)`

 SetMinWordsNil sets the value for MinWords to be an explicit nil

### UnsetMinWords
`func (o *StringFieldPropertiesDto) UnsetMinWords()`

UnsetMinWords ensures that no value is present for MinWords, not even an explicit nil
### GetMaxWords

`func (o *StringFieldPropertiesDto) GetMaxWords() int32`

GetMaxWords returns the MaxWords field if non-nil, zero value otherwise.

### GetMaxWordsOk

`func (o *StringFieldPropertiesDto) GetMaxWordsOk() (*int32, bool)`

GetMaxWordsOk returns a tuple with the MaxWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWords

`func (o *StringFieldPropertiesDto) SetMaxWords(v int32)`

SetMaxWords sets MaxWords field to given value.

### HasMaxWords

`func (o *StringFieldPropertiesDto) HasMaxWords() bool`

HasMaxWords returns a boolean if a field has been set.

### SetMaxWordsNil

`func (o *StringFieldPropertiesDto) SetMaxWordsNil(b bool)`

 SetMaxWordsNil sets the value for MaxWords to be an explicit nil

### UnsetMaxWords
`func (o *StringFieldPropertiesDto) UnsetMaxWords()`

UnsetMaxWords ensures that no value is present for MaxWords, not even an explicit nil
### GetAllowedValues

`func (o *StringFieldPropertiesDto) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *StringFieldPropertiesDto) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *StringFieldPropertiesDto) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *StringFieldPropertiesDto) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *StringFieldPropertiesDto) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *StringFieldPropertiesDto) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetSchemaIds

`func (o *StringFieldPropertiesDto) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *StringFieldPropertiesDto) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *StringFieldPropertiesDto) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *StringFieldPropertiesDto) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *StringFieldPropertiesDto) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *StringFieldPropertiesDto) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil
### GetIsUnique

`func (o *StringFieldPropertiesDto) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *StringFieldPropertiesDto) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *StringFieldPropertiesDto) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *StringFieldPropertiesDto) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsEmbeddable

`func (o *StringFieldPropertiesDto) GetIsEmbeddable() bool`

GetIsEmbeddable returns the IsEmbeddable field if non-nil, zero value otherwise.

### GetIsEmbeddableOk

`func (o *StringFieldPropertiesDto) GetIsEmbeddableOk() (*bool, bool)`

GetIsEmbeddableOk returns a tuple with the IsEmbeddable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbeddable

`func (o *StringFieldPropertiesDto) SetIsEmbeddable(v bool)`

SetIsEmbeddable sets IsEmbeddable field to given value.

### HasIsEmbeddable

`func (o *StringFieldPropertiesDto) HasIsEmbeddable() bool`

HasIsEmbeddable returns a boolean if a field has been set.

### GetInlineEditable

`func (o *StringFieldPropertiesDto) GetInlineEditable() bool`

GetInlineEditable returns the InlineEditable field if non-nil, zero value otherwise.

### GetInlineEditableOk

`func (o *StringFieldPropertiesDto) GetInlineEditableOk() (*bool, bool)`

GetInlineEditableOk returns a tuple with the InlineEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEditable

`func (o *StringFieldPropertiesDto) SetInlineEditable(v bool)`

SetInlineEditable sets InlineEditable field to given value.

### HasInlineEditable

`func (o *StringFieldPropertiesDto) HasInlineEditable() bool`

HasInlineEditable returns a boolean if a field has been set.

### GetCreateEnum

`func (o *StringFieldPropertiesDto) GetCreateEnum() bool`

GetCreateEnum returns the CreateEnum field if non-nil, zero value otherwise.

### GetCreateEnumOk

`func (o *StringFieldPropertiesDto) GetCreateEnumOk() (*bool, bool)`

GetCreateEnumOk returns a tuple with the CreateEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnum

`func (o *StringFieldPropertiesDto) SetCreateEnum(v bool)`

SetCreateEnum sets CreateEnum field to given value.

### HasCreateEnum

`func (o *StringFieldPropertiesDto) HasCreateEnum() bool`

HasCreateEnum returns a boolean if a field has been set.

### GetContentType

`func (o *StringFieldPropertiesDto) GetContentType() StringContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *StringFieldPropertiesDto) GetContentTypeOk() (*StringContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *StringFieldPropertiesDto) SetContentType(v StringContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *StringFieldPropertiesDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEditor

`func (o *StringFieldPropertiesDto) GetEditor() StringFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *StringFieldPropertiesDto) GetEditorOk() (*StringFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *StringFieldPropertiesDto) SetEditor(v StringFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *StringFieldPropertiesDto) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


