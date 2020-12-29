/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// WebhookRuleActionDto struct for WebhookRuleActionDto
type WebhookRuleActionDto struct {
	ActionType string `json:"actionType"`
	// The url to the webhook.
	Url string `json:"url"`
	// The shared secret that is used to calculate the signature.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// The optional custom request payload.
	Payload *string `json:"payload,omitempty"`
}
