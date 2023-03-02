/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// AssetFoldersDto struct for AssetFoldersDto
type AssetFoldersDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The total number of assets.
	Total int64 `json:"total,omitempty"`
	// The assets folders.
	Items []AssetFolderDto `json:"items"`
	// The path to the current folder.
	Path []AssetFolderDto `json:"path"`
}