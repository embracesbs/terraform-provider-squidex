package squidex

import (
	"log"
	"context"
	"encoding/json"
	"regexp"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"

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
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-z0-9]+(\-[a-z0-9]+)*$`), "name may only contain a-z 0-9 and - and not start with -."),
							Description: "The name of the field. Must be unique within the schema. Only [a-z0-9] and may contain dashes - but not start with them.",
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
							Description: "Determines the optional partitioning of the field.",
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
			if field["name"] != nil {
				// name is a required field
				squidexfields[i].Name = field["name"].(string)
			}
			if field["partitioning"] != nil {
				partitioning := field["partitioning"].(string)
				squidexfields[i].Partitioning = &partitioning
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
					if properties["tags"] != nil {
						tags := interfaceSliceToStringSlice(properties["tags"].([]interface{}))
						squidexProperties.Tags = &tags
					}
					if properties["field_type"] != nil {
						// field_type is a required field
						squidexProperties.FieldType = properties["field_type"].(string)
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
							if properties["tags"] != nil {
								tags := interfaceSliceToStringSlice(properties["tags"].([]interface{}))
								squidexProperties.Tags = &tags
							}
							if properties["field_type"] != nil {
								// field_type is a required field
								squidexProperties.FieldType = properties["field_type"].(string)
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
	
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	
	log.Println("create dto: start")
	dto, _ := getCreateSchemaDtoFromData(data)
	source, _ := json.Marshal(dto)
	log.Printf("create dto: \n%s", string(source))

	// TODO: handle all errors
	result, resp, err := client.SchemasApi.SchemasPostSchema(ctx, appName, dto)

	common.HandleAPIError(resp)

	if err != nil {
		return diag.FromErr(err)
	}

	// prevent drift and set ALL values from the result
	// TODO: setDataFromSchemaDetailsDto() on all data
	if err = setDataFromSchemaDetailsDto(data, result); err != nil {
		return diag.FromErr(err)
	}

	id := result.Id

	data.SetId(id)

	return diags
}

func resourceSchemaRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	result, resp, err := client.SchemasApi.SchemasGetSchema(ctx, appName, name)

	common.HandleAPIError(resp)

	if err != nil {
		// log.Printf("", response.)
		return diag.FromErr(err)
	}

	//prevent drift and set ALL values from the get
	if err = setDataFromSchemaDetailsDto(data, result); err != nil {
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
	dto := mapCreateSchemaDtoToSynchronizeSchemaDto(createDto)
		
	result, resp, err := client.SchemasApi.SchemasPutSchemaSync(ctx, appName, name, dto)
	
	common.HandleAPIError(resp)

	if err != nil {
		// TODO: handle all errors
		// TODO: update - display correct error response and message
		return diag.FromErr(err)
	}
	
	// prevent drift and set ALL values from the result
	// to test this, do once with, and once without, and compare results
	if err = setDataFromSchemaDetailsDto(data, result); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSchemaDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	resp, err := client.SchemasApi.SchemasDeleteSchema(ctx, appName, name)
	
	common.HandleAPIError(resp)
	
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
