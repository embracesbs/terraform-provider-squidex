/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// WorkflowsDto struct for WorkflowsDto
type WorkflowsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The workflow.
	Items []WorkflowDto `json:"items"`
	// The errros that should be fixed.
	Errors []string `json:"errors"`
}