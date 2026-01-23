package provider

import (
	_ "embed"
	"path/filepath"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/AdrianSilaghi/pulumi-danubedata/provider/shim"
)

const (
	// mainPkg is the name of this package.
	mainPkg = "danubedata"
	// mainMod is the module name for the main package.
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-danubedata/bridge-metadata.json
var metadata []byte

// Provider returns the DanubeData Pulumi provider.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		// Use the PF bridge to wrap the Terraform Plugin Framework provider.
		P: pf.ShimProvider(shim.NewProvider("dev")()),

		Name:        "danubedata",
		DisplayName: "DanubeData",
		Publisher:   "AdrianSilaghi",
		LogoURL:     "",
		PluginDownloadURL: "github://api.github.com/AdrianSilaghi/pulumi-danubedata",
		Description: "A Pulumi provider for managing DanubeData cloud infrastructure resources.",
		Keywords: []string{
			"pulumi",
			"danubedata",
			"category/cloud",
			"kind/native",
		},
		License:    "Apache-2.0",
		Homepage:   "https://danubedata.ro",
		Repository: "https://github.com/AdrianSilaghi/pulumi-danubedata",
		GitHubOrg:  "AdrianSilaghi",

		MetadataInfo: tfbridge.NewProviderMetadata(metadata),

		Config: map[string]*tfbridge.SchemaInfo{
			"api_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DANUBEDATA_API_TOKEN"},
				},
				Secret: tfbridge.True(),
			},
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DANUBEDATA_BASE_URL"},
					Value:   "https://danubedata.ro/api/v1",
				},
			},
		},

		Resources: map[string]*tfbridge.ResourceInfo{
			// Compute
			"danubedata_vps":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vps")},
			"danubedata_serverless": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Serverless")},

			// Data Services
			"danubedata_cache":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cache")},
			"danubedata_database": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Database")},

			// Storage
			"danubedata_storage_bucket":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageBucket")},
			"danubedata_storage_access_key": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageAccessKey")},

			// Security
			"danubedata_ssh_key":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SshKey")},
			"danubedata_firewall": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Firewall")},

			// Backup
			"danubedata_vps_snapshot": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpsSnapshot")},
		},

		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Static data sources
			"danubedata_vps_images": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpsImages"),
			},
			"danubedata_cache_providers": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCacheProviders"),
			},
			"danubedata_database_providers": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDatabaseProviders"),
			},
			"danubedata_ssh_keys": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSshKeys"),
			},

			// Resource listing data sources
			"danubedata_vpss": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpss"),
			},
			"danubedata_databases": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDatabases"),
			},
			"danubedata_caches": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCaches"),
			},
			"danubedata_firewalls": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFirewalls"),
			},
			"danubedata_serverless_containers": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getServerlessContainers"),
			},
			"danubedata_storage_buckets": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageBuckets"),
			},
			"danubedata_storage_access_keys": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageAccessKeys"),
			},
			"danubedata_vps_snapshots": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpsSnapshots"),
			},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@danubedata/pulumi",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
			},
		},

		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_danubedata",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				"github.com/AdrianSilaghi/pulumi-danubedata",
				"sdk",
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},

		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace: "DanubeData",
			Namespaces: map[string]string{
				mainPkg: "DanubeData",
			},
		},
	}

	// MustComputeTokens will compute Pulumi tokens for all resources and data sources.
	prov.MustComputeTokens(tokens.SingleModule(mainPkg, mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
