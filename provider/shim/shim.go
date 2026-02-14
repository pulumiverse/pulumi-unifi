package shim

import (
	"github.com/filipowm/terraform-provider-unifi/internal/provider"
	x "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	providerFunc := provider.New("dev") // TODO: pass correct version later
	return providerFunc()
}

func NewProviderV2() x.Provider {
	providerFunc := provider.NewV2("dev") // TODO: pass correct version later
	return providerFunc()
}
