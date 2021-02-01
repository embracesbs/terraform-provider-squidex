package squidex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/common"
)

func resourceLanguage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLanguageCreate,
		ReadContext:   resourceLanguageRead,
		UpdateContext: resourceLanguageUpdate,
		DeleteContext: resourceLanguageDelete,
		Schema: map[string]*schema.Schema{
			"invalidated_state": {
				Type: schema.TypeBool,
				Computed: true,
				Description: "Hidden field to invalidate state on response errors.",
			},
			"app_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"language": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"language": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: common.ValidateLanguage,
						},
						"is_master": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"language_backup": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"language": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_master": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

const ( 
	Sleep = "sleep" 
	Update = "update" 
	Create = "create" 
)

func resourceLanguageRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	//todo: inplement read thingy
	var diags diag.Diagnostics
	
	data.Set("invalidated_state", false)

	return diags
}

func resourceLanguageCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	result := data.Get("language").([]interface{})
	appName := data.Get("app_name").(string)

	var diags diag.Diagnostics

	client := meta.(providerConfig).Client

	var hasEn = false

	for _, element := range result {

		rs := element.(map[string]interface{})

		if rs["language"].(string) == "en" {
			hasEn = true
		} else {
			_, _, err := client.AppsApi.AppLanguagesPostLanguage(ctx, appName, squidexclient.AddLanguageDto{
				Language: rs["language"].(string),
			})

			if err != nil {
				data.Set("invalidated_state", true)
				return diag.Errorf("failed to add language")
			}
		}

		if rs["is_master"].(bool) == true {

			bl := true

			_, _, err := client.AppsApi.AppLanguagesPutLanguage(ctx, appName, rs["language"].(string), squidexclient.UpdateLanguageDto{
				IsMaster: &bl,
			})

			if err != nil {
				data.Set("invalidated_state", true)
				return diag.Errorf("failed to set master")
			}
		}

		//Set backup for compare upate
		data.Set("language_backup", result)

		data.SetId("1")

	}

	if hasEn == false {
		_, _, err := client.AppsApi.AppLanguagesDeleteLanguage(ctx, appName, "en")

		if err != nil {
			data.Set("invalidated_state", true)
			return diag.FromErr(err)
		}
	}

	return diags

}

func resourceLanguageUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)
	result := data.Get("language").([]interface{})
	oldResult := data.Get("language_backup").([]interface{})

	var diags diag.Diagnostics

	client := meta.(providerConfig).Client

	for _, element := range result {
		rs := element.(map[string]interface{})
		name := rs["language"].(string)
		isActive := rs["is_master"].(bool)
		todo := CheckUpdate(rs, result, oldResult)

		if todo == Create {
			_, _, err := client.AppsApi.AppLanguagesPostLanguage(ctx, appName, squidexclient.AddLanguageDto{
				Language: name,
			})

			if err != nil {
				data.Set("invalidated_state", true)
				return diag.Errorf("failed to perform update :  create " + name)
			}

			_, _, err_update := client.AppsApi.AppLanguagesPutLanguage(ctx, appName, name, squidexclient.UpdateLanguageDto{
				IsMaster: &isActive,
			})

			if err_update != nil {
				data.Set("invalidated_state", true)
				return diag.Errorf("failed to perform update :  set active " + name)
			}
		}

		if todo == Update {

			_, _, err_update := client.AppsApi.AppLanguagesPutLanguage(ctx, appName, name, squidexclient.UpdateLanguageDto{
				IsMaster: &isActive,
			})

			if err_update != nil  {

				return diag.Errorf("failed to perform update :  Update " + name)
			}

		}
	}


	for _, element := range oldResult {

		rs := element.(map[string]interface{})
		name := rs["language"].(string)
		isActive := rs["is_master"].(bool)

		if CheckDelete(rs,result) == false {

			if isActive == true {
				data.Set("invalidated_state", true)
				return diag.Errorf("cannot delete language that is master " + name)
			} 


			client.AppsApi.AppLanguagesDeleteLanguage(ctx, appName, name)

		}


	}



	return diags

}

func resourceLanguageDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	appName := data.Get("app_name").(string)

	var diags diag.Diagnostics

	result := data.Get("language").([]interface{})

	client := meta.(providerConfig).Client

	existing, _, err := client.AppsApi.AppLanguagesGetLanguages(ctx, appName)

	if err != nil {
		return diag.FromErr(err)
	}

	for _, element := range result {

		rs := element.(map[string]interface{})

		if rs["language"].(string) != "en" {

			if Same(existing, rs["language"].(string), rs["is_master"].(bool)) == true && rs["is_master"].(bool) == true {

				if Same(existing, "en", false) != true {

					_, _, err := client.AppsApi.AppLanguagesPostLanguage(ctx, appName, squidexclient.AddLanguageDto{
						Language: "en",
					})

					if err != nil {
						return diag.Errorf("failed to recreate default en language")
					}

				}
				bl := true
				client.AppsApi.AppLanguagesPutLanguage(ctx, appName, "en", squidexclient.UpdateLanguageDto{
					IsMaster: &bl,
				})
			}

			_, _, err := client.AppsApi.AppLanguagesDeleteLanguage(ctx, appName, rs["language"].(string))

			if err != nil {
				return diag.Errorf("failed to delete: " + rs["language"].(string))
			}

		}

	}

	data.SetId("")

	return diags
}

func CheckDelete(o map[string]interface{}, n []interface{}) bool {
	oldName := o["language"].(string)

	var found = false

	for _, element := range n {
		rs := element.(map[string]interface{})

		name := rs["language"].(string)

		if oldName == name {

			found = true

		}

	}

	return found
}

func CheckUpdate(n map[string]interface{}, na []interface{}, o []interface{}) string {

	updateName := n["language"].(string)
	updateIsActive := n["is_master"].(bool)

	for _, element := range o {
		rs := element.(map[string]interface{})

		name := rs["language"].(string)
		isActive := rs["is_master"].(bool)

		if name == updateName {

			if updateIsActive != isActive {
				return Update
			}
		}else{
			return Create
		}

	}

	return Sleep
}

func Same(s squidexclient.AppLanguagesDto, b string, c bool) bool {

	for _, a := range s.Items {

		if a.Iso2Code == b && a.IsMaster == c {
			return true
		}
	}
	return false
}
