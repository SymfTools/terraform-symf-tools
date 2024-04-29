package main

import (
	"github.com/Boltairex/terraform-symf-tools/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: json.Provider,
	})
}