/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CommentsDto struct for CommentsDto
type CommentsDto struct {
	// The created comments including the updates.
	CreatedComments *[]CommentDto `json:"createdComments,omitempty"`
	// The updates comments since the last version.
	UpdatedComments *[]CommentDto `json:"updatedComments,omitempty"`
	// The deleted comments since the last version.
	DeletedComments *[]string `json:"deletedComments,omitempty"`
	// The current version.
	Version int64 `json:"version,omitempty"`
}