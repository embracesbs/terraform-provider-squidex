/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// ReorderFieldsDto struct for ReorderFieldsDto
type ReorderFieldsDto struct {
	// The field ids in the target order.
	FieldIds []int64 `json:"fieldIds"`
}