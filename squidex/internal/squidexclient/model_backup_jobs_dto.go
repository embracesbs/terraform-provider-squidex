/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// BackupJobsDto struct for BackupJobsDto
type BackupJobsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The backups.
	Items []BackupJobDto `json:"items"`
}
