package squidex

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://docs.squidex.io/02-documentation/concepts/schemas
func resourceSchema() *schema.Resource {
	fieldPropertiesDto := map[string]*schema.Schema{
		"label": {
			Type: schema.TypeString,
			Optional: true,
			Default: "",
			Description: "Optional label for the editor.",
		},
		"hints": {
			Type: schema.TypeString,
			Optional: true,
			Default: "",
			Description: "Hints to describe the schema.",
		},
		"placeholder": {
			Type: schema.TypeString,
			Optional: true,
			Default: "",
			Description: "Placeholder to show when no value has been entered.",
		},
		"required": {
			Type: schema.TypeBool,
			Optional: true,
			Default: false,
			Description: "Indicates if the field is required.",
		},
		"half_width": {
			Type: schema.TypeBool,
			Optional: true,
			Default: false,
			Description: "Indicates if the field should be rendered with half width only.",
		},
		"editor_url": {
			Type: schema.TypeString,
			Optional: true,
			Default: "",
			Description: "Optional url to the editor.",
		},
		"unique": {
			Type: schema.TypeBool,
			Optional: true,
			Default: false,
			Description: "Indicates if the field value must be unique. Ignored for nested fields and localized fields.",
		},
		"editor": {
			Type: schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
					"Checkbox",
					"Toggle",
					"DateTime",
					"Date",
					"Map",
					"Input",
					"Radio",
					"Stars",
					"List",
					"Checkboxes",
					"Tags",
					"Slug",
					"TextArea",
					"RichText",
					"Markdown",
					"Color",
					"Html",
					"StockPhoto",
					"Dropdown",
					"Separator"}, false),
			Description: "The editor that is used to manage this field. Possible Values are; \nBoolean: Checkbox Toggle, \nDateTime: DateTime Date, \nGeoLocation: Map, \nNumber: Input Radio Dropdown Stars, \nReferences: List Dropdown Checkboxes Tags, \nString: Input Slug TextArea RichText Markdown Dropdown Radio Color Html StockPhoto, \nTags: Tags Checkboxes Dropdown, \nUI: Separator",
		},
