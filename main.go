package main

import (
	"log"
	"os"

	"github.com/embracesbs/terraform-provider-squidex/squidex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Printf("[TRACE] Running version %s!", meta.SDKVersionString())

	opts := &plugin.ServeOpts{
		ProviderFunc: squidex.Provider,
	}

	/// Uncomment to debug the plugin
	// if true {
	// 	err := plugin.Debug(context.Background(), "terraform.embracecloud.nl/embracecloud/squidex", opts)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	return
	// }

	plugin.Serve(opts)
}
