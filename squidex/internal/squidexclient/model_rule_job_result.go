/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// RuleJobResult the model 'RuleJobResult'
type RuleJobResult string

// List of RuleJobResult
const (
	PENDING RuleJobResult = "Pending"
	SUCCESS RuleJobResult = "Success"
	RETRY   RuleJobResult = "Retry"

	CANCELLED RuleJobResult = "Cancelled"
)
