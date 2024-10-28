//go:build nodejs || all
// +build nodejs all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumiverse/unifi",
		},
	})

	return baseJS
}

// func TestNetworkExampleTypescript(t *testing.T) {
//     test := getJSBaseOptions(t).
//         With(integration.ProgramTestOptions{
//             Dir: filepath.Join(getCwd(t), "network", "typescript"),
//         })
//     integration.ProgramTest(t, &test)
// }
