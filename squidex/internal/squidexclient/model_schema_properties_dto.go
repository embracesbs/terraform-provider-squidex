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

// SchemaPropertiesDto struct for SchemaPropertiesDto
type SchemaPropertiesDto struct {
	// Optional label for the editor.
	Label NullableString `json:"label,omitempty"`
	// Hints to describe the schema.
	Hints NullableString `json:"hints,omitempty"`
	// The url to a the sidebar plugin for content lists.
	ContentsSidebarUrl NullableString `json:"contentsSidebarUrl,omitempty"`
	// The url to a the sidebar plugin for content items.
	ContentSidebarUrl NullableString `json:"contentSidebarUrl,omitempty"`
	// The url to the editor plugin.
	ContentEditorUrl NullableString `json:"contentEditorUrl,omitempty"`
	// True to validate the content items on publish.
	ValidateOnPublish *bool `json:"validateOnPublish,omitempty"`
	// Tags for automation processes.
	Tags []string `json:"tags,omitempty"`
}

// NewSchemaPropertiesDto instantiates a new SchemaPropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaPropertiesDto() *SchemaPropertiesDto {
	this := SchemaPropertiesDto{}
	return &this
}

// NewSchemaPropertiesDtoWithDefaults instantiates a new SchemaPropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaPropertiesDtoWithDefaults() *SchemaPropertiesDto {
	this := SchemaPropertiesDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *SchemaPropertiesDto) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *SchemaPropertiesDto) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *SchemaPropertiesDto) UnsetLabel() {
	o.Label.Unset()
}

// GetHints returns the Hints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetHints() string {
	if o == nil || isNil(o.Hints.Get()) {
		var ret string
		return ret
	}
	return *o.Hints.Get()
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetHintsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hints.Get(), o.Hints.IsSet()
}

// HasHints returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasHints() bool {
	if o != nil && o.Hints.IsSet() {
		return true
	}

	return false
}

// SetHints gets a reference to the given NullableString and assigns it to the Hints field.
func (o *SchemaPropertiesDto) SetHints(v string) {
	o.Hints.Set(&v)
}
// SetHintsNil sets the value for Hints to be an explicit nil
func (o *SchemaPropertiesDto) SetHintsNil() {
	o.Hints.Set(nil)
}

// UnsetHints ensures that no value is present for Hints, not even an explicit nil
func (o *SchemaPropertiesDto) UnsetHints() {
	o.Hints.Unset()
}

// GetContentsSidebarUrl returns the ContentsSidebarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetContentsSidebarUrl() string {
	if o == nil || isNil(o.ContentsSidebarUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ContentsSidebarUrl.Get()
}

// GetContentsSidebarUrlOk returns a tuple with the ContentsSidebarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetContentsSidebarUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContentsSidebarUrl.Get(), o.ContentsSidebarUrl.IsSet()
}

// HasContentsSidebarUrl returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasContentsSidebarUrl() bool {
	if o != nil && o.ContentsSidebarUrl.IsSet() {
		return true
	}

	return false
}

// SetContentsSidebarUrl gets a reference to the given NullableString and assigns it to the ContentsSidebarUrl field.
func (o *SchemaPropertiesDto) SetContentsSidebarUrl(v string) {
	o.ContentsSidebarUrl.Set(&v)
}
// SetContentsSidebarUrlNil sets the value for ContentsSidebarUrl to be an explicit nil
func (o *SchemaPropertiesDto) SetContentsSidebarUrlNil() {
	o.ContentsSidebarUrl.Set(nil)
}

// UnsetContentsSidebarUrl ensures that no value is present for ContentsSidebarUrl, not even an explicit nil
func (o *SchemaPropertiesDto) UnsetContentsSidebarUrl() {
	o.ContentsSidebarUrl.Unset()
}

// GetContentSidebarUrl returns the ContentSidebarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetContentSidebarUrl() string {
	if o == nil || isNil(o.ContentSidebarUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ContentSidebarUrl.Get()
}

// GetContentSidebarUrlOk returns a tuple with the ContentSidebarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetContentSidebarUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContentSidebarUrl.Get(), o.ContentSidebarUrl.IsSet()
}

// HasContentSidebarUrl returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasContentSidebarUrl() bool {
	if o != nil && o.ContentSidebarUrl.IsSet() {
		return true
	}

	return false
}

