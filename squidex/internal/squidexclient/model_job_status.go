/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// JobStatus the model 'JobStatus'
type JobStatus string

// List of JobStatus
const (
	CREATED JobStatus = "Created"
	STARTED JobStatus = "Started"
	COMPLETED JobStatus = "Completed"
	FAILED JobStatus = "Failed"
)
