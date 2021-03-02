/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// StringFieldPropertiesDto struct for StringFieldPropertiesDto
type StringFieldPropertiesDto struct {
	// Optional label for the editor.
	Label *string `json:"label,omitempty"`
	// Hints to describe the schema.
	Hints *string `json:"hints,omitempty"`
	// Placeholder to show when no value has been entered.
	Placeholder *string `json:"placeholder,omitempty"`
	// Indicates if the field is required.
	IsRequired bool `json:"isRequired,omitempty"`
	// Indicates if the field should be rendered with half width only.
	IsHalfWidth bool `json:"isHalfWidth,omitempty"`
	// Optional url to the editor.
	EditorUrl *string `json:"editorUrl,omitempty"`
	// Tags for automation processes.
	Tags      *[]string `json:"tags,omitempty"`
	FieldType string    `json:"fieldType"`
	// The default value for the field value.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The pattern to enforce a specific format for the field value.
	Pattern *string `json:"pattern,omitempty"`
	// The validation message for the pattern.
	PatternMessage *string `json:"patternMessage,omitempty"`
	// The minimum allowed length for the field value.
	MinLength *int32 `json:"minLength,omitempty"`
	// The maximum allowed length for the field value.
	MaxLength *int32 `json:"maxLength,omitempty"`
	// The minimum allowed of normal characters for the field value.
	MinCharacters *int32 `json:"minCharacters,omitempty"`
	// The maximum allowed of normal characters for the field value.
	MaxCharacters *int32 `json:"maxCharacters,omitempty"`
	// The minimum allowed number of words for the field value.
	MinWords *int32 `json:"minWords,omitempty"`
	// The maximum allowed number of words for the field value.
	MaxWords *int32 `json:"maxWords,omitempty"`
	// The allowed values for the field value.
	AllowedValues *[]string `json:"allowedValues,omitempty"`
	// Indicates if the field value must be unique. Ignored for nested fields and localized fields.
	IsUnique bool `json:"isUnique,omitempty"`
	// Indicates that the inline editor is enabled for this field.
	InlineEditable bool `json:"inlineEditable,omitempty"`
	// How the string content should be interpreted.
	ContentType StringContentType `json:"contentType,omitempty"`
	// The editor that is used to manage this field.
	Editor StringFieldEditor `json:"editor,omitempty"`
}
