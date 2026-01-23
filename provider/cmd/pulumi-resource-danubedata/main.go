package main

import (
	"context"
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"

	danubedata "github.com/AdrianSilaghi/pulumi-danubedata/provider"
)

//go:embed schema.json
var pulumiSchema []byte

func main() {
	meta := tfbridge.ProviderMetadata{PackageSchema: pulumiSchema}
	tfbridge.Main(context.Background(), "danubedata", danubedata.Provider(), meta)
}
