# CreateAppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the app. | 
**Template** | Pointer to **NullableString** | Initialize the app with the inbuilt template. | [optional] 

## Methods

### NewCreateAppDto

`func NewCreateAppDto(name string, ) *CreateAppDto`

NewCreateAppDto instantiates a new CreateAppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppDtoWithDefaults

`func NewCreateAppDtoWithDefaults() *CreateAppDto`

NewCreateAppDtoWithDefaults instantiates a new CreateAppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppDto) SetName(v string)`

SetName sets Name field to given value.


### GetTemplate

`func (o *CreateAppDto) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateAppDto) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateAppDto) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CreateAppDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CreateAppDto) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CreateAppDto) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


