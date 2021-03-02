/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UiFieldPropertiesDto struct for UiFieldPropertiesDto
type UiFieldPropertiesDto struct {
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
	// The editor that is used to manage this field.
	Editor UiFieldEditor `json:"editor,omitempty"`
}
