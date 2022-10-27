package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/embracesbs/terraform-provider-squidex/squidex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	log.SetOutput(os.Stderr)
	log.Printf("[TRACE] Running version %s!", meta.SDKVersionString())

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: squidex.Provider,
	}

	/// debug the plugin
	if debugMode {
		err := plugin.Debug(context.Background(), "terraform.embracecloud.nl/embracecloud/squidex", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