// TODO: test these new fields on all crud mappings:
		"min_items": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum allowed items for the field value.",
		},
		"max_items": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum allowed items for the field value.",
		},
		"preview_mode": {
			Type: schema.TypeString,
			Optional: true,
			Default: "ImageAndFileName",
			Description: "(Asset field_type only) The preview mode for the asset.",
			ValidateFunc: validation.StringInSlice([]string{
					"ImageAndFileName",
					"Image",
					"FileName"}, false),
		},
		"default_values": {
			Type: schema.TypeMap,
			Optional: true,
			Description: "The language specific default value as a list of asset ids.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
/*	DefaultValues The language specific default value as a list of asset ids.
	DefaultValue *interface{} `json:"defaultValue,omitempty"`
		Options: 
		- lists: {"nl-NL": ["asset1","tag2"]}
		- string: {"nl-NL": "asset1"}
		- int/double: {"nl-NL": 10.9}
		- bool: {"nl-NL": true}
	*/ 
			// TODO: test if strings are accepted on api for all field_type's
			// how about arrays?
		},
		"default_value": {
			Type: schema.TypeList,
			/*	DefaultValue The default value for the field value.
	DefaultValues *map[string]interface{} `json:"defaultValues,omitempty"`
		Options:
		- lists: ["asset1","tag2"]
		- string: "asset1"
		- int/double: 10.9
		- bool: true
	*/
			// TODO: test if strings are accepted on api for all field_type's
			Optional: true,
			Description: "The language specific default value as a list of asset ids.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"min_size": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum file size in bytes.",
		},
		"max_size": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum file size in bytes.",
		},
		"min_width": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum image width in pixels.",
		},
		"max_width": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum image width in pixels.",
		},
		"min_height": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum image height in pixels.",
		},
		"max_height": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum image height in pixels.",
		},
		"aspect_width": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The image aspect width in pixels.",
		},
		"aspect_height": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The image aspect height in pixels.",
		},
		"must_be_image": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "Defines if the asset must be an image.",
		},
		"resolve_first": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "True to resolve first asset in the content list.",
		},
		"allowed_extensions": {
			Type: schema.TypeList,
			Optional: true,
			Description: "The allowed file extensions.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allow_duplicates": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "True, if duplicate values are allowed.",
		},
		"inline_editable": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "Indicates that the inline editor is enabled for this field.",
		},
		"max_value": {
			/*	MaxValue The maximum allowed value for the field value.
			MaxValue *interface{} `json:"maxValue,omitempty"`
				Options:
				string
				float32
			*/
			Type: schema.TypeString,
			Optional: true,
			Description: "The maximum allowed value for the field value.",
		},
		"min_value": {
			/*	
			MinValue *interface{} `json:"minValue,omitempty"`
				Options:
				string
				float32
			*/
			Type: schema.TypeString,
			Optional: true,
			// ? empty string doesn't parse correct for DateTime fields Default: nil,
			Description: "The minimum allowed value for the field value.",
		},
		"calculated_default_value": {
			Type: schema.TypeString,
			Optional: true,
			Description: "(DateTime field_type only) The calculated default value for the field value.",
			ValidateFunc: validation.StringInSlice([]string{
					"Now",
					"Today"}, false),
		},
		"allowed_values": {
			/*	AllowedValues The allowed values for the field value.
				Options:
				[]string
				[]float32
			*/
			Type: schema.TypeList,
			Optional: true,
			Description: "The allowed values for the field value.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"resolve_reference": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "True to resolve references in the content list. (Required for fieldtype References)",
		},
		"must_be_published": {
			Type: schema.TypeBool,
			Optional: true,
			Description: "True when all references must be published. (Required for fieldtype References)",
		},
		"schema_ids": {
			Type: schema.TypeList,
			Optional: true,
			Description: "The id of the referenced schemas.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"pattern": {
			Type: schema.TypeString,
			Optional: true,
			Description: "The pattern to enforce a specific format for the field value.",
		},
		"pattern_message": {
			Type: schema.TypeString,
			Optional: true,
			Description: "The validation message for the pattern.",
		},
		"min_length": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum allowed length for the field value.",
		},
		"max_length": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum allowed length for the field value.",
		},
		"min_characters": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum allowed of normal characters for the field value.",
		},
		"max_characters": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum allowed of normal characters for the field value.",
		},
		"min_words": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The minimum allowed number of words for the field value.",
		},
		"max_words": {
			Type: schema.TypeInt,
			Optional: true,
			Description: "The maximum allowed number of words for the field value.",
		},
		"content_type": {
			Type: schema.TypeString,
			Optional: true,
			Default: "Unspecified",
			Description: "How the string content should be interpreted. (Required for fieldtype string)",
			ValidateFunc: validation.StringInSlice([]string{
					"Unspecified",
					"Markdown",
					"Html"}, false),
		},
		"tags": {
			Type: schema.TypeList,
			Optional: true,
			Description: "Tags for automation processes.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		// Maybe make FieldType OneOfMany() but it's enums are not in the client
		"field_type": {
			Type: schema.TypeString,
			Optional: true,
			Default: "String",
			Description: "Type of the field.",
			ValidateFunc: validation.StringInSlice([]string{
					"Assets",
					"Array",
					"Boolean",
					"DateTime",
					"Geolocation",
					"Json",
					"Number",
					"References",
					"String",
					"Tags",
					"UI"}, false),
		},
	}
	return &schema.Resource{
		ReadContext:   resourceSchemaRead,
		CreateContext: resourceSchemaCreate,
		UpdateContext: resourceSchemaUpdate,
		DeleteContext: resourceSchemaDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"invalidated_state": {
				Type: schema.TypeBool,
				Computed: true,
				Description: "Hidden field to invalidate state on response errors.",
			},
			"app_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-z0-9]+(\-[a-z0-9]+)*$`), "app_name may only contain a-z 0-9 and - and not start with -."),
				Description: "Name of the application. Can not be changed later. Only [a-z0-9] and may contain dashes - but not start with them.",
			},
			"properties": {
				Type: schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Description: "The optional properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "Optional label for the editor.",
						},
						"hints": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "Hints to describe the schema.",
						},
						"contents_sidebar_url": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The url to a the sidebar plugin for content lists.",
						},
						"content_sidebar_url": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The url to a the sidebar plugin for content items.",
						},
						"tags": {
							Type: schema.TypeList,
							Optional: true,
							Description: "Tags for automation processes.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"scripts": {
				Type: schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Description: "The optional scripts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The script that is executed for each query when querying contents.",
						},
						"create": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The script that is executed when creating a content.",
						},
						"update": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The script that is executed when updating a content.",
						},
						"delete": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The script that is executed when deleting a content.",
						},
						"change": {
							Type: schema.TypeString,
							Optional: true,
							Default: "",
							Description: "The script that is executed when change a content status.",
						},
					},
				},
			},
			"fields_in_references": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "The names of the fields that should be used in references.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"fields_in_list": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "The names of the fields that should be shown in lists, including meta fields.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"fields": {
				Type: schema.TypeList,
				Optional: true,
				Description: "Optional fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9]+(\-[a-zA-Z0-9]+)*$`), "name may only contain a-z A-Z 0-9 and - and not start with -."),
							Description: "The name of the field. Must be unique within the schema. Only [a-zA-Z0-9] and may contain dashes - but not start with them.",
						},
						"hidden": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
							Description: "Hide field in the Api",
						},
						"locked": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
							Description: "Lock field and prevent changes.",
						},
						"disabled": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
							Description: "Disable field in the UI.",
						},
						"partitioning": {
							Type: schema.TypeString,
							Optional: true,
							Default: "invariant",
							Description: "Determines the optional partitioning of the field. Possible values are; invariant, language.",
							ValidateFunc: validation.StringInSlice([]string{					
								"invariant",
								"language"}, false),
						},
						"properties": {
							Type: schema.TypeList,
							Required: true,
							MaxItems: 1,
							Description: "Determines the optional partitioning of the field.",
							Elem: &schema.Resource{
								Schema: fieldPropertiesDto,
							},
						},
						"nested": {
							Type: schema.TypeList,
							Optional: true,
							Description: "The nested fields.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString,
										Required: true,
										Description: "The name of the field. Must be unique within the schema.",
									},
									"hidden": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
										Description: "Defines if the field is hidden.",
									},
									"locked": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
										Description: "Defines if the field is locked.",
									},
									"disabled": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
										Description: "Defines if the field is disabled.",
									},
									"properties": {
										Type: schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Description: "The field properties.",
										Elem: &schema.Resource{
											Schema: fieldPropertiesDto,
										},
									},
								},
							},
						},
					},
				},
			},
			"preview_urls": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: "The optional preview urls.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Default: "",
				Description: "The category.",
			},
			"published": {
				Type:     schema.TypeBool,
				Optional: true,
				Default: true,
				Description: "Set it to true to autopublish the schema.",
			},
			// TODO: add validation for all fields
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the schema. Only [a-z0-9] and may contain dashes - but not start with them.",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-z0-9]+(\-[a-z0-9]+)*$`), "Name may only contain a-z 0-9 and - and not start with -."),
			},
			"singleton": {
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
				Description: "Set to true to allow a single content item only.",
			},
		},
	}
}

