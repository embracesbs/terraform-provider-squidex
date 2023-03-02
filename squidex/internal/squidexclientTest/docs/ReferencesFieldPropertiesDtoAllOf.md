# ReferencesFieldPropertiesDtoAllOf

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

### NewReferencesFieldPropertiesDtoAllOf

`func NewReferencesFieldPropertiesDtoAllOf() *ReferencesFieldPropertiesDtoAllOf`

NewReferencesFieldPropertiesDtoAllOf instantiates a new ReferencesFieldPropertiesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencesFieldPropertiesDtoAllOfWithDefaults

`func NewReferencesFieldPropertiesDtoAllOfWithDefaults() *ReferencesFieldPropertiesDtoAllOf`

NewReferencesFieldPropertiesDtoAllOfWithDefaults instantiates a new ReferencesFieldPropertiesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValues() map[string][]string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string][]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ReferencesFieldPropertiesDtoAllOf) SetDefaultValues(v map[string][]string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ReferencesFieldPropertiesDtoAllOf) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValue() []string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetDefaultValueOk() (*[]string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ReferencesFieldPropertiesDtoAllOf) SetDefaultValue(v []string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ReferencesFieldPropertiesDtoAllOf) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ReferencesFieldPropertiesDtoAllOf) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ReferencesFieldPropertiesDtoAllOf) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMinItems

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *ReferencesFieldPropertiesDtoAllOf) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *ReferencesFieldPropertiesDtoAllOf) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *ReferencesFieldPropertiesDtoAllOf) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *ReferencesFieldPropertiesDtoAllOf) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *ReferencesFieldPropertiesDtoAllOf) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *ReferencesFieldPropertiesDtoAllOf) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *ReferencesFieldPropertiesDtoAllOf) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *ReferencesFieldPropertiesDtoAllOf) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetAllowDuplicates

`func (o *ReferencesFieldPropertiesDtoAllOf) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *ReferencesFieldPropertiesDtoAllOf) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *ReferencesFieldPropertiesDtoAllOf) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetResolveReference

`func (o *ReferencesFieldPropertiesDtoAllOf) GetResolveReference() bool`

GetResolveReference returns the ResolveReference field if non-nil, zero value otherwise.

### GetResolveReferenceOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetResolveReferenceOk() (*bool, bool)`

GetResolveReferenceOk returns a tuple with the ResolveReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveReference

`func (o *ReferencesFieldPropertiesDtoAllOf) SetResolveReference(v bool)`

SetResolveReference sets ResolveReference field to given value.

### HasResolveReference

`func (o *ReferencesFieldPropertiesDtoAllOf) HasResolveReference() bool`

HasResolveReference returns a boolean if a field has been set.

### GetMustBePublished

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMustBePublished() bool`

GetMustBePublished returns the MustBePublished field if non-nil, zero value otherwise.

### GetMustBePublishedOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetMustBePublishedOk() (*bool, bool)`

GetMustBePublishedOk returns a tuple with the MustBePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBePublished

`func (o *ReferencesFieldPropertiesDtoAllOf) SetMustBePublished(v bool)`

SetMustBePublished sets MustBePublished field to given value.

### HasMustBePublished

`func (o *ReferencesFieldPropertiesDtoAllOf) HasMustBePublished() bool`

HasMustBePublished returns a boolean if a field has been set.

### GetEditor

`func (o *ReferencesFieldPropertiesDtoAllOf) GetEditor() ReferencesFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetEditorOk() (*ReferencesFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *ReferencesFieldPropertiesDtoAllOf) SetEditor(v ReferencesFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *ReferencesFieldPropertiesDtoAllOf) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetSchemaIds

`func (o *ReferencesFieldPropertiesDtoAllOf) GetSchemaIds() []string`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *ReferencesFieldPropertiesDtoAllOf) GetSchemaIdsOk() (*[]string, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *ReferencesFieldPropertiesDtoAllOf) SetSchemaIds(v []string)`

SetSchemaIds sets SchemaIds field to given value.

### HasSchemaIds

`func (o *ReferencesFieldPropertiesDtoAllOf) HasSchemaIds() bool`

HasSchemaIds returns a boolean if a field has been set.

### SetSchemaIdsNil

`func (o *ReferencesFieldPropertiesDtoAllOf) SetSchemaIdsNil(b bool)`

 SetSchemaIdsNil sets the value for SchemaIds to be an explicit nil

### UnsetSchemaIds
`func (o *ReferencesFieldPropertiesDtoAllOf) UnsetSchemaIds()`

UnsetSchemaIds ensures that no value is present for SchemaIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


