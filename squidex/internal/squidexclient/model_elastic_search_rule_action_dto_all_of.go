/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// ElasticSearchRuleActionDtoAllOf struct for ElasticSearchRuleActionDtoAllOf
type ElasticSearchRuleActionDtoAllOf struct {
	// The url to the elastic search instance or cluster.
	Host string `json:"host"`
	// The name of the index.
	IndexName string `json:"indexName"`
	// The optional username.
	Username *string `json:"username,omitempty"`
	// The optional password.
	Password *string `json:"password,omitempty"`
	// The optional custom document.
	Document *string `json:"document,omitempty"`
	// The condition when to delete the document.
	Delete *string `json:"delete,omitempty"`
}