func setDataFromSchemaDetailsDto(data *schema.ResourceData, schema squidexclient.SchemaDetailsDto) (error) {
	log.Printf("ids: %s %s", data.Id(), schema.Id)
	data.SetId(schema.Id)

	data.Set("name", schema.Name)
	data.Set("published", schema.IsPublished)
	data.Set("singleton", schema.IsSingleton)
	data.Set("category", schema.Category)

	if (squidexclient.SchemaPropertiesDto{}) == schema.Properties {
		data.Set("properties", nil)
	} else {
		properties := make(map[string]interface{})
		properties["label"] = schema.Properties.Label
		properties["hints"] = schema.Properties.Hints
		properties["contents_sidebar_url"] = schema.Properties.ContentsSidebarUrl
		properties["content_sidebar_url"] = schema.Properties.ContentSidebarUrl
		// TODO: properties - tags -> how to set array?
		properties["tags"] = schema.Properties.Tags
		data.Set("properties", []interface{}{properties})
	}

	if (squidexclient.SchemaScriptsDto{}) == schema.Scripts {
		data.Set("scripts", nil)
	} else {
		scripts := make(map[string]interface{})
		scripts["query"] = schema.Scripts.Query
		scripts["update"] = schema.Scripts.Update
		scripts["delete"] = schema.Scripts.Delete
		scripts["create"] = schema.Scripts.Create
		scripts["change"] = schema.Scripts.Change
		data.Set("scripts", []interface{}{scripts})
	}

	fields := make([]map[string]interface{}, len(schema.Fields))
	for i, v := range schema.Fields {
		fields[i] = make(map[string]interface{})
		fields[i]["name"] = v.Name
		fields[i]["hidden"] = &v.IsHidden
		fields[i]["locked"] = &v.IsLocked
		fields[i]["disabled"] = &v.IsDisabled
		fields[i]["partitioning"] = v.Partitioning
		
		if (squidexclient.FieldPropertiesDto{}) == v.Properties {
			fields[i]["properties"] = nil
		} else {
			properties := make(map[string]interface{})
			properties["hints"] = v.Properties.Hints
			properties["editor_url"] = v.Properties.EditorURL
			properties["min_items"] = v.Properties.MinItems
			properties["max_items"] = v.Properties.MaxItems
			properties["preview_mode"] = v.Properties.PreviewMode
			properties["default_values"] = v.Properties.DefaultValues
			properties["default_value"] = v.Properties.DefaultValue
			properties["min_size"] = v.Properties.MinSize
			properties["max_size"] = v.Properties.MaxSize
			properties["min_width"] = v.Properties.MinWidth
			properties["max_width"] = v.Properties.MaxWidth
			properties["min_height"] = v.Properties.MinHeight
			properties["max_height"] = v.Properties.MaxHeight
			properties["aspect_width"] = v.Properties.AspectWidth
			properties["aspect_height"] = v.Properties.AspectHeight
			properties["must_be_image"] = v.Properties.MustBeImage
			properties["resolve_first"] = v.Properties.ResolveFirst
			properties["allowed_extensions"] = v.Properties.AllowedExtensions
			properties["allow_duplicates"] = v.Properties.AllowDuplicates
			properties["inline_editable"] = v.Properties.InlineEditable
			properties["max_value"] = v.Properties.MaxValue
			
			log.Printf("v.Properties.MinValue: %s", v.Properties.MinValue)
			if (v.Properties.MinValue != nil) { 
				properties["min_value"] = v.Properties.MinValue
			}

			properties["calculated_default_value"] = v.Properties.CalculatedDefaultValue
			properties["allowed_values"] = v.Properties.AllowedValues
			properties["resolve_reference"] = v.Properties.ResolveReference
			properties["must_be_published"] = v.Properties.MustBePublished
			properties["schema_ids"] = v.Properties.SchemaIds
			properties["pattern"] = v.Properties.Pattern
			properties["pattern_message"] = v.Properties.PatternMessage
			properties["min_length"] = v.Properties.MinLength
			properties["max_length"] = v.Properties.MaxLength
			properties["min_characters"] = v.Properties.MinCharacters
			properties["max_characters"] = v.Properties.MaxCharacters
			properties["min_words"] = v.Properties.MinWords
			properties["max_words"] = v.Properties.MaxWords
			properties["content_type"] = v.Properties.ContentType
			properties["unique"] = v.Properties.Unique
			properties["editor"] = v.Properties.Editor
			properties["field_type"] = v.Properties.FieldType
			properties["half_width"] = &v.Properties.IsHalfWidth
			properties["required"] = &v.Properties.IsRequired
			properties["label"] = v.Properties.Label
			properties["placeholder"] = v.Properties.Placeholder
			properties["tags"] = v.Properties.Tags
			fields[i]["properties"] = []interface{}{properties}
		}
		
		if v.Nested == nil {
			fields[i]["nested"] = nil
		} else {
			nesteds := make([]map[string]interface{}, len(*v.Nested))
			for i, nested := range *v.Nested { 
				nesteds[i] = make(map[string]interface{})
				nesteds[i]["name"] = nested.Name
				nesteds[i]["hidden"] = nested.IsHidden
				nesteds[i]["locked"] = nested.IsLocked
				nesteds[i]["disabled"] = nested.IsDisabled

				properties := make(map[string]interface{})
				properties["hints"] = nested.Properties.Hints
				properties["editor_url"] = nested.Properties.EditorURL
				properties["min_items"] = nested.Properties.MinItems
				properties["max_items"] = nested.Properties.MaxItems
				properties["preview_mode"] = nested.Properties.PreviewMode
				properties["default_values"] = nested.Properties.DefaultValues
				properties["default_value"] = nested.Properties.DefaultValue
				properties["min_size"] = nested.Properties.MinSize
				properties["max_size"] = nested.Properties.MaxSize
				properties["min_width"] = nested.Properties.MinWidth
				properties["max_width"] = nested.Properties.MaxWidth
				properties["min_height"] = nested.Properties.MinHeight
				properties["max_height"] = nested.Properties.MaxHeight
				properties["aspect_width"] = nested.Properties.AspectWidth
				properties["aspect_height"] = nested.Properties.AspectHeight
				properties["must_be_image"] = nested.Properties.MustBeImage
				properties["resolve_first"] = nested.Properties.ResolveFirst
				properties["allowed_extensions"] = nested.Properties.AllowedExtensions
				properties["allow_duplicates"] = nested.Properties.AllowDuplicates
				properties["inline_editable"] = nested.Properties.InlineEditable
				properties["max_value"] = nested.Properties.MaxValue
				properties["min_value"] = nested.Properties.MinValue
				properties["calculated_default_value"] = nested.Properties.CalculatedDefaultValue
				properties["allowed_values"] = nested.Properties.AllowedValues
				properties["resolve_reference"] = nested.Properties.ResolveReference
				properties["must_be_published"] = nested.Properties.MustBePublished
				properties["schema_ids"] = nested.Properties.SchemaIds
				properties["pattern"] = nested.Properties.Pattern
				properties["pattern_message"] = nested.Properties.PatternMessage
				properties["min_length"] = nested.Properties.MinLength
				properties["max_length"] = nested.Properties.MaxLength
				properties["min_characters"] = nested.Properties.MinCharacters
				properties["max_characters"] = nested.Properties.MaxCharacters
				properties["min_words"] = nested.Properties.MinWords
				properties["max_words"] = nested.Properties.MaxWords
				properties["content_type"] = nested.Properties.ContentType
				properties["unique"] = nested.Properties.Unique
				properties["editor"] = nested.Properties.Editor
				properties["field_type"] = nested.Properties.FieldType
				properties["half_width"] = nested.Properties.IsHalfWidth
				properties["required"] = nested.Properties.IsRequired
				properties["label"] = nested.Properties.Label
				properties["placeholder"] = nested.Properties.Placeholder
				tags := *nested.Properties.Tags
				properties["tags"] = tags
				nesteds[i]["properties"] = []interface{}{properties}
			}
			fields[i]["nested"] = nesteds
		}
	}
	data.Set("fields", fields)
	
	data.Set("fields_in_references", schema.FieldsInReferences)
	data.Set("fields_in_list", schema.FieldsInLists)
	
	data.Set("preview_urls", stringMapToInterfaceMap(schema.PreviewUrls))
	
	return nil
}

