package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/kofan/terraform-provider-cloudflare/cloudflare"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudflare.Provider})
}
