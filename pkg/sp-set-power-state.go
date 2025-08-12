package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpSetPowerState(c *api.Client, spDn string, powerState string) (*mo.LsPower, error) {
	var out mo.LsPowerMo
	lsPowerMo := mo.LsPowerMo{
		LsPower: mo.LsPower{
			Dn:    spDn + "/power",
			State: powerState,
		},
	}
	req := api.ConfigConfMoRequest{
		Cookie:         c.Cookie,
		Dn:             spDn,
		InHierarchical: "false",
		InConfig:       lsPowerMo,
	}
	if err := c.ConfigConfMo(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve powerstate: %w", err)
	}
	return &out.LsPower, nil
}
