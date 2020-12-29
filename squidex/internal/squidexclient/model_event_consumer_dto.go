/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// EventConsumerDto struct for EventConsumerDto
type EventConsumerDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	IsStopped bool `json:"isStopped,omitempty"`
	IsResetting bool `json:"isResetting,omitempty"`
	Name string `json:"name,omitempty"`
	Error *string `json:"error,omitempty"`
	Position *string `json:"position,omitempty"`
}
