# UpsertSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**SchemaPropertiesDto**](SchemaPropertiesDto.md) |  | [optional] 
**Scripts** | Pointer to [**SchemaScriptsDto**](SchemaScriptsDto.md) |  | [optional] 
**FieldsInReferences** | Pointer to **[]string** |  | [optional] 
**FieldsInLists** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to [**[]UpsertSchemaFieldDto**](UpsertSchemaFieldDto.md) | Optional fields. | [optional] 
**PreviewUrls** | Pointer to **map[string]string** | The optional preview urls. | [optional] 
**FieldRules** | Pointer to [**[]FieldRuleDto**](FieldRuleDto.md) | The optional field Rules. | [optional] 
**Category** | Pointer to **NullableString** | The category. | [optional] 
**IsPublished** | Pointer to **bool** | Set it to true to autopublish the schema. | [optional] 

## Methods

### NewUpsertSchemaDto

`func NewUpsertSchemaDto() *UpsertSchemaDto`

NewUpsertSchemaDto instantiates a new UpsertSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertSchemaDtoWithDefaults

`func NewUpsertSchemaDtoWithDefaults() *UpsertSchemaDto`

NewUpsertSchemaDtoWithDefaults instantiates a new UpsertSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *UpsertSchemaDto) GetProperties() SchemaPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpsertSchemaDto) GetPropertiesOk() (*SchemaPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpsertSchemaDto) SetProperties(v SchemaPropertiesDto)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UpsertSchemaDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScripts

`func (o *UpsertSchemaDto) GetScripts() SchemaScriptsDto`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *UpsertSchemaDto) GetScriptsOk() (*SchemaScriptsDto, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *UpsertSchemaDto) SetScripts(v SchemaScriptsDto)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *UpsertSchemaDto) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetFieldsInReferences

`func (o *UpsertSchemaDto) GetFieldsInReferences() []string`

GetFieldsInReferences returns the FieldsInReferences field if non-nil, zero value otherwise.

### GetFieldsInReferencesOk

`func (o *UpsertSchemaDto) GetFieldsInReferencesOk() (*[]string, bool)`

GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInReferences

`func (o *UpsertSchemaDto) SetFieldsInReferences(v []string)`

SetFieldsInReferences sets FieldsInReferences field to given value.

### HasFieldsInReferences

`func (o *UpsertSchemaDto) HasFieldsInReferences() bool`

HasFieldsInReferences returns a boolean if a field has been set.

### GetFieldsInLists

`func (o *UpsertSchemaDto) GetFieldsInLists() []string`

GetFieldsInLists returns the FieldsInLists field if non-nil, zero value otherwise.

### GetFieldsInListsOk

`func (o *UpsertSchemaDto) GetFieldsInListsOk() (*[]string, bool)`

GetFieldsInListsOk returns a tuple with the FieldsInLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInLists

`func (o *UpsertSchemaDto) SetFieldsInLists(v []string)`

SetFieldsInLists sets FieldsInLists field to given value.

### HasFieldsInLists

`func (o *UpsertSchemaDto) HasFieldsInLists() bool`

HasFieldsInLists returns a boolean if a field has been set.

### GetFields

`func (o *UpsertSchemaDto) GetFields() []UpsertSchemaFieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpsertSchemaDto) GetFieldsOk() (*[]UpsertSchemaFieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpsertSchemaDto) SetFields(v []UpsertSchemaFieldDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UpsertSchemaDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *UpsertSchemaDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *UpsertSchemaDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPreviewUrls

`func (o *UpsertSchemaDto) GetPreviewUrls() map[string]string`

GetPreviewUrls returns the PreviewUrls field if non-nil, zero value otherwise.

### GetPreviewUrlsOk

`func (o *UpsertSchemaDto) GetPreviewUrlsOk() (*map[string]string, bool)`

GetPreviewUrlsOk returns a tuple with the PreviewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrls

`func (o *UpsertSchemaDto) SetPreviewUrls(v map[string]string)`

SetPreviewUrls sets PreviewUrls field to given value.

### HasPreviewUrls

`func (o *UpsertSchemaDto) HasPreviewUrls() bool`

HasPreviewUrls returns a boolean if a field has been set.

### SetPreviewUrlsNil

`func (o *UpsertSchemaDto) SetPreviewUrlsNil(b bool)`

 SetPreviewUrlsNil sets the value for PreviewUrls to be an explicit nil

### UnsetPreviewUrls
`func (o *UpsertSchemaDto) UnsetPreviewUrls()`

UnsetPreviewUrls ensures that no value is present for PreviewUrls, not even an explicit nil
### GetFieldRules

`func (o *UpsertSchemaDto) GetFieldRules() []FieldRuleDto`

GetFieldRules returns the FieldRules field if non-nil, zero value otherwise.

### GetFieldRulesOk

`func (o *UpsertSchemaDto) GetFieldRulesOk() (*[]FieldRuleDto, bool)`

GetFieldRulesOk returns a tuple with the FieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRules

`func (o *UpsertSchemaDto) SetFieldRules(v []FieldRuleDto)`

SetFieldRules sets FieldRules field to given value.

### HasFieldRules

`func (o *UpsertSchemaDto) HasFieldRules() bool`

HasFieldRules returns a boolean if a field has been set.

### SetFieldRulesNil

`func (o *UpsertSchemaDto) SetFieldRulesNil(b bool)`

 SetFieldRulesNil sets the value for FieldRules to be an explicit nil

### UnsetFieldRules
`func (o *UpsertSchemaDto) UnsetFieldRules()`

UnsetFieldRules ensures that no value is present for FieldRules, not even an explicit nil
### GetCategory

`func (o *UpsertSchemaDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpsertSchemaDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpsertSchemaDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpsertSchemaDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpsertSchemaDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpsertSchemaDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsPublished

`func (o *UpsertSchemaDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *UpsertSchemaDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *UpsertSchemaDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *UpsertSchemaDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


