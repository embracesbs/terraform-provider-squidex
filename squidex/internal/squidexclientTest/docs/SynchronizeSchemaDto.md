# SynchronizeSchemaDto

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
**NoFieldDeletion** | Pointer to **bool** | True, when fields should not be deleted. | [optional] 
**NoFieldRecreation** | Pointer to **bool** | True, when fields with different types should not be recreated. | [optional] 

## Methods

### NewSynchronizeSchemaDto

`func NewSynchronizeSchemaDto() *SynchronizeSchemaDto`

NewSynchronizeSchemaDto instantiates a new SynchronizeSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynchronizeSchemaDtoWithDefaults

`func NewSynchronizeSchemaDtoWithDefaults() *SynchronizeSchemaDto`

NewSynchronizeSchemaDtoWithDefaults instantiates a new SynchronizeSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SynchronizeSchemaDto) GetProperties() SchemaPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SynchronizeSchemaDto) GetPropertiesOk() (*SchemaPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SynchronizeSchemaDto) SetProperties(v SchemaPropertiesDto)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SynchronizeSchemaDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScripts

`func (o *SynchronizeSchemaDto) GetScripts() SchemaScriptsDto`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *SynchronizeSchemaDto) GetScriptsOk() (*SchemaScriptsDto, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *SynchronizeSchemaDto) SetScripts(v SchemaScriptsDto)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *SynchronizeSchemaDto) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetFieldsInReferences

`func (o *SynchronizeSchemaDto) GetFieldsInReferences() []string`

GetFieldsInReferences returns the FieldsInReferences field if non-nil, zero value otherwise.

### GetFieldsInReferencesOk

`func (o *SynchronizeSchemaDto) GetFieldsInReferencesOk() (*[]string, bool)`

GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInReferences

`func (o *SynchronizeSchemaDto) SetFieldsInReferences(v []string)`

SetFieldsInReferences sets FieldsInReferences field to given value.

### HasFieldsInReferences

`func (o *SynchronizeSchemaDto) HasFieldsInReferences() bool`

HasFieldsInReferences returns a boolean if a field has been set.

### GetFieldsInLists

`func (o *SynchronizeSchemaDto) GetFieldsInLists() []string`

GetFieldsInLists returns the FieldsInLists field if non-nil, zero value otherwise.

### GetFieldsInListsOk

`func (o *SynchronizeSchemaDto) GetFieldsInListsOk() (*[]string, bool)`

GetFieldsInListsOk returns a tuple with the FieldsInLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInLists

`func (o *SynchronizeSchemaDto) SetFieldsInLists(v []string)`

SetFieldsInLists sets FieldsInLists field to given value.

### HasFieldsInLists

`func (o *SynchronizeSchemaDto) HasFieldsInLists() bool`

HasFieldsInLists returns a boolean if a field has been set.

### GetFields

`func (o *SynchronizeSchemaDto) GetFields() []UpsertSchemaFieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SynchronizeSchemaDto) GetFieldsOk() (*[]UpsertSchemaFieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SynchronizeSchemaDto) SetFields(v []UpsertSchemaFieldDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SynchronizeSchemaDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *SynchronizeSchemaDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *SynchronizeSchemaDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPreviewUrls

`func (o *SynchronizeSchemaDto) GetPreviewUrls() map[string]string`

GetPreviewUrls returns the PreviewUrls field if non-nil, zero value otherwise.

### GetPreviewUrlsOk

`func (o *SynchronizeSchemaDto) GetPreviewUrlsOk() (*map[string]string, bool)`

GetPreviewUrlsOk returns a tuple with the PreviewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrls

`func (o *SynchronizeSchemaDto) SetPreviewUrls(v map[string]string)`

SetPreviewUrls sets PreviewUrls field to given value.

### HasPreviewUrls

`func (o *SynchronizeSchemaDto) HasPreviewUrls() bool`

HasPreviewUrls returns a boolean if a field has been set.

### SetPreviewUrlsNil

`func (o *SynchronizeSchemaDto) SetPreviewUrlsNil(b bool)`

 SetPreviewUrlsNil sets the value for PreviewUrls to be an explicit nil

### UnsetPreviewUrls
`func (o *SynchronizeSchemaDto) UnsetPreviewUrls()`

UnsetPreviewUrls ensures that no value is present for PreviewUrls, not even an explicit nil
### GetFieldRules

`func (o *SynchronizeSchemaDto) GetFieldRules() []FieldRuleDto`

GetFieldRules returns the FieldRules field if non-nil, zero value otherwise.

### GetFieldRulesOk

`func (o *SynchronizeSchemaDto) GetFieldRulesOk() (*[]FieldRuleDto, bool)`

GetFieldRulesOk returns a tuple with the FieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRules

`func (o *SynchronizeSchemaDto) SetFieldRules(v []FieldRuleDto)`

SetFieldRules sets FieldRules field to given value.

### HasFieldRules

`func (o *SynchronizeSchemaDto) HasFieldRules() bool`

HasFieldRules returns a boolean if a field has been set.

### SetFieldRulesNil

`func (o *SynchronizeSchemaDto) SetFieldRulesNil(b bool)`

 SetFieldRulesNil sets the value for FieldRules to be an explicit nil

### UnsetFieldRules
`func (o *SynchronizeSchemaDto) UnsetFieldRules()`

UnsetFieldRules ensures that no value is present for FieldRules, not even an explicit nil
### GetCategory

`func (o *SynchronizeSchemaDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SynchronizeSchemaDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SynchronizeSchemaDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SynchronizeSchemaDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *SynchronizeSchemaDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *SynchronizeSchemaDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsPublished

`func (o *SynchronizeSchemaDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *SynchronizeSchemaDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *SynchronizeSchemaDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *SynchronizeSchemaDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetNoFieldDeletion

`func (o *SynchronizeSchemaDto) GetNoFieldDeletion() bool`

GetNoFieldDeletion returns the NoFieldDeletion field if non-nil, zero value otherwise.

### GetNoFieldDeletionOk

`func (o *SynchronizeSchemaDto) GetNoFieldDeletionOk() (*bool, bool)`

GetNoFieldDeletionOk returns a tuple with the NoFieldDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFieldDeletion

`func (o *SynchronizeSchemaDto) SetNoFieldDeletion(v bool)`

SetNoFieldDeletion sets NoFieldDeletion field to given value.

### HasNoFieldDeletion

`func (o *SynchronizeSchemaDto) HasNoFieldDeletion() bool`

HasNoFieldDeletion returns a boolean if a field has been set.

### GetNoFieldRecreation

`func (o *SynchronizeSchemaDto) GetNoFieldRecreation() bool`

GetNoFieldRecreation returns the NoFieldRecreation field if non-nil, zero value otherwise.

### GetNoFieldRecreationOk

`func (o *SynchronizeSchemaDto) GetNoFieldRecreationOk() (*bool, bool)`

GetNoFieldRecreationOk returns a tuple with the NoFieldRecreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFieldRecreation

`func (o *SynchronizeSchemaDto) SetNoFieldRecreation(v bool)`

SetNoFieldRecreation sets NoFieldRecreation field to given value.

### HasNoFieldRecreation

`func (o *SynchronizeSchemaDto) HasNoFieldRecreation() bool`

HasNoFieldRecreation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


