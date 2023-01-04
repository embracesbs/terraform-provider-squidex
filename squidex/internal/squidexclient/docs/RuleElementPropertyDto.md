# RuleElementPropertyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Editor** | [**RuleFieldEditor**](RuleFieldEditor.md) |  | 
**Name** | **string** | The name of the editor. | 
**Display** | **string** | The label to use. | 
**Options** | Pointer to **[]string** | The options, if the editor is a dropdown. | [optional] 
**Description** | Pointer to **NullableString** | The optional description. | [optional] 
**IsFormattable** | Pointer to **bool** | Indicates if the property is formattable. | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the property is required. | [optional] 

## Methods

### NewRuleElementPropertyDto

`func NewRuleElementPropertyDto(editor RuleFieldEditor, name string, display string, ) *RuleElementPropertyDto`

NewRuleElementPropertyDto instantiates a new RuleElementPropertyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleElementPropertyDtoWithDefaults

`func NewRuleElementPropertyDtoWithDefaults() *RuleElementPropertyDto`

NewRuleElementPropertyDtoWithDefaults instantiates a new RuleElementPropertyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditor

`func (o *RuleElementPropertyDto) GetEditor() RuleFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *RuleElementPropertyDto) GetEditorOk() (*RuleFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *RuleElementPropertyDto) SetEditor(v RuleFieldEditor)`

SetEditor sets Editor field to given value.


### GetName

`func (o *RuleElementPropertyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleElementPropertyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleElementPropertyDto) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *RuleElementPropertyDto) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RuleElementPropertyDto) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RuleElementPropertyDto) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetOptions

`func (o *RuleElementPropertyDto) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RuleElementPropertyDto) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RuleElementPropertyDto) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RuleElementPropertyDto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *RuleElementPropertyDto) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *RuleElementPropertyDto) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDescription

`func (o *RuleElementPropertyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleElementPropertyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleElementPropertyDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleElementPropertyDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RuleElementPropertyDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RuleElementPropertyDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFormattable

`func (o *RuleElementPropertyDto) GetIsFormattable() bool`

GetIsFormattable returns the IsFormattable field if non-nil, zero value otherwise.

### GetIsFormattableOk

`func (o *RuleElementPropertyDto) GetIsFormattableOk() (*bool, bool)`

GetIsFormattableOk returns a tuple with the IsFormattable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFormattable

`func (o *RuleElementPropertyDto) SetIsFormattable(v bool)`

SetIsFormattable sets IsFormattable field to given value.

### HasIsFormattable

`func (o *RuleElementPropertyDto) HasIsFormattable() bool`

HasIsFormattable returns a boolean if a field has been set.

### GetIsRequired

`func (o *RuleElementPropertyDto) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *RuleElementPropertyDto) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *RuleElementPropertyDto) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *RuleElementPropertyDto) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


