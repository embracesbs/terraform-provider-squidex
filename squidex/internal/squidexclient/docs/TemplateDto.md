# TemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Name** | **string** | The name of the template. | 
**Title** | **string** | The title of the template. | 
**Description** | **string** | The description of the template. | 
**IsStarter** | Pointer to **bool** | True, if the template is a starter. | [optional] 

## Methods

### NewTemplateDto

`func NewTemplateDto(links map[string]ResourceLink, name string, title string, description string, ) *TemplateDto`

NewTemplateDto instantiates a new TemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDtoWithDefaults

`func NewTemplateDtoWithDefaults() *TemplateDto`

NewTemplateDtoWithDefaults instantiates a new TemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetName

`func (o *TemplateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDto) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *TemplateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TemplateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsStarter

`func (o *TemplateDto) GetIsStarter() bool`

GetIsStarter returns the IsStarter field if non-nil, zero value otherwise.

### GetIsStarterOk

`func (o *TemplateDto) GetIsStarterOk() (*bool, bool)`

GetIsStarterOk returns a tuple with the IsStarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarter

`func (o *TemplateDto) SetIsStarter(v bool)`

SetIsStarter sets IsStarter field to given value.

### HasIsStarter

`func (o *TemplateDto) HasIsStarter() bool`

HasIsStarter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


