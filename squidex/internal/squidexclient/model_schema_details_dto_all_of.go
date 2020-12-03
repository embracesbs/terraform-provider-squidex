/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// SchemaDetailsDtoAllOf struct for SchemaDetailsDtoAllOf
type SchemaDetailsDtoAllOf struct {
	// The scripts.

	PreviewUrls map[string]string `json:"previewUrls"`
	// The name of fields that are used in content lists.

	// The list of fields.
	Fields []FieldDto `json:"fields"`
}