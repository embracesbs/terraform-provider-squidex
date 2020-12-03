/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
// TweetRuleActionDtoAllOf struct for TweetRuleActionDtoAllOf
type TweetRuleActionDtoAllOf struct {
	//  The generated access token.
	AccessToken string `json:"accessToken"`
	//  The generated access secret.
	AccessSecret string `json:"accessSecret"`
	// The text that is sent as tweet to twitter.
	Text string `json:"text"`
}
