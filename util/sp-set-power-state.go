package util

import (
	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

func SpSetPowerState(c *api.Client, spDn string, powerState string) (lsPower *mo.LsPower, err error) {
	var out mo.LsPowerMo
	lsPowerMo := mo.LsPowerMo {
			LsPower: mo.LsPower {
				    Dn: spDn + "/power",
				    State: powerState,
			},
	}
	req := api.ConfigConfMoRequest {
		    Cookie: c.Cookie,
		    Dn: spDn,
		    InHierarchical: "false",
		    InConfig: lsPowerMo,
	}
	if err = c.ConfigConfMo(req, &out); err == nil {
		lsPower = &out.LsPower
	}
	return
}
