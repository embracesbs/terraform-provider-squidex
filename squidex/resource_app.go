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
		},
	}
}

func resourceAppRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceAppCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	name := data.Get("name").(string)

	result, _, err := client.AppsApi.AppsPostApp(ctx, squidexclient.CreateAppDto{
		Name: name,
	})

	if err == nil {
		//todo: error log in output
	}

	data.SetId(result.Id)

	return nil
}

func resourceAppUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAppDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
