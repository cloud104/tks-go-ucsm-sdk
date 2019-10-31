package util

import (
	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

func SptInstantiate(c *api.Client, sptDn string, spOrg string, spNamePrefix string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	req := api.LsInstantiateNTemplateRequest {
		    Cookie: c.Cookie,
		    Dn: sptDn,
		    InTargetOrg: spOrg,
		    InServerNamePrefixOrEmpty: spNamePrefix,
		    InNumberOf: "1",
		    InHierarchical: "false",
	}
	if err = c.LsInstantiateNTemplate(req, &out); err == nil {
		lsServer = &out.LsServer
	}
	return
}