// TODO: move to utils
func interfaceSliceToStringSlice(iv []interface{}) []string {
	var sv []string
	for _, i := range iv {
		sv = append(sv, i.(string))
	}
	return sv
}

// TODO: move to utils
func interfaceMapToStringMap(iv map[string]interface{}) map[string]string {
	sv := make(map[string]string)
	for k, v := range iv {
		sv[k] = v.(string)
	}
	return sv
}

// TODO: move to utils
func stringMapToInterfaceMap(iv map[string]string) map[string]interface{} {
	sv := make(map[string]interface{})
	for k, v := range iv {
		sv[k] = v
	}
	return sv
}

func stringArrayToNonNilStringArray(iv []string) []string {
	if iv == nil {
		return make([]string, 0)
	}
	sv := make([]string, len(iv))
	for _, v := range iv {
		sv = append(sv, v)
	}
	return sv
}

func getCreateSchemaDtoFromData(data *schema.ResourceData) (squidexclient.CreateSchemaDto, error) {
	squidexschema := squidexclient.CreateSchemaDto{
		IsPublished: data.Get("published").(bool),
		Name: data.Get("name").(string),
		IsSingleton: data.Get("singleton").(bool),
	}
	
	if category, ok := data.GetOk("category"); ok {
		x := category.(string)
		squidexschema.Category = &x
	}
	
	if v, ok := data.GetOk("properties"); ok {
		properties := v.([]interface{})[0].(map[string]interface{})
		squidexschema.Properties = new(squidexclient.SchemaPropertiesDto)
		if p, ok := properties["label"]; ok { 
			x := p.(string)
			squidexschema.Properties.Label = &x
		}
		if p, ok := properties["hints"]; ok { 
			x := p.(string)			
			squidexschema.Properties.Hints = &x
		}
		if p, ok := properties["contents_sidebar_url"]; ok { 
			x := p.(string)
			squidexschema.Properties.ContentsSidebarUrl = &x
		}
		if p, ok := properties["content_sidebar_url"]; ok { 
			x := p.(string)
			squidexschema.Properties.ContentSidebarUrl = &x
		}
		if p, ok := properties["tags"]; ok { 
			tags := interfaceSliceToStringSlice(p.([]interface{}))
			squidexschema.Properties.Tags = &tags
		}
	}

	if v, ok := data.GetOk("scripts"); ok {
		scripts := v.([]interface{})[0].(map[string]interface{})
		squidexschema.Scripts = new(squidexclient.SchemaScriptsDto)
		if p, ok := scripts["query"]; ok { 
			x := p.(string)
			squidexschema.Scripts.Query = &x
		}
		if p, ok := scripts["create"]; ok { 
			x := p.(string)
			squidexschema.Scripts.Create = &x
		}
		if p, ok := scripts["update"]; ok { 
			x := p.(string)
			squidexschema.Scripts.Update = &x
		}
		if p, ok := scripts["delete"]; ok { 
			x := p.(string)
			squidexschema.Scripts.Delete = &x
		}
		if p, ok := scripts["change"]; ok { 
			x := p.(string)
			squidexschema.Scripts.Change = &x
		}
	}
	if v, ok := data.GetOk("fields_in_references"); ok {
		fieldsInReferences := interfaceSliceToStringSlice(v.([]interface{}))
		squidexschema.FieldsInReferences = &fieldsInReferences		
	}
	if v, ok := data.GetOk("fields_in_list"); ok {
		fieldsInLists := interfaceSliceToStringSlice(v.([]interface{}))
		squidexschema.FieldsInLists = &fieldsInLists		
	}
	if v, ok := data.GetOk("preview_urls"); ok {
		previewUrls := interfaceMapToStringMap(v.(map[string]interface {}))
		squidexschema.PreviewUrls = &previewUrls
		//panic: interface conversion: interface {} is map[string]interface {}, not map[string]string
	}
	if v, ok := data.GetOk("fields"); ok {
		fields := v.([]interface{})
		squidexfields := make([]squidexclient.UpsertSchemaFieldDto, len(fields))
		for i, v := range fields {
			field := v.(map[string]interface{})
			
			// required field:
			partitioning := field["partitioning"].(string)
			squidexfields[i].Partitioning = &partitioning

			if field["name"] != nil {
				// name is a required field
				squidexfields[i].Name = field["name"].(string)
			}
			if field["hidden"] != nil {
				squidexfields[i].IsHidden = field["hidden"].(bool)
			}
			if field["locked"] != nil {
				squidexfields[i].IsLocked = field["locked"].(bool)
			}
			if field["disabled"] != nil {
				squidexfields[i].IsDisabled = field["disabled"].(bool)
			}
			if field["properties"] != nil {
				properties := field["properties"].([]interface{})
				if len(properties) == 1 {
					properties := properties[0].(map[string]interface{})
					squidexProperties := squidexclient.FieldPropertiesDto{}

					// required field:
					fieldType := properties["field_type"].(string)
					squidexProperties.FieldType = fieldType

					if properties["label"] != nil {
						label := properties["label"].(string)
						squidexProperties.Label = &label
					}
					if properties["hints"] != nil {
						hints := properties["hints"].(string)
						squidexProperties.Hints = &hints
					}
					if properties["placeholder"] != nil {
						placeholder := properties["placeholder"].(string)
						squidexProperties.Placeholder = &placeholder
					}
					if properties["required"] != nil {
						squidexProperties.IsRequired = properties["required"].(bool)
					}
					if properties["half_width"] != nil {
						squidexProperties.IsHalfWidth = properties["half_width"].(bool)
					}
					if properties["editor_url"] != nil {
						editorURL := properties["editor_url"].(string)
						squidexProperties.EditorURL = &editorURL
					}
					if properties["min_items"] != nil {
						minitems := properties["min_items"].(int)
						squidexProperties.MinItems = &minitems
					}
					if properties["max_items"] != nil {
						maxitems := properties["max_items"].(int)
						squidexProperties.MaxItems = &maxitems
					}
					if properties["preview_mode"] != nil {
						previewMode := properties["preview_mode"].(string)
						squidexProperties.PreviewMode = &previewMode
					}
					if properties["default_values"] != nil {
						defaultvalues := defaultValuesToInterface(fieldType, partitioning, properties["default_values"])
						squidexProperties.DefaultValues = &defaultvalues
					}
					if properties["default_value"] != nil {
						defaultvalue := defaultValueToInterface(fieldType, properties["default_value"])
						squidexProperties.DefaultValue = &defaultvalue
					}
					if properties["min_size"] != nil {
						minsize := properties["min_size"].(int)
						squidexProperties.MinSize = &minsize
					}
					if properties["max_size"] != nil {
						maxsize := properties["max_size"].(int)
						squidexProperties.MaxSize = &maxsize
					}
					if properties["min_width"] != nil {
						minwidth := properties["min_width"].(int)
						squidexProperties.MinWidth = &minwidth
					}
					if properties["max_width"] != nil {
						maxwidth := properties["max_width"].(int)
						squidexProperties.MaxWidth = &maxwidth
					}
					if properties["min_height"] != nil {
						minheight := properties["min_height"].(int)
						squidexProperties.MinHeight = &minheight
					}
					if properties["max_height"] != nil {
						maxheight := properties["max_height"].(int)
						squidexProperties.MaxHeight = &maxheight
					}
					if properties["aspect_width"] != nil {
						aspectwidth := properties["aspect_width"].(int)
						squidexProperties.AspectWidth = &aspectwidth
					}
					if properties["aspect_height"] != nil {
						aspectheight := properties["aspect_height"].(int)
						squidexProperties.AspectHeight = &aspectheight
					}
					if properties["must_be_image"] != nil {
						mustbeimage := properties["must_be_image"].(bool)
						squidexProperties.MustBeImage = &mustbeimage
					}
					if properties["resolve_first"] != nil {
						resolvefirst := properties["resolve_first"].(bool)
						squidexProperties.ResolveFirst = &resolvefirst
					}
					if properties["allowed_extensions"] != nil {
						allowedextensions := interfaceSliceToStringSlice(properties["allowed_extensions"].([]interface{}))
						squidexProperties.AllowedExtensions = &allowedextensions
					}
					if properties["allow_duplicates"] != nil {
						val := properties["allow_duplicates"].(bool)
						squidexProperties.AllowDuplicates = &val
					}
					if properties["inline_editable"] != nil {
						val := properties["inline_editable"].(bool)
						squidexProperties.InlineEditable = &val
					}
					if properties["max_value"] != nil {
						val := properties["max_value"]
						squidexProperties.MaxValue = &val
					}
					if properties["min_value"] != nil {
						val := properties["min_value"]
						squidexProperties.MinValue = &val
					}
					if properties["calculated_default_value"] != nil {
						val := properties["calculated_default_value"].(string)
						squidexProperties.CalculatedDefaultValue = &val
					}
					if properties["allowed_values"] != nil {
						val := properties["allowed_values"].([]interface{})
						squidexProperties.AllowedValues = &val
					}
					if properties["resolve_reference"] != nil {
						val := properties["resolve_reference"].(bool)
						squidexProperties.ResolveReference = &val
					}
					if properties["must_be_published"] != nil {
						val := properties["must_be_published"].(bool)
						squidexProperties.MustBePublished = &val
					}
					if properties["schema_ids"] != nil {
						val := interfaceSliceToStringSlice(properties["schema_ids"].([]interface{}))
						squidexProperties.SchemaIds = &val
					}
					if properties["pattern"] != nil {
						val := properties["pattern"].(string)
						squidexProperties.Pattern = &val
					}
					if properties["pattern_message"] != nil {
						val := properties["pattern_message"].(string)
						squidexProperties.PatternMessage = &val
					}
					if properties["min_length"] != nil {
						val := properties["min_length"].(int)
						squidexProperties.MinLength = &val
					}
					if properties["max_length"] != nil {
						val := properties["max_length"].(int)
						squidexProperties.MaxLength = &val
					}
					if properties["min_characters"] != nil {
						val := properties["min_characters"].(int)
						squidexProperties.MinCharacters = &val
					}
					if properties["max_characters"] != nil {
						val := properties["max_characters"].(int)
						squidexProperties.MaxCharacters = &val
					}
					if properties["min_words"] != nil {
						val := properties["min_words"].(int)
						squidexProperties.MinWords = &val
					}
					if properties["max_words"] != nil {
						val := properties["max_words"].(int)
						squidexProperties.MaxWords = &val
					}
					if properties["content_type"] != nil {
						val := properties["content_type"].(string)
						squidexProperties.ContentType = &val
					}
					if properties["unique"] != nil {
						unique := properties["unique"].(bool)
						squidexProperties.Unique = &unique
					}
					if properties["editor"] != nil {
						editor := properties["editor"].(string)
						squidexProperties.Editor = &editor
					}
					if properties["tags"] != nil {
						tags := interfaceSliceToStringSlice(properties["tags"].([]interface{}))
						squidexProperties.Tags = &tags
					}
					squidexfields[i].Properties = squidexProperties
				}
			}
			// Only array fields can have nested fields.
			if squidexfields[i].Properties.FieldType == "Array" && field["nested"] != nil {
				nesteds := field["nested"].([]interface{})
				squidexNested := make([]squidexclient.UpsertSchemaNestedFieldDto, len(nesteds))
				for i, v := range nesteds {
					nested := v.(map[string]interface{})
					if nested["name"] != nil {
						squidexNested[i].Name = nested["name"].(string)
					}
					if nested["hidden"] != nil {
						squidexNested[i].IsHidden = nested["hidden"].(bool)
					}
					if nested["locked"] != nil {
						squidexNested[i].IsLocked = nested["locked"].(bool)
					}
					if nested["disabled"] != nil {
						squidexNested[i].IsDisabled = nested["disabled"].(bool)
					}
					if nested["properties"] != nil {
						properties := nested["properties"].([]interface{})
						if len(properties) == 1 {
							properties := properties[0].(map[string]interface{})
							squidexProperties := squidexclient.FieldPropertiesDto{}
							
							// required field:
							fieldType := properties["field_type"].(string)	
							squidexProperties.FieldType = fieldType

							if properties["label"] != nil {
								label := properties["label"].(string)
								squidexProperties.Label = &label
							}
							if properties["hints"] != nil {
								hints := properties["hints"].(string)
								squidexProperties.Hints = &hints
							}
							if properties["placeholder"] != nil {
								placeholder := properties["placeholder"].(string)
								squidexProperties.Placeholder = &placeholder
							}
							if properties["required"] != nil {
								squidexProperties.IsRequired = properties["required"].(bool)
							}
							if properties["half_width"] != nil {
								squidexProperties.IsHalfWidth = properties["half_width"].(bool)
							}
							if properties["editor_url"] != nil {
								editorURL := properties["editor_url"].(string)
								squidexProperties.EditorURL = &editorURL
							}

							if properties["min_items"] != nil {
								minitems := properties["min_items"].(int)
								squidexProperties.MinItems = &minitems
							}
							if properties["max_items"] != nil {
								maxitems := properties["max_items"].(int)
								squidexProperties.MaxItems = &maxitems
							}
							if properties["preview_mode"] != nil {
								previewMode := properties["preview_mode"].(string)
								squidexProperties.PreviewMode = &previewMode
							}
							if properties["default_values"] != nil {
								defaultvalues := defaultValuesToInterface(fieldType, partitioning, properties["default_values"])
								squidexProperties.DefaultValues = &defaultvalues
							}
							if properties["default_value"] != nil {
								defaultvalue := defaultValueToInterface(fieldType, properties["default_value"])
								squidexProperties.DefaultValue = &defaultvalue
							}
							if properties["min_size"] != nil {
								minsize := properties["min_size"].(int)
								squidexProperties.MinSize = &minsize
							}
							if properties["max_size"] != nil {
								maxsize := properties["max_size"].(int)
								squidexProperties.MaxSize = &maxsize
							}
							if properties["min_width"] != nil {
								minwidth := properties["min_width"].(int)
								squidexProperties.MinWidth = &minwidth
							}
							if properties["max_width"] != nil {
								maxwidth := properties["max_width"].(int)
								squidexProperties.MaxWidth = &maxwidth
							}
							if properties["min_height"] != nil {
								minheight := properties["min_height"].(int)
								squidexProperties.MinHeight = &minheight
							}
							if properties["max_height"] != nil {
								maxheight := properties["max_height"].(int)
								squidexProperties.MaxHeight = &maxheight
							}
							if properties["aspect_width"] != nil {
								aspectwidth := properties["aspect_width"].(int)
								squidexProperties.AspectWidth = &aspectwidth
							}
							if properties["aspect_height"] != nil {
								aspectheight := properties["aspect_height"].(int)
								squidexProperties.AspectHeight = &aspectheight
							}
							if properties["must_be_image"] != nil {
								mustbeimage := properties["must_be_image"].(bool)
								squidexProperties.MustBeImage = &mustbeimage
							}
							if properties["resolve_first"] != nil {
								resolvefirst := properties["resolve_first"].(bool)
								squidexProperties.ResolveFirst = &resolvefirst
							}
							if properties["unique"] != nil {
								unique := properties["unique"].(bool)
								squidexProperties.Unique = &unique
							}
							if properties["editor"] != nil {
								editor := properties["editor"].(string)
								squidexProperties.Editor = &editor
							}
							if properties["tags"] != nil {
								tags := interfaceSliceToStringSlice(properties["tags"].([]interface{}))
								squidexProperties.Tags = &tags
							}
							squidexNested[i].Properties = squidexProperties
						}
					}
				}
				squidexfields[i].Nested = &squidexNested
			}
		}
		squidexschema.Fields = &squidexfields
	}
	return squidexschema, nil
}

