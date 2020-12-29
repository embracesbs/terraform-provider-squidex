/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CreateSchemaDto struct for CreateSchemaDto
type CreateSchemaDto struct {
	// The optional properties.
	Properties *OneOfSchemaPropertiesDto `json:"properties,omitempty"`
	// The optional scripts.
	Scripts *OneOfSchemaScriptsDto `json:"scripts,omitempty"`
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
	// The name of the schema.
	Name string `json:"name"`
	// Set to true to allow a single content item only.
	IsSingleton bool `json:"isSingleton,omitempty"`
}
