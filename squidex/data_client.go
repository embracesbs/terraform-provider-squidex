package squidex

import (
	"context"
	"strings"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClient() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClientRead,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceClientRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(providerConfig).Client

	var diags diag.Diagnostics

	appName := d.Get("app_name").(string)
	clientName := d.Get("client_name").(string)

	result, response, err := client.AppsApi.AppClientsGetClients(ctx, appName)

	err = common.HandleAPIError(response, err)

	var resultItem *squidexclient.ClientDto
	for i := range result.Items {
		if strings.EqualFold(strings.ToLower(result.Items[i].Name), strings.ToLower(clientName)) {
			resultItem = &result.Items[i]
			break
		}
	}

	if resultItem == nil {
		return diag.Errorf("Not Found: Client with id %s", clientName)
	} else {
		d.Set("secret", resultItem.Secret)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
