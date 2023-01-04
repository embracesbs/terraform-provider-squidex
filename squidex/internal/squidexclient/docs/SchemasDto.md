# SchemasDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | Pointer to [**[]SchemaDto**](SchemaDto.md) | The schemas. | [optional] 

## Methods

### NewSchemasDto

`func NewSchemasDto(links map[string]ResourceLink, ) *SchemasDto`

NewSchemasDto instantiates a new SchemasDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasDtoWithDefaults

`func NewSchemasDtoWithDefaults() *SchemasDto`

NewSchemasDtoWithDefaults instantiates a new SchemasDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SchemasDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchemasDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchemasDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *SchemasDto) GetItems() []SchemaDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SchemasDto) GetItemsOk() (*[]SchemaDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SchemasDto) SetItems(v []SchemaDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *SchemasDto) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


