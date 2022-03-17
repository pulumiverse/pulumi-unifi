// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unifi

import (
	"fmt"
	"path/filepath"

	provider "github.com/paultyng/terraform-provider-unifi/shim"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-unifi/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "unifi"
	// modules:
	mainMod = "index" // the unifi module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(provider.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "unifi",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Unifi",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/pulumiverse/.github/main/assets/mascot.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing Unifi network resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "unifi", "category/network"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumiverse",
		Repository: "https://github.com/pulumiverse/pulumi-unifi",
		// The GitHub Org for the provider - defaults to `terraform-providers`
		GitHubOrg: "paultyng",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"unifi_device":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Device")},
			"unifi_dynamic_dns":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DynamicDNS")},
			"unifi_firewall_group": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallGroup")},
			"unifi_firewall_rule":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallRule")},
			"unifi_network":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Network")},
			"unifi_port_forward":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PortForward")},
			"unifi_port_profile":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PortProfile")},
			"unifi_setting_mgmt":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SettingMgmt")},
			"unifi_setting_usg":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SettingUSG")},
			"unifi_site":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Site")},
			"unifi_static_route":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StaticRoute")},
			"unifi_user":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"unifi_user_group":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserGroup")},
			"unifi_wlan":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Wlan")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"unifi_ap_group":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApGroup")},
			"unifi_network":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetwork")},
			"unifi_port_profile":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPortProfile")},
			"unifi_radius_profile": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRadiusProfile")},
			"unifi_user":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"unifi_user_group":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserGroup")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
