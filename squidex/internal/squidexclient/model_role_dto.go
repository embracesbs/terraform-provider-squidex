/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient
// RoleDto struct for RoleDto
type RoleDto struct {
	// The links.
	Links map[string]ResourceLink `json:"_links"`
	// The role name.
	Name string `json:"name"`
	// The number of clients with this role.
	NumClients int32 `json:"numClients,omitempty"`
	// The number of contributors with this role.
	NumContributors int32 `json:"numContributors,omitempty"`
	// Indicates if the role is an builtin default role.
	IsDefaultRole bool `json:"isDefaultRole,omitempty"`
	// Associated list of permissions.
	Permissions []string `json:"permissions"`
}
