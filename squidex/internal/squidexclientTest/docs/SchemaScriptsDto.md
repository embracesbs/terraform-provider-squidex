# SchemaScriptsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** | The script that is executed for each content when querying contents. | [optional] 
**QueryPre** | Pointer to **NullableString** | The script that is executed for all contents when querying contents. | [optional] 
**Create** | Pointer to **NullableString** | The script that is executed when creating a content. | [optional] 
**Update** | Pointer to **NullableString** | The script that is executed when updating a content. | [optional] 
**Delete** | Pointer to **NullableString** | The script that is executed when deleting a content. | [optional] 
**Change** | Pointer to **NullableString** | The script that is executed when change a content status. | [optional] 

## Methods

### NewSchemaScriptsDto

`func NewSchemaScriptsDto() *SchemaScriptsDto`

NewSchemaScriptsDto instantiates a new SchemaScriptsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaScriptsDtoWithDefaults

`func NewSchemaScriptsDtoWithDefaults() *SchemaScriptsDto`

NewSchemaScriptsDtoWithDefaults instantiates a new SchemaScriptsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SchemaScriptsDto) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SchemaScriptsDto) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SchemaScriptsDto) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SchemaScriptsDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *SchemaScriptsDto) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *SchemaScriptsDto) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetQueryPre

`func (o *SchemaScriptsDto) GetQueryPre() string`

GetQueryPre returns the QueryPre field if non-nil, zero value otherwise.

### GetQueryPreOk

`func (o *SchemaScriptsDto) GetQueryPreOk() (*string, bool)`

GetQueryPreOk returns a tuple with the QueryPre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPre

`func (o *SchemaScriptsDto) SetQueryPre(v string)`

SetQueryPre sets QueryPre field to given value.

### HasQueryPre

`func (o *SchemaScriptsDto) HasQueryPre() bool`

HasQueryPre returns a boolean if a field has been set.

### SetQueryPreNil

`func (o *SchemaScriptsDto) SetQueryPreNil(b bool)`

 SetQueryPreNil sets the value for QueryPre to be an explicit nil

### UnsetQueryPre
`func (o *SchemaScriptsDto) UnsetQueryPre()`

UnsetQueryPre ensures that no value is present for QueryPre, not even an explicit nil
### GetCreate

`func (o *SchemaScriptsDto) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SchemaScriptsDto) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SchemaScriptsDto) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SchemaScriptsDto) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreateNil

`func (o *SchemaScriptsDto) SetCreateNil(b bool)`

 SetCreateNil sets the value for Create to be an explicit nil

### UnsetCreate
`func (o *SchemaScriptsDto) UnsetCreate()`

UnsetCreate ensures that no value is present for Create, not even an explicit nil
### GetUpdate

`func (o *SchemaScriptsDto) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SchemaScriptsDto) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SchemaScriptsDto) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SchemaScriptsDto) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *SchemaScriptsDto) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *SchemaScriptsDto) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetDelete

`func (o *SchemaScriptsDto) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *SchemaScriptsDto) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *SchemaScriptsDto) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *SchemaScriptsDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *SchemaScriptsDto) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *SchemaScriptsDto) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetChange

`func (o *SchemaScriptsDto) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *SchemaScriptsDto) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *SchemaScriptsDto) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *SchemaScriptsDto) HasChange() bool`

HasChange returns a boolean if a field has been set.

### SetChangeNil

`func (o *SchemaScriptsDto) SetChangeNil(b bool)`

 SetChangeNil sets the value for Change to be an explicit nil

### UnsetChange
`func (o *SchemaScriptsDto) UnsetChange()`

UnsetChange ensures that no value is present for Change, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


