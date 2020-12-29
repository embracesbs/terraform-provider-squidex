/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// NotificationRuleActionDtoAllOf struct for NotificationRuleActionDtoAllOf
type NotificationRuleActionDtoAllOf struct {
	// The user id or email.
	User string `json:"user"`
	// The text to send.
	Text string `json:"text"`
	// The optional url to attach to the notification.
	Url *string `json:"url,omitempty"`
	// An optional client name.
	Client *string `json:"client,omitempty"`
}
