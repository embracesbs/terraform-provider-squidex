/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// RuleEventsDto struct for RuleEventsDto
type RuleEventsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The rule events.
	Items []RuleEventDto `json:"items"`
	// The total number of rule events.
	Total int64 `json:"total,omitempty"`
}
