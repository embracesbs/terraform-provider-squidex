/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
// AssetsDto struct for AssetsDto
type AssetsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The total number of assets.
	Total int64 `json:"total,omitempty"`
	// The assets.
	Items []AssetDto `json:"items"`
}
