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

// CreateSchemaDto struct for CreateSchemaDto
type CreateSchemaDto struct {
	Properties *SchemaPropertiesDto `json:"properties,omitempty"`
	Scripts *SchemaScriptsDto `json:"scripts,omitempty"`
	FieldsInReferences []string `json:"fieldsInReferences,omitempty"`
	FieldsInLists []string `json:"fieldsInLists,omitempty"`
	// Optional fields.
	Fields []UpsertSchemaFieldDto `json:"fields,omitempty"`
	// The optional preview urls.
	PreviewUrls map[string]string `json:"previewUrls,omitempty"`
	// The optional field Rules.
	FieldRules []FieldRuleDto `json:"fieldRules,omitempty"`
	// The category.
	Category NullableString `json:"category,omitempty"`
	// Set it to true to autopublish the schema.
	IsPublished *bool `json:"isPublished,omitempty"`
	// The name of the schema.
	Name string `json:"name"`
	Type *SchemaType `json:"type,omitempty"`
	// Set to true to allow a single content item only.
	// Deprecated
	IsSingleton *bool `json:"isSingleton,omitempty"`
}

// NewCreateSchemaDto instantiates a new CreateSchemaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSchemaDto(name string) *CreateSchemaDto {
	this := CreateSchemaDto{}
	this.Name = name
	return &this
}

// NewCreateSchemaDtoWithDefaults instantiates a new CreateSchemaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSchemaDtoWithDefaults() *CreateSchemaDto {
	this := CreateSchemaDto{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetProperties() SchemaPropertiesDto {
	if o == nil || isNil(o.Properties) {
		var ret SchemaPropertiesDto
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetPropertiesOk() (*SchemaPropertiesDto, bool) {
	if o == nil || isNil(o.Properties) {
    return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given SchemaPropertiesDto and assigns it to the Properties field.
func (o *CreateSchemaDto) SetProperties(v SchemaPropertiesDto) {
	o.Properties = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetScripts() SchemaScriptsDto {
	if o == nil || isNil(o.Scripts) {
		var ret SchemaScriptsDto
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetScriptsOk() (*SchemaScriptsDto, bool) {
	if o == nil || isNil(o.Scripts) {
    return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasScripts() bool {
	if o != nil && !isNil(o.Scripts) {
		return true
	}

	return false
}

// SetScripts gets a reference to the given SchemaScriptsDto and assigns it to the Scripts field.
func (o *CreateSchemaDto) SetScripts(v SchemaScriptsDto) {
	o.Scripts = &v
}

// GetFieldsInReferences returns the FieldsInReferences field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetFieldsInReferences() []string {
	if o == nil || isNil(o.FieldsInReferences) {
		var ret []string
		return ret
	}
	return o.FieldsInReferences
}

// GetFieldsInReferencesOk returns a tuple with the FieldsInReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetFieldsInReferencesOk() ([]string, bool) {
	if o == nil || isNil(o.FieldsInReferences) {
    return nil, false
	}
	return o.FieldsInReferences, true
}

// HasFieldsInReferences returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasFieldsInReferences() bool {
	if o != nil && !isNil(o.FieldsInReferences) {
		return true
	}

	return false
}

// SetFieldsInReferences gets a reference to the given []string and assigns it to the FieldsInReferences field.
func (o *CreateSchemaDto) SetFieldsInReferences(v []string) {
	o.FieldsInReferences = v
}

// GetFieldsInLists returns the FieldsInLists field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetFieldsInLists() []string {
	if o == nil || isNil(o.FieldsInLists) {
		var ret []string
		return ret
	}
	return o.FieldsInLists
}

// GetFieldsInListsOk returns a tuple with the FieldsInLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetFieldsInListsOk() ([]string, bool) {
	if o == nil || isNil(o.FieldsInLists) {
    return nil, false
	}
	return o.FieldsInLists, true
}

// HasFieldsInLists returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasFieldsInLists() bool {
	if o != nil && !isNil(o.FieldsInLists) {
		return true
	}

	return false
}

// SetFieldsInLists gets a reference to the given []string and assigns it to the FieldsInLists field.
func (o *CreateSchemaDto) SetFieldsInLists(v []string) {
	o.FieldsInLists = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSchemaDto) GetFields() []UpsertSchemaFieldDto {
	if o == nil {
		var ret []UpsertSchemaFieldDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSchemaDto) GetFieldsOk() ([]UpsertSchemaFieldDto, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []UpsertSchemaFieldDto and assigns it to the Fields field.
func (o *CreateSchemaDto) SetFields(v []UpsertSchemaFieldDto) {
	o.Fields = v
}

// GetPreviewUrls returns the PreviewUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSchemaDto) GetPreviewUrls() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.PreviewUrls
}

// GetPreviewUrlsOk returns a tuple with the PreviewUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSchemaDto) GetPreviewUrlsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.PreviewUrls) {
    return nil, false
	}
	return &o.PreviewUrls, true
}

// HasPreviewUrls returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasPreviewUrls() bool {
	if o != nil && isNil(o.PreviewUrls) {
		return true
	}

	return false
}

// SetPreviewUrls gets a reference to the given map[string]string and assigns it to the PreviewUrls field.
func (o *CreateSchemaDto) SetPreviewUrls(v map[string]string) {
	o.PreviewUrls = v
}

// GetFieldRules returns the FieldRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSchemaDto) GetFieldRules() []FieldRuleDto {
	if o == nil {
		var ret []FieldRuleDto
		return ret
	}
	return o.FieldRules
}

// GetFieldRulesOk returns a tuple with the FieldRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSchemaDto) GetFieldRulesOk() ([]FieldRuleDto, bool) {
	if o == nil || isNil(o.FieldRules) {
    return nil, false
	}
	return o.FieldRules, true
}

