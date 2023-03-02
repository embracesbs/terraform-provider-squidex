/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// ComponentsFieldPropertiesDtoAllOf struct for ComponentsFieldPropertiesDtoAllOf
type ComponentsFieldPropertiesDtoAllOf struct {
	// The minimum allowed items for the field value.
	MinItems *int32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems *int32 `json:"maxItems,omitempty"`
	// The id of the embedded schemas.
	SchemaIds *[]string `json:"schemaIds,omitempty"`
	// The fields that must be unique.
	UniqueFields *[]string `json:"uniqueFields,omitempty"`
}
