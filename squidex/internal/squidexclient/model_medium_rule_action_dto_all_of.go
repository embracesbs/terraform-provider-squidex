/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 5.3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// MediumRuleActionDtoAllOf struct for MediumRuleActionDtoAllOf
type MediumRuleActionDtoAllOf struct {
	// The self issued access token.
	AccessToken string `json:"accessToken"`
	// The title, used for the url.
	Title string `json:"title"`
	// The content, either html or markdown.
	Content string `json:"content"`
	// The original home of this content, if it was originally published elsewhere.
	CanonicalUrl *string `json:"canonicalUrl,omitempty"`
	// The optional comma separated list of tags.
	Tags *string `json:"tags,omitempty"`
	// Optional publication id.
	PublicationId *string `json:"publicationId,omitempty"`
	// Indicates whether the content is markdown or html.
	IsHtml bool `json:"isHtml,omitempty"`
}
