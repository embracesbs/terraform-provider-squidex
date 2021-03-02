/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// ErrorDto struct for ErrorDto
type ErrorDto struct {
	// Error message.
	Message string `json:"message"`
	// The optional trace id.
	TraceId *string `json:"traceId,omitempty"`
	// Link to the error details.
	Type *string `json:"type,omitempty"`
	// Detailed error messages.
	Details *[]string `json:"details,omitempty"`
	// Status code of the http response.
	StatusCode int32 `json:"statusCode,omitempty"`
}
