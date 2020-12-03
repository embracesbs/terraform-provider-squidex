package main

import (
	"github.com/embracesbs/terraform-provider-squidex/squidex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: squidex.Provider,
	})
}
