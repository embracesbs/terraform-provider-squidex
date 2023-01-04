# ContentChangedRuleTriggerDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]ContentChangedRuleTriggerSchemaDto**](ContentChangedRuleTriggerSchemaDto.md) | The schema settings. | [optional] 
**HandleAll** | Pointer to **bool** | Determines whether the trigger should handle all content changes events. | [optional] 

## Methods

### NewContentChangedRuleTriggerDtoAllOf

`func NewContentChangedRuleTriggerDtoAllOf() *ContentChangedRuleTriggerDtoAllOf`

NewContentChangedRuleTriggerDtoAllOf instantiates a new ContentChangedRuleTriggerDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentChangedRuleTriggerDtoAllOfWithDefaults

`func NewContentChangedRuleTriggerDtoAllOfWithDefaults() *ContentChangedRuleTriggerDtoAllOf`

NewContentChangedRuleTriggerDtoAllOfWithDefaults instantiates a new ContentChangedRuleTriggerDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ContentChangedRuleTriggerDtoAllOf) GetSchemas() []ContentChangedRuleTriggerSchemaDto`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ContentChangedRuleTriggerDtoAllOf) GetSchemasOk() (*[]ContentChangedRuleTriggerSchemaDto, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ContentChangedRuleTriggerDtoAllOf) SetSchemas(v []ContentChangedRuleTriggerSchemaDto)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ContentChangedRuleTriggerDtoAllOf) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemasNil

`func (o *ContentChangedRuleTriggerDtoAllOf) SetSchemasNil(b bool)`

 SetSchemasNil sets the value for Schemas to be an explicit nil

### UnsetSchemas
`func (o *ContentChangedRuleTriggerDtoAllOf) UnsetSchemas()`

UnsetSchemas ensures that no value is present for Schemas, not even an explicit nil
### GetHandleAll

`func (o *ContentChangedRuleTriggerDtoAllOf) GetHandleAll() bool`

GetHandleAll returns the HandleAll field if non-nil, zero value otherwise.

### GetHandleAllOk

`func (o *ContentChangedRuleTriggerDtoAllOf) GetHandleAllOk() (*bool, bool)`

GetHandleAllOk returns a tuple with the HandleAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleAll

`func (o *ContentChangedRuleTriggerDtoAllOf) SetHandleAll(v bool)`

SetHandleAll sets HandleAll field to given value.

### HasHandleAll

`func (o *ContentChangedRuleTriggerDtoAllOf) HasHandleAll() bool`

HasHandleAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


