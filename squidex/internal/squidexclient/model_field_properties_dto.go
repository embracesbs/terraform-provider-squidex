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

// FieldPropertiesDto struct for FieldPropertiesDto
type FieldPropertiesDto struct {
	
	// Optional label for the editor.
	Label NullableString `json:"label,omitempty"`
	// Hints to describe the field.
	Hints NullableString `json:"hints,omitempty"`
	// Placeholder to show when no value has been entered.
	Placeholder NullableString `json:"placeholder,omitempty"`
	// Indicates if the field is required.
	IsRequired *bool `json:"isRequired,omitempty"`
	// Indicates if the field is required when publishing.
	IsRequiredOnPublish *bool `json:"isRequiredOnPublish,omitempty"`
	// Indicates if the field should be rendered with half width only.
	IsHalfWidth *bool `json:"isHalfWidth,omitempty"`
	// Optional url to the editor.
	EditorUrl NullableString `json:"editorUrl,omitempty"`
	// Tags for automation processes.
	Tags []string `json:"tags,omitempty"`
	FieldType string `json:"fieldType"`

	// Editor
	Editor *string `json:"editor,omitempty"`
	// MinItems The minimum allowed items for the field value.
	MinItems *int `json:"minItems,omitempty"`
	// MaxItems The maximum allowed items for the field value.
	MaxItems *int `json:"maxItems,omitempty"`
	// PreviewMode The preview mode for the asset.
	/*	Options:
		- asset(ImageAndFileName Image FileName)
	*/
	PreviewMode *string `json:"previewMode,omitempty"`
	/*	DefaultValues The language specific default value as a list of asset ids.
		Options:
		- lists: {"nl-NL": ["asset1","tag2"]}
		- string: {"nl-NL": "asset1"}
		- int/double: {"nl-NL": 10.9}
		- bool: {"nl-NL": true}
	*/
	DefaultValues *map[string]interface{} `json:"defaultValues,omitempty"`
	/*	DefaultValue The default value for the field value.
		Options:
		- lists: ["asset1","tag2"]
		- string: "asset1"
		- int/double: 10.9
		- bool: true
	*/
	DefaultValue *interface{} `json:"defaultValue,omitempty"`
	// MinSize The minimum file size in bytes.
	MinSize *int `json:"minSize,omitempty"`
	// MaxSize The maximum file size in bytes.
	MaxSize *int `json:"maxSize,omitempty"`
	// MinWidth The minimum image width in pixels.
	MinWidth *int `json:"minWidth,omitempty"`
	// MaxWidth The maximum image width in pixels.
	MaxWidth *int `json:"maxWidth,omitempty"`
	// MinHeight The minimum image height in pixels.
	MinHeight *int `json:"minHeight,omitempty"`
	// MaxHeight The maximum image height in pixels.
	MaxHeight *int `json:"maxHeight,omitempty"`
	// AspectWidth The image aspect width in pixels.
	AspectWidth *int `json:"aspectWidth,omitempty"`
	// AspectHeight The image aspect height in pixels.
	AspectHeight *int `json:"aspectHeight,omitempty"`
	// MustBeImage Defines if the asset must be an image.
	MustBeImage *bool `json:"mustBeImage,omitempty"`
	// ResolveFirst True to resolve first asset in the content list.
	ResolveFirst *bool `json:"resolveFirst,omitempty"`
	// AllowedExtensions The allowed file extensions.
	AllowedExtensions *[]string `json:"allowedExtensions,omitempty"`
	// AllowDuplicates True, if duplicate values are allowed.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// InlineEditable Indicates that the inline editor is enabled for this field.
	InlineEditable *bool `json:"inlineEditable,omitempty"`
	/*	MaxValue The maximum allowed value for the field value.
		Options:
		string
		float32
	*/
	MaxValue *interface{} `json:"maxValue,omitempty"`
	/* MinValue The minimum allowed value for the field value.
	Options:
	string
	float32
	*/
	MinValue *interface{} `json:"minValue,omitempty"`
	// CalculatedDefaultValue The calculated default value for the field value.
	// Options: Now Today (or leave empty for none)
	CalculatedDefaultValue *string `json:"calculatedDefaultValue,omitempty"`
	/*	AllowedValues The allowed values for the field value.
			Options:
		 	[]string
		 	[]float32
	*/
	AllowedValues *[]interface{} `json:"allowedValues,omitempty"`
	// Unique Indicates if the field value must be unique. Ignored for nested fields and localized fields.
	Unique *bool `json:"isUnique,omitempty"`
	// ResolveReference True to resolve references in the content list. (Required for fieldtype References)
	ResolveReference *bool `json:"resolveReference,omitempty"`
	// MustBePublished True when all references must be published. (Required for fieldtype References)
	MustBePublished *bool `json:"mustBePublished,omitempty"`
	// schemaIds The id of the referenced schemas.
	SchemaIds *[]string `json:"schemaIds,omitempty"`
	// Pattern The pattern to enforce a specific format for the field value.
	Pattern *string `json:"pattern,omitempty"`
	// PatternMessage The validation message for the pattern.
	PatternMessage *string `json:"patternMessage,omitempty"`
	// MinLength The minimum allowed length for the field value.
	MinLength *int `json:"minLength,omitempty"`
	// MaxLength The maximum allowed length for the field value.
	MaxLength *int `json:"maxLength,omitempty"`
	// MinCharacters The minimum allowed of normal characters for the field value.
	MinCharacters *int `json:"minCharacters,omitempty"`
	// MaxCharacters The maximum allowed of normal characters for the field value.
	MaxCharacters *int `json:"maxCharacters,omitempty"`
	// MinWords The minimum allowed number of words for the field value.
	MinWords *int `json:"minWords,omitempty"`
	// MaxWords The maximum allowed number of words for the field value.
	MaxWords *int `json:"maxWords,omitempty"`
	/*	ContentType How the string content should be interpreted. (Required for fieldtype string)
		Options:
		Unspecified
		Markdown
		Html
	*/
	ContentType *string `json:"contentType,omitempty"`
}

