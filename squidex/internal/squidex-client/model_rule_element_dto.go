/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidex-client
// RuleElementDto struct for RuleElementDto
type RuleElementDto struct {
	// Describes the action or trigger type.
	Description string `json:"description"`
	// The label for the action or trigger type.
	Display string `json:"display"`
	// Optional title.
	Title *string `json:"title,omitempty"`
	// The color for the icon.
	IconColor *string `json:"iconColor,omitempty"`
	// The image for the icon.
	IconImage *string `json:"iconImage,omitempty"`
	// The optional link to the product that is integrated.
	ReadMore *string `json:"readMore,omitempty"`
	// The properties.
	Properties []RuleElementPropertyDto `json:"properties"`
}
