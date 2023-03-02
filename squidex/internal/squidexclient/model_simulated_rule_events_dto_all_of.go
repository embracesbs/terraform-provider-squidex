/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// SimulatedRuleEventsDtoAllOf struct for SimulatedRuleEventsDtoAllOf
type SimulatedRuleEventsDtoAllOf struct {
	// The simulated rule events.
	Items []SimulatedRuleEventDto `json:"items"`
	// The total number of simulated rule events.
	Total int64 `json:"total,omitempty"`
}