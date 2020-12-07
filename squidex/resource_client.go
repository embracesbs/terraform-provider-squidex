package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClient() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceClientRead,
		CreateContext: resourceClientCreate,
		UpdateContext: resourceClientUpdate,
		DeleteContext: resourceClientDelete,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceClientRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceClientCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	result, _, err := client.AppsApi.AppClientsPostClient(ctx, appName, squidexclient.CreateClientDto{
		Id: name,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	id := result.Items[0].Id

	data.SetId(id)

	//resourceClientUpdate(ctx, data, meta)

	return diags
}

func resourceClientUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)
	role := data.Get("role").(*string)

	_, _, err := client.AppsApi.AppClientsPutClient(ctx, appName, name, squidexclient.UpdateClientDto{
		Name: name,
		Role: role,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceClientDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	_, _, err := client.AppsApi.AppClientsDeleteClient(ctx, appName, name)

	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
