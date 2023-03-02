/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// AssetsFieldPropertiesDto struct for AssetsFieldPropertiesDto
type AssetsFieldPropertiesDto struct {
	// Optional label for the editor.
	Label *string `json:"label,omitempty"`
	// Hints to describe the field.
	Hints *string `json:"hints,omitempty"`
	// Placeholder to show when no value has been entered.
	Placeholder *string `json:"placeholder,omitempty"`
	// Indicates if the field is required.
	IsRequired bool `json:"isRequired,omitempty"`
	// Indicates if the field is required when publishing.
	IsRequiredOnPublish bool `json:"isRequiredOnPublish,omitempty"`
	// Indicates if the field should be rendered with half width only.
	IsHalfWidth bool `json:"isHalfWidth,omitempty"`
	// Optional url to the editor.
	EditorUrl *string `json:"editorUrl,omitempty"`
	// Tags for automation processes.
	Tags          *[]string           `json:"tags,omitempty"`
	FieldType     string              `json:"fieldType"`
	PreviewMode   AssetPreviewMode    `json:"previewMode,omitempty"`
	DefaultValues map[string][]string `json:"defaultValues,omitempty"`
	// The default value as a list of asset ids.
	DefaultValue *[]string `json:"defaultValue,omitempty"`
	// The initial id to the folder.
	FolderId *string `json:"folderId,omitempty"`
	// The minimum allowed items for the field value.
	MinItems *int32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems *int32 `json:"maxItems,omitempty"`
	// The minimum file size in bytes.
	MinSize *int32 `json:"minSize,omitempty"`
	// The maximum file size in bytes.
	MaxSize *int32 `json:"maxSize,omitempty"`
	// The minimum image width in pixels.
	MinWidth *int32 `json:"minWidth,omitempty"`
	// The maximum image width in pixels.
	MaxWidth *int32 `json:"maxWidth,omitempty"`
	// The minimum image height in pixels.
	MinHeight *int32 `json:"minHeight,omitempty"`
	// The maximum image height in pixels.
	MaxHeight *int32 `json:"maxHeight,omitempty"`
	// The image aspect width in pixels.
	AspectWidth *int32 `json:"aspectWidth,omitempty"`
	// The image aspect height in pixels.
	AspectHeight *int32    `json:"aspectHeight,omitempty"`
	ExpectedType AssetType `json:"expectedType,omitempty"`
	// True to resolve first asset in the content list.
	ResolveFirst bool `json:"resolveFirst,omitempty"`
	// True to resolve first image in the content list.
	MustBeImage bool `json:"mustBeImage,omitempty"`
	// True to resolve first image in the content list.
	ResolveImage bool `json:"resolveImage,omitempty"`
	// The allowed file extensions.
	AllowedExtensions *[]string `json:"allowedExtensions,omitempty"`
	// True, if duplicate values are allowed.
	AllowDuplicates bool `json:"allowDuplicates,omitempty"`
}
