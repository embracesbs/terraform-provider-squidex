package squidex

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
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
			"invalidated_state": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hidden field to invalidate state on response errors.",
			},
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"contributor_email": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The emailaddress of the user to add to the app.",
				Required:    true,
				ForceNew:    true,
			},
			"role": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The role of the contributor.",
				Required:    true,
			},
			// Disabled as input, because new users must have it 'true' and updating is always with 'false'
			//"invite": &schema.Schema{
			//	Type:     schema.TypeBool,
			//	Description: "Set to true to (email) invite the user, on creation, if he does not exist.",
			//	Optional: true,
			//	Default: false,
			//},
		},
	}
}

func resourceContributorRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorID := data.Id()

	data.Set("invalidated_state", false)

	result, response, err := client.AppsApi.AppContributorsGetContributors(ctx, appName)

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ContributorDto
	for i := range result.Items {
		if strings.EqualFold(result.Items[i].ContributorId, contributorID) {
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

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorEmail := data.Get("contributor_email").(string)
	role := data.Get("role").(string)
	invite := true // required for new users, only existing users may be added as contributer without invite data.Get("invite").(bool)

	result, response, err := client.AppsApi.AppContributorsPostContributor(ctx, appName, squidexclient.AssignContributorDto{
		ContributorId: contributorEmail,
		Role:          &role,
		Invite:        invite,
	})

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ContributorDto
	for i := range result.Items {
		if strings.EqualFold(result.Items[i].ContributorEmail, contributorEmail) {
			resultItem = &result.Items[i]
			break
		}
	}

	if resultItem == nil {
		data.Set("invalidated_state", true)
		source, _ := json.Marshal(result)
		return diag.Errorf("Not Found: Contributor with email %s\nApi Response: %s", contributorEmail, string(source))
	}

	data.SetId(resultItem.ContributorId)

	// there are no new values from the server, resourceContributorUpdate(ctx, data, meta)

	return diags
}

func resourceContributorUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	contributorID := data.Id()
	role := data.Get("role").(string)
	invite := false // no invites send for updating role

	// there is no update method, just use the create to set it again, but no invite!
	result, response, err := client.AppsApi.AppContributorsPostContributor(ctx, appName, squidexclient.AssignContributorDto{
		ContributorId: contributorID,
		Role:          &role,
		Invite:        invite,
	})

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.ContributorDto
	for i := range result.Items {
		if strings.EqualFold(result.Items[i].ContributorId, contributorID) {
			resultItem = &result.Items[i]
			break
		}
	}

	if resultItem == nil {
		data.Set("invalidated_state", true)
		source, _ := json.Marshal(result)
		return diag.Errorf("Not Found: Contributor with id %s\nApi Response: %s", contributorID, string(source))
	}

	data.SetId(resultItem.ContributorId)

	// there are no new values from the server, so no drifting, resourceContributorRead(ctx, data, meta)

	return diags

}

func resourceContributorDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)
	contributorID := data.Id()

	client := meta.(providerConfig).Client
	var diags diag.Diagnostics

	_, response, err := client.AppsApi.AppContributorsDeleteContributor(ctx, appName, contributorID)

	err = common.HandleAPIError(response, err, true)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
