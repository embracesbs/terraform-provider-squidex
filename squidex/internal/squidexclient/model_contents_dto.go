/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ContentsDto struct for ContentsDto
type ContentsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The total number of content items.
	Total int64 `json:"total,omitempty"`
	// The content items.
	Items []ContentDto `json:"items"`
	// The possible statuses.
	Statuses []StatusInfoDto `json:"statuses"`
}
