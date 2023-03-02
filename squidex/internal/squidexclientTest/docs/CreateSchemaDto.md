# CreateSchemaDto

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
**Name** | **string** | The name of the schema. | 
**Type** | Pointer to [**SchemaType**](SchemaType.md) |  | [optional] 
**IsSingleton** | Pointer to **bool** | Set to true to allow a single content item only. | [optional] 

## Methods

### NewCreateSchemaDto

`func NewCreateSchemaDto(name string, ) *CreateSchemaDto`

NewCreateSchemaDto instantiates a new CreateSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSchemaDtoWithDefaults

`func NewCreateSchemaDtoWithDefaults() *CreateSchemaDto`

NewCreateSchemaDtoWithDefaults instantiates a new CreateSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *CreateSchemaDto) GetProperties() SchemaPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateSchemaDto) GetPropertiesOk() (*SchemaPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateSchemaDto) SetProperties(v SchemaPropertiesDto)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateSchemaDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScripts

`func (o *CreateSchemaDto) GetScripts() SchemaScriptsDto`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *CreateSchemaDto) GetScriptsOk() (*SchemaScriptsDto, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *CreateSchemaDto) SetScripts(v SchemaScriptsDto)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *CreateSchemaDto) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetFieldsInReferences

`func (o *CreateSchemaDto) GetFieldsInReferences() []string`

GetFieldsInReferences returns the FieldsInReferences field if non-nil, zero value otherwise.

### GetFieldsInReferencesOk

`func (o *CreateSchemaDto) GetFieldsInReferencesOk() (*[]string, bool)`

GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInReferences

`func (o *CreateSchemaDto) SetFieldsInReferences(v []string)`

SetFieldsInReferences sets FieldsInReferences field to given value.

### HasFieldsInReferences

`func (o *CreateSchemaDto) HasFieldsInReferences() bool`

HasFieldsInReferences returns a boolean if a field has been set.

### GetFieldsInLists

`func (o *CreateSchemaDto) GetFieldsInLists() []string`

GetFieldsInLists returns the FieldsInLists field if non-nil, zero value otherwise.

### GetFieldsInListsOk

`func (o *CreateSchemaDto) GetFieldsInListsOk() (*[]string, bool)`

GetFieldsInListsOk returns a tuple with the FieldsInLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInLists

`func (o *CreateSchemaDto) SetFieldsInLists(v []string)`

SetFieldsInLists sets FieldsInLists field to given value.

### HasFieldsInLists

`func (o *CreateSchemaDto) HasFieldsInLists() bool`

HasFieldsInLists returns a boolean if a field has been set.

### GetFields

`func (o *CreateSchemaDto) GetFields() []UpsertSchemaFieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateSchemaDto) GetFieldsOk() (*[]UpsertSchemaFieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateSchemaDto) SetFields(v []UpsertSchemaFieldDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateSchemaDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *CreateSchemaDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *CreateSchemaDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPreviewUrls

`func (o *CreateSchemaDto) GetPreviewUrls() map[string]string`

GetPreviewUrls returns the PreviewUrls field if non-nil, zero value otherwise.

### GetPreviewUrlsOk

`func (o *CreateSchemaDto) GetPreviewUrlsOk() (*map[string]string, bool)`

GetPreviewUrlsOk returns a tuple with the PreviewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrls

`func (o *CreateSchemaDto) SetPreviewUrls(v map[string]string)`

SetPreviewUrls sets PreviewUrls field to given value.

### HasPreviewUrls

`func (o *CreateSchemaDto) HasPreviewUrls() bool`

HasPreviewUrls returns a boolean if a field has been set.

### SetPreviewUrlsNil

`func (o *CreateSchemaDto) SetPreviewUrlsNil(b bool)`

 SetPreviewUrlsNil sets the value for PreviewUrls to be an explicit nil

### UnsetPreviewUrls
`func (o *CreateSchemaDto) UnsetPreviewUrls()`

UnsetPreviewUrls ensures that no value is present for PreviewUrls, not even an explicit nil
### GetFieldRules

`func (o *CreateSchemaDto) GetFieldRules() []FieldRuleDto`

GetFieldRules returns the FieldRules field if non-nil, zero value otherwise.

### GetFieldRulesOk

`func (o *CreateSchemaDto) GetFieldRulesOk() (*[]FieldRuleDto, bool)`

GetFieldRulesOk returns a tuple with the FieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRules

`func (o *CreateSchemaDto) SetFieldRules(v []FieldRuleDto)`

SetFieldRules sets FieldRules field to given value.

### HasFieldRules

`func (o *CreateSchemaDto) HasFieldRules() bool`

HasFieldRules returns a boolean if a field has been set.

### SetFieldRulesNil

`func (o *CreateSchemaDto) SetFieldRulesNil(b bool)`

 SetFieldRulesNil sets the value for FieldRules to be an explicit nil

### UnsetFieldRules
`func (o *CreateSchemaDto) UnsetFieldRules()`

UnsetFieldRules ensures that no value is present for FieldRules, not even an explicit nil
### GetCategory

`func (o *CreateSchemaDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateSchemaDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateSchemaDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateSchemaDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateSchemaDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateSchemaDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsPublished

`func (o *CreateSchemaDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CreateSchemaDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CreateSchemaDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *CreateSchemaDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetName

`func (o *CreateSchemaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSchemaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSchemaDto) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateSchemaDto) GetType() SchemaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSchemaDto) GetTypeOk() (*SchemaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSchemaDto) SetType(v SchemaType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSchemaDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSingleton

`func (o *CreateSchemaDto) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *CreateSchemaDto) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *CreateSchemaDto) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *CreateSchemaDto) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


