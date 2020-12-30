/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// ContributorsDtoAllOf struct for ContributorsDtoAllOf
type ContributorsDtoAllOf struct {
	// The contributors.
	Items []ContributorDto `json:"items"`
	// The maximum number of allowed contributors.
	MaxContributors int32 `json:"maxContributors,omitempty"`
	// The metadata to provide information about this request.
	Meta ContributorsMetadata `json:"_meta,omitempty"`
}
