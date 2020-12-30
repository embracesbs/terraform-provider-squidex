/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
import (
	"time"
)
// BackupJobDto struct for BackupJobDto
type BackupJobDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the backup job.
	Id string `json:"id,omitempty"`
	// The time when the job has been started.
	Started time.Time `json:"started,omitempty"`
	// The time when the job has been stopped.
	Stopped *time.Time `json:"stopped,omitempty"`
	// The number of handled events.
	HandledEvents int32 `json:"handledEvents,omitempty"`
	// The number of handled assets.
	HandledAssets int32 `json:"handledAssets,omitempty"`
	// The status of the operation.
	Status JobStatus `json:"status,omitempty"`
}
