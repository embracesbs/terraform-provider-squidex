/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// PrerenderRuleActionDtoAllOf struct for PrerenderRuleActionDtoAllOf
type PrerenderRuleActionDtoAllOf struct {
	// The prerender token from your account.
	Token string `json:"token"`
	// The url to recache.
	Url string `json:"url"`
}