// NewFieldPropertiesDto instantiates a new FieldPropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldPropertiesDto(fieldType string) *FieldPropertiesDto {
	this := FieldPropertiesDto{}
	this.FieldType = fieldType
	return &this
}

// NewFieldPropertiesDtoWithDefaults instantiates a new FieldPropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldPropertiesDtoWithDefaults() *FieldPropertiesDto {
	this := FieldPropertiesDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldPropertiesDto) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldPropertiesDto) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *FieldPropertiesDto) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *FieldPropertiesDto) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *FieldPropertiesDto) UnsetLabel() {
	o.Label.Unset()
}

// GetHints returns the Hints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldPropertiesDto) GetHints() string {
	if o == nil || isNil(o.Hints.Get()) {
		var ret string
		return ret
	}
	return *o.Hints.Get()
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldPropertiesDto) GetHintsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hints.Get(), o.Hints.IsSet()
}

// HasHints returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasHints() bool {
	if o != nil && o.Hints.IsSet() {
		return true
	}

	return false
}

// SetHints gets a reference to the given NullableString and assigns it to the Hints field.
func (o *FieldPropertiesDto) SetHints(v string) {
	o.Hints.Set(&v)
}
// SetHintsNil sets the value for Hints to be an explicit nil
func (o *FieldPropertiesDto) SetHintsNil() {
	o.Hints.Set(nil)
}

// UnsetHints ensures that no value is present for Hints, not even an explicit nil
func (o *FieldPropertiesDto) UnsetHints() {
	o.Hints.Unset()
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldPropertiesDto) GetPlaceholder() string {
	if o == nil || isNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldPropertiesDto) GetPlaceholderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *FieldPropertiesDto) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *FieldPropertiesDto) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *FieldPropertiesDto) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *FieldPropertiesDto) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPropertiesDto) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
    return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *FieldPropertiesDto) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsRequiredOnPublish returns the IsRequiredOnPublish field value if set, zero value otherwise.
func (o *FieldPropertiesDto) GetIsRequiredOnPublish() bool {
	if o == nil || isNil(o.IsRequiredOnPublish) {
		var ret bool
		return ret
	}
	return *o.IsRequiredOnPublish
}

