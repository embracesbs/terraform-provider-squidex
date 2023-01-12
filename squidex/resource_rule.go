package squidex

import (
	"context"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRule() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceRuleRead,
		CreateContext: resourceRuleCreate,
		UpdateContext: resourceRuleUpdate,
		DeleteContext: resourceRuleDelete,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Url used for a webhook rule.",
				Optional:    true,
			},
			"method": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Method used for a webhook rule.",
			},
			"trigger_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the rules trigger.",
			},
			"action_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the rules action.",
			},
			"shared_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Shared secret for a webhook rule.",
			},
			"handle_all": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Handle all",
			},
		},
	}
}

func resourceRuleRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	name := data.Id()

	data.Set("invalidated_state", false)

	result, response, err := client.RulesApi.RulesGetRules(ctx, appName)

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		return diag.FromErr(err)
	}

	var resultItem *squidexclient.RuleDto
	for i := range result.Items {
		if result.Items[i].Id == name {
			resultItem = &result.Items[i]
			break
		}
	}

	if resultItem == nil {
		return diag.Errorf("Not Found: Rule with name %s", name)
	}

	return diags
}

func resourceRuleCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	url := data.Get("url").(string)
	method := data.Get("method").(string)
	triggerType := data.Get("trigger_type").(string)
	actionType := data.Get("action_type").(string)
	sharedSecret := data.Get("shared_secret").(string)
	handleAll := data.Get("handle_all").(bool)

	createdRule, response, err := client.RulesApi.RulesPostRule(ctx, appName, squidexclient.CreateRuleDto{
		Trigger: squidexclient.RuleTriggerDto{
			TriggerType: triggerType,
			HandleAll:   handleAll,
		},
		Action: squidexclient.RuleAction{
			ActionType:   actionType,
			Url:          url,
			Method:       method,
			SharedSecret: sharedSecret,
		},
	})

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	data.SetId(createdRule.Id)

	resourceRuleUpdate(ctx, data, meta)

	return diags
}

func resourceRuleUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(providerConfig).Client

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	id := data.Id()
	name := data.Get("name").(string)
	url := data.Get("url").(string)
	method := data.Get("method").(string)
	triggerType := data.Get("trigger_type").(string)
	actionType := data.Get("action_type").(string)
	sharedSecret := data.Get("shared_secret").(string)
	handleAll := data.Get("handle_all").(bool)

	_, response, err := client.RulesApi.RulesPutRule(ctx, appName, id, squidexclient.UpdateRuleDto{
		Name: &name,
		Trigger: squidexclient.RuleTriggerDto{
			TriggerType: triggerType,
			HandleAll:   handleAll,
		},
		Action: squidexclient.RuleAction{
			ActionType:   actionType,
			Url:          url,
			Method:       method,
			SharedSecret: sharedSecret,
		},
		IsEnabled: true,
	})

	err = common.HandleAPIError(response, err, false)

	if err != nil {
		data.Set("invalidated_state", true)
		return diag.FromErr(err)
	}

	// prevent drift:
	resourceRuleRead(ctx, data, meta)

	return diags

}

func resourceRuleDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)
	id := data.Id()

	client := meta.(providerConfig).Client
	var diags diag.Diagnostics
	response, err := client.RulesApi.RulesDeleteRule(ctx, appName, id)

	err = common.HandleAPIError(response, err, true)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
