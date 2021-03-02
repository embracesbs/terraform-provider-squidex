/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UpsertSchemaFieldDto struct for UpsertSchemaFieldDto
type UpsertSchemaFieldDto struct {
	// The name of the field. Must be unique within the schema.
	Name string `json:"name"`
	// Defines if the field is hidden.
	IsHidden bool `json:"isHidden,omitempty"`
	// Defines if the field is locked.
	IsLocked bool `json:"isLocked,omitempty"`
	// Defines if the field is disabled.
	IsDisabled bool `json:"isDisabled,omitempty"`
	// Determines the optional partitioning of the field.
	Partitioning *string `json:"partitioning,omitempty"`
	// The field properties.
	Properties FieldPropertiesDto `json:"properties"`
	// The nested fields.
	Nested *[]UpsertSchemaNestedFieldDto `json:"nested,omitempty"`
}
