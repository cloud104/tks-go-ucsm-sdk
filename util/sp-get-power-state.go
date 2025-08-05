package util

import (
	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SpGetPowerState(c *api.Client, spDn string) (powerState string, err error) {
	var out mo.LsPowerMo
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassId:        "lsPower",
		InHierarchical: "false",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		powerState = out.LsPower.State
	}
	return
}
