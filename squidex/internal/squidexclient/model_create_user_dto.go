/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// CreateUserDto struct for CreateUserDto
type CreateUserDto struct {
	// The email of the user. Unique value.
	Email string `json:"email"`
	// The display name (usually first name and last name) of the user.
	DisplayName string `json:"displayName"`
	// The password of the user.
	Password string `json:"password"`
	// Additional permissions for the user.
	Permissions []string `json:"permissions"`
}
