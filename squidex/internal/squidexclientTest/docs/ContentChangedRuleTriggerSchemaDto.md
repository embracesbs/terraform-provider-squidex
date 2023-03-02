# ContentChangedRuleTriggerSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaId** | Pointer to **string** | The id of the schema. | [optional] 
**Condition** | Pointer to **NullableString** | Javascript condition when to trigger. | [optional] 

## Methods

### NewContentChangedRuleTriggerSchemaDto

`func NewContentChangedRuleTriggerSchemaDto() *ContentChangedRuleTriggerSchemaDto`

NewContentChangedRuleTriggerSchemaDto instantiates a new ContentChangedRuleTriggerSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentChangedRuleTriggerSchemaDtoWithDefaults

`func NewContentChangedRuleTriggerSchemaDtoWithDefaults() *ContentChangedRuleTriggerSchemaDto`

NewContentChangedRuleTriggerSchemaDtoWithDefaults instantiates a new ContentChangedRuleTriggerSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaId

`func (o *ContentChangedRuleTriggerSchemaDto) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ContentChangedRuleTriggerSchemaDto) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ContentChangedRuleTriggerSchemaDto) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *ContentChangedRuleTriggerSchemaDto) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetCondition

`func (o *ContentChangedRuleTriggerSchemaDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ContentChangedRuleTriggerSchemaDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ContentChangedRuleTriggerSchemaDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ContentChangedRuleTriggerSchemaDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ContentChangedRuleTriggerSchemaDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ContentChangedRuleTriggerSchemaDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


