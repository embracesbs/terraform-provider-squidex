/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

import (
	"time"
)

// SchemaDtoAllOf struct for SchemaDtoAllOf
type SchemaDtoAllOf struct {
	// The id of the schema.
	Id string `json:"id,omitempty"`
	// The user that has created the schema.
	CreatedBy string `json:"createdBy"`
	// The user that has updated the schema.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the schema. Unique within the app.
	Name string     `json:"name"`
	Type SchemaType `json:"type,omitempty"`
	// The name of the category.
	Category   *string             `json:"category,omitempty"`
	Properties SchemaPropertiesDto `json:"properties"`
	// Indicates if the schema is a singleton.
	IsSingleton bool `json:"isSingleton,omitempty"`
	// Indicates if the schema is published.
	IsPublished bool `json:"isPublished,omitempty"`
	// The date and time when the schema has been created.
	Created time.Time `json:"created,omitempty"`
	// The date and time when the schema has been modified last.
	LastModified time.Time `json:"lastModified,omitempty"`
	// The version of the schema.
	Version int64            `json:"version,omitempty"`
	Scripts SchemaScriptsDto `json:"scripts"`
	// The preview Urls.
	PreviewUrls        map[string]string `json:"previewUrls"`
	FieldsInLists      []string          `json:"fieldsInLists"`
	FieldsInReferences []string          `json:"fieldsInReferences"`
	// The field rules.
	FieldRules []FieldRuleDto `json:"fieldRules,omitempty"`
	// The list of fields.
	Fields []FieldDto `json:"fields"`
}
