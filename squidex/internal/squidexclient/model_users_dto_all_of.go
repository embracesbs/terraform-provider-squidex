/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// UsersDtoAllOf struct for UsersDtoAllOf
type UsersDtoAllOf struct {
	// The total number of users.
	Total int64 `json:"total,omitempty"`
	// The users.
	Items []UserDto `json:"items"`
}
