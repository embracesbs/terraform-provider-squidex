package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"

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
			"token_endpoint": &schema.Schema{
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
			"squidex_app":       resourceApp(),
			"squidex_client":    resourceClient(),
			"squidex_languages": resourceLanguage(),
			"squidex_schema": resourceSchema(),
			"squidex_role": resourceRole(),
			"squidex_contributor": resourceContributor(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	url := d.Get("url").(string)
	tokenEndpoint := d.Get("token_endpoint").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	var diags diag.Diagnostics

	config := &squidexclient.Configuration{
		BasePath:   url,
		HTTPClient: common.NewClient(clientID, clientSecret, tokenEndpoint, "squidex-api"),
	}

	return squidexclient.NewAPIClient(config), diags

}
