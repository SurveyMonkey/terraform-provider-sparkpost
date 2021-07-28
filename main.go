package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/SurveyMonkey/terraform-provider-sparkpost/internal/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: provider.New(version)}

	plugin.Serve(opts)
}