// SetContentSidebarUrl gets a reference to the given NullableString and assigns it to the ContentSidebarUrl field.
func (o *SchemaPropertiesDto) SetContentSidebarUrl(v string) {
	o.ContentSidebarUrl.Set(&v)
}
// SetContentSidebarUrlNil sets the value for ContentSidebarUrl to be an explicit nil
func (o *SchemaPropertiesDto) SetContentSidebarUrlNil() {
	o.ContentSidebarUrl.Set(nil)
}

// UnsetContentSidebarUrl ensures that no value is present for ContentSidebarUrl, not even an explicit nil
func (o *SchemaPropertiesDto) UnsetContentSidebarUrl() {
	o.ContentSidebarUrl.Unset()
}

// GetContentEditorUrl returns the ContentEditorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetContentEditorUrl() string {
	if o == nil || isNil(o.ContentEditorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ContentEditorUrl.Get()
}

// GetContentEditorUrlOk returns a tuple with the ContentEditorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetContentEditorUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContentEditorUrl.Get(), o.ContentEditorUrl.IsSet()
}

// HasContentEditorUrl returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasContentEditorUrl() bool {
	if o != nil && o.ContentEditorUrl.IsSet() {
		return true
	}

	return false
}

// SetContentEditorUrl gets a reference to the given NullableString and assigns it to the ContentEditorUrl field.
func (o *SchemaPropertiesDto) SetContentEditorUrl(v string) {
	o.ContentEditorUrl.Set(&v)
}
// SetContentEditorUrlNil sets the value for ContentEditorUrl to be an explicit nil
func (o *SchemaPropertiesDto) SetContentEditorUrlNil() {
	o.ContentEditorUrl.Set(nil)
}

// UnsetContentEditorUrl ensures that no value is present for ContentEditorUrl, not even an explicit nil
func (o *SchemaPropertiesDto) UnsetContentEditorUrl() {
	o.ContentEditorUrl.Unset()
}

// GetValidateOnPublish returns the ValidateOnPublish field value if set, zero value otherwise.
func (o *SchemaPropertiesDto) GetValidateOnPublish() bool {
	if o == nil || isNil(o.ValidateOnPublish) {
		var ret bool
		return ret
	}
	return *o.ValidateOnPublish
}

// GetValidateOnPublishOk returns a tuple with the ValidateOnPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaPropertiesDto) GetValidateOnPublishOk() (*bool, bool) {
	if o == nil || isNil(o.ValidateOnPublish) {
    return nil, false
	}
	return o.ValidateOnPublish, true
}

// HasValidateOnPublish returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasValidateOnPublish() bool {
	if o != nil && !isNil(o.ValidateOnPublish) {
		return true
	}

	return false
}

// SetValidateOnPublish gets a reference to the given bool and assigns it to the ValidateOnPublish field.
func (o *SchemaPropertiesDto) SetValidateOnPublish(v bool) {
	o.ValidateOnPublish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaPropertiesDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaPropertiesDto) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SchemaPropertiesDto) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SchemaPropertiesDto) SetTags(v []string) {
	o.Tags = v
}

func (o SchemaPropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Hints.IsSet() {
		toSerialize["hints"] = o.Hints.Get()
	}
	if o.ContentsSidebarUrl.IsSet() {
		toSerialize["contentsSidebarUrl"] = o.ContentsSidebarUrl.Get()
	}
	if o.ContentSidebarUrl.IsSet() {
		toSerialize["contentSidebarUrl"] = o.ContentSidebarUrl.Get()
	}
	if o.ContentEditorUrl.IsSet() {
		toSerialize["contentEditorUrl"] = o.ContentEditorUrl.Get()
	}
	if !isNil(o.ValidateOnPublish) {
		toSerialize["validateOnPublish"] = o.ValidateOnPublish
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaPropertiesDto struct {
	value *SchemaPropertiesDto
	isSet bool
}

func (v NullableSchemaPropertiesDto) Get() *SchemaPropertiesDto {
	return v.value
}

func (v *NullableSchemaPropertiesDto) Set(val *SchemaPropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaPropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaPropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaPropertiesDto(val *SchemaPropertiesDto) *NullableSchemaPropertiesDto {
	return &NullableSchemaPropertiesDto{value: val, isSet: true}
}

func (v NullableSchemaPropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaPropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


