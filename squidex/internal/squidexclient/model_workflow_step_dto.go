/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// WorkflowStepDto struct for WorkflowStepDto
type WorkflowStepDto struct {
	// The transitions.
	Transitions map[string]WorkflowTransitionDto `json:"transitions"`
	// The optional color.
	Color *string `json:"color,omitempty"`
	// Indicates if updates should not be allowed.
	NoUpdate bool `json:"noUpdate,omitempty"`
	// Optional expression that must evaluate to true when you want to prevent updates.
	NoUpdateExpression *string `json:"noUpdateExpression,omitempty"`
	// Optional list of roles to restrict the updates for users with these roles.
	NoUpdateRoles *[]string `json:"noUpdateRoles,omitempty"`
}
