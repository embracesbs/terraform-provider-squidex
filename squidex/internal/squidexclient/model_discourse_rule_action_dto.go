/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// DiscourseRuleActionDto struct for DiscourseRuleActionDto
type DiscourseRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The url to the discourse server.
	Url string `json:"url"`
	// The api key to authenticate to your discourse server.
	ApiKey string `json:"apiKey"`
	// The api username to authenticate to your discourse server.
	ApiUsername string `json:"apiUsername"`
	// The text as markdown.
	Text string `json:"text"`
	// The optional title when creating new topics.
	Title *string `json:"title,omitempty"`
	// The optional topic id.
	Topic *int32 `json:"topic,omitempty"`
	// The optional category id.
	Category *int32 `json:"category,omitempty"`
}
