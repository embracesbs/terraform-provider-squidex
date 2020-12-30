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
// SchemaDetailsDto struct for SchemaDetailsDto
type SchemaDetailsDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The id of the schema.
	Id string `json:"id,omitempty"`
	// The name of the schema. Unique within the app.
	Name string `json:"name"`
	// The name of the category.
	Category *string `json:"category,omitempty"`
	// The schema properties.
	Properties SchemaPropertiesDto `json:"properties"`
	// Indicates if the schema is a singleton.
	IsSingleton bool `json:"isSingleton,omitempty"`
	// Indicates if the schema is published.
	IsPublished bool `json:"isPublished,omitempty"`
	// The user that has created the schema.
	CreatedBy string `json:"createdBy"`
	// The user that has updated the schema.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The date and time when the schema has been created.
	Created time.Time `json:"created,omitempty"`
	// The date and time when the schema has been modified last.
	LastModified time.Time `json:"lastModified,omitempty"`
	// The version of the schema.
	Version int64 `json:"version,omitempty"`
	// The scripts.
	Scripts SchemaScriptsDto `json:"scripts"`
	// The preview Urls.
	PreviewUrls map[string]string `json:"previewUrls"`
	// The name of fields that are used in content lists.
	FieldsInLists []string `json:"fieldsInLists"`
	// The name of fields that are used in content references.
	FieldsInReferences []string `json:"fieldsInReferences"`
	// The field rules.
	FieldRules []FieldRuleDto `json:"fieldRules,omitempty"`
	// The list of fields.
	Fields []FieldDto `json:"fields"`
}
