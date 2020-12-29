/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// AzureQueueRuleActionDto struct for AzureQueueRuleActionDto
type AzureQueueRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The connection string to the storage account.
	ConnectionString string `json:"connectionString"`
	// The name of the queue.
	Queue string `json:"queue"`
}
