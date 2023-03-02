/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// BulkUpdateDto struct for BulkUpdateDto
type BulkUpdateDto struct {
	// The contents to update or insert.
	Jobs []BulkUpdateJobDto `json:"jobs"`
	// True to automatically publish the content.
	Publish bool `json:"publish,omitempty"`
	// True to turn off scripting for faster inserts. Default: true.
	DoNotScript bool `json:"doNotScript,omitempty"`
	// True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true.
	OptimizeValidation bool `json:"optimizeValidation,omitempty"`
}