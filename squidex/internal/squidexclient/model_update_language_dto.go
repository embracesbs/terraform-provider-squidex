/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// UpdateLanguageDto struct for UpdateLanguageDto
type UpdateLanguageDto struct {
	// Set the value to true to make the language the master.
	IsMaster *bool `json:"isMaster,omitempty"`
	// Set the value to true to make the language optional.
	IsOptional bool `json:"isOptional,omitempty"`
	// Optional fallback languages.
	Fallback *[]string `json:"fallback,omitempty"`
}
