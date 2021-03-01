package main

import (
	"log"
	"os"

	"github.com/embracesbs/terraform-provider-squidex/squidex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
    log.SetOutput(os.Stderr)
    log.Println("[INFO] Does this even work?")

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: squidex.Provider,
	})
}
