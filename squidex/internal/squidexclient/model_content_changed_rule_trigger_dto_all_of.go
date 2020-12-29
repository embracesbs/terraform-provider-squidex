/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ContentChangedRuleTriggerDtoAllOf struct for ContentChangedRuleTriggerDtoAllOf
type ContentChangedRuleTriggerDtoAllOf struct {
	// The schema settings.
	Schemas []ContentChangedRuleTriggerSchemaDto `json:"schemas"`
	// Determines whether the trigger should handle all content changes events.
	HandleAll bool `json:"handleAll,omitempty"`
}
