/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// CreateSchemaDtoAllOf struct for CreateSchemaDtoAllOf
type CreateSchemaDtoAllOf struct {
	// The name of the schema.
	Name string `json:"name"`
	// Set to true to allow a single content item only.
	IsSingleton bool `json:"isSingleton,omitempty"`
}