// HasFieldRules returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasFieldRules() bool {
	if o != nil && isNil(o.FieldRules) {
		return true
	}

	return false
}

// SetFieldRules gets a reference to the given []FieldRuleDto and assigns it to the FieldRules field.
func (o *CreateSchemaDto) SetFieldRules(v []FieldRuleDto) {
	o.FieldRules = v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSchemaDto) GetCategory() string {
	if o == nil || isNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSchemaDto) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *CreateSchemaDto) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *CreateSchemaDto) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *CreateSchemaDto) UnsetCategory() {
	o.Category.Unset()
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetIsPublished() bool {
	if o == nil || isNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetIsPublishedOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublished) {
    return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasIsPublished() bool {
	if o != nil && !isNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *CreateSchemaDto) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetName returns the Name field value
func (o *CreateSchemaDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSchemaDto) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateSchemaDto) GetType() SchemaType {
	if o == nil || isNil(o.Type) {
		var ret SchemaType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSchemaDto) GetTypeOk() (*SchemaType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SchemaType and assigns it to the Type field.
func (o *CreateSchemaDto) SetType(v SchemaType) {
	o.Type = &v
}

// GetIsSingleton returns the IsSingleton field value if set, zero value otherwise.
// Deprecated
func (o *CreateSchemaDto) GetIsSingleton() bool {
	if o == nil || isNil(o.IsSingleton) {
		var ret bool
		return ret
	}
	return *o.IsSingleton
}

// GetIsSingletonOk returns a tuple with the IsSingleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateSchemaDto) GetIsSingletonOk() (*bool, bool) {
	if o == nil || isNil(o.IsSingleton) {
    return nil, false
	}
	return o.IsSingleton, true
}

// HasIsSingleton returns a boolean if a field has been set.
func (o *CreateSchemaDto) HasIsSingleton() bool {
	if o != nil && !isNil(o.IsSingleton) {
		return true
	}

	return false
}

// SetIsSingleton gets a reference to the given bool and assigns it to the IsSingleton field.
// Deprecated
func (o *CreateSchemaDto) SetIsSingleton(v bool) {
	o.IsSingleton = &v
}

func (o CreateSchemaDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.Scripts) {
		toSerialize["scripts"] = o.Scripts
	}
	if !isNil(o.FieldsInReferences) {
		toSerialize["fieldsInReferences"] = o.FieldsInReferences
	}
	if !isNil(o.FieldsInLists) {
		toSerialize["fieldsInLists"] = o.FieldsInLists
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.PreviewUrls != nil {
		toSerialize["previewUrls"] = o.PreviewUrls
	}
	if o.FieldRules != nil {
		toSerialize["fieldRules"] = o.FieldRules
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if !isNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.IsSingleton) {
		toSerialize["isSingleton"] = o.IsSingleton
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSchemaDto struct {
	value *CreateSchemaDto
	isSet bool
}

func (v NullableCreateSchemaDto) Get() *CreateSchemaDto {
	return v.value
}

func (v *NullableCreateSchemaDto) Set(val *CreateSchemaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSchemaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSchemaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSchemaDto(val *CreateSchemaDto) *NullableCreateSchemaDto {
	return &NullableCreateSchemaDto{value: val, isSet: true}
}

func (v NullableCreateSchemaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSchemaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

