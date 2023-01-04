# RuleEventsDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]RuleEventDto**](RuleEventDto.md) | The rule events. | 
**Total** | Pointer to **int64** | The total number of rule events. | [optional] 

## Methods

### NewRuleEventsDtoAllOf

`func NewRuleEventsDtoAllOf(items []RuleEventDto, ) *RuleEventsDtoAllOf`

NewRuleEventsDtoAllOf instantiates a new RuleEventsDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEventsDtoAllOfWithDefaults

`func NewRuleEventsDtoAllOfWithDefaults() *RuleEventsDtoAllOf`

NewRuleEventsDtoAllOfWithDefaults instantiates a new RuleEventsDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RuleEventsDtoAllOf) GetItems() []RuleEventDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RuleEventsDtoAllOf) GetItemsOk() (*[]RuleEventDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RuleEventsDtoAllOf) SetItems(v []RuleEventDto)`

SetItems sets Items field to given value.


### GetTotal

`func (o *RuleEventsDtoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RuleEventsDtoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RuleEventsDtoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RuleEventsDtoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


