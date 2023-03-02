# QueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | The optional list of ids to query. | [optional] 
**OData** | Pointer to **NullableString** | The optional odata query. | [optional] 
**Q** | Pointer to **interface{}** | The optional json query. | [optional] 
**ParentId** | Pointer to **NullableString** | The parent id (for assets). | [optional] 

## Methods

### NewQueryDto

`func NewQueryDto() *QueryDto`

NewQueryDto instantiates a new QueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDtoWithDefaults

`func NewQueryDtoWithDefaults() *QueryDto`

NewQueryDtoWithDefaults instantiates a new QueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *QueryDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *QueryDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *QueryDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *QueryDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *QueryDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *QueryDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetOData

`func (o *QueryDto) GetOData() string`

GetOData returns the OData field if non-nil, zero value otherwise.

### GetODataOk

`func (o *QueryDto) GetODataOk() (*string, bool)`

GetODataOk returns a tuple with the OData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOData

`func (o *QueryDto) SetOData(v string)`

SetOData sets OData field to given value.

### HasOData

`func (o *QueryDto) HasOData() bool`

HasOData returns a boolean if a field has been set.

### SetODataNil

`func (o *QueryDto) SetODataNil(b bool)`

 SetODataNil sets the value for OData to be an explicit nil

### UnsetOData
`func (o *QueryDto) UnsetOData()`

UnsetOData ensures that no value is present for OData, not even an explicit nil
### GetQ

`func (o *QueryDto) GetQ() interface{}`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *QueryDto) GetQOk() (*interface{}, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *QueryDto) SetQ(v interface{})`

SetQ sets Q field to given value.

### HasQ

`func (o *QueryDto) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQNil

`func (o *QueryDto) SetQNil(b bool)`

 SetQNil sets the value for Q to be an explicit nil

### UnsetQ
`func (o *QueryDto) UnsetQ()`

UnsetQ ensures that no value is present for Q, not even an explicit nil
### GetParentId

`func (o *QueryDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QueryDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QueryDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QueryDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *QueryDto) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *QueryDto) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