func defaultValuesToInterface(fieldType string, partitioning string, defaultValue interface{}) interface{}{
	if  partitioning != "language" ||
		defaultValue == nil || 
		fieldType == "Array" || 
		fieldType == "Geolocation" || 
		fieldType == "Json" || 
		fieldType == "UI" {
		return nil
	}
	// defaultValues is always a map[string]string in the schema
	v := interfaceMapToStringMap(defaultValue.(map[string]interface{}))
	// where map key is the language as 'nl-NL'
	switch fieldType {
		case "Assets", "References", "Tags":
			target := make(map[string][]string)
			for key, value := range v {
				value = strings.TrimLeft(value, "[")
				value = strings.TrimRight(value, "]")
				value := strings.Replace(value, "\", \"","\",\"", -1)
				valueSlice := strings.Split(value, "\",\"")
				// message := ("Error converting default_values to []string. Expected {\"nl-NL\" = \"[\"string1\", \"string2\"]\"], got %s", value)
				target[key] = valueSlice
			}
			return target
		case "Boolean":
			target := make(map[string]bool)
			for key, value := range v {
				result, err := strconv.ParseBool(value)
				if err != nil {
					log.Panicf("Error converting default_values to boolean. Expected {\"nl-NL\" = \"true\"], got %s %s", key, value)
				}
				target[key] = result
			}
			return target
		case "DateTime", "String":
			return v
		case "Number":
			target := make(map[string]float64)
			for key, value := range v {
				result, err := strconv.ParseFloat(value, 64)
				if err != nil {
					log.Panicf("Error converting default_values to float32. Expected {\"nl-NL\" = \"1234.5678\"], got %s %s", key, value)
				}
				target[key] = result
			}
			return target
		default:
			// Any unknown field_type is ignored
			return nil
	}
}
func defaultValueToInterface(fieldType string, defaultValue interface{}) interface{}{
	if defaultValue == nil || 
		fieldType == "Array" || 
		fieldType == "Geolocation" || 
		fieldType == "Json" || 
		fieldType == "UI" {
		return nil
	}
	// defaultValue is always a string slice in the schema
	v := interfaceSliceToStringSlice(defaultValue.([]interface{}))
	if len(v) == 0 {
		return nil
	}
	switch fieldType {
		case "Assets", "References", "Tags":
			return v
		case "Boolean":
			result, err := strconv.ParseBool(v[0])
			if err != nil {
				log.Panicf("Error converting default_value to boolean. Expected [\"true\"], got %s", v[0])
			}
			return result
		case "DateTime", "String":
			return v[0]
		case "Number":
			result, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Panicf("Error converting default_value to float32. Expected [\"1234.5678\"], got %s", v[0])
			}
			return result
		default:
			// Any unknown field_type is ignored
			return nil
	}
}

