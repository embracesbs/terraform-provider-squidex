
package squidexclient
// FieldPropertiesDto struct for FieldPropertiesDto
// For field descriptions, requirements etc. check out the full documentation on https://cloud.squidex.io/api/docs (ensure correct version)
// All ...DefaultValues fields: if partitioning=language, then there is a default for each language, otherwise use defaultvalue field
type FieldPropertiesDto struct {

	// common fields:

	// Label Optional label for the editor.
	Label *string `json:"label,omitempty"`
	// Hints to describe the schema.
	Hints *string `json:"hints,omitempty"`
	// Placeholder to show when no value has been entered.
	Placeholder *string `json:"placeholder,omitempty"`
	// IsRequired Indicates if the field is required.
	IsRequired bool `json:"isRequired,omitempty"`
	// IsHalfWidth Indicates if the field should be rendered with half width only.
	IsHalfWidth bool `json:"isHalfWidth,omitempty"`
	// EditorUrl Optional url to the editor.
	EditorURL *string `json:"editorUrl,omitempty"`
	// Tags for automation processes.
	Tags *[]string `json:"tags,omitempty"`
	// FieldType, depending the field type there are other fields required (prefixed with fieldType) other are ignored.
	FieldType string `json:"fieldType"`
	
	// Editor
	Editor *string `json:"editor,omitempty"`
	// MinItems The minimum allowed items for the field value.
	MinItems *int `json:"minItems,omitempty"`
	// MaxItems The maximum allowed items for the field value.
	MaxItems *int `json:"maxItems,omitempty"`
	// PreviewMode The preview mode for the asset.
	/*	Options: 
		- asset(ImageAndFileName Image FileName)
	*/
	PreviewMode *string `json:"previewMode,omitempty"`
	/*	DefaultValues The language specific default value as a list of asset ids.
		Options: 
		- lists: {"nl-NL": ["asset1","tag2"]}
		- string: {"nl-NL": "asset1"}
		- int/double: {"nl-NL": 10.9}
		- bool: {"nl-NL": true}
	*/ 
	DefaultValues *map[string]interface{} `json:"defaultValues,omitempty"`
	/*	DefaultValue The default value for the field value.
		Options:
		- lists: ["asset1","tag2"]
		- string: "asset1"
		- int/double: 10.9
		- bool: true
	*/
	DefaultValue *interface{} `json:"defaultValue,omitempty"`
	// MinSize The minimum file size in bytes.
	MinSize *int `json:"minSize,omitempty"`
	// MaxSize The maximum file size in bytes.
	MaxSize *int `json:"maxSize,omitempty"`
	// MinWidth The minimum image width in pixels.
	MinWidth *int `json:"minWidth,omitempty"`
	// MaxWidth The maximum image width in pixels.
	MaxWidth *int `json:"maxWidth,omitempty"`
	// MinHeight The minimum image height in pixels.
	MinHeight *int `json:"minHeight,omitempty"`
	// MaxHeight The maximum image height in pixels.
	MaxHeight *int `json:"maxHeight,omitempty"`
	// AspectWidth The image aspect width in pixels.
	AspectWidth *int `json:"aspectWidth,omitempty"`
	// AspectHeight The image aspect height in pixels.
	AspectHeight *int `json:"aspectHeight,omitempty"`
	// MustBeImage Defines if the asset must be an image.
	MustBeImage *bool `json:"mustBeImage,omitempty"`
	// ResolveFirst True to resolve first asset in the content list.
	ResolveFirst *bool `json:"resolveFirst,omitempty"`
	// AllowedExtensions The allowed file extensions.
	AllowedExtensions *[]string `json:"allowedExtensions,omitempty"`
	// AllowDuplicates True, if duplicate values are allowed.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// InlineEditable Indicates that the inline editor is enabled for this field.
	InlineEditable *bool `json:"inlineEditable,omitempty"`
	/*	MaxValue The maximum allowed value for the field value.
		Options:
		string
		float32
	*/
	MaxValue *interface{} `json:"maxValue,omitempty"`
	/* MinValue The minimum allowed value for the field value.
		Options:
		string
		float32
	*/
	MinValue *interface{} `json:"minValue,omitempty"`
	// CalculatedDefaultValue The calculated default value for the field value.
	// Options: Now Today (or leave empty for none)
	CalculatedDefaultValue *string `json:"calculatedDefaultValue,omitempty"`
	/*	AllowedValues The allowed values for the field value.
		Options:
	 	[]string
	 	[]float32
	*/
	AllowedValues *[]interface{} `json:"allowedValues,omitempty"`
	// Unique Indicates if the field value must be unique. Ignored for nested fields and localized fields.
	Unique *bool `json:"isUnique,omitempty"`
	// ResolveReference True to resolve references in the content list. (Required for fieldtype References)
	ResolveReference *bool `json:"resolveReference,omitempty"`
	// MustBePublished True when all references must be published. (Required for fieldtype References)
	MustBePublished *bool `json:"mustBePublished,omitempty"`
	// schemaIds The id of the referenced schemas.
	SchemaIds *[]string `json:"schemaIds,omitempty"`
	// Pattern The pattern to enforce a specific format for the field value.
	Pattern *string `json:"pattern,omitempty"`
	// PatternMessage The validation message for the pattern.
	PatternMessage *string `json:"patternMessage,omitempty"`
	// MinLength The minimum allowed length for the field value.
	MinLength *int `json:"minLength,omitempty"`
	// MaxLength The maximum allowed length for the field value.
	MaxLength *int `json:"maxLength,omitempty"`
	// MinCharacters The minimum allowed of normal characters for the field value.
	MinCharacters *int `json:"minCharacters,omitempty"`
	// MaxCharacters The maximum allowed of normal characters for the field value.
	MaxCharacters *int `json:"maxCharacters,omitempty"`
	// MinWords The minimum allowed number of words for the field value.
	MinWords *int `json:"minWords,omitempty"`
	// MaxWords The maximum allowed number of words for the field value.
	MaxWords *int `json:"maxWords,omitempty"`
	/*	ContentType How the string content should be interpreted. (Required for fieldtype string)
		Options: 
		Unspecified 
		Markdown 
		Html
	*/
	ContentType *string `json:"contentType,omitempty"`
}
