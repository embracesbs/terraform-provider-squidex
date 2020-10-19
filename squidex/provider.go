package squidex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"squidex_language": resourceLanguage(),
			"squidex_client":   resourceClient(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"squidex_languages": dataSourceLanguages(),
			"squidex_clients":   dataSourceClients(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	appName := d.Get("app_name").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (clientID != "") && (clientSecret != "") {
		client, err := NewClient(&url, &appName, &clientID, &clientSecret)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return client, diags
	}

	client, err := NewClient(nil, nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}
