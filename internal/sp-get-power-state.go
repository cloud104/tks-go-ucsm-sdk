package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetPowerState(c *api.Client, spDn string) (powerState string, err error) {
	var out mo.LsPowerMo
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "lsPower",
		InHierarchical: "false",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		powerState = out.LsPower.State
	}
	return powerState, fmt.Errorf("failed to resolve power state: %w", err)
}
