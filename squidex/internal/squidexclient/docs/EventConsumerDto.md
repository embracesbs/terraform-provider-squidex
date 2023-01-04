# EventConsumerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**IsStopped** | Pointer to **bool** |  | [optional] 
**IsResetting** | Pointer to **bool** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventConsumerDto

`func NewEventConsumerDto(links map[string]ResourceLink, ) *EventConsumerDto`

NewEventConsumerDto instantiates a new EventConsumerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventConsumerDtoWithDefaults

`func NewEventConsumerDtoWithDefaults() *EventConsumerDto`

NewEventConsumerDtoWithDefaults instantiates a new EventConsumerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EventConsumerDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventConsumerDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventConsumerDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetIsStopped

`func (o *EventConsumerDto) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *EventConsumerDto) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *EventConsumerDto) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.

### HasIsStopped

`func (o *EventConsumerDto) HasIsStopped() bool`

HasIsStopped returns a boolean if a field has been set.

### GetIsResetting

`func (o *EventConsumerDto) GetIsResetting() bool`

GetIsResetting returns the IsResetting field if non-nil, zero value otherwise.

### GetIsResettingOk

`func (o *EventConsumerDto) GetIsResettingOk() (*bool, bool)`

GetIsResettingOk returns a tuple with the IsResetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResetting

`func (o *EventConsumerDto) SetIsResetting(v bool)`

SetIsResetting sets IsResetting field to given value.

### HasIsResetting

`func (o *EventConsumerDto) HasIsResetting() bool`

HasIsResetting returns a boolean if a field has been set.

### GetCount

`func (o *EventConsumerDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventConsumerDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventConsumerDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventConsumerDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *EventConsumerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventConsumerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventConsumerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventConsumerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetError

`func (o *EventConsumerDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventConsumerDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventConsumerDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventConsumerDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EventConsumerDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EventConsumerDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPosition

`func (o *EventConsumerDto) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EventConsumerDto) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EventConsumerDto) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *EventConsumerDto) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *EventConsumerDto) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *EventConsumerDto) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


