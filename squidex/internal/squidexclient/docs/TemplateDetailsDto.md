# TemplateDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Details** | **string** | The details of the template. | 

## Methods

### NewTemplateDetailsDto

`func NewTemplateDetailsDto(links map[string]ResourceLink, details string, ) *TemplateDetailsDto`

NewTemplateDetailsDto instantiates a new TemplateDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDetailsDtoWithDefaults

`func NewTemplateDetailsDtoWithDefaults() *TemplateDetailsDto`

NewTemplateDetailsDtoWithDefaults instantiates a new TemplateDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateDetailsDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateDetailsDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateDetailsDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetDetails

`func (o *TemplateDetailsDto) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TemplateDetailsDto) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TemplateDetailsDto) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


