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
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"schema_delete_allow": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Allow deleting schemas.",
			},
			// TODO: add settings to allow/forbid 1. deleting schemas 2. deleting fields of schemas 3. recreating fields of schemas
			// if not allowed, only delete terraform references, making it no longer under terraform control but no changes to server are made.
			// when they are needed again, they need to be imported first
			// this will serve as a safeguard for accidental removal of data
			// needs discussion
		},
		ResourcesMap: map[string]*schema.Resource{
			"squidex_app":         resourceApp(),
			"squidex_client":      resourceClient(),
			"squidex_languages":   resourceLanguage(),
			"squidex_schema":      resourceSchema(),
			"squidex_role":        resourceRole(),
			"squidex_contributor": resourceContributor(),
			"squidex_rule":        resourceRule(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

type providerConfig struct {
	Client              *squidexclient.APIClient
	SchemaDeleteAllowed bool
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

	return providerConfig{
		Client:              squidexclient.NewAPIClient(config),
		SchemaDeleteAllowed: d.Get("schema_delete_allow").(bool),
	}, diags
	// return squidexclient.NewAPIClient(config), diags

}
