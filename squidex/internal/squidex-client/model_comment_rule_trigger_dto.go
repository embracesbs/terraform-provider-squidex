/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
// CommentRuleTriggerDto struct for CommentRuleTriggerDto
type CommentRuleTriggerDto struct {
	TriggerType string `json:"triggerType"`
	// Javascript condition when to trigger.
	Condition *string `json:"condition,omitempty"`
}
