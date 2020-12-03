/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
import (
	"time"
)
// ContentDto struct for ContentDto
type ContentDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The if of the content item.
	Id string `json:"id,omitempty"`
	// The user that has created the content item.
	CreatedBy string `json:"createdBy"`
	// The user that has updated the content item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The data of the content item.
	Data interface{} `json:"data"`
	// The pending changes of the content item.
	DataDraft *interface{} `json:"dataDraft,omitempty"`
	// The reference data for the frontend UI.
	ReferenceData *OneOfNamedContentData `json:"referenceData,omitempty"`
	// Indicates if the draft data is pending.
	IsPending bool `json:"isPending,omitempty"`
	// The scheduled status.
	ScheduleJob *OneOfScheduleJobDto `json:"scheduleJob,omitempty"`
	// The date and time when the content item has been created.
	Created time.Time `json:"created,omitempty"`
	// The date and time when the content item has been modified last.
	LastModified time.Time `json:"lastModified,omitempty"`
	// The status of the content.
	Status string `json:"status,omitempty"`
	// The color of the status.
	StatusColor string `json:"statusColor,omitempty"`
	// The name of the schema.
	SchemaName *string `json:"schemaName,omitempty"`
	// The display name of the schema.
	SchemaDisplayName *string `json:"schemaDisplayName,omitempty"`
	// The reference fields.
	ReferenceFields *[]FieldDto `json:"referenceFields,omitempty"`
	// The version of the content.
	Version int64 `json:"version,omitempty"`
}
