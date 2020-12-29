/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// ReferencesFieldPropertiesDtoAllOf struct for ReferencesFieldPropertiesDtoAllOf
type ReferencesFieldPropertiesDtoAllOf struct {
	// The minimum allowed items for the field value.
	MinItems *int32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems *int32 `json:"maxItems,omitempty"`
	// True, if duplicate values are allowed.
	AllowDuplicates bool `json:"allowDuplicates,omitempty"`
	// True to resolve references in the content list.
	ResolveReference bool `json:"resolveReference,omitempty"`
	// The editor that is used to manage this field.

	// The id of the referenced schemas.
	SchemaIds *[]string `json:"schemaIds,omitempty"`
}
