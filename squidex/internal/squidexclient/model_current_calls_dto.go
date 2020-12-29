/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CurrentCallsDto struct for CurrentCallsDto
type CurrentCallsDto struct {
	// The number of calls.
	Count int64 `json:"count,omitempty"`
	// The number of maximum allowed calls.
	MaxAllowed int64 `json:"maxAllowed,omitempty"`
}
