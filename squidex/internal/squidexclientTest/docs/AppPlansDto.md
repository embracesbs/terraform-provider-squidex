# AppPlansDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | [**[]PlanDto**](PlanDto.md) | The available plans. | 
**CurrentPlanId** | Pointer to **NullableString** | The current plan id. | [optional] 
**PlanOwner** | Pointer to **NullableString** | The plan owner. | [optional] 
**HasPortal** | Pointer to **bool** | Indicates if there is a billing portal. | [optional] 

## Methods

### NewAppPlansDto

`func NewAppPlansDto(plans []PlanDto, ) *AppPlansDto`

NewAppPlansDto instantiates a new AppPlansDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPlansDtoWithDefaults

`func NewAppPlansDtoWithDefaults() *AppPlansDto`

NewAppPlansDtoWithDefaults instantiates a new AppPlansDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *AppPlansDto) GetPlans() []PlanDto`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AppPlansDto) GetPlansOk() (*[]PlanDto, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AppPlansDto) SetPlans(v []PlanDto)`

SetPlans sets Plans field to given value.


### GetCurrentPlanId

`func (o *AppPlansDto) GetCurrentPlanId() string`

GetCurrentPlanId returns the CurrentPlanId field if non-nil, zero value otherwise.

### GetCurrentPlanIdOk

`func (o *AppPlansDto) GetCurrentPlanIdOk() (*string, bool)`

GetCurrentPlanIdOk returns a tuple with the CurrentPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlanId

`func (o *AppPlansDto) SetCurrentPlanId(v string)`

SetCurrentPlanId sets CurrentPlanId field to given value.

### HasCurrentPlanId

`func (o *AppPlansDto) HasCurrentPlanId() bool`

HasCurrentPlanId returns a boolean if a field has been set.

### SetCurrentPlanIdNil

`func (o *AppPlansDto) SetCurrentPlanIdNil(b bool)`

 SetCurrentPlanIdNil sets the value for CurrentPlanId to be an explicit nil

### UnsetCurrentPlanId
`func (o *AppPlansDto) UnsetCurrentPlanId()`

UnsetCurrentPlanId ensures that no value is present for CurrentPlanId, not even an explicit nil
### GetPlanOwner

`func (o *AppPlansDto) GetPlanOwner() string`

GetPlanOwner returns the PlanOwner field if non-nil, zero value otherwise.

### GetPlanOwnerOk

`func (o *AppPlansDto) GetPlanOwnerOk() (*string, bool)`

GetPlanOwnerOk returns a tuple with the PlanOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanOwner

`func (o *AppPlansDto) SetPlanOwner(v string)`

SetPlanOwner sets PlanOwner field to given value.

### HasPlanOwner

`func (o *AppPlansDto) HasPlanOwner() bool`

HasPlanOwner returns a boolean if a field has been set.

### SetPlanOwnerNil

`func (o *AppPlansDto) SetPlanOwnerNil(b bool)`

 SetPlanOwnerNil sets the value for PlanOwner to be an explicit nil

### UnsetPlanOwner
`func (o *AppPlansDto) UnsetPlanOwner()`

UnsetPlanOwner ensures that no value is present for PlanOwner, not even an explicit nil
### GetHasPortal

`func (o *AppPlansDto) GetHasPortal() bool`

GetHasPortal returns the HasPortal field if non-nil, zero value otherwise.

### GetHasPortalOk

`func (o *AppPlansDto) GetHasPortalOk() (*bool, bool)`

GetHasPortalOk returns a tuple with the HasPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPortal

`func (o *AppPlansDto) SetHasPortal(v bool)`

SetHasPortal sets HasPortal field to given value.

### HasHasPortal

`func (o *AppPlansDto) HasHasPortal() bool`

HasHasPortal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


