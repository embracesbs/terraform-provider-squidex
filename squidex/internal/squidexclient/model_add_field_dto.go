/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// AddFieldDto struct for AddFieldDto
type AddFieldDto struct {
	// The name of the field. Must be unique within the schema.
	Name string `json:"name"`
	// Determines the optional partitioning of the field.
	Partitioning *string `json:"partitioning,omitempty"`
	// The field properties.
	Properties FieldPropertiesDto `json:"properties"`
}
