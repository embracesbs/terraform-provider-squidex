/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// CommentRuleActionDto struct for CommentRuleActionDto
type CommentRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The comment text.
	Text string `json:"text"`
	// An optional client name.
	Client *string `json:"client,omitempty"`
}
