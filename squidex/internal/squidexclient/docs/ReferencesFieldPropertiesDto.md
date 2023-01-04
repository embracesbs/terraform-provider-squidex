# ReferencesFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to **map[string][]string** |  | [optional] 
**DefaultValue** | Pointer to **[]string** | The default value as a list of content ids. | [optional] 
**MinItems** | Pointer to **NullableInt32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | The maximum allowed items for the field value. | [optional] 
**AllowDuplicates** | Pointer to **bool** | True, if duplicate values are allowed. | [optional] 
**ResolveReference** | Pointer to **bool** | True to resolve references in the content list. | [optional] 
**MustBePublished** | Pointer to **bool** | True when all references must be published. | [optional] 
**Editor** | Pointer to [**ReferencesFieldEditor**](ReferencesFieldEditor.md) |  | [optional] 
**SchemaIds** | Pointer to **[]string** | The id of the referenced schemas. | [optional] 

## Methods

### NewReferencesFieldPropertiesDto

`func NewReferencesFieldPropertiesDto() *ReferencesFieldPropertiesDto`

NewReferencesFieldPropertiesDto instantiates a new ReferencesFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencesFieldPropertiesDtoWithDefaults

`func NewReferencesFieldPropertiesDtoWithDefaults() *ReferencesFieldPropertiesDto`

NewReferencesFieldPropertiesDtoWithDefaults instantiates a new ReferencesFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *ReferencesFieldPropertiesDto) GetDefaultValues() map[string][]string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ReferencesFieldPropertiesDto) GetDefaultValuesOk() (*map[string][]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ReferencesFieldPropertiesDto) SetDefaultValues(v map[string][]string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ReferencesFieldPropertiesDto) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ReferencesFieldPropertiesDto) GetDefaultValue() []string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ReferencesFieldPropertiesDto) GetDefaultValueOk() (*[]string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ReferencesFieldPropertiesDto) SetDefaultValue(v []string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ReferencesFieldPropertiesDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ReferencesFieldPropertiesDto) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ReferencesFieldPropertiesDto) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMinItems

`func (o *ReferencesFieldPropertiesDto) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *ReferencesFieldPropertiesDto) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *ReferencesFieldPropertiesDto) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *ReferencesFieldPropertiesDto) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *ReferencesFieldPropertiesDto) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *ReferencesFieldPropertiesDto) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *ReferencesFieldPropertiesDto) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *ReferencesFieldPropertiesDto) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *ReferencesFieldPropertiesDto) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *ReferencesFieldPropertiesDto) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *ReferencesFieldPropertiesDto) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *ReferencesFieldPropertiesDto) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetAllowDuplicates

`func (o *ReferencesFieldPropertiesDto) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *ReferencesFieldPropertiesDto) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *ReferencesFieldPropertiesDto) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *ReferencesFieldPropertiesDto) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetResolveReference

`func (o *ReferencesFieldPropertiesDto) GetResolveReference() bool`

GetResolveReference returns the ResolveReference field if non-nil, zero value otherwise.

### GetResolveReferenceOk

`func (o *ReferencesFieldPropertiesDto) GetResolveReferenceOk() (*bool, bool)`

GetResolveReferenceOk returns a tuple with the ResolveReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveReference

`func (o *ReferencesFieldPropertiesDto) SetResolveReference(v bool)`

SetResolveReference sets ResolveReference field to given value.

### HasResolveReference

`func (o *ReferencesFieldPropertiesDto) HasResolveReference() bool`

HasResolveReference returns a boolean if a field has been set.

### GetMustBePublished

`func (o *ReferencesFieldPropertiesDto) GetMustBePublished() bool`

GetMustBePublished returns the MustBePublished field if non-nil, zero value otherwise.

### GetMustBePublishedOk

`func (o *ReferencesFieldPropertiesDto) GetMustBePublishedOk() (*bool, bool)`

GetMustBePublishedOk returns a tuple with the MustBePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBePublished

`func (o *ReferencesFieldPropertiesDto) SetMustBePublished(v bool)`

SetMustBePublished sets MustBePublished field to given value.

### HasMustBePublished

`func (o *ReferencesFieldPropertiesDto) HasMustBePublished() bool`

HasMustBePublished returns a boolean if a field has been set.

### GetEditor

`func (o *ReferencesFieldPropertiesDto) GetEditor() ReferencesFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *ReferencesFieldPropertiesDto) GetEditorOk() (*ReferencesFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *ReferencesFieldPropertiesDto) SetEditor(v ReferencesFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *ReferencesFieldPropertiesDto) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetSchemaIds

`func (o *ReferencesFieldPropertiesDto) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *ReferencesFieldPropertiesDto) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *ReferencesFieldPropertiesDto) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *ReferencesFieldPropertiesDto) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *ReferencesFieldPropertiesDto) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *ReferencesFieldPropertiesDto) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


