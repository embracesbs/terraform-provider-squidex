/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// UpdateAppDto struct for UpdateAppDto
type UpdateAppDto struct {
	// The optional label of your app.
	Label *string `json:"label,omitempty"`
	// The optional description of your app.
	Description *string `json:"description,omitempty"`
}
