/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ContentsDtoAllOf struct for ContentsDtoAllOf
type ContentsDtoAllOf struct {
	// The total number of content items.
	Total int64 `json:"total,omitempty"`
	// The content items.
	Items []ContentDto `json:"items"`
	// The possible statuses.
	Statuses []StatusInfoDto `json:"statuses"`
}
