// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package setting

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "unifi:setting/mgmt:Mgmt":
		r = &Mgmt{}
	case "unifi:setting/radius:Radius":
		r = &Radius{}
	case "unifi:setting/uSG:USG":
		r = &USG{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"unifi",
		"setting/mgmt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"unifi",
		"setting/radius",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"unifi",
		"setting/uSG",
		&module{version},
	)
}
