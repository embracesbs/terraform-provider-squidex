/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CreateContentRuleActionDto struct for CreateContentRuleActionDto
type CreateContentRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The content data.
	Data string `json:"data"`
	// The name of the schema.
	Schema string `json:"schema"`
	// An optional client name.
	Client *string `json:"client,omitempty"`
	// Publish the content.
	Publish bool `json:"publish,omitempty"`
}