package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
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
			"invalidated_state": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hidden field to invalidate state on response errors.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAppRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// TODO: all read stuff
	data.Set("invalidated_state", false)
	return diags
}

func resourceAppCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	name := data.Get("name").(string)

	result, _, err := client.AppsApi.AppsPostApp(ctx, squidexclient.CreateAppDto{
		Name: name,
	})

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	data.SetId(result.Id)

	resourceAppUpdate(ctx, data, meta)

	return diags
}

func resourceAppUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	name := data.Get("name").(string)
	label := data.Get("label").(string)
	description := data.Get("description").(string)

	_, _, err := client.AppsApi.AppsPutApp(ctx, name, squidexclient.UpdateAppDto{
		Label:       &label,
		Description: &description,
	})

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	return diags

}

func resourceAppDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := data.Get("name").(string)

	client := meta.(providerConfig).Client
	var diags diag.Diagnostics
	response, err := client.AppsApi.AppsDeleteApp(ctx, name)

	err = common.HandleAPIError(response, err, true)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