// GetIsRequiredOnPublishOk returns a tuple with the IsRequiredOnPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPropertiesDto) GetIsRequiredOnPublishOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequiredOnPublish) {
    return nil, false
	}
	return o.IsRequiredOnPublish, true
}

// HasIsRequiredOnPublish returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasIsRequiredOnPublish() bool {
	if o != nil && !isNil(o.IsRequiredOnPublish) {
		return true
	}

	return false
}

// SetIsRequiredOnPublish gets a reference to the given bool and assigns it to the IsRequiredOnPublish field.
func (o *FieldPropertiesDto) SetIsRequiredOnPublish(v bool) {
	o.IsRequiredOnPublish = &v
}

// GetIsHalfWidth returns the IsHalfWidth field value if set, zero value otherwise.
func (o *FieldPropertiesDto) GetIsHalfWidth() bool {
	if o == nil || isNil(o.IsHalfWidth) {
		var ret bool
		return ret
	}
	return *o.IsHalfWidth
}

// GetIsHalfWidthOk returns a tuple with the IsHalfWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPropertiesDto) GetIsHalfWidthOk() (*bool, bool) {
	if o == nil || isNil(o.IsHalfWidth) {
    return nil, false
	}
	return o.IsHalfWidth, true
}

// HasIsHalfWidth returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasIsHalfWidth() bool {
	if o != nil && !isNil(o.IsHalfWidth) {
		return true
	}

	return false
}

// SetIsHalfWidth gets a reference to the given bool and assigns it to the IsHalfWidth field.
func (o *FieldPropertiesDto) SetIsHalfWidth(v bool) {
	o.IsHalfWidth = &v
}

// GetEditorUrl returns the EditorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldPropertiesDto) GetEditorUrl() string {
	if o == nil || isNil(o.EditorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.EditorUrl.Get()
}

// GetEditorUrlOk returns a tuple with the EditorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldPropertiesDto) GetEditorUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EditorUrl.Get(), o.EditorUrl.IsSet()
}

// HasEditorUrl returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasEditorUrl() bool {
	if o != nil && o.EditorUrl.IsSet() {
		return true
	}

	return false
}

// SetEditorUrl gets a reference to the given NullableString and assigns it to the EditorUrl field.
func (o *FieldPropertiesDto) SetEditorUrl(v string) {
	o.EditorUrl.Set(&v)
}
// SetEditorUrlNil sets the value for EditorUrl to be an explicit nil
func (o *FieldPropertiesDto) SetEditorUrlNil() {
	o.EditorUrl.Set(nil)
}

// UnsetEditorUrl ensures that no value is present for EditorUrl, not even an explicit nil
func (o *FieldPropertiesDto) UnsetEditorUrl() {
	o.EditorUrl.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldPropertiesDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldPropertiesDto) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FieldPropertiesDto) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FieldPropertiesDto) SetTags(v []string) {
	o.Tags = v
}

// GetFieldType returns the FieldType field value
func (o *FieldPropertiesDto) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *FieldPropertiesDto) GetFieldTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *FieldPropertiesDto) SetFieldType(v string) {
	o.FieldType = v
}

func (o FieldPropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Hints.IsSet() {
		toSerialize["hints"] = o.Hints.Get()
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !isNil(o.IsRequiredOnPublish) {
		toSerialize["isRequiredOnPublish"] = o.IsRequiredOnPublish
	}
	if !isNil(o.IsHalfWidth) {
		toSerialize["isHalfWidth"] = o.IsHalfWidth
	}
	if o.EditorUrl.IsSet() {
		toSerialize["editorUrl"] = o.EditorUrl.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["fieldType"] = o.FieldType
	}
	return json.Marshal(toSerialize)
}

type NullableFieldPropertiesDto struct {
	value *FieldPropertiesDto
	isSet bool
}

func (v NullableFieldPropertiesDto) Get() *FieldPropertiesDto {
	return v.value
}

func (v *NullableFieldPropertiesDto) Set(val *FieldPropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldPropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldPropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldPropertiesDto(val *FieldPropertiesDto) *NullableFieldPropertiesDto {
	return &NullableFieldPropertiesDto{value: val, isSet: true}
}

func (v NullableFieldPropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldPropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


