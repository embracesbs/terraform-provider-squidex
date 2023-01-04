# RuleEventsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Items** | [**[]RuleEventDto**](RuleEventDto.md) | The rule events. | 
**Total** | Pointer to **int64** | The total number of rule events. | [optional] 

## Methods

### NewRuleEventsDto

`func NewRuleEventsDto(links map[string]ResourceLink, items []RuleEventDto, ) *RuleEventsDto`

NewRuleEventsDto instantiates a new RuleEventsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEventsDtoWithDefaults

`func NewRuleEventsDtoWithDefaults() *RuleEventsDto`

NewRuleEventsDtoWithDefaults instantiates a new RuleEventsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RuleEventsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RuleEventsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RuleEventsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *RuleEventsDto) GetItems() []RuleEventDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RuleEventsDto) GetItemsOk() (*[]RuleEventDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RuleEventsDto) SetItems(v []RuleEventDto)`

SetItems sets Items field to given value.


### GetTotal

`func (o *RuleEventsDto) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RuleEventsDto) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RuleEventsDto) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RuleEventsDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


