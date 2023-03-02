# SimulatedRuleEventsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]SimulatedRuleEventDto**](SimulatedRuleEventDto.md) | The simulated rule events. | 
**Total** | Pointer to **int64** | The total number of simulated rule events. | [optional] 

## Methods

### NewSimulatedRuleEventsDto

`func NewSimulatedRuleEventsDto(links map[string]ResourceLink, items []SimulatedRuleEventDto, ) *SimulatedRuleEventsDto`

NewSimulatedRuleEventsDto instantiates a new SimulatedRuleEventsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatedRuleEventsDtoWithDefaults

`func NewSimulatedRuleEventsDtoWithDefaults() *SimulatedRuleEventsDto`

NewSimulatedRuleEventsDtoWithDefaults instantiates a new SimulatedRuleEventsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SimulatedRuleEventsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SimulatedRuleEventsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SimulatedRuleEventsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *SimulatedRuleEventsDto) GetItems() []SimulatedRuleEventDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SimulatedRuleEventsDto) GetItemsOk() (*[]SimulatedRuleEventDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SimulatedRuleEventsDto) SetItems(v []SimulatedRuleEventDto)`

SetItems sets Items field to given value.


### GetTotal

`func (o *SimulatedRuleEventsDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SimulatedRuleEventsDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SimulatedRuleEventsDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SimulatedRuleEventsDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


