/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UpsertCommentDto struct for UpsertCommentDto
type UpsertCommentDto struct {
	// The comment text.
	Text string `json:"text"`
	// The url where the comment is created.
	Url *string `json:"url,omitempty"`
}