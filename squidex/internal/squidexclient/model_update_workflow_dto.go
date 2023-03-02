/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UpdateWorkflowDto struct for UpdateWorkflowDto
type UpdateWorkflowDto struct {
	// The name of the workflow.
	Name *string `json:"name,omitempty"`
	// The workflow steps.
	Steps map[string]WorkflowStepDto `json:"steps"`
	// The schema ids.
	SchemaIds *[]string `json:"schemaIds,omitempty"`
	// The initial step.
	Initial string `json:"initial"`
}
