module github.com/pulumiverse/pulumi-unifi/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220505215311-795430389fa7
	github.com/paultyng/terraform-provider-unifi/shim => ./shim
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	github.com/paultyng/terraform-provider-unifi/shim v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.25.2
	github.com/pulumi/pulumi/sdk/v3 v3.35.3
)
