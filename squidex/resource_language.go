package squidex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
)

func resourceLanguage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLanguageCreate,
		ReadContext:   resourceLanguageRead,
		UpdateContext: resourceLanguageUpdate,
		DeleteContext: resourceLanguageDelete,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_master": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceLanguageRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	//todo: inplement read thingy
	var diags diag.Diagnostics

	return diags
}

func resourceLanguageCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	language := data.Get("language").(string)

	result, _, err := client.AppsApi.AppLanguagesPostLanguage(ctx, appName, squidexclient.AddLanguageDto{
		Language: language,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Items[0].EnglishName)

	resourceLanguageUpdate(ctx, data, meta)

	return diags
}

func resourceLanguageUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	language := data.Get("language").(string)
	isMaster := data.Get("is_master").(*bool)

	_, _, err := client.AppsApi.AppLanguagesPutLanguage(ctx, appName, language, squidexclient.UpdateLanguageDto{
		IsMaster: isMaster,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}

func resourceLanguageDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*squidexclient.APIClient)

	var diags diag.Diagnostics

	appName := data.Get("app_name").(string)
	language := data.Get("language").(string)

	_, _, err := client.AppsApi.AppLanguagesDeleteLanguage(ctx, appName, language)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags
}
