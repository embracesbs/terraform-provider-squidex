/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// GraphQlPostDto struct for GraphQlPostDto
type GraphQlPostDto struct {
	OperationName string      `json:"operationName,omitempty"`
	Query         string      `json:"query,omitempty"`
	Variables     interface{} `json:"variables,omitempty"`
}
