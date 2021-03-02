/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// WebhookRuleActionDtoAllOf struct for WebhookRuleActionDtoAllOf
type WebhookRuleActionDtoAllOf struct {
	// The url to the webhook.
	Url string `json:"url"`
	// The type of the request.
	Method WebhookMethod `json:"method"`
	// Leave it empty to use the full event as body.
	Payload *string `json:"payload,omitempty"`
	// The mime type of the payload.
	PayloadType *string `json:"payloadType,omitempty"`
	// The message headers in the format '[Key]=[Value]', one entry per line.
	Headers *string `json:"headers,omitempty"`
	// The shared secret that is used to calculate the payload signature.
	SharedSecret *string `json:"sharedSecret,omitempty"`
}
