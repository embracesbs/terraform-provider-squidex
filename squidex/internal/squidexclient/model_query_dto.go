/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// QueryDto struct for QueryDto
type QueryDto struct {
	// The optional list of ids to query.
	Ids *[]string `json:"ids,omitempty"`
	// The optional odata query.
	OData *string `json:"oData,omitempty"`
	// The optional json query.
	Q *interface{} `json:"q,omitempty"`
	// The parent id (for assets).
	ParentId *string `json:"parentId,omitempty"`
}
