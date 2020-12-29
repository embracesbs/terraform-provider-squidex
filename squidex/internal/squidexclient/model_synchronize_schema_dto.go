/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// SynchronizeSchemaDto struct for SynchronizeSchemaDto
type SynchronizeSchemaDto struct {
	// The optional properties.
	Properties *SchemaPropertiesDto `json:"properties,omitempty"`
	// The optional scripts.
	Scripts *SchemaScriptsDto `json:"scripts,omitempty"`
	// The names of the fields that should be used in references.
	FieldsInReferences *[]string `json:"fieldsInReferences,omitempty"`
	// The names of the fields that should be shown in lists, including meta fields.
	FieldsInLists *[]string `json:"fieldsInLists,omitempty"`
	// Optional fields.
	Fields *[]UpsertSchemaFieldDto `json:"fields,omitempty"`
	// The optional preview urls.
	PreviewUrls *map[string]string `json:"previewUrls,omitempty"`
	// The category.
	Category *string `json:"category,omitempty"`
	// Set it to true to autopublish the schema.
	IsPublished bool `json:"isPublished,omitempty"`
	// True, when fields should not be deleted.
	NoFieldDeletion bool `json:"noFieldDeletion,omitempty"`
	// True, when fields with different types should not be recreated.
	NoFieldRecreation bool `json:"noFieldRecreation,omitempty"`
}
