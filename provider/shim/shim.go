package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paultyng/terraform-provider-unifi/internal/provider"
)

func NewProvider() *schema.Provider {
	providerFunc := provider.New("dev") // TODO: pass correct version later
	return providerFunc()
}
