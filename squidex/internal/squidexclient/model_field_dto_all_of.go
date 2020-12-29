/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// FieldDtoAllOf struct for FieldDtoAllOf
type FieldDtoAllOf struct {
	// The id of the field.
	FieldId int64 `json:"fieldId,omitempty"`
	// The name of the field. Must be unique within the schema.
	Name string `json:"name"`
	// Defines if the field is hidden.
	IsHidden bool `json:"isHidden,omitempty"`
	// Defines if the field is locked.
	IsLocked bool `json:"isLocked,omitempty"`
	// Defines if the field is disabled.
	IsDisabled bool `json:"isDisabled,omitempty"`
	// Defines the partitioning of the field.
	Partitioning string `json:"partitioning"`
	// The field properties.

	Nested *[]NestedFieldDto `json:"nested,omitempty"`
}
