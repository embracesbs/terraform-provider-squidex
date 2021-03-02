/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// TagsFieldPropertiesDtoAllOf struct for TagsFieldPropertiesDtoAllOf
type TagsFieldPropertiesDtoAllOf struct {
	// The minimum allowed items for the field value.
	MinItems *int32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems *int32 `json:"maxItems,omitempty"`
	// The allowed values for the field value.
	AllowedValues *[]string `json:"allowedValues,omitempty"`
	// The editor that is used to manage this field.
	Editor TagsFieldEditor `json:"editor,omitempty"`
}
