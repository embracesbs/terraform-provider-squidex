/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// BulkResultDto struct for BulkResultDto
type BulkResultDto struct {
	Error ErrorDto `json:"error,omitempty"`
	// The index of the bulk job where the result belongs to. The order can change.
	JobIndex int32 `json:"jobIndex,omitempty"`
	// The id of the entity that has been handled successfully or not.
	Id *string `json:"id,omitempty"`
	// The id of the entity that has been handled successfully or not.
	ContentId *string `json:"contentId,omitempty"`
}
