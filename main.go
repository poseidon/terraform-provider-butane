package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/poseidon/terraform-provider-butane/internal"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: internal.Provider,
	})
}
