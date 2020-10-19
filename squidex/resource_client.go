package squidex

import (
	"context"

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
	apiClient := meta.(*APIClient)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := data.Id()

	client, err := apiClient.GetClientByName(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(client.Name)

	if err := data.Set("name", client.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("role", client.Role); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceClientCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	newClient := &Client{
		Name: data.Get("name").(string),
	}

	client, err := apiClient.CreateClient(newClient)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newClient.Name)

	if err := data.Set("name", client.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("role", client.Role); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceClientUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*APIClient)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	updateClient := &Client{
		Name: data.Get("name").(string),
		Role: data.Get("role").(string),
	}

	client, err := apiClient.UpdateClient(updateClient)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(updateClient.Name)

	if err := data.Set("name", client.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("role", client.Role); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceClientDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*APIClient)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := data.Id()

	err := apiClient.DeleteClient(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
