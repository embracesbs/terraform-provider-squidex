# SchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | Pointer to **string** | The id of the schema. | [optional] 
**CreatedBy** | **string** | The user that has created the schema. | 
**LastModifiedBy** | **string** | The user that has updated the schema. | 
**Name** | **string** | The name of the schema. Unique within the app. | 
**Type** | Pointer to [**SchemaType**](SchemaType.md) |  | [optional] 
**Category** | Pointer to **NullableString** | The name of the category. | [optional] 
**Properties** | [**SchemaPropertiesDto**](SchemaPropertiesDto.md) |  | 
**IsSingleton** | Pointer to **bool** | Indicates if the schema is a singleton. | [optional] 
**IsPublished** | Pointer to **bool** | Indicates if the schema is published. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the schema has been created. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the schema has been modified last. | [optional] 
**Version** | Pointer to **int64** | The version of the schema. | [optional] 
**Scripts** | [**SchemaScriptsDto**](SchemaScriptsDto.md) |  | 
**PreviewUrls** | **map[string]string** | The preview Urls. | 
**FieldsInLists** | **[]string** |  | 
**FieldsInReferences** | **[]string** |  | 
**FieldRules** | Pointer to [**[]FieldRuleDto**](FieldRuleDto.md) | The field rules. | [optional] 
**Fields** | [**[]FieldDto**](FieldDto.md) | The list of fields. | 

## Methods

### NewSchemaDto

`func NewSchemaDto(links map[string]ResourceLink, createdBy string, lastModifiedBy string, name string, properties SchemaPropertiesDto, scripts SchemaScriptsDto, previewUrls map[string]string, fieldsInLists []string, fieldsInReferences []string, fields []FieldDto, ) *SchemaDto`

NewSchemaDto instantiates a new SchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDtoWithDefaults

`func NewSchemaDtoWithDefaults() *SchemaDto`

NewSchemaDtoWithDefaults instantiates a new SchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SchemaDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchemaDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchemaDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *SchemaDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SchemaDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SchemaDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SchemaDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModifiedBy

`func (o *SchemaDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *SchemaDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *SchemaDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetName

`func (o *SchemaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaDto) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SchemaDto) GetType() SchemaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaDto) GetTypeOk() (*SchemaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaDto) SetType(v SchemaType)`

SetType sets Type field to given value.

### HasType

`func (o *SchemaDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *SchemaDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SchemaDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SchemaDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SchemaDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *SchemaDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *SchemaDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetProperties

`func (o *SchemaDto) GetProperties() SchemaPropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaDto) GetPropertiesOk() (*SchemaPropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SchemaDto) SetProperties(v SchemaPropertiesDto)`

SetProperties sets Properties field to given value.


### GetIsSingleton

`func (o *SchemaDto) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *SchemaDto) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *SchemaDto) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *SchemaDto) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.

### GetIsPublished

`func (o *SchemaDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *SchemaDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *SchemaDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *SchemaDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCreated

`func (o *SchemaDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SchemaDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SchemaDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SchemaDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *SchemaDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SchemaDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SchemaDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SchemaDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetScripts

`func (o *SchemaDto) GetScripts() SchemaScriptsDto`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *SchemaDto) GetScriptsOk() (*SchemaScriptsDto, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *SchemaDto) SetScripts(v SchemaScriptsDto)`

SetScripts sets Scripts field to given value.


### GetPreviewUrls

`func (o *SchemaDto) GetPreviewUrls() map[string]string`

GetPreviewUrls returns the PreviewUrls field if non-nil, zero value otherwise.

### GetPreviewUrlsOk

`func (o *SchemaDto) GetPreviewUrlsOk() (*map[string]string, bool)`

GetPreviewUrlsOk returns a tuple with the PreviewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrls

`func (o *SchemaDto) SetPreviewUrls(v map[string]string)`

SetPreviewUrls sets PreviewUrls field to given value.


### GetFieldsInLists

`func (o *SchemaDto) GetFieldsInLists() []string`

GetFieldsInLists returns the FieldsInLists field if non-nil, zero value otherwise.

### GetFieldsInListsOk

`func (o *SchemaDto) GetFieldsInListsOk() (*[]string, bool)`

GetFieldsInListsOk returns a tuple with the FieldsInLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInLists

`func (o *SchemaDto) SetFieldsInLists(v []string)`

SetFieldsInLists sets FieldsInLists field to given value.


### GetFieldsInReferences

`func (o *SchemaDto) GetFieldsInReferences() []string`

GetFieldsInReferences returns the FieldsInReferences field if non-nil, zero value otherwise.

### GetFieldsInReferencesOk

`func (o *SchemaDto) GetFieldsInReferencesOk() (*[]string, bool)`

GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsInReferences

`func (o *SchemaDto) SetFieldsInReferences(v []string)`

SetFieldsInReferences sets FieldsInReferences field to given value.


### GetFieldRules

`func (o *SchemaDto) GetFieldRules() []FieldRuleDto`

GetFieldRules returns the FieldRules field if non-nil, zero value otherwise.

### GetFieldRulesOk

`func (o *SchemaDto) GetFieldRulesOk() (*[]FieldRuleDto, bool)`

GetFieldRulesOk returns a tuple with the FieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRules

`func (o *SchemaDto) SetFieldRules(v []FieldRuleDto)`

SetFieldRules sets FieldRules field to given value.

### HasFieldRules

`func (o *SchemaDto) HasFieldRules() bool`

HasFieldRules returns a boolean if a field has been set.

### GetFields

`func (o *SchemaDto) GetFields() []FieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SchemaDto) GetFieldsOk() (*[]FieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SchemaDto) SetFields(v []FieldDto)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


