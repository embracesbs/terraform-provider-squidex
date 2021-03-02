/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// CreateAppDto struct for CreateAppDto
type CreateAppDto struct {
	// The name of the app.
	Name string `json:"name"`
	// Initialize the app with the inbuilt template.
	Template *string `json:"template,omitempty"`
}
