package squidex

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Languages struct {
	Items []Language  `json:"items"`
	Links interface{} `json:"_links"`
}

type Language struct {
	Iso2Code    string      `json:"iso2Code"`
	EnglishName string      `json:"englishName"`
	Fallback    []string    `json:"fallback"`
	IsMaster    bool        `json:"isMaster"`
	IsOptional  bool        `json:"isOptional"`
	Links       interface{} `json:"_links"`
}

func dataSourceAppLanguages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppLanguagesRead,
		Schema: map[string]*schema.Schema{
			"languages": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_2_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_master": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_optional": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAppLanguagesRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	u, err := url.Parse(client.HostURL)
	if err != nil {
		return diag.FromErr(err)
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "languages")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	body, err := client.doRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)
	if err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		log.Println(err)
		return diag.FromErr(err)
	}

	if err := data.Set("languages", flattenLanguagesData(&languages.Items)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenLanguagesData(languageItems *[]Language) []interface{} {
	if languageItems != nil {
		ois := make([]interface{}, len(*languageItems), len(*languageItems))

		for i, language := range *languageItems {
			oi := make(map[string]interface{})

			oi["iso_2_code"] = language.Iso2Code
			oi["is_master"] = language.IsMaster
			oi["is_optional"] = language.IsOptional

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
