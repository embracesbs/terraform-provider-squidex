/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CreateRuleDto struct for CreateRuleDto
type CreateRuleDto struct {
	// The trigger properties.
	Trigger OneOfRuleTriggerDto `json:"trigger"`
	// The action properties.
	Action OneOfRuleAction `json:"action"`
}
