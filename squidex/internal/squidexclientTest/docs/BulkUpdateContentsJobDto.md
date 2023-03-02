# BulkUpdateContentsJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**QueryJsonDto**](QueryJsonDto.md) |  | [optional] 
**Id** | Pointer to **NullableString** | An optional id of the content to update. | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **NullableString** | The new status when the type is set to &#39;ChangeStatus&#39; or &#39;Upsert&#39;. | [optional] 
**DueTime** | Pointer to **NullableTime** | The due time. | [optional] 
**Type** | Pointer to [**BulkUpdateContentType**](BulkUpdateContentType.md) |  | [optional] 
**Schema** | Pointer to **NullableString** | The optional schema id or name. | [optional] 
**Patch** | Pointer to **bool** | Makes the update as patch. | [optional] 
**Permanent** | Pointer to **bool** | True to delete the content permanently. | [optional] 
**ExpectedCount** | Pointer to **int64** | The number of expected items. Set it to a higher number to update multiple items when a query is defined. | [optional] 
**ExpectedVersion** | Pointer to **int64** | The expected version. | [optional] 

## Methods

### NewBulkUpdateContentsJobDto

`func NewBulkUpdateContentsJobDto() *BulkUpdateContentsJobDto`

NewBulkUpdateContentsJobDto instantiates a new BulkUpdateContentsJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateContentsJobDtoWithDefaults

`func NewBulkUpdateContentsJobDtoWithDefaults() *BulkUpdateContentsJobDto`

NewBulkUpdateContentsJobDtoWithDefaults instantiates a new BulkUpdateContentsJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *BulkUpdateContentsJobDto) GetQuery() QueryJsonDto`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkUpdateContentsJobDto) GetQueryOk() (*QueryJsonDto, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkUpdateContentsJobDto) SetQuery(v QueryJsonDto)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BulkUpdateContentsJobDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetId

`func (o *BulkUpdateContentsJobDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateContentsJobDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateContentsJobDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUpdateContentsJobDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BulkUpdateContentsJobDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BulkUpdateContentsJobDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *BulkUpdateContentsJobDto) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkUpdateContentsJobDto) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkUpdateContentsJobDto) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *BulkUpdateContentsJobDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *BulkUpdateContentsJobDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkUpdateContentsJobDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkUpdateContentsJobDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkUpdateContentsJobDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BulkUpdateContentsJobDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BulkUpdateContentsJobDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDueTime

`func (o *BulkUpdateContentsJobDto) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *BulkUpdateContentsJobDto) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *BulkUpdateContentsJobDto) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *BulkUpdateContentsJobDto) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *BulkUpdateContentsJobDto) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *BulkUpdateContentsJobDto) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil
### GetType

`func (o *BulkUpdateContentsJobDto) GetType() BulkUpdateContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkUpdateContentsJobDto) GetTypeOk() (*BulkUpdateContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkUpdateContentsJobDto) SetType(v BulkUpdateContentType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkUpdateContentsJobDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSchema

`func (o *BulkUpdateContentsJobDto) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BulkUpdateContentsJobDto) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BulkUpdateContentsJobDto) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *BulkUpdateContentsJobDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *BulkUpdateContentsJobDto) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *BulkUpdateContentsJobDto) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetPatch

`func (o *BulkUpdateContentsJobDto) GetPatch() bool`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *BulkUpdateContentsJobDto) GetPatchOk() (*bool, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *BulkUpdateContentsJobDto) SetPatch(v bool)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *BulkUpdateContentsJobDto) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPermanent

`func (o *BulkUpdateContentsJobDto) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *BulkUpdateContentsJobDto) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *BulkUpdateContentsJobDto) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *BulkUpdateContentsJobDto) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.

### GetExpectedCount

`func (o *BulkUpdateContentsJobDto) GetExpectedCount() int64`

GetExpectedCount returns the ExpectedCount field if non-nil, zero value otherwise.

### GetExpectedCountOk

`func (o *BulkUpdateContentsJobDto) GetExpectedCountOk() (*int64, bool)`

GetExpectedCountOk returns a tuple with the ExpectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCount

`func (o *BulkUpdateContentsJobDto) SetExpectedCount(v int64)`

SetExpectedCount sets ExpectedCount field to given value.

### HasExpectedCount

`func (o *BulkUpdateContentsJobDto) HasExpectedCount() bool`

HasExpectedCount returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *BulkUpdateContentsJobDto) GetExpectedVersion() int64`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *BulkUpdateContentsJobDto) GetExpectedVersionOk() (*int64, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *BulkUpdateContentsJobDto) SetExpectedVersion(v int64)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *BulkUpdateContentsJobDto) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


