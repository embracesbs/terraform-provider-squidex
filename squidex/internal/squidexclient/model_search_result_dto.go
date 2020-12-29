/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// SearchResultDto struct for SearchResultDto
type SearchResultDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The name of the search result.
	Name string `json:"name"`
	// The type of the search result.
	Type SearchResultType `json:"type"`
	// An optional label.
	Label *string `json:"label,omitempty"`
}
