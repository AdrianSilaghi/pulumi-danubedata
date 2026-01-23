package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"

	danubedata "github.com/AdrianSilaghi/pulumi-danubedata/provider"
)

func main() {
	tfgen.Main("danubedata", danubedata.Provider())
}
