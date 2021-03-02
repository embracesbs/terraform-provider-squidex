/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UsageRuleTriggerDto struct for UsageRuleTriggerDto
type UsageRuleTriggerDto struct {
	TriggerType string `json:"triggerType"`
	// The number of monthly api calls.
	Limit int32 `json:"limit,omitempty"`
	// The number of days to check or null for the current month.
	NumDays *int32 `json:"numDays,omitempty"`
}
