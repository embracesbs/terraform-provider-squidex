/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// RuleResult the model 'RuleResult'
type RuleResult string

// List of RuleResult
const (
	RuleResultPENDING RuleResult = "Pending"
	RuleResultSUCCESS RuleResult = "Success"
	RuleResultFAILED  RuleResult = "Failed"
	RuleResultTIMEOUT RuleResult = "Timeout"
)