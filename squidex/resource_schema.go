package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceSchemaRead,
		CreateContext: resourceSchemaCreate,
		UpdateContext: resourceSchemaUpdate,
		DeleteContext: resourceSchemaDelete,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"published": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"singleton": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceSchemaRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceSchemaCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)
	published := data.Get("published").(bool)
	singleton := data.Get("singleton").(bool)

	result, _, err := client.SchemasApi.SchemasPostSchema(ctx, appName, squidexclient.CreateSchemaDto{
		Name: name,
		IsPublished: published,
		IsSingleton: singleton,
	})

	if err != nil {
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

	_, _, err := client.SchemasApi.SchemasPutSchema(ctx, appName, name, squidexclient.UpdateSchemaDto{
		// TODO: add fields
	})

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

	_, err := client.SchemasApi.SchemasDeleteSchema(ctx, appName, name)

	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
