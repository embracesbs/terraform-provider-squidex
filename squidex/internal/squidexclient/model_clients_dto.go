/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ClientsDto struct for ClientsDto
type ClientsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The clients.
	Items []ClientDto `json:"items"`
}
