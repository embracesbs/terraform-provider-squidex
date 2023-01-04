# RuleDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the rule. | [optional] 
**CreatedBy** | **string** | The user that has created the rule. | 
**LastModifiedBy** | **string** | The user that has updated the rule. | 
**Created** | Pointer to **time.Time** | The date and time when the rule has been created. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the rule has been modified last. | [optional] 
**Version** | Pointer to **int64** | The version of the rule. | [optional] 
**IsEnabled** | Pointer to **bool** | Determines if the rule is enabled. | [optional] 
**Name** | Pointer to **NullableString** | Optional rule name. | [optional] 
**Trigger** | [**RuleTriggerDto**](RuleTriggerDto.md) |  | 
**Action** | [**RuleAction**](RuleAction.md) |  | 
**NumSucceeded** | Pointer to **int32** | The number of completed executions. | [optional] 
**NumFailed** | Pointer to **int32** | The number of failed executions. | [optional] 
**LastExecuted** | Pointer to **NullableTime** | The date and time when the rule was executed the last time. | [optional] 

## Methods

### NewRuleDtoAllOf

`func NewRuleDtoAllOf(createdBy string, lastModifiedBy string, trigger RuleTriggerDto, action RuleAction, ) *RuleDtoAllOf`

NewRuleDtoAllOf instantiates a new RuleDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleDtoAllOfWithDefaults

`func NewRuleDtoAllOfWithDefaults() *RuleDtoAllOf`

NewRuleDtoAllOfWithDefaults instantiates a new RuleDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleDtoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleDtoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleDtoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleDtoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RuleDtoAllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RuleDtoAllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RuleDtoAllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModifiedBy

`func (o *RuleDtoAllOf) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *RuleDtoAllOf) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *RuleDtoAllOf) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetCreated

`func (o *RuleDtoAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RuleDtoAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RuleDtoAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RuleDtoAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *RuleDtoAllOf) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RuleDtoAllOf) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RuleDtoAllOf) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RuleDtoAllOf) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersion

`func (o *RuleDtoAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleDtoAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleDtoAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RuleDtoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsEnabled

`func (o *RuleDtoAllOf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RuleDtoAllOf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RuleDtoAllOf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *RuleDtoAllOf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *RuleDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleDtoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleDtoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RuleDtoAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RuleDtoAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTrigger

`func (o *RuleDtoAllOf) GetTrigger() RuleTriggerDto`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RuleDtoAllOf) GetTriggerOk() (*RuleTriggerDto, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RuleDtoAllOf) SetTrigger(v RuleTriggerDto)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *RuleDtoAllOf) GetAction() RuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleDtoAllOf) GetActionOk() (*RuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleDtoAllOf) SetAction(v RuleAction)`

SetAction sets Action field to given value.


### GetNumSucceeded

`func (o *RuleDtoAllOf) GetNumSucceeded() int32`

GetNumSucceeded returns the NumSucceeded field if non-nil, zero value otherwise.

### GetNumSucceededOk

`func (o *RuleDtoAllOf) GetNumSucceededOk() (*int32, bool)`

GetNumSucceededOk returns a tuple with the NumSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSucceeded

`func (o *RuleDtoAllOf) SetNumSucceeded(v int32)`

SetNumSucceeded sets NumSucceeded field to given value.

### HasNumSucceeded

`func (o *RuleDtoAllOf) HasNumSucceeded() bool`

HasNumSucceeded returns a boolean if a field has been set.

### GetNumFailed

`func (o *RuleDtoAllOf) GetNumFailed() int32`

GetNumFailed returns the NumFailed field if non-nil, zero value otherwise.

### GetNumFailedOk

`func (o *RuleDtoAllOf) GetNumFailedOk() (*int32, bool)`

GetNumFailedOk returns a tuple with the NumFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailed

`func (o *RuleDtoAllOf) SetNumFailed(v int32)`

SetNumFailed sets NumFailed field to given value.

### HasNumFailed

`func (o *RuleDtoAllOf) HasNumFailed() bool`

HasNumFailed returns a boolean if a field has been set.

### GetLastExecuted

`func (o *RuleDtoAllOf) GetLastExecuted() time.Time`

GetLastExecuted returns the LastExecuted field if non-nil, zero value otherwise.

### GetLastExecutedOk

`func (o *RuleDtoAllOf) GetLastExecutedOk() (*time.Time, bool)`

GetLastExecutedOk returns a tuple with the LastExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecuted

`func (o *RuleDtoAllOf) SetLastExecuted(v time.Time)`

SetLastExecuted sets LastExecuted field to given value.

### HasLastExecuted

`func (o *RuleDtoAllOf) HasLastExecuted() bool`

HasLastExecuted returns a boolean if a field has been set.

### SetLastExecutedNil

`func (o *RuleDtoAllOf) SetLastExecutedNil(b bool)`

 SetLastExecutedNil sets the value for LastExecuted to be an explicit nil

### UnsetLastExecuted
`func (o *RuleDtoAllOf) UnsetLastExecuted()`

UnsetLastExecuted ensures that no value is present for LastExecuted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