func mapCreateSchemaDtoToSynchronizeSchemaDto(createSchema squidexclient.CreateSchemaDto) (squidexclient.SynchronizeSchemaDto) {
	// use syncchronize with NoFieldDeletion & NoFieldRecreation to true (and message why it's not allowed on errors)
	dto := squidexclient.SynchronizeSchemaDto{
		NoFieldDeletion: true,
		NoFieldRecreation: true,
		Category: createSchema.Category,
		Fields: createSchema.Fields,
		FieldsInLists: createSchema.FieldsInLists,
		FieldsInReferences: createSchema.FieldsInReferences,
		IsPublished: createSchema.IsPublished,
		PreviewUrls: createSchema.PreviewUrls,
		Properties: createSchema.Properties,
		Scripts: createSchema.Scripts,
	}
	return dto
}

func resourceSchemaCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)
	
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	
	createDto, _ := getCreateSchemaDtoFromData(data)

	// TODO: handle all errors
	result, response, err := client.SchemasApi.SchemasPostSchema(ctx, appName, createDto)

	err = common.HandleAPIError(response, err)

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	// prevent drift and set ALL values from the result
	err = setDataFromSchemaDetailsDto(data, result)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Id)

	return diags
}

func resourceSchemaRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	data.Set("invalidated_state", false)

	result, response, err := client.SchemasApi.SchemasGetSchema(ctx, appName, name)

	err = common.HandleAPIError(response, err)

	if err != nil {
		return diag.FromErr(err)
	}

	//prevent drift and set ALL values from the get
	err = setDataFromSchemaDetailsDto(data, result)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSchemaUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	createDto, err := getCreateSchemaDtoFromData(data)
	if err != nil {
		return diag.FromErr(err)
	}
	syncDto := mapCreateSchemaDtoToSynchronizeSchemaDto(createDto)
		
	result, response, err := client.SchemasApi.SchemasPutSchemaSync(ctx, appName, name, syncDto)
	
	err = common.HandleAPIError(response, err)

	if err != nil {
		data.Set("invalidated_state", true)
		// TODO: handle all errors
		// TODO: update - display correct error response and message
		return diag.FromErr(err)
	}
	
	// prevent drift and set ALL values from the result
	// to test this, do once with, and once without, and compare results
	err = setDataFromSchemaDetailsDto(data, result);
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSchemaDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	response, err := client.SchemasApi.SchemasDeleteSchema(ctx, appName, name)
	
	err = common.HandleAPIError(response, err)
	
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")
	

	return diags
}
