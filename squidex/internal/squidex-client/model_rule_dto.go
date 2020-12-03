/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
import (
	"time"
)
// RuleDto struct for RuleDto
type RuleDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the rule.
	Id string `json:"id,omitempty"`
	// The user that has created the rule.
	CreatedBy string `json:"createdBy"`
	// The user that has updated the rule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The date and time when the rule has been created.
	Created time.Time `json:"created,omitempty"`
	// The date and time when the rule has been modified last.
	LastModified time.Time `json:"lastModified,omitempty"`
	// The version of the rule.
	Version int64 `json:"version,omitempty"`
	// Determines if the rule is enabled.
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Optional rule name.
	Name *string `json:"name,omitempty"`
	// The trigger properties.
	Trigger OneOfRuleTriggerDto `json:"trigger"`
	// The action properties.
	Action OneOfRuleAction `json:"action"`
	// The number of completed executions.
	NumSucceeded int32 `json:"numSucceeded,omitempty"`
	// The number of failed executions.
	NumFailed int32 `json:"numFailed,omitempty"`
	// The date and time when the rule was executed the last time.
	LastExecuted *time.Time `json:"lastExecuted,omitempty"`
}
