package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceRoleRead,
		CreateContext: resourceRoleCreate,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		Schema: map[string]*schema.Schema{
			"invalidated_state": {
				Type: schema.TypeBool,
				Computed: true,
				Description: "Hidden field to invalidate state on response errors.",
			},
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permissions": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
		},
	}
}

func resourceRoleRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Id()
	
	data.Set("invalidated_state", false)

	result, response, err := client.AppsApi.AppRolesGetRoles(ctx, appName)

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.RoleDto
	for i := range result.Items {
    	if result.Items[i].Name == name {
			resultItem = &result.Items[i]
        	break
    	}
	}

	if resultItem == nil {
		return diag.Errorf("Not Found: Role with name %s", name)
	}

	if err := data.Set("permissions", resultItem.Permissions); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("properties", resultItem.Properties); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceRoleCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	_, response, err := client.AppsApi.AppRolesPostRole(ctx, appName, squidexclient.AddRoleDto{
		Name: name,
	})

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	data.SetId(name)

	resourceRoleUpdate(ctx, data, meta)

	return diags
}

func resourceRoleUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)
	permissions := toStringArray(data.Get("permissions").([]interface{}))
	properties := data.Get("properties").(map[string]interface{})
	
	_, response, err := client.AppsApi.AppRolesPutRole(ctx, appName, name, squidexclient.UpdateRoleDto{
		Permissions: permissions,
		Properties: &properties,
	})

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	// prevent drift:
	resourceRoleRead(ctx, data, meta)

	return diags

}

func resourceRoleDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)
	name := data.Get("name").(string)

	client := meta.(providerConfig).Client
	var diags diag.Diagnostics
	_, response, err := client.AppsApi.AppRolesDeleteRole(ctx, appName, name)

	err = common.HandleAPIError(response, err)
	
	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}

// TODO: move to helper class
func toStringArray(list []interface{}) []string {
	vs := make([]string, 0, len(list))
	for _, v := range list {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
