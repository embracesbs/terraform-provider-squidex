/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// WorkflowsDtoAllOf struct for WorkflowsDtoAllOf
type WorkflowsDtoAllOf struct {
	// The workflow.
	Items []WorkflowDto `json:"items"`
	// The errros that should be fixed.
	Errors []string `json:"errors"`
}
