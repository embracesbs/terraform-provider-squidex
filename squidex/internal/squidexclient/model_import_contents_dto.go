/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ImportContentsDto struct for ImportContentsDto
type ImportContentsDto struct {
	// The data to import.
	Datas []NamedContentData `json:"datas"`
	// True to automatically publish the content.
	Publish bool `json:"publish,omitempty"`
	// True to turn off scripting for faster inserts. Default: true.
	DoNotScript bool `json:"doNotScript,omitempty"`
	// True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true.
	OptimizeValidation bool `json:"optimizeValidation,omitempty"`
}