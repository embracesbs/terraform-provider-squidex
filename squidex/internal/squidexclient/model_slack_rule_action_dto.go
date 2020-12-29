/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// SlackRuleActionDto struct for SlackRuleActionDto
type SlackRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The slack webhook url.
	WebhookUrl string `json:"webhookUrl"`
	// The text that is sent as message to slack.
	Text string `json:"text"`
}
