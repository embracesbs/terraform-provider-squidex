package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAppRead,
		CreateContext: resourceAppCreate,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAppRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags
}

func resourceAppCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	name := data.Get("name").(string)

	result, _, err := client.AppsApi.AppsPostApp(ctx, squidexclient.CreateAppDto{
		Name: name,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Id)

	resourceAppUpdate(ctx, data, meta)

	return diags
}

func resourceAppUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	name := data.Get("name").(string)
	description := data.Get("description").(*string)

	_, _, err := client.AppsApi.AppsUpdateApp(ctx, name, squidexclient.UpdateAppDto{
		Description: description,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}

func resourceAppDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := data.Get("name").(string)

	client := meta.(*squidexclient.APIClient)
	var diags diag.Diagnostics
	_, err := client.AppsApi.AppsDeleteApp(ctx, name)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
