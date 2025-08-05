package util

import (
	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SpSetIscsiBoot(c *api.Client, spDn string, iscsiVnicName string, iscsiInitiatorName string, iscsiVnicAddr mo.VnicIPv4IscsiAddr, iscsiTargets []mo.VnicIScsiStaticTargetIf) (vnicIscsi *mo.VnicIScsi, err error) {
	var out mo.VnicIScsiPairs
	vnicIScsiPairs := mo.VnicIScsiPairs{
		Pairs: []mo.VnicIScsiPair{
			{
				Key: spDn + "/iscsi-" + iscsiVnicName,
				VnicIScsi: mo.VnicIScsi{
					InitiatorName: iscsiInitiatorName,
					VnicVlan: mo.VnicVlan{
						VnicIPv4If: mo.VnicIPv4If{
							VnicIPv4IscsiAddr: iscsiVnicAddr,
							VnicIPv4Dhcp: mo.VnicIPv4Dhcp{
								Status: "deleted",
							},
						},
						VnicIScsiStaticTargets: iscsiTargets,
					},
				},
			},
		},
	}
	req := api.ConfigConfMosRequest{
		Cookie:         c.Cookie,
		InHierarchical: "true",
		InConfigs:      vnicIScsiPairs,
	}
	if err = c.ConfigConfMos(req, &out); err == nil {
		for _, pair := range out.Pairs {
			if pair.VnicIScsi.Dn == spDn+"/iscsi-"+iscsiVnicName {
				vnicIscsi = &pair.VnicIScsi
			}
		}
	}
	return
}
