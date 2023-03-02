/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// AssetScriptsDtoAllOf struct for AssetScriptsDtoAllOf
type AssetScriptsDtoAllOf struct {
	// The script that is executed when creating an asset.
	Create *string `json:"create,omitempty"`
	// The script that is executed when updating a content.
	Update *string `json:"update,omitempty"`
	// The script that is executed when annotating a content.
	Annotate *string `json:"annotate,omitempty"`
	// The script that is executed when moving a content.
	Move *string `json:"move,omitempty"`
	// The script that is executed when deleting a content.
	Delete *string `json:"delete,omitempty"`
	// The version of the app.
	Version int64 `json:"version,omitempty"`
}
