# RuleElementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Describes the action or trigger type. | 
**Display** | **string** | The label for the action or trigger type. | 
**Title** | Pointer to **NullableString** | Optional title. | [optional] 
**IconColor** | Pointer to **NullableString** | The color for the icon. | [optional] 
**IconImage** | Pointer to **NullableString** | The image for the icon. | [optional] 
**ReadMore** | Pointer to **NullableString** | The optional link to the product that is integrated. | [optional] 
**Properties** | [**[]RuleElementPropertyDto**](RuleElementPropertyDto.md) | The properties. | 

## Methods

### NewRuleElementDto

`func NewRuleElementDto(description string, display string, properties []RuleElementPropertyDto, ) *RuleElementDto`

NewRuleElementDto instantiates a new RuleElementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleElementDtoWithDefaults

`func NewRuleElementDtoWithDefaults() *RuleElementDto`

NewRuleElementDtoWithDefaults instantiates a new RuleElementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RuleElementDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleElementDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleElementDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplay

`func (o *RuleElementDto) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RuleElementDto) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RuleElementDto) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetTitle

`func (o *RuleElementDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleElementDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleElementDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RuleElementDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RuleElementDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RuleElementDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetIconColor

`func (o *RuleElementDto) GetIconColor() string`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *RuleElementDto) GetIconColorOk() (*string, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *RuleElementDto) SetIconColor(v string)`

SetIconColor sets IconColor field to given value.

### HasIconColor

`func (o *RuleElementDto) HasIconColor() bool`

HasIconColor returns a boolean if a field has been set.

### SetIconColorNil

`func (o *RuleElementDto) SetIconColorNil(b bool)`

 SetIconColorNil sets the value for IconColor to be an explicit nil

### UnsetIconColor
`func (o *RuleElementDto) UnsetIconColor()`

UnsetIconColor ensures that no value is present for IconColor, not even an explicit nil
### GetIconImage

`func (o *RuleElementDto) GetIconImage() string`

GetIconImage returns the IconImage field if non-nil, zero value otherwise.

### GetIconImageOk

`func (o *RuleElementDto) GetIconImageOk() (*string, bool)`

GetIconImageOk returns a tuple with the IconImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconImage

`func (o *RuleElementDto) SetIconImage(v string)`

SetIconImage sets IconImage field to given value.

### HasIconImage

`func (o *RuleElementDto) HasIconImage() bool`

HasIconImage returns a boolean if a field has been set.

### SetIconImageNil

`func (o *RuleElementDto) SetIconImageNil(b bool)`

 SetIconImageNil sets the value for IconImage to be an explicit nil

### UnsetIconImage
`func (o *RuleElementDto) UnsetIconImage()`

UnsetIconImage ensures that no value is present for IconImage, not even an explicit nil
### GetReadMore

`func (o *RuleElementDto) GetReadMore() string`

GetReadMore returns the ReadMore field if non-nil, zero value otherwise.

### GetReadMoreOk

`func (o *RuleElementDto) GetReadMoreOk() (*string, bool)`

GetReadMoreOk returns a tuple with the ReadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadMore

`func (o *RuleElementDto) SetReadMore(v string)`

SetReadMore sets ReadMore field to given value.

### HasReadMore

`func (o *RuleElementDto) HasReadMore() bool`

HasReadMore returns a boolean if a field has been set.

### SetReadMoreNil

`func (o *RuleElementDto) SetReadMoreNil(b bool)`

 SetReadMoreNil sets the value for ReadMore to be an explicit nil

### UnsetReadMore
`func (o *RuleElementDto) UnsetReadMore()`

UnsetReadMore ensures that no value is present for ReadMore, not even an explicit nil
### GetProperties

`func (o *RuleElementDto) GetProperties() []RuleElementPropertyDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RuleElementDto) GetPropertiesOk() (*[]RuleElementPropertyDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RuleElementDto) SetProperties(v []RuleElementPropertyDto)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


