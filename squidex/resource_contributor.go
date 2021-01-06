package squidex

import (
	"context"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceContributor() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceContributorRead,
		CreateContext: resourceContributorCreate,
		UpdateContext: resourceContributorUpdate,
		DeleteContext: resourceContributorDelete,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"contributor_id": &schema.Schema{
				Type:     schema.TypeString,
				Description: "The id of the user to add or delete to the app.",
				Required: true,
				ForceNew: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Description: "The role of the contributor.",
				Optional: true,
			},
			"invite": &schema.Schema{
				Type:     schema.TypeBool,
				Description: "Set to true to (email) invite the user, on creation, if he does not exist.",
				Required: true,
				Default: false,
			},
		},
	}
}

func resourceContributorRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorID := data.Id()
	
	result, _, err := client.AppsApi.AppContributorsGetContributors(ctx, appName)

	if err != nil {
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ContributorDto
	for i := range result.Items {
    	if result.Items[i].ContributorId == contributorID {
			resultItem = &result.Items[i]
        	break
    	}
	}

	if resultItem == nil {
		return diag.Errorf("Not Found: Contributor with Id %s", contributorID)
	}

	if err := data.Set("role", resultItem.Role); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceContributorCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorID := data.Get("contributor_id").(string)
	role := data.Get("role").(string)
	invite := data.Get("invite").(bool)

	_, _, err := client.AppsApi.AppContributorsPostContributor(ctx, appName, squidexclient.AssignContributorDto{
		ContributorId: contributorID,
		Role: &role,
		Invite: invite,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(contributorID)

	// there are no new values from the server, resourceContributorUpdate(ctx, data, meta)

	return diags
}

func resourceContributorUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorID := data.Get("contributor_id").(string)
	role := data.Get("role").(string)
	invite := false // no invites send for updating role
	
	// there is no update method, just use the create to set it again, but no invite!
	_, _, err := client.AppsApi.AppContributorsPostContributor(ctx, appName, squidexclient.AssignContributorDto{
		ContributorId: contributorID,
		Role: &role,
		Invite: invite,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	// there are no new values from the server, so no drifting, resourceContributorRead(ctx, data, meta)

	return diags

}

func resourceContributorDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)
	contributorID := data.Get("contributor_id").(string)

	client := meta.(*squidexclient.APIClient)
	var diags diag.Diagnostics
	_, _, err := client.AppsApi.AppContributorsDeleteContributor(ctx, appName, contributorID)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
