package squidex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLanguage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLanguageCreate,
		ReadContext:   resourceLanguageRead,
		UpdateContext: resourceLanguageUpdate,
		DeleteContext: resourceLanguageDelete,
		Schema: map[string]*schema.Schema{
			"iso_2_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_master": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_optional": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceLanguageCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	newLanguage := &Language{
		Iso2Code: data.Get("iso_2_code").(string),
		IsMaster: data.Get("is_master").(bool),
	}

	language, err := client.CreateLanguage(newLanguage)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("iso_2_code", language.Iso2Code); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("is_master", language.IsMaster); err != nil {
		return diag.FromErr(err)
	}
	data.SetId(language.Iso2Code)

	return diags
}

func resourceLanguageRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	languageIso2Code := data.Id()

	language, err := client.GetLanguage(languageIso2Code)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("iso_2_code", language.Iso2Code); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("is_master", language.IsMaster); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceLanguageUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	updateLanguage := &Language{
		Iso2Code: data.Get("iso_2_code").(string),
		IsMaster: data.Get("is_master").(bool),
	}

	language, err := client.UpdateLanguage(updateLanguage)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("iso_2_code", language.Iso2Code); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("is_master", language.IsMaster); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceLanguageDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	languageIso2Code := data.Id()

	err := client.DeleteLanguage(languageIso2Code)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
