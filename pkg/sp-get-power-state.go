package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetPowerState(c *api.Client, spDn string) (string, error) {
	var out mo.LsPowerMo
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "lsPower",
		InHierarchical: "false",
	}
	if err := c.ConfigResolveChildren(req, &out); err != nil {
		return "", fmt.Errorf("power state is not available for service profile %s", spDn)
	}
	return out.LsPower.State, nil
}
