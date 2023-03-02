/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// TagsFieldPropertiesDto struct for TagsFieldPropertiesDto
type TagsFieldPropertiesDto struct {
	// Optional label for the editor.
	Label *string `json:"label,omitempty"`
	// Hints to describe the field.
	Hints *string `json:"hints,omitempty"`
	// Placeholder to show when no value has been entered.
	Placeholder *string `json:"placeholder,omitempty"`
	// Indicates if the field is required.
	IsRequired bool `json:"isRequired,omitempty"`
	// Indicates if the field is required when publishing.
	IsRequiredOnPublish bool `json:"isRequiredOnPublish,omitempty"`
	// Indicates if the field should be rendered with half width only.
	IsHalfWidth bool `json:"isHalfWidth,omitempty"`
	// Optional url to the editor.
	EditorUrl *string `json:"editorUrl,omitempty"`
	// Tags for automation processes.
	Tags          *[]string           `json:"tags,omitempty"`
	FieldType     string              `json:"fieldType"`
	DefaultValues map[string][]string `json:"defaultValues,omitempty"`
	// The default value.
	DefaultValue *[]string `json:"defaultValue,omitempty"`
	// The minimum allowed items for the field value.
	MinItems *int32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems *int32 `json:"maxItems,omitempty"`
	// The allowed values for the field value.
	AllowedValues *[]string `json:"allowedValues,omitempty"`
	// Indicates whether GraphQL Enum should be created.
	CreateEnum bool            `json:"createEnum,omitempty"`
	Editor     TagsFieldEditor `json:"editor,omitempty"`
}
