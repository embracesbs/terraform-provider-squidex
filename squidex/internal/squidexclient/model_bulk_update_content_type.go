/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// BulkUpdateContentType the model 'BulkUpdateContentType'
type BulkUpdateContentType string

// List of BulkUpdateContentType
const (
	BulkUpdateContentTypeUPSERT        BulkUpdateContentType = "Upsert"
	BulkUpdateContentTypeCHANGE_STATUS BulkUpdateContentType = "ChangeStatus"
	BulkUpdateContentTypeCREATE        BulkUpdateContentType = "Create"
	BulkUpdateContentTypeDELETE        BulkUpdateContentType = "Delete"
	BulkUpdateContentTypePATCH         BulkUpdateContentType = "Patch"
	BulkUpdateContentTypeUPDATE        BulkUpdateContentType = "Update"
	BulkUpdateContentTypeVALIDATE      BulkUpdateContentType = "Validate"
)
