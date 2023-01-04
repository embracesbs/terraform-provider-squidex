# SimulatedRuleEventsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SimulatedRuleEventDto**](SimulatedRuleEventDto.md) | The simulated rule events. | 
**Total** | Pointer to **int64** | The total number of simulated rule events. | [optional] 

## Methods

### NewSimulatedRuleEventsDtoAllOf

`func NewSimulatedRuleEventsDtoAllOf(items []SimulatedRuleEventDto, ) *SimulatedRuleEventsDtoAllOf`

NewSimulatedRuleEventsDtoAllOf instantiates a new SimulatedRuleEventsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatedRuleEventsDtoAllOfWithDefaults

`func NewSimulatedRuleEventsDtoAllOfWithDefaults() *SimulatedRuleEventsDtoAllOf`

NewSimulatedRuleEventsDtoAllOfWithDefaults instantiates a new SimulatedRuleEventsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SimulatedRuleEventsDtoAllOf) GetItems() []SimulatedRuleEventDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SimulatedRuleEventsDtoAllOf) GetItemsOk() (*[]SimulatedRuleEventDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SimulatedRuleEventsDtoAllOf) SetItems(v []SimulatedRuleEventDto)`

SetItems sets Items field to given value.


### GetTotal

`func (o *SimulatedRuleEventsDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SimulatedRuleEventsDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SimulatedRuleEventsDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SimulatedRuleEventsDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


