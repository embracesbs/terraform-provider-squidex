/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// BulkUpdateAssetsJobDto struct for BulkUpdateAssetsJobDto
type BulkUpdateAssetsJobDto struct {
	// An optional id of the asset to update.
	Id   string              `json:"id,omitempty"`
	Type BulkUpdateAssetType `json:"type,omitempty"`
	// The parent folder id.
	ParentId string `json:"parentId,omitempty"`
	// The new name of the asset.
	FileName *string `json:"fileName,omitempty"`
	// The new slug of the asset.
	Slug *string `json:"slug,omitempty"`
	// True, when the asset is not public.
	IsProtected *bool `json:"isProtected,omitempty"`
	// The new asset tags.
	Tags *[]string `json:"tags,omitempty"`
	// The asset metadata.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// True to delete the asset permanently.
	Permanent bool `json:"permanent,omitempty"`
	// The expected version.
	ExpectedVersion int64 `json:"expectedVersion,omitempty"`
}
