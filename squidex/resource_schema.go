package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
// https://docs.squidex.io/02-documentation/concepts/schemas
func resourceSchema() *schema.Resource {
	fieldPropertiesDto := map[string]*schema.Schema{
		"label": {
			Type: schema.TypeString,
			Optional: true,
			Description: "Optional label for the editor.",
		},
		"hints": {
			Type: schema.TypeString,
			Optional: true,
			Description: "Hints to describe the schema.",
		},
		"placeholder": {
			Type: schema.TypeString,
			Optional: true,
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
			Required: true,
			Description: "",
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
				Description: "",
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
							Description: "Optional label for the editor.",
						},
						"hints": {
							Type: schema.TypeString,
							Optional: true,
							Description: "Hints to describe the schema.",
						},
						"contents_sidebar_url": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The url to a the sidebar plugin for content lists.",
						},
						"content_sidebar_url": {
							Type: schema.TypeString,
							Optional: true,
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
							Description: "The script that is executed for each query when querying contents.",
						},
						"create": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The script that is executed when creating a content.",
						},
						"update": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The script that is executed when updating a content.",
						},
						"delete": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The script that is executed when deleting a content.",
						},
						"change": {
							Type: schema.TypeString,
							Optional: true,
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
						"partitioning": {
							Type: schema.TypeString,
							Optional: true,
							Description: "Determines the optional partitioning of the field.",
						},
						"properties": {
							Type: schema.TypeList,
							Optional: true,
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
				Description: "The category.",
			},
			"published": {
				Type:     schema.TypeBool,
				Required: true,
				Description: "Set it to true to autopublish the schema.",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the schema.",
			},
			"singleton": {
				Type:     schema.TypeBool,
				Required: true,
				Description: "Set to true to allow a single content item only.",
			},
		},
	}
}

func resourceSchemaRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	result, _, err := client.SchemasApi.SchemasGetSchema(ctx, appName, name)
	
	if err != nil {
		return diag.FromErr(err)
	}
	
	// prevent drift and set ALL values from the get
	if err = setDataFromSchemaDetailsDto(data, result); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setDataFromSchemaDetailsDto(data *schema.ResourceData, schema squidexclient.SchemaDetailsDto) (error) {
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
		// properties["tags"] = schema.Properties.Tags
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
		fields[i]["name"] = v.Name
		fields[i]["hidden"] = v.IsHidden
		fields[i]["locked"] = v.IsLocked
		fields[i]["disabled"] = v.IsDisabled
		fields[i]["partitioning"] = v.Partitioning
		
		if (squidexclient.FieldPropertiesDto{}) == v.Properties {
			fields[i]["properties"] = nil
		} else {
			properties := make(map[string]interface{})
			properties["hints"] = v.Properties.Hints
			// TODO: fields - properties
			fields[i]["properties"] = []interface{}{properties}
		}
		
		if v.Nested == nil {
			fields[i]["nested"] = nil
		} else {
			// nested := make([]map[string]interface{}, len(*v.Nested))
			// for i, v := range nested { }
			// TODO: fields - nested
			// fields[i]["nested"] = []interface{}{nested}
		}
	}
	data.Set("fields", []interface{}{fields})
	
	data.Set("fields_in_references", []interface{}{schema.FieldsInReferences})
	data.Set("fields_in_lists", []interface{}{schema.FieldsInLists})
	data.Set("preview_urls", []interface{}{schema.PreviewUrls})
	
	return nil
}

func interfaceSliceToStringSlice(iv []interface{}) []string {
	var sv []string
	for _, i := range iv {
		sv = append(sv, i.(string))
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
// panic: interface conversion: interface {} is []interface {}, not []string
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
		fieldsInReferences := v.([]string)
		squidexschema.FieldsInReferences = &fieldsInReferences		
	}
	if v, ok := data.GetOk("fields_in_lists"); ok {
		fieldsInLists := v.([]string)
		squidexschema.FieldsInLists = &fieldsInLists		
	}
	if v, ok := data.GetOk("preview_urls"); ok {
		previewUrls := v.(map[string]string)
		squidexschema.PreviewUrls = &previewUrls
	}
	if v, ok := data.GetOk("fields"); ok {
		fields := v.([]squidexclient.UpsertSchemaFieldDto)
		squidexschema.Fields = &fields
		// TODO: get fields from data
	}
		//for _, _ := range fields {
			// append(squidexschema.Fields, squidexclient.UpsertSchemaFieldDto{
			// })

			/*
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
			"partitioning": {
				Type: schema.TypeString,
				Optional: true,
				Description: "Determines the optional partitioning of the field.",
			},
			"properties": {
				Type: schema.TypeList,
				Optional: true,
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
			*/

	return squidexschema, nil
}

func getUpdateSchemaDtoFromData(data *schema.ResourceData) (squidexclient.UpdateSchemaDto, error) {
	squidexschema := squidexclient.UpdateSchemaDto{}
	if v, ok := data.GetOk("properties"); ok {
		properties := v.([]interface{})[0].(map[string]interface{})
		if p, ok := properties["label"]; ok { 
			x := p.(string)
			squidexschema.Label = &x
		}
		if p, ok := properties["hints"]; ok { 
			x := p.(string)
			squidexschema.Hints = &x
		}
		if p, ok := properties["contents_sidebar_url"]; ok { 
			x := p.(string)
			squidexschema.ContentsSidebarUrl = &x
		}
		if p, ok := properties["content_sidebar_url"]; ok { 
			x := p.(string)
			squidexschema.ContentSidebarUrl = &x
		}
		if p, ok := properties["tags"]; ok { 
			tags := interfaceSliceToStringSlice(p.([]interface{}))
			squidexschema.Tags = &tags
		}
	}
	return squidexschema, nil
}

func resourceSchemaCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	
	dto, _ := getCreateSchemaDtoFromData(data)
	// TODO: handle all errors
	result, _, err := client.SchemasApi.SchemasPostSchema(ctx, appName, dto)

	if err != nil {
		return diag.FromErr(err)
	}

	// prevent drift and set ALL values from the result
	if err = setDataFromSchemaDetailsDto(data, result); err != nil {
		return diag.FromErr(err)
	}

	id := result.Id

	data.SetId(id)

	return diags
}

func resourceSchemaUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	dto, _ := getUpdateSchemaDtoFromData(data)
	// TODO: handle all errors
	result, _, err := client.SchemasApi.SchemasPutSchema(ctx, appName, name, dto)

	if err != nil {
		return diag.FromErr(err)
	}
	
	// prevent drift and set ALL values from the result
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

	_, err := client.SchemasApi.SchemasDeleteSchema(ctx, appName, name)

	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
