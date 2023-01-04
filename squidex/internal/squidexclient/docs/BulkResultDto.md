# BulkResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorDto**](ErrorDto.md) |  | [optional] 
**JobIndex** | Pointer to **int32** | The index of the bulk job where the result belongs to. The order can change. | [optional] 
**Id** | Pointer to **NullableString** | The id of the entity that has been handled successfully or not. | [optional] 
**ContentId** | Pointer to **NullableString** | The id of the entity that has been handled successfully or not. | [optional] 

## Methods

### NewBulkResultDto

`func NewBulkResultDto() *BulkResultDto`

NewBulkResultDto instantiates a new BulkResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResultDtoWithDefaults

`func NewBulkResultDtoWithDefaults() *BulkResultDto`

NewBulkResultDtoWithDefaults instantiates a new BulkResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BulkResultDto) GetError() ErrorDto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResultDto) GetErrorOk() (*ErrorDto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResultDto) SetError(v ErrorDto)`

SetError sets Error field to given value.

### HasError

`func (o *BulkResultDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetJobIndex

`func (o *BulkResultDto) GetJobIndex() int32`

GetJobIndex returns the JobIndex field if non-nil, zero value otherwise.

### GetJobIndexOk

`func (o *BulkResultDto) GetJobIndexOk() (*int32, bool)`

GetJobIndexOk returns a tuple with the JobIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIndex

`func (o *BulkResultDto) SetJobIndex(v int32)`

SetJobIndex sets JobIndex field to given value.

### HasJobIndex

`func (o *BulkResultDto) HasJobIndex() bool`

HasJobIndex returns a boolean if a field has been set.

### GetId

`func (o *BulkResultDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkResultDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkResultDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkResultDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BulkResultDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BulkResultDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetContentId

`func (o *BulkResultDto) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *BulkResultDto) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *BulkResultDto) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *BulkResultDto) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### SetContentIdNil

`func (o *BulkResultDto) SetContentIdNil(b bool)`

 SetContentIdNil sets the value for ContentId to be an explicit nil

### UnsetContentId
`func (o *BulkResultDto) UnsetContentId()`

UnsetContentId ensures that no value is present for ContentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


