package util

import (
	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SptInstantiate(c *api.Client, sptDn string, spOrg string, spName string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	req := api.LsInstantiateNNamedTemplateRequest{
		Cookie:      c.Cookie,
		Dn:          sptDn,
		InTargetOrg: spOrg,
		InNameSet: []api.Dn{
			{
				Value: spName,
			},
		},
		InHierarchical: "false",
	}
	if err = c.LsInstantiateNNamedTemplate(req, &out); err == nil {
		lsServer = &out.LsServer
	}
	return
}
