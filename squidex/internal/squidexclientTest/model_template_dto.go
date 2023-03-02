/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
)

// TemplateDto struct for TemplateDto
type TemplateDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The name of the template.
	Name string `json:"name"`
	// The title of the template.
	Title string `json:"title"`
	// The description of the template.
	Description string `json:"description"`
	// True, if the template is a starter.
	IsStarter *bool `json:"isStarter,omitempty"`
}

// NewTemplateDto instantiates a new TemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDto(links map[string]ResourceLink, name string, title string, description string) *TemplateDto {
	this := TemplateDto{}
	this.Links = links
	this.Name = name
	this.Title = title
	this.Description = description
	return &this
}

// NewTemplateDtoWithDefaults instantiates a new TemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDtoWithDefaults() *TemplateDto {
	this := TemplateDto{}
	return &this
}

// GetLinks returns the Links field value
func (o *TemplateDto) GetLinks() map[string]ResourceLink {
	if o == nil {
		var ret map[string]ResourceLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetLinksOk() (*map[string]ResourceLink, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TemplateDto) SetLinks(v map[string]ResourceLink) {
	o.Links = v
}

// GetName returns the Name field value
func (o *TemplateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateDto) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *TemplateDto) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TemplateDto) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *TemplateDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TemplateDto) SetDescription(v string) {
	o.Description = v
}

// GetIsStarter returns the IsStarter field value if set, zero value otherwise.
func (o *TemplateDto) GetIsStarter() bool {
	if o == nil || isNil(o.IsStarter) {
		var ret bool
		return ret
	}
	return *o.IsStarter
}

// GetIsStarterOk returns a tuple with the IsStarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetIsStarterOk() (*bool, bool) {
	if o == nil || isNil(o.IsStarter) {
    return nil, false
	}
	return o.IsStarter, true
}

// HasIsStarter returns a boolean if a field has been set.
func (o *TemplateDto) HasIsStarter() bool {
	if o != nil && !isNil(o.IsStarter) {
		return true
	}

	return false
}

// SetIsStarter gets a reference to the given bool and assigns it to the IsStarter field.
func (o *TemplateDto) SetIsStarter(v bool) {
	o.IsStarter = &v
}

func (o TemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsStarter) {
		toSerialize["isStarter"] = o.IsStarter
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateDto struct {
	value *TemplateDto
	isSet bool
}

func (v NullableTemplateDto) Get() *TemplateDto {
	return v.value
}

func (v *NullableTemplateDto) Set(val *TemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDto(val *TemplateDto) *NullableTemplateDto {
	return &NullableTemplateDto{value: val, isSet: true}
}

func (v NullableTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

