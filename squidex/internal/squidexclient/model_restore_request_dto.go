/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// RestoreRequestDto struct for RestoreRequestDto
type RestoreRequestDto struct {
	// The name of the app.
	Name string `json:"name,omitempty"`
	// The url to the restore file.
	Url string `json:"url"`
}
