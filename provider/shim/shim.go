package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	tfshim "github.com/AdrianSilaghi/terraform-provider-danubedata/shim"
)

// NewProvider returns the DanubeData Terraform provider factory for bridging to Pulumi.
func NewProvider(version string) func() provider.Provider {
	return tfshim.NewProvider(version)
}
