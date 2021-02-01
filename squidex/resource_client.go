package squidex

import (
	"context"
	"strings"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
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
			"invalidated_state": {
				Type: schema.TypeBool,
				Computed: true,
				Description: "Hidden field to invalidate state on response errors.",
			},
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

func setDataFromClientDto(data *schema.ResourceData, client *squidexclient.ClientDto) (error) {
	data.SetId(client.Id)
	data.Set("name", client.Name)
	data.Set("role", client.Role)

	return nil
}

func resourceClientRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	id := data.Id()
	
	data.Set("invalidated_state", false)
	
	result, response, err := client.AppsApi.AppClientsGetClients(ctx, appName)

	err = common.HandleAPIError(response, err)

	if err != nil {
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ClientDto
	for i := range result.Items {
    	if strings.EqualFold(result.Items[i].Id, id) {
			resultItem = &result.Items[i]
        	break
    	}
	}

	if resultItem == nil {
		return diag.Errorf("Not Found: Client with id %s", id)
	}

	err = setDataFromClientDto(data, resultItem)

	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: client setdata after read, and investigate for drifting
	return diags
}

func resourceClientCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	result, response, err := client.AppsApi.AppClientsPostClient(ctx, appName, squidexclient.CreateClientDto{
		Id: name,
	})

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ClientDto
	for i := range result.Items {
    	if strings.EqualFold(result.Items[i].Name, name) {
			resultItem = &result.Items[i]
        	break
    	}
	}

	if resultItem == nil {
		data.Set("invalidated_state", true)
		return diag.Errorf("Not Found: Client with name %s", name)
	}

	data.SetId(resultItem.Id)

	resourceClientUpdate(ctx, data, meta)

	return diags
}

func resourceClientUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	id := data.Id()
	name := data.Get("name").(string)
	role := data.Get("role").(string)

	_, response, err := client.AppsApi.AppClientsPutClient(ctx, appName, id, squidexclient.UpdateClientDto{
		Name: &name,
		Role: &role,
	})

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	return diags
}

func resourceClientDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	id := data.Get("id").(string)

	_, response, err := client.AppsApi.AppClientsDeleteClient(ctx, appName, id)

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
