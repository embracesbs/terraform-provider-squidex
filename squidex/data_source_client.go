package squidex

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClients() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClientsRead,
		Schema: map[string]*schema.Schema{
			"clients": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClientsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	u, err := url.Parse(client.HostURL)
	if err != nil {
		return diag.FromErr(err)
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "clients")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	body, err := client.doRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	clients := Clients{}
	err = json.Unmarshal(body, &clients)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("clients", flattenClientsData(&clients.Items)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenClientsData(clientItems *[]Client) []interface{} {
	if clientItems == nil {
		return make([]interface{}, 0)
	}

	clients := make([]interface{}, len(*clientItems), len(*clientItems))

	for i, clientItem := range *clientItems {
		client := make(map[string]interface{})

		client["name"] = clientItem.Name
		client["secret"] = clientItem.Secret
		client["role"] = clientItem.Role

		clients[i] = client
	}

	return clients
}
