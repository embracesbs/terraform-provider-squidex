# UIFieldPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Editor** | Pointer to [**UIFieldEditor**](UIFieldEditor.md) |  | [optional] 

## Methods

### NewUIFieldPropertiesDto

`func NewUIFieldPropertiesDto() *UIFieldPropertiesDto`

NewUIFieldPropertiesDto instantiates a new UIFieldPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIFieldPropertiesDtoWithDefaults

`func NewUIFieldPropertiesDtoWithDefaults() *UIFieldPropertiesDto`

NewUIFieldPropertiesDtoWithDefaults instantiates a new UIFieldPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditor

`func (o *UIFieldPropertiesDto) GetEditor() UIFieldEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *UIFieldPropertiesDto) GetEditorOk() (*UIFieldEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *UIFieldPropertiesDto) SetEditor(v UIFieldEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *UIFieldPropertiesDto) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

